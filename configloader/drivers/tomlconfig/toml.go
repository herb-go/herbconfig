package tomlconfig

import "github.com/BurntSushi/toml"
import "github.com/herb-go/herbconfig/configloader"

const DefaultConfigLoaderName = "toml"

type TOMLParser struct {
}

func (p *TOMLParser) Parse(data []byte) (configloader.Part, error) {
	var m = interface{}(nil)
	err := toml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return configloader.NewMapPart(m), nil
}

var Parser = &TOMLParser{}

func RegisterDefaultLoader() {
	loader := configloader.NewConfigLoader()
	loader.SetAssemblerConfig(configloader.NewCommonConfig())
	loader.SetParser(Parser)
	configloader.RegisterConfigLoader(DefaultConfigLoaderName, loader)
}

var Config = configloader.NewCommonConfig()

var Assembler = configloader.EmptyAssembler.WithConfig(Config)

func init() {
	RegisterDefaultLoader()
}
