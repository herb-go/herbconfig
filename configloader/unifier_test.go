package configloader

import (
	"errors"
	"reflect"
	"testing"
)

var dummyUnifier = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	return true, nil
})
var errUnifier = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	return false, errors.New("err")
})

func TestCommonUnifier(t *testing.T) {
	defer func() {
		InitCommon()
	}()
	InitCommon()
	CommonUnifiers.Insert(testType, errUnifier)
	c := NewCommonConfig()
	c.Checkers.Insert(dummyTypeChecker)
	a := EmptyAssembler.WithConfig(c)
	v := &testDummyStruct{}
	ok, err := c.Unifiers.Unify(a, reflect.ValueOf(v))
	if err == nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
	CommonUnifiers.Insert(testType, dummyUnifier)
	ok, err = c.Unifiers.Unify(a, reflect.ValueOf(v))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal(ok)
	}
}
