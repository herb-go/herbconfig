package timedurationtype

import (
	"reflect"
	"time"

	"github.com/herb-go/herbconfig/configloader"
)

var TypeTimeDuration = configloader.Type("configloader.types.timeduration")

var rtypeTimeDuration = reflect.TypeOf(time.Duration(0))

//UnifierTimeDuration unifier for time field
var UnifierTimeDuration = configloader.UnifierFunc(func(a *configloader.Assembler, rv reflect.Value) (bool, error) {
	str := ""
	strptr := &str
	ok, err := configloader.UnifierString.Unify(a, reflect.ValueOf(strptr).Elem())
	if ok == false || err != nil {
		return false, nil

	}
	if *strptr == "" {
		return false, nil
	}
	d, err := time.ParseDuration(*strptr)
	if err != nil {
		return false, nil
	}
	err = configloader.SetValue(rv, reflect.ValueOf(d))
	if err != nil {
		return false, err
	}
	return true, nil
})

//TypeCheckerTimeDuration type checker for int64
var TypeCheckerTimeDuration = &configloader.Checker{
	Type: TypeTimeDuration,
	Checker: func(a *configloader.Assembler, rt reflect.Type) (bool, error) {
		return rt == rtypeTimeDuration, nil
	},
}

func RegisterType() {
	configloader.CommonTypeCheckers.Insert(TypeCheckerTimeDuration)
	configloader.CommonUnifiers.Append(TypeTimeDuration, UnifierTimeDuration)

}
func init() {
	RegisterType()
}
