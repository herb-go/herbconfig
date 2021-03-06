package loader

import (
	"errors"
	"reflect"
	"testing"
)

var testType = Type("loader_test.test")

type testDummyStruct struct {
	Value string
}

var dummyTypeChecker = &Checker{
	Type: testType,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return true, nil
	},
}

var errTypeChecker = &Checker{
	Type: testType,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return false, errors.New("err")
	},
}

func TestCommonTypeCheckers(t *testing.T) {
	defer func() {
		InitCommon()
	}()
	InitCommon()
	c := NewCommonConfig()
	a := EmptyAssembler.WithConfig(c)
	rv := reflect.ValueOf(testDummyStruct{})
	tp, err := c.Checkers.CheckType(a, rv)
	if err != nil {
		t.Fatal(err)
	}
	if tp != TypeStruct {
		t.Fatal(tp)
	}
	CommonTypeCheckers.Insert(errTypeChecker)
	tp, err = c.Checkers.CheckType(a, rv)
	if err == nil {
		t.Fatal(err)
	}
	if tp != TypeUnkonwn {
		t.Fatal(tp)
	}
	CommonTypeCheckers.Insert(dummyTypeChecker)
	tp, err = c.Checkers.CheckType(a, rv)
	if err != nil {
		t.Fatal(err)
	}
	if tp != testType {
		t.Fatal(tp)
	}
}
