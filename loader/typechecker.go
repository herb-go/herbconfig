package loader

import "reflect"

type TypeChecker interface {
	CheckType(a *Assembler, rv reflect.Value) (Type, error)
}

type Checker struct {
	Type    Type
	Checker func(a *Assembler, rv reflect.Value) (bool, error)
}

func (c *Checker) CheckType(a *Assembler, rv reflect.Value) (Type, error) {
	ok, err := c.Checker(a, rv)
	if err != nil {
		return TypeUnkonwn, err
	}
	if ok {
		return c.Type, nil
	}
	return TypeUnkonwn, nil
}

//TypeCheckers type checkers list in order type
type TypeCheckers []TypeChecker

//Append append checkers to last of given type checker.
func (c *TypeCheckers) Append(checkers ...TypeChecker) *TypeCheckers {
	*c = append(*c, checkers...)
	return c
}

//AppendWith append with given TypeCheckers
func (c *TypeCheckers) AppendWith(checkers *TypeCheckers) *TypeCheckers {
	return c.Append(*checkers...)
}

//Insert insert checkers to first of given type checker.
func (c *TypeCheckers) Insert(checkers ...TypeChecker) *TypeCheckers {
	*c = TypeCheckers(append(checkers, *c...))
	return c
}

//InsertWith insert with given TypeCheckers
func (c *TypeCheckers) InsertWith(checkers *TypeCheckers) *TypeCheckers {
	return c.Insert(*checkers...)
}

//CheckType check type with given assembler and reflect type.
//Return type and any error if raised.
func (c *TypeCheckers) CheckType(a *Assembler, rv reflect.Value) (Type, error) {
	for _, v := range *c {
		t, err := v.CheckType(a, rv)
		if err != nil {
			return TypeUnkonwn, err
		}
		if t != TypeUnkonwn {
			return t, nil
		}
	}
	return TypeUnkonwn, nil
}

//NewTypeCheckers create new type checkers
func NewTypeCheckers() *TypeCheckers {
	return &TypeCheckers{}
}

//TypeCheckerString type checker for string.
var TypeCheckerString = &Checker{
	Type: TypeString,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.String, nil
	},
}

//TypeCheckerBool type checker for bool.
var TypeCheckerBool = &Checker{
	Type: TypeBool,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Bool, nil
	},
}

//TypeCheckerInt type checker for int.
var TypeCheckerInt = &Checker{
	Type: TypeInt,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Int, nil
	},
}

//TypeCheckerUint type checker for uint.
var TypeCheckerUint = &Checker{
	Type: TypeUint,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Uint, nil
	},
}

//TypeCheckerInt64 type checker for int64
var TypeCheckerInt64 = &Checker{
	Type: TypeInt64,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Int64, nil
	},
}

//TypeCheckerUint64 type checker for uint64
var TypeCheckerUint64 = &Checker{
	Type: TypeUint64,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Uint64, nil
	},
}

//TypeCheckerFloat32 type checker for float32
var TypeCheckerFloat32 = &Checker{
	Type: TypeFloat32,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Float32, nil
	},
}

//TypeCheckerFloat64 type checker for float64
var TypeCheckerFloat64 = &Checker{
	Type: TypeFloat64,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Float64, nil
	},
}

//TypeCheckerStringKeyMap type checker for string key map.
var TypeCheckerStringKeyMap = &Checker{
	Type: TypeMap,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Map && rv.Type().Key().Kind() == reflect.String, nil
	},
}

//TypeCheckerSlice type checker for slice
var TypeCheckerSlice = &Checker{
	Type: TypeSlice,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Slice, nil
	},
}

//TypeCheckerStruct type checker for struct
var TypeCheckerStruct = &Checker{
	Type: TypeStruct,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Struct, nil
	},
}

//TypeCheckerEmptyInterface type checker for empty interface.
var TypeCheckerEmptyInterface = &Checker{
	Type: TypeEmptyInterface,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Interface && rv.Type().NumMethod() == 0, nil
	},
}

//TypeCheckerLazyLoadFunc type checker for lazy load func.
var TypeCheckerLazyLoadFunc = &Checker{
	Type: TypeLazyLoadFunc,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		lt := a.Config().TagLazyLoad
		if lt == "" {
			return false, nil
		}
		step := a.Step()
		if step == nil || step.Type() != TypeStructField {
			return false, nil
		}
		field := step.Interface().(reflect.StructField)
		tag, err := a.Config().GetTag(rv.Type(), field)
		if err != nil {
			return false, err
		}
		ok := rv.Type().Kind() == reflect.Func && tag != nil && tag.Flags[lt] != ""
		//init value for avoid panic when config data omit.
		if ok && rv.IsNil() {
			rv.Set(NopLazyLoadFunc)
		}
		return ok, nil
	},
}

//TypeCheckerLazyLoader type checker for lazy loader.
var TypeCheckerLazyLoader = &Checker{
	Type: TypeLazyLoader,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		lt := a.Config().TagLazyLoad
		if lt == "" {
			return false, nil
		}
		step := a.Step()
		if step == nil || step.Type() != TypeStructField {
			return false, nil
		}
		field := step.Interface().(reflect.StructField)
		tag, err := a.Config().GetTag(rv.Type(), field)
		if err != nil {
			return false, err
		}
		ok := rv.Type().Kind() == reflect.Interface && tag != nil && tag.Flags[lt] != ""
		//init value for avoid panic when config data omit.
		if ok && rv.IsNil() {
			rv.Set(NopLazyLoader)
		}
		return ok, nil
	},
}

//TypeCheckerPtr type checker for pointer
var TypeCheckerPtr = &Checker{
	Type: TypePtr,
	Checker: func(a *Assembler, rv reflect.Value) (bool, error) {
		return rv.Type().Kind() == reflect.Ptr, nil
	},
}

//DefaultCommonTypeCheckers default common type checkers
func DefaultCommonTypeCheckers() *TypeCheckers {
	return NewTypeCheckers().Append(
		TypeCheckerBool,
		TypeCheckerString,
		TypeCheckerInt,
		TypeCheckerUint,
		TypeCheckerInt64,
		TypeCheckerUint64,
		TypeCheckerFloat32,
		TypeCheckerFloat64,
		TypeCheckerStringKeyMap,
		TypeCheckerSlice,
		TypeCheckerStruct,
		TypeCheckerEmptyInterface,
		TypeCheckerPtr,
		TypeCheckerLazyLoadFunc,
		TypeCheckerLazyLoader,
	)
}

//CommonTypeCheckers common type checkers used in NewCommonConfig
var CommonTypeCheckers = NewTypeCheckers().AppendWith(DefaultCommonTypeCheckers())
