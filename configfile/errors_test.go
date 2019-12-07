package configfile

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrors(t *testing.T) {
	if GetErrorType(nil) != "" {
		t.Fatal(GetErrorType(nil))
	}
	err := fmt.Errorf("err:%s", "err")
	if GetErrorType(err) != "" {
		t.Fatal(err)
	}
	err = NewFileObjectSchemeError("testid")
	if GetErrorType(err) != ErrTypeFileObjectSchemeNotavaliable {
		t.Fatal(err)
	}
	errmsg := err.Error()
	if !strings.Contains(errmsg, "testid") {
		t.Fatal(errmsg)
	}
	err = NewFileObjectNotWriteableError("testid")
	if GetErrorType(err) != ErrTypeFileObjectNotWriteable {
		t.Fatal(err)
	}
	errmsg = err.Error()
	if !strings.Contains(errmsg, "testid") {
		t.Fatal(errmsg)
	}

}
