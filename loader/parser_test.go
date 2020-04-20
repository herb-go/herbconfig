package loader_test

import (
	"encoding/json"
	"testing"

	"github.com/herb-go/herbconfig/loader"

	_ "github.com/herb-go/herbconfig/loader/drivers/jsonconfig"
)

type testStruct struct {
	Value string
}

func TestLoader(t *testing.T) {
	data := testStruct{
		Value: "test",
	}
	bs, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	loader := loader.NewLoader("json", bs)
	v := &testStruct{}
	err = loader(v)
	if err != nil {
		t.Fatal(err)
	}
	if v.Value != "test" {
		t.Fatal(v)
	}
}
