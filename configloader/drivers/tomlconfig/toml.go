package tomlconfig

import "github.com/BurntSushi/toml"
import "github.com/herb-go/herbconfig/configloader"

var Unmarshaler = func(data []byte, v interface{}) error {
	var m = interface{}(nil)
	err := toml.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	a := Assembler.WithPart(configloader.NewMapPart(m))
	_, err = a.Assemble(v)
	return err
}

var Config = configloader.NewCommonConfig()

var Assembler = configloader.EmptyAssembler.WithConfig(Config)
