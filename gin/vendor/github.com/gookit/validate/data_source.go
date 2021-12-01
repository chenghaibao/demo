package validate

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/filter"
	"github.com/gookit/goutil/strutil"
)

type sourceType uint8

const (
	// from user setting, unmarshal JSON
	sourceMap sourceType = iota + 1
	// from URL.Values, PostForm. contains Files data
	sourceForm
	// from user setting
	sourceStruct
)

// 0: top level field
// 1: field at anonymous struct
// 2: field at non-anonymous struct
const (
	fieldAtTopStruct int8 = iota
	fieldAtAnonymous
	fieldAtSubStruct
)

var timeType = reflect.TypeOf(time.Time{})

// data (Un)marshal func
var (
	Marshal   MarshalFunc   = json.Marshal
	Unmarshal UnmarshalFunc = json.Unmarshal
)

type (
	// MarshalFunc define
	MarshalFunc func(v interface{}) ([]byte, error)
	// UnmarshalFunc define
	UnmarshalFunc func(data []byte, v interface{}) error
)

// DataFace data source interface definition
//
// current has three data source:
// - map
// - form
// - struct
type DataFace interface {
	Type() uint8
	Get(key string) (interface{}, bool)
	Set(field string, val interface{}) (interface{}, error)
	// validation instance create func
	Create(err ...error) *Validation
	Validation(err ...error) *Validation
}

/*************************************************************
 * Map Data
 *************************************************************/

// MapData definition
type MapData struct {
	// Map the source map data
	Map map[string]interface{}
	// from reflect Map
	value reflect.Value
	// bodyJSON from the original JSON bytes/string.
	// available for FromJSONBytes(), FormJSON().
	bodyJSON []byte
	// map field reflect.Value caches
	// fields map[string]reflect.Value
}

/*************************************************************
 * Map data operate
 *************************************************************/

// Type get
func (d *MapData) Type() uint8 {
	return uint8(sourceMap)
}

// Set value by key
func (d *MapData) Set(field string, val interface{}) (interface{}, error) {
	d.Map[field] = val
	return val, nil
}

// Get value by key
func (d *MapData) Get(field string) (interface{}, bool) {
	// if fv, ok := d.fields[field]; ok {
	// 	return fv, true
	// }

	return filter.GetByPath(field, d.Map)
}

// Create a Validation from data
func (d *MapData) Create(err ...error) *Validation {
	return d.Validation(err...)
}

// Validation create from data
func (d *MapData) Validation(err ...error) *Validation {
	if len(err) > 0 {
		return NewValidation(d).WithError(err[0])
	}
	return NewValidation(d)
}

// BindJSON binds v to the JSON data in the request body.
// It calls json.Unmarshal and sets the value of v.
func (d *MapData) BindJSON(ptr interface{}) error {
	if len(d.bodyJSON) == 0 {
		return nil
	}
	return Unmarshal(d.bodyJSON, ptr)
}

/*************************************************************
 * Struct Data
 *************************************************************/

// ConfigValidationFace definition. you can do something on create Validation.
type ConfigValidationFace interface {
	ConfigValidation(v *Validation)
}

// FieldTranslatorFace definition. you can custom field translates.
// Usage:
// 	type User struct {
// 		Name string `json:"name" validate:"required|minLen:5"`
// 	}
//
// 	func (u *User) Translates() map[string]string {
// 		return MS{
// 			"Name": "User name",
// 		}
// 	}
type FieldTranslatorFace interface {
	Translates() map[string]string
}

// CustomMessagesFace definition. you can custom validator error messages.
// Usage:
// 	type User struct {
// 		Name string `json:"name" validate:"required|minLen:5"`
// 	}
//
// 	func (u *User) Messages() map[string]string {
// 		return MS{
// 			"Name.required": "oh! User name is required",
// 		}
// 	}
type CustomMessagesFace interface {
	Messages() map[string]string
}

