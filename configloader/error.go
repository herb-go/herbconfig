package configloader

import (
	"errors"
	"fmt"
)

//ErrNotSetable err raised when value can not set
var ErrNotSetable = errors.New("value cannot set")

//ErrNotAssignable err raseid when given value not assignable
var ErrNotAssignable = errors.New("value is not assignable")

//AssemblerError assembler error with assembler info
type AssemblerError struct {
	a   *Assembler
	err error
}

//Unwrap unwrap error
func (e *AssemblerError) Unwrap() error {
	return e.err
}

//NewAssemblerError create new assemble error
func NewAssemblerError(a *Assembler, err error) error {
	if err == nil {
		return nil
	}
	return &AssemblerError{
		a:   a,
		err: err,
	}
}

//Error show error with assembler info
func (e *AssemblerError) Error() string {
	return fmt.Sprintf("configloader: error: %s (%s)", e.err.Error(), ConvertPathToString(e.a.Path()))
}

//ErrConfigLoaderNotRegistered error raised when config loader not registered.
var ErrConfigLoaderNotRegistered = errors.New("configloader not registered")