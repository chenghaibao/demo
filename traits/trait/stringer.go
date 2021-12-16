package trait

import (
	"fmt"
	"reflect"
	"strings"
)

type stringer interface {
	fmt.Stringer
	setStringer(i interface{})
}

type Stringer struct {
	self interface{}
}

func (str *Stringer) setStringer(i interface{}) {
	str.self = i
}

// String returns a string representation of an embedding struct.
func (str *Stringer) String() string {
	v := reflect.ValueOf(str.self).Elem()
	t := reflect.TypeOf(str.self).Elem()

	return recursiveToString(v, t, "")
}


func recursiveToString(v reflect.Value, t reflect.Type, ptr string) string {
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			// filter out the traits package internals
			if v.Field(i).Type().PkgPath() == traitsPackage {
				values[i] = ""
			} else {
				values[i] = recursiveToString(v.Field(i), v.Field(i).Type(), "")
			}

		case reflect.Ptr:
			if v.Field(i).Elem().Kind() == reflect.Struct {
				// filter out the traits package internals
				if v.Field(i).Type().PkgPath() == traitsPackage {
					values[i] = ""
				} else {
					values[i] = recursiveToString(v.Field(i).Elem(), v.Field(i).Elem().Type(), "&")
				}
			} else {
				values[i] = fmt.Sprintf("%+v", v.Field(i).Elem())
			}

		case reflect.Interface:
			values[i] = fmt.Sprintf("i[%+v]", v.Field(i))

		case reflect.Chan:
			values[i] = fmt.Sprintf("chan[%+v]", v.Field(i))

		default:
			values[i] = fmt.Sprintf("%+v", v.Field(i))
		}
	}

	values = DeleteEmpty(values)
	return ptr + t.Name() + "(" + strings.Join(values, ",") + ")"
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}