// StructData definition
type StructData struct {
	// source struct data, from user setting
	src interface{}
	// max depth for parse sub-struct. TODO WIP ...
	depth int
	// from reflect source Struct
	value reflect.Value
	// source struct reflect.Type
	valueTpy reflect.Type
	// field names in the src struct
	// 0:common field 1:anonymous field 2:nonAnonymous field
	fieldNames map[string]int8
	// cache field value info
	fieldValues map[string]reflect.Value
	// TODO field reflect values cache
	fieldRftValues map[string]interface{}
	// FieldTag name in the struct tags. for define filed translate
	FieldTag string
	// MessageTag define error message for the field.
	MessageTag string
	// FilterTag name in the struct tags.
	FilterTag string
	// ValidateTag name in the struct tags.
	ValidateTag string
}

// StructOption definition
// type StructOption struct {
// 	// ValidateTag in the struct tags.
// 	ValidateTag string
// 	// MethodName  string
// }

var (
	cmFaceType = reflect.TypeOf(new(CustomMessagesFace)).Elem()
	ftFaceType = reflect.TypeOf(new(FieldTranslatorFace)).Elem()
	cvFaceType = reflect.TypeOf(new(ConfigValidationFace)).Elem()
)

// Type get
func (d *StructData) Type() uint8 {
	return uint8(sourceStruct)
}

// Create a Validation from the StructData
func (d *StructData) Validation(err ...error) *Validation {
	return d.Create(err...)
}

// Validation create from the StructData
func (d *StructData) Create(err ...error) *Validation {
	v := NewValidation(d)
	if len(err) > 0 && err[0] != nil {
		return v.WithError(err[0])
	}

	// collect field filter/validate rules from struct tags
	d.parseRulesFromTag(v)

	// if has custom config func
	if d.valueTpy.Implements(cvFaceType) {
		fv := d.value.MethodByName("ConfigValidation")
		fv.Call([]reflect.Value{reflect.ValueOf(v)})
	}

	// collect custom field translates config
	if d.valueTpy.Implements(ftFaceType) {
		fv := d.value.MethodByName("Translates")
		vs := fv.Call(nil)
		v.WithTranslates(vs[0].Interface().(map[string]string))
	}

	// collect custom error messages config
	// if reflect.PtrTo(d.valueTpy).Implements(cmFaceType) {
	if d.valueTpy.Implements(cmFaceType) {
		fv := d.value.MethodByName("Messages")
		vs := fv.Call(nil)
		v.WithMessages(vs[0].Interface().(map[string]string))
	}

	// for struct, default update source value
	v.UpdateSource = true
	return v
}

