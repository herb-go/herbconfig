package jsonconfig

import (
	"encoding/json"

	"github.com/herb-go/herbconfig/configloader"
)

var Unmarshaler = func(data []byte, v interface{}) error {
	var m = interface{}(nil)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	a := Assembler.WithPart(configloader.NewMapPart(m))
	_, err = a.Assemble(v)
	return err
}

var Config = configloader.NewCommonConfig()

var Assembler = configloader.EmptyAssembler.WithConfig(Config)

type Loader struct {
	config    *configloader.Config
	assembler *configloader.Assembler
}

func (l *Loader) WithConfig(c *configloader.Config) *Loader {
	l.config = c
	l.assembler = configloader.EmptyAssembler.WithConfig(c)
	return l
}

func NewLoader() *Loader {
	return &Loader{
		assembler: configloader.EmptyAssembler,
	}
}

func NewCommonLoader() *Loader {
	return NewLoader().WithConfig(configloader.NewCommonConfig())
}
