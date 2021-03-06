package jsonconfig

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/herb-go/herbconfig/loader"
	"github.com/herb-go/herbconfig/loader/example"
)

func TestJSON(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buffer).Encode(example.ExampleData)
	if err != nil {
		t.Fatal(err)
	}
	var data = &example.ExampleStruct{}
	bytes, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	bytes = []byte("//commtent\n" + string(bytes))
	err = loader.LoadConfig(DefaultConfigLoaderName, bytes, data)
	if err != nil {
		t.Fatal(err)
	}
	if !data.Equal(example.ExampleData) {
		t.Fatal(data)
	}

}