// parse and collect rules from struct tags.
func (d *StructData) parseRulesFromTag(v *Validation) {
	var recursiveFunc func(vv reflect.Value, vt reflect.Type, preStrName string, parentIsAnonymous bool)
	if d.ValidateTag == "" {
		d.ValidateTag = gOpt.ValidateTag
	}

	if d.FilterTag == "" {
		d.FilterTag = gOpt.FilterTag
	}

	fMap := make(map[string]string, 0)

	vv := d.value
	vt := d.valueTpy
	recursiveFunc = func(vv reflect.Value, vt reflect.Type, preStrName string, parentIsAnonymous bool) {
		for i := 0; i < vt.NumField(); i++ {
			fValue := removeValuePtr(vv).Field(i)
			fv := vt.Field(i)
			ft := vt.Field(i).Type
			ft = removeTypePtr(ft)

			// skip don't exported field
			name := fv.Name
			if name[0] >= 'a' && name[0] <= 'z' {
				continue
			}

			if preStrName == "" {
				d.fieldNames[name] = fieldAtTopStruct
			} else {
				name = preStrName + "." + name
				if parentIsAnonymous {
					d.fieldNames[name] = fieldAtAnonymous
				} else {
					d.fieldNames[name] = fieldAtSubStruct
				}
			}

			// validate rule
			vRule := fv.Tag.Get(d.ValidateTag)
			if vRule != "" {
				v.StringRule(name, vRule)
			}

			// filter rule
			fRule := fv.Tag.Get(d.FilterTag)
			if fRule != "" {
				v.FilterRule(name, fRule)
			}

			// load field translate name. eg: `json:"user_name"`
			if gOpt.FieldTag != "" {
				fName := fv.Tag.Get(gOpt.FieldTag)
				if fName != "" {
					fMap[name] = fName
				}
			}

			// load custom error messages.
			// eg: `message:"required:name is required|minLen:name min len is %d"`
			if gOpt.MessageTag != "" {
				errMsg := fv.Tag.Get(gOpt.MessageTag)
				if errMsg != "" {
					d.loadMessagesFromTag(v.trans, name, vRule, errMsg)
				}
			}

			// collect rules from sub-struct and from arrays/slices elements
			// TODO should use ft == timeType check time.Time
			if ft != timeType {
				if fValue.Type().Kind() == reflect.Ptr && fValue.IsNil() {
					continue
				}

				switch ft.Kind() {
				case reflect.Struct:
					recursiveFunc(fValue, ft, name, fv.Anonymous)

				case reflect.Array, reflect.Slice:
					fValue = removeValuePtr(fValue)
					for j := 0; j < fValue.Len(); j++ {
						elemValue := removeValuePtr(fValue.Index(j))
						elemType := removeTypePtr(elemValue.Type())

						arrayName := fmt.Sprintf("%s.%d", name, j)
						if elemType.Kind() == reflect.Struct {
							recursiveFunc(elemValue, elemType, arrayName, fv.Anonymous)
						}
					}

				case reflect.Map:
					fValue = removeValuePtr(fValue)
					for _, key := range fValue.MapKeys() {
						key = removeValuePtr(key)
						elemValue := removeValuePtr(fValue.MapIndex(key))
						elemType := removeTypePtr(elemValue.Type())

						kind := key.Kind()
						format := "%s."
						val := key.Interface()
						switch {
						case kind == reflect.String:
							format += "%s"
							val = strings.ReplaceAll(val.(string), "\"", "")
						case kind >= reflect.Int && kind <= reflect.Uint64:
							format += "%d"
						case kind >= reflect.Float32 && kind <= reflect.Complex128:
							format += "%f"
						default:
							format += "%#v"
						}

						arrayName := fmt.Sprintf(format, name, val)
						if elemType.Kind() == reflect.Struct {
							recursiveFunc(elemValue, elemType, arrayName, fv.Anonymous)
						}
					}

				}
			}
		}
	}

	recursiveFunc(removeValuePtr(vv), vt, "", false)

	if len(fMap) > 0 {
		v.trans.AddFieldMap(fMap)
	}
}

// eg: `message:"required:name is required|minLen:name min len is %d"`
func (d *StructData) loadMessagesFromTag(trans *Translator, field, vRule, vMsg string) {
	var msgKey, vName string

	// only one message, use for first validator.
	// eg: `message:"name is required"`
	if !strings.ContainsRune(vMsg, '|') {
		// eg: `message:"required:name is required"`
		if strings.ContainsRune(vMsg, ':') {
			nodes := strings.SplitN(vMsg, ":", 2)
			vName = strings.TrimSpace(nodes[0])
			// first is validator name
			vMsg = strings.TrimSpace(nodes[1])
		}

		if vName == "" {
			// eg `validate:"required|date"`
			vName = vRule
			if strings.ContainsRune(vRule, '|') {
				nodes := strings.SplitN(vRule, "|", 2)
				// use first validator name
				vName = nodes[0]
			}

			// has params for validator: "minLen:5"
			if strings.ContainsRune(vName, ':') {
				nodes := strings.SplitN(vRule, ":", 2)
				// use first validator name
				vName = nodes[0]
			}
		}

		// if rName, has := validatorAliases[validator]; has {
		// 	msgKey = field + "." + rName
		// } else {
		msgKey = field + "." + vName
		// }

		trans.AddMessage(msgKey, vMsg)
		return
	}

	// multi message for validators
	// eg: `message:"required:name is required | minLen:name min len is %d"`
	msgNodes := strings.Split(vMsg, "|")
	for _, validatorWithMsg := range msgNodes {
		// validatorWithMsg eg: "required:name is required"
		nodes := strings.SplitN(validatorWithMsg, ":", 2)

		validator := nodes[0]
		if rName, has := validatorAliases[validator]; has {
			msgKey = field + "." + rName
		} else {
			msgKey = field + "." + validator
		}

		trans.AddMessage(msgKey, strings.TrimSpace(nodes[1]))
	}
}

