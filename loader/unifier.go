package loader

import (
	"reflect"
	"strings"
)

func indirectReflectValue(rv reflect.Value) reflect.Value {
	if rv.IsZero() {
		return rv
	}
	if rv.Type().Kind() == reflect.Ptr {
		return indirectReflectValue(rv.Elem())
	}
	return rv
}

//SetValue set src value to dst.
//return any error if rasied
func SetValue(dst, src reflect.Value) error {
	if !dst.CanSet() {
		return ErrNotSetable
	}
	if !src.Type().AssignableTo(dst.Type()) {
		return ErrNotAssignable
	}
	dst.Set(src)
	return nil
}

//Unifier unifier interface
type Unifier interface {
	//Unify unify value from assembler to reflect value
	//Return whether unify successed or any error rasied
	Unify(a *Assembler, rv reflect.Value) (bool, error)
}
type Unifiers []Unifier

func (u *Unifiers) Append(unifier ...Unifier) *Unifiers {
	*u = append(*u, unifier...)
	return u
}
func (u *Unifiers) Insert(unifier ...Unifier) *Unifiers {
	*u = Unifiers(append(unifier, *u...))
	return u
}
func (u *Unifiers) Unify(a *Assembler, rv reflect.Value) (bool, error) {
	for k := range *u {
		ok, err := (*u)[k].Unify(a, rv)
		if err != nil {
			return false, NewAssemblerError(a, err)
		}
		if ok {
			return ok, nil
		}
	}
	return false, nil
}

func NewUnifiers() *Unifiers {
	return &Unifiers{}
}

//GroupedUnifiers unifier map grouped by type
type GroupedUnifiers map[Type][]Unifier

//Unify unify value from assembler to reflect value
//Return whether unify successed or any error rasied
func (u *GroupedUnifiers) Unify(a *Assembler, rv reflect.Value) (bool, error) {
	tp, err := a.CheckType(rv)
	if err != nil {
		return false, err
	}
	if tp == TypeUnkonwn {
		return false, nil
	}
	unifiers, ok := (*u)[tp]
	if ok == false {
		return false, nil
	}
	for k := range unifiers {
		result, err := unifiers[k].Unify(a, rv)
		if err != nil {
			return false, err
		}
		if result {
			return true, nil
		}
	}
	return false, nil

}

//Append append unifier to last by given type
func (u *GroupedUnifiers) Append(tp Type, unifier Unifier) *GroupedUnifiers {
	m := (*u)
	v := m[tp]
	v = append(v, unifier)
	m[tp] = v
	*u = m
	return u
}

//AppendWith append with given unifiers
func (u *GroupedUnifiers) AppendWith(unifiers *GroupedUnifiers) *GroupedUnifiers {
	for k, v := range *unifiers {
		for i := range v {
			u.Append(k, v[i])
		}
	}
	return u
}

//Insert insert unifier to first by given type
func (u *GroupedUnifiers) Insert(tp Type, unifier Unifier) *GroupedUnifiers {
	m := (*u)
	v := []Unifier{unifier}
	v = append(v, m[tp]...)
	m[tp] = v
	*u = m
	return u
}

//InsertWith insert with given unifiers
func (u *GroupedUnifiers) InsertWith(unifiers *GroupedUnifiers) *GroupedUnifiers {
	for k, v := range *unifiers {
		for i := range v {
			u.Append(k, v[i])
		}
	}
	return u
}

//NewGroupedUnifiers create new unifiers.
func NewGroupedUnifiers() *GroupedUnifiers {
	return &GroupedUnifiers{}
}

//String interface
type String interface {
	String() string
}

//UnifierFunc unifier func type
type UnifierFunc func(a *Assembler, rv reflect.Value) (bool, error)

//Unify unify value from assembler to reflect value
//Return whether unify successed or any error rasied
func (f UnifierFunc) Unify(a *Assembler, rv reflect.Value) (bool, error) {
	return f(a, rv)
}

//UnifierBool unifier for bool field
var UnifierBool = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	v, err := a.Part().Value()
	if err != nil {
		return false, err
	}
	s, ok := v.(bool)
	if ok {
		err = SetValue(rv, reflect.ValueOf(s))
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
})

//UnifierString unifier for string field
var UnifierString = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	v, err := a.Part().Value()
	if err != nil {
		return false, err
	}
	s, ok := v.(string)
	if ok {
		err = SetValue(rv, reflect.ValueOf(s).Convert(rv.Type()))
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if !a.Config().DisableConvertStringInterface {
		i, ok := v.(String)
		if ok {
			err = SetValue(rv, reflect.ValueOf(i))
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
})

//UnifierNumber unifier for number field
var UnifierNumber = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	v, err := a.Part().Value()
	if err != nil {
		return false, err
	}
	av := reflect.ValueOf(v)
	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		if rv.Kind() == av.Kind() {
			err = SetValue(rv, av)
			if err != nil {
				return false, err
			}
			return true, nil
		}
		err = SetValue(rv, reflect.ValueOf(v).Convert(rv.Type()))
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
})

