package source

import (
	"fmt"
)

type ErrorType string

const ErrTypeFileObjectSchemeNotavaliable = ErrorType("schemenotavaliable")

type FileObjectError struct {
	Type   ErrorType
	Msg    string
	FileID string
}

func (f *FileObjectError) Error() string {
	return fmt.Sprintf(f.Msg, f.FileID)
}
func NewFileObjectSchemeError(id string) error {
	return &FileObjectError{
		Type:   ErrTypeFileObjectSchemeNotavaliable,
		Msg:    "file scheme of file object \"%s\" is not avaliable",
		FileID: id,
	}
}

func GetErrorType(err error) ErrorType {
	if err == nil {
		return ""
	}
	e, ok := err.(*FileObjectError)
	if ok == false {
		return ""
	}
	return e.Type
}