/*************************************************************
 * Struct data operate
 *************************************************************/

// Get value by field name
func (d *StructData) Get(field string) (interface{}, bool) {
	var fv reflect.Value
	field = strutil.UpperFirst(field)

	// want get sub struct field.
	if strings.ContainsRune(field, '.') {
		fieldNodes := strings.Split(field, ".")

		if len(fieldNodes) < 2 {
			return nil, false
		}

		topLevelField, ok := d.valueTpy.FieldByName(fieldNodes[0])
		if !ok {
			return nil, false
		}

		kind := removeTypePtr(topLevelField.Type).Kind()
		if kind != reflect.Struct && kind != reflect.Array && kind != reflect.Slice && kind != reflect.Map {
			return nil, false
		}

		fv = removeValuePtr(d.value.FieldByName(fieldNodes[0]))
		if !fv.IsValid() {
			return nil, false
		}

		fieldNodes = fieldNodes[1:]
		lastIndex := len(fieldNodes) - 1

		for i, fieldNode := range fieldNodes {
			fieldNode = strings.ReplaceAll(fieldNode, "\"", "") // for strings as keys

			kind := fv.Type().Kind()
			switch kind {
			case reflect.Array, reflect.Slice:
				index, _ := strconv.Atoi(fieldNode)
				fv = fv.Index(index)
			case reflect.Map:
				fv = fv.MapIndex(reflect.ValueOf(fieldNode))
			default:
				fv = fv.FieldByName(fieldNode)
			}

			fv = removeValuePtr(fv)

			if !fv.IsValid() {
				return nil, false
			}

			if IsZero(fv) || (fv.Kind() == reflect.Ptr && fv.IsNil()) {
				return nil, false
			}

			if i < lastIndex && fv.Type().Kind() != reflect.Struct {
				return nil, false
			}
		}

		d.fieldNames[field] = fieldAtSubStruct
	} else {
		// field at top struct
		fv = d.value.FieldByName(field)

		// is it a pointer
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() { // fix: top-field is nil
				return nil, false
			}

			fv = removeValuePtr(fv)
		}

		if !fv.IsValid() { // field not exists
			return nil, false
		}
	}

	// check can interface
	if fv.CanInterface() {
		// up: if is zero value, as not exist.
		if IsZero(fv) {
			return nil, false
		}

		// cache field value info
		d.fieldValues[field] = fv
		return fv.Interface(), true
	}
	return nil, false
}

