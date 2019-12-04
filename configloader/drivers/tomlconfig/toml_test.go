package tomlconfig

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/BurntSushi/toml"

	"github.com/herb-go/herbconfig/configloader/example"
)

func TestToml(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})
	err := toml.NewEncoder(buffer).Encode(example.ExampleData)
	if err != nil {
		t.Fatal(err)
	}
	var data = &example.ExampleStruct{}
	bytes, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Fatal(err)
	}
	err = Unmarshaler(bytes, data)
	if err != nil {
		t.Fatal(err)
	}
	if !data.Equal(example.ExampleData) {
		t.Fatal(data)
	}
}