//UnifierSlice unifier for slice field
var UnifierSlice = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	iter, err := a.Part().Iter()
	if err != nil {
		return false, err
	}
	if iter == nil {
		return false, nil
	}
	sv := reflect.MakeSlice(rv.Type(), 0, 0)
	for iter != nil {
		if iter.Step.Type() == TypeArray {
			v := reflect.New(rv.Type().Elem()).Elem()
			_, err = a.Config().Unifiers.Unify(a.WithChild(iter.Part, iter.Step), v)
			if err != nil {
				return false, err
			}
			sv = reflect.Append(sv, v)
		}
		iter, err = iter.Next()
		if err != nil {
			return false, err
		}
	}
	err = SetValue(rv, sv)
	if err != nil {
		return false, err
	}
	return true, nil
})

//UnifierMap unifier for map field
var UnifierMap = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	iter, err := a.Part().Iter()
	if err != nil || iter == nil {
		return false, err
	}

	mv := reflect.MakeMap(rv.Type())
	for iter != nil {
		miv := reflect.New(rv.Type().Elem()).Elem()
		_, err = a.Config().Unifiers.Unify(a.WithChild(iter.Part, iter.Step), miv)
		if err != nil {
			return false, err
		}
		mv.SetMapIndex(reflect.ValueOf(iter.Step.Interface()), miv)
		iter, err = iter.Next()
		if err != nil {
			return false, err
		}
	}
	err = SetValue(rv, mv)
	if err != nil {
		return false, err
	}
	return true, nil
})

func convertIterToArray(iter *PartIter) ([]interface{}, error) {
	a := []interface{}{}
	for iter != nil {
		pv, err := iter.Part.Value()
		if err != nil {
			return nil, err
		}
		a = append(a, pv)
		iter, err = iter.Next()
		if err != nil {
			return nil, err
		}
	}
	return a, nil
}

func convertIterToStringMap(iter *PartIter) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	for iter != nil {
		pv, err := iter.Part.Value()
		if err != nil {
			return nil, err
		}
		m[iter.Step.String()] = pv
		iter, err = iter.Next()
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
func convertIterToInterfaceMap(iter *PartIter) (map[interface{}]interface{}, error) {
	m := map[interface{}]interface{}{}
	for iter != nil {
		pv, err := iter.Part.Value()
		if err != nil {
			return nil, err
		}
		m[iter.Step.Interface()] = pv
		iter, err = iter.Next()
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
func convertIter(i *PartIter) (interface{}, error) {
	switch i.Step.Type() {
	case TypeArray:
		return convertIterToArray(i)
	case TypeString:
		return convertIterToStringMap(i)
	case TypeEmptyInterface:
		return convertIterToInterfaceMap(i)
	}
	return nil, nil
}

//UnifierEmptyInterface unifier for empty interface field
var UnifierEmptyInterface = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	iter, err := a.Part().Iter()
	if err != nil {
		return false, err
	}
	if iter == nil {
		v, err := a.Part().Value()
		if err != nil {
			return false, err
		}
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8,
			reflect.String, reflect.Bool,
			reflect.Float32, reflect.Float64,
			reflect.Map, reflect.Slice:
			err = SetValue(rv, reflect.ValueOf(v))
			if err != nil {
				return false, err
			}
			return true, nil
		}
	} else {
		val, err := convertIter(iter)
		if err != nil {
			return false, err
		}
		if val == nil {
			return false, nil
		}
		err = SetValue(rv, reflect.ValueOf(val))
		if err != nil {
			return false, err
		}
	}
	return false, nil
})

type structData struct {
	assembler  *Assembler
	valuemap   map[string]Part
	civaluemap map[string]Part
}