// Set value by field name.
// Notice: `StructData.src` the incoming struct must be a pointer to set the value
func (d *StructData) Set(field string, val interface{}) (newVal interface{}, err error) {
	field = strutil.UpperFirst(field)
	if !d.HasField(field) { // field not found
		return nil, ErrNoField
	}

	fv, ok := d.fieldValues[field]
	if !ok {
		f := d.fieldNames[field]
		switch f {
		case fieldAtTopStruct:
			fv = d.value.FieldByName(field)
		case fieldAtAnonymous:
		case fieldAtSubStruct:
			fieldNodes := strings.Split(field, ".")
			if len(fieldNodes) < 2 {
				return nil, ErrInvalidData
			}

			fv = d.value.FieldByName(fieldNodes[0])
			fieldNodes = fieldNodes[1:]

			for _, fieldNode := range fieldNodes {
				switch fv.Type().Kind() {
				case reflect.Array, reflect.Slice:
					index, err := strconv.Atoi(fieldNode)
					if err != nil {
						return nil, ErrInvalidData
					}

					fv = fv.Index(index)
				case reflect.Map:
					fv = fv.MapIndex(reflect.ValueOf(fieldNode))
				default:
					fv = removeValuePtr(fv.FieldByName(fieldNode))
				}
			}
		default:
			return nil, ErrNoField
		}
	}

	// check whether the value of v can be changed.
	if !fv.CanSet() {
		return nil, ErrSetValue
	}

	// Notice: need convert value type
	rftVal := reflect.ValueOf(val)

	// check whether can direct convert type
	if rftVal.Type().ConvertibleTo(fv.Type()) {
		fv.Set(rftVal.Convert(fv.Type()))
		return val, nil
	}

	// try manual convert type
	srcKind, err := basicKind(rftVal)
	if err != nil {
		return nil, err
	}

	newVal, err = convertType(val, srcKind, fv.Kind())
	if err != nil {
		return nil, err
	}

	// update field value
	fv.Set(reflect.ValueOf(newVal))
	return
}

// FuncValue get func value in the src struct
func (d *StructData) FuncValue(name string) (reflect.Value, bool) {
	fv := d.value.MethodByName(filter.UpperFirst(name))
	return fv, fv.IsValid()
}

// HasField in the src struct
func (d *StructData) HasField(field string) bool {
	if _, ok := d.fieldNames[field]; ok {
		return true
	}

	// has field, cache it
	if _, ok := d.valueTpy.FieldByName(field); ok {
		d.fieldNames[field] = fieldAtTopStruct
		return true
	}

	return false
}

/*************************************************************
 * Form Data
 *************************************************************/

// FormData obtained from the request body or url query parameters or user custom setting.
type FormData struct {
	// Form holds any basic key-value string data
	// This includes all fields from urlencoded form,
	// and the form fields only (not files) from a multipart form
	Form url.Values
	// Files holds files from a multipart form only.
	// For any other type of request, it will always
	// be empty. Files only supports one file per key,
	// since this is by far the most common use. If you
	// need to have more than one file per key, parse the
	// files manually using r.MultipartForm.File.
	Files map[string]*multipart.FileHeader
	// jsonBodies holds the original body of the request.
	// Only available for json requests.
	jsonBodies []byte
}

func newFormData() *FormData {
	return &FormData{
		Form:  make(map[string][]string),
		Files: make(map[string]*multipart.FileHeader),
	}
}

/*************************************************************
 * Form data operate
 *************************************************************/

// Type get
func (d *FormData) Type() uint8 {
	return uint8(sourceForm)
}

// Create a Validation from data
func (d *FormData) Create(err ...error) *Validation {
	return d.Validation(err...)
}

// Validation create from data
func (d *FormData) Validation(err ...error) *Validation {
	if len(err) > 0 && err[0] != nil {
		return NewValidation(d).WithError(err[0])
	}
	return NewValidation(d)
}

// Add adds the value to key. It appends to any existing values associated with key.
func (d *FormData) Add(key string, value string) {
	d.Form.Add(key, value)
}

// AddValues to Data.Form
func (d *FormData) AddValues(values url.Values) {
	for key, vals := range values {
		for _, val := range vals {
			d.Form.Add(key, val)
		}
	}
}

// AddFiles adds the multipart form files to data
func (d *FormData) AddFiles(filesMap map[string][]*multipart.FileHeader) {
	for key, files := range filesMap {
		if len(files) != 0 {
			d.AddFile(key, files[0])
		}
	}
}

// AddFile adds the multipart form file to data with the given key.
func (d *FormData) AddFile(key string, file *multipart.FileHeader) {
	d.Files[key] = file
}

// Del deletes the values associated with key.
func (d *FormData) Del(key string) {
	d.Form.Del(key)
}

