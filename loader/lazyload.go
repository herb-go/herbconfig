package loader

import (
	"reflect"
)

//AssemblerLazyLoader assembler lazy loader struct
type AssemblerLazyLoader struct {
	Assembler *Assembler
}

//LazyLoadConfig lazeload data into interface.
//Return any error if raised
func (l *AssemblerLazyLoader) LazyLoadConfig(v interface{}) error {
	if l.Assembler == nil {
		return nil
	}
	_, err := l.Assembler.Assemble(v)
	return err
}

//NewLazyLoader create new assembler lazy loader
func NewLazyLoader() *AssemblerLazyLoader {
	return &AssemblerLazyLoader{}
}

var nopAssemblerLazyLoader = NewLazyLoader()

//NopLazyLoader no op assembler lazy loader
var NopLazyLoader = reflect.ValueOf(nopAssemblerLazyLoader)

//NopLazyLoadFunc no op assembler lazy load func
var NopLazyLoadFunc = reflect.ValueOf(nopAssemblerLazyLoader.LazyLoadConfig)

//LazyLoaderFunc lazy loader func interface
type LazyLoaderFunc func(v interface{}) error

//LazyLoader lazy loader interface
type LazyLoader interface {
	//LazyLoad lazeload data into interface.
	//Return any error if raised
	LazyLoadConfig(v interface{}) error
}
