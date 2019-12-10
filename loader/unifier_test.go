package loader

import (
	"errors"
	"reflect"
	"testing"
)

var dummyUnifier = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	return true, nil
})
var nilUnifier = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	rv.SetString("test")
	return true, nil
})
var errUnifier = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	return false, errors.New("err")
})

func TestNilUnifier(t *testing.T) {
	defer func() {
		InitCommon()
	}()
	InitCommon()
	CommonUnifiers.Insert(testType, errUnifier)
	CommonUnifiers.Insert(testType, nilUnifier)
	c := NewCommonConfig()
	c.Checkers.Insert(dummyTypeChecker)
	c.Checkers.Insert(TypeCheckerPtr)
	a := EmptyAssembler.WithConfig(c)
	var v string
	ok, err := a.Assemble(&v)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal(ok)
	}
	if v != "test" {
		t.Fatal(v)
	}
}
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

func TestEmptyMap(t *testing.T) {
	defer func() {
		InitCommon()
	}()
	InitCommon()
	var m map[string]interface{}
	var i interface{} = m
	c := NewCommonConfig()
	a := EmptyAssembler.WithConfig(c).WithPart(NewMapPart(i))
	v := &testDummyStruct{}
	ok, err := a.Assemble(v)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
}