// DelFile deletes the file associated with key (if any).
// If there is no file associated with key, it does nothing.
func (d *FormData) DelFile(key string) {
	delete(d.Files, key)
}

// Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux") sorted by key.
// Any files in d will be ignored because there is no direct way to convert a file to a
// URL encoded value.
func (d *FormData) Encode() string {
	return d.Form.Encode()
}

// Set sets the key to value. It replaces any existing values.
func (d *FormData) Set(field string, val interface{}) (newVal interface{}, err error) {
	newVal = val
	switch val.(type) {
	case string:
		d.Form.Set(field, val.(string))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		newVal = strutil.MustString(val)
		d.Form.Set(field, newVal.(string))
	default:
		err = fmt.Errorf("set value failure for field: %s", field)
	}
	return
}

// Get value by key
func (d FormData) Get(key string) (interface{}, bool) {
	// get form value
	if vs, ok := d.Form[key]; ok && len(vs) > 0 {
		return vs[0], true
	}

	// get uploaded file
	if fh, ok := d.Files[key]; ok {
		return fh, true
	}

	return nil, false
}

// String value get by key
func (d FormData) String(key string) string {
	return d.Form.Get(key)
}

// Strings value get by key
func (d FormData) Strings(key string) []string {
	return d.Form[key]
}

// GetFile returns the multipart form file associated with key, if any, as a *multipart.FileHeader.
// If there is no file associated with key, it returns nil. If you just want the body of the
// file, use GetFileBytes.
func (d FormData) GetFile(key string) *multipart.FileHeader {
	return d.Files[key]
}

// Has key in the Data
func (d FormData) Has(key string) bool {
	if vs, ok := d.Form[key]; ok && len(vs) > 0 {
		return true
	}

	if _, ok := d.Files[key]; ok {
		return true
	}

	return false
}

// HasField returns true iff data.Form[key] exists. When parsing a request body, the key
// is considered to be in existence if it was provided in the request body, even if its value
// is empty.
func (d FormData) HasField(key string) bool {
	_, found := d.Form[key]
	return found
}

// HasFile returns true iff data.Files[key] exists. When parsing a request body, the key
// is considered to be in existence if it was provided in the request body, even if the file
// is empty.
func (d FormData) HasFile(key string) bool {
	_, found := d.Files[key]
	return found
}

// Int returns the first element in data[key] converted to an int.
func (d FormData) Int(key string) int {
	if val := d.String(key); val != "" {
		iVal, _ := strconv.Atoi(val)
		return iVal
	}

	return 0
}

// Int64 returns the first element in data[key] converted to an int64.
func (d FormData) Int64(key string) int64 {
	if val := d.String(key); val != "" {
		i64, _ := strconv.ParseInt(val, 10, 0)
		return i64
	}

	return 0
}

// Float returns the first element in data[key] converted to a float.
func (d FormData) Float(key string) float64 {
	if val := d.String(key); val != "" {
		result, _ := strconv.ParseFloat(val, 0)
		return result
	}

	return 0
}

// Bool returns the first element in data[key] converted to a bool.
func (d FormData) Bool(key string) bool {
	if val := d.String(key); val != "" {
		blVal, _ := filter.Bool(val)
		return blVal
	}

	return false
}

// FileBytes returns the body of the file associated with key. If there is no
// file associated with key, it returns nil (not an error). It may return an error if
// there was a problem reading the file. If you need to know whether or not the file
// exists (i.e. whether it was provided in the request), use the FileExists method.
func (d FormData) FileBytes(field string) ([]byte, error) {
	fh, found := d.Files[field]
	if !found {
		return nil, nil
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

// FileMimeType get File Mime Type name. eg "image/png"
func (d FormData) FileMimeType(field string) (mime string) {
	fh, found := d.Files[field]
	if !found {
		return
	}

	if file, err := fh.Open(); err == nil {
		var buf [sniffLen]byte
		n, _ := io.ReadFull(file, buf[:])
		mime = http.DetectContentType(buf[:n])
	}
	return
}
