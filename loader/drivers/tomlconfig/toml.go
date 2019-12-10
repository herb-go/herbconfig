package tomlconfig

import (
	"github.com/BurntSushi/toml"
	"github.com/herb-go/herbconfig/loader"
)

const DefaultConfigLoaderName = "toml"

type TOMLParser struct {
}

func (p *TOMLParser) Parse(data []byte) (loader.Part, error) {
	var m = interface{}(nil)
	err := toml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return loader.NewMapPart(m), nil
}

var Parser = &TOMLParser{}

func RegisterDefaultLoader() {
	l := loader.NewConfigLoader()
	l.SetAssemblerConfig(loader.NewCommonConfig())
	l.SetParser(Parser)
	loader.RegisterConfigLoader(DefaultConfigLoaderName, l)
}

var Config = loader.NewCommonConfig()

var Assembler = loader.EmptyAssembler.WithConfig(Config)

func init() {
	RegisterDefaultLoader()
}