//LoadValues load values form assembler to struct data
//Return whether load successed and any error if raised
//Load will fail if iter is nil
func (d *structData) LoadValues() (bool, error) {
	a := d.assembler

	iter, err := a.Part().Iter()
	if err != nil {
		return false, err
	}
	if iter == nil {
		return false, nil
	}

	d.valuemap = map[string]Part{}
	d.civaluemap = map[string]Part{}
	ci := !a.Config().CaseSensitive
	for iter != nil {

		d.valuemap[iter.Step.String()] = iter.Part
		if ci {
			d.civaluemap[strings.ToLower(iter.Step.String())] = iter.Part
		}
		iter, err = iter.Next()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func indirectReflectType(rt reflect.Type) reflect.Type {
	if rt.Kind() == reflect.Ptr {
		return indirectReflectType(rt.Elem())
	}
	return rt
}

//IsAnonymous return if given field with tag is anonymous
func (d *structData) IsAnonymous(field reflect.StructField, tag *Tag) bool {
	if tag.Ignored || tag.Name != "" {
		return false
	}
	c := d.assembler.Config()
	if c.TagAnonymous != "" && tag.Flags[c.TagAnonymous] != "" {
		return true
	}
	if d.valuemap[field.Name] != nil {
		return false
	}
	ci := !c.CaseSensitive
	if ci && d.civaluemap[strings.ToLower(field.Name)] != nil {
		return false
	}
	return true
}

//WalkStruct walk struct fields of given reflect value and set field values
func (d *structData) WalkStruct(rv reflect.Value) (bool, error) {
	if rv.Type().Kind() == reflect.Ptr {
		elemv := reflect.New(rv.Type().Elem())
		ok, err := d.WalkStruct(elemv.Elem())
		if ok == false || err != nil {
			return ok, err
		}
		err = SetValue(rv, elemv)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if rv.Type().Kind() != reflect.Struct {
		return false, nil
	}
	a := d.assembler
	rt := rv.Type()
	fl := rt.NumField()
	ci := !a.Config().CaseSensitive
	value := reflect.Indirect(rv)
	for i := 0; i < fl; i++ {
		var part Part
		var ok bool
		field := rt.Field(i)
		if field.PkgPath != "" {
			continue
		}
		step := NewFieldStep(&field)
		fv := value.Field(i)
		//Call checktype to verify type and init value if necessary(for example,lazyloader and lazyloadfunc)
		tp, err := a.WithChild(nil, step).CheckType(fv)
		if err != nil {
			return false, err
		}
		if tp == TypeUnkonwn {
			continue
		}
		tag, err := a.Config().GetTag(rt, field)
		if err != nil {
			return false, err
		}
		if tag.Ignored {
			continue
		}
		if d.IsAnonymous(field, tag) {
			_, err := d.WalkStruct(indirectReflectValue(fv))
			if err != nil {
				return false, err
			}
			continue
		}
		if err != nil {
			return false, err
		}
		if tag.Name != "" {
			part, ok = d.valuemap[tag.Name]
		}
		if !ok {
			part, ok = d.valuemap[field.Name]
		}
		if !ok && ci {
			part, ok = d.civaluemap[strings.ToLower(field.Name)]
		}
		if !ok {
			continue
		}
		_, err = a.Config().Unifiers.Unify(a.WithChild(part, step), fv)
		if err != nil {
			return false, err
		}

	}
	err := SetValue(rv, value)
	if err != nil {
		return false, err
	}
	return true, nil

}

//newStructData create new struct data
func newStructData() *structData {
	return &structData{
		valuemap:   map[string]Part{},
		civaluemap: map[string]Part{},
	}
}

//UnifierStruct unifier for struct
var UnifierStruct = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	v, err := a.Part().Value()
	if err != nil {
		return false, err
	}
	if v == nil {
		return true, nil
	}

	sd := newStructData()
	sd.assembler = a
	ok, err := sd.LoadValues()
	if err != nil {
		return false, err
	}
	if ok == false {
		prv := reflect.Indirect(reflect.ValueOf(v))
		if prv.Kind() != reflect.Struct {
			return false, nil
		}
		err = SetValue(rv, prv)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return sd.WalkStruct(rv)
})

//UnifierLazyLoadFunc unifier for lazyload func
var UnifierLazyLoadFunc = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	l := NewLazyLoader()
	l.Assembler = a
	err := SetValue(rv, reflect.ValueOf(l.LazyLoadConfig))
	if err != nil {
		return false, err
	}
	return true, nil
})

//UnifierLazyLoader unifier for lazy loader
var UnifierLazyLoader = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	l := NewLazyLoader()
	l.Assembler = a
	err := SetValue(rv, reflect.ValueOf(l))
	if err != nil {
		return false, err
	}
	return true, nil
})

//UnifierPtr unifier for pointer
var UnifierPtr = UnifierFunc(func(a *Assembler, rv reflect.Value) (bool, error) {
	av, err := a.Part().Value()
	if err != nil {
		return false, err
	}
	if av == nil {
		return true, nil
	}
	v := reflect.New(rv.Type().Elem())
	err = SetValue(rv, v)
	if err != nil {
		return false, err
	}
	return a.Config().Unifiers.Unify(a, rv.Elem())
})

//DefaultCommonUnifiers return default common unifiers
func DefaultCommonUnifiers() *GroupedUnifiers {
	var u = NewGroupedUnifiers()
	u.Append(TypeBool, UnifierBool)
	u.Append(TypeString, UnifierString)
	u.Append(TypeInt, UnifierNumber)
	u.Append(TypeUint, UnifierNumber)
	u.Append(TypeInt64, UnifierNumber)
	u.Append(TypeUint64, UnifierNumber)
	u.Append(TypeFloat32, UnifierNumber)
	u.Append(TypeFloat64, UnifierNumber)
	u.Append(TypeSlice, UnifierSlice)
	u.Append(TypeMap, UnifierMap)
	u.Append(TypeStruct, UnifierStruct)
	u.Append(TypePtr, UnifierPtr)
	u.Append(TypeEmptyInterface, UnifierEmptyInterface)
	u.Append(TypeLazyLoadFunc, UnifierLazyLoadFunc)
	u.Append(TypeLazyLoader, UnifierLazyLoader)
	return u
}

//CommonUnifiers common unifiers user in NewCommonConfig
var CommonUnifiers = NewGroupedUnifiers().AppendWith(DefaultCommonUnifiers())
