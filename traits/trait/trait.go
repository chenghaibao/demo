package trait

import (
	"reflect"
)

var traitsPackage = reflect.TypeOf(Stringer{}).PkgPath()

func Init(obj interface{}) {
	if o, ok := obj.(stringer); ok {
		o.setStringer(obj)
	}
}