package json

import "encoding/json"
import "github.com/herb-go/unmarshaler"

var Unmarshaler = func(data []byte, v interface{}) error {
	var m = interface{}(nil)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	a := Assembler.WithPart(unmarshaler.NewMapPart(m))
	_, err = a.Assemble(v)
	return err
}

var Config = unmarshaler.NewCommonConfig()

var Assembler = unmarshaler.EmptyAssembler.WithConfig(Config)
