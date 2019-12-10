package timedurationtype

import (
	"reflect"
	"time"

	"github.com/herb-go/herbconfig/loader"
)

var TypeTimeDuration = loader.Type("loader.types.timeduration")

var rtypeTimeDuration = reflect.TypeOf(time.Duration(0))

//UnifierTimeDuration unifier for time field
var UnifierTimeDuration = loader.UnifierFunc(func(a *loader.Assembler, rv reflect.Value) (bool, error) {
	str := ""
	strptr := &str
	ok, err := loader.UnifierString.Unify(a, reflect.ValueOf(strptr).Elem())
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
	err = loader.SetValue(rv, reflect.ValueOf(d))
	if err != nil {
		return false, err
	}
	return true, nil
})

//TypeCheckerTimeDuration type checker for int64
var TypeCheckerTimeDuration = &loader.Checker{
	Type: TypeTimeDuration,
	Checker: func(a *loader.Assembler, rt reflect.Type) (bool, error) {
		return rt == rtypeTimeDuration, nil
	},
}

func RegisterType() {
	loader.CommonTypeCheckers.Insert(TypeCheckerTimeDuration)
	loader.CommonUnifiers.Append(TypeTimeDuration, UnifierTimeDuration)

}
func init() {
	RegisterType()
}
