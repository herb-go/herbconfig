package jsonconfig

import (
	"encoding/json"

	"github.com/herb-go/herbconfig/configloader"
)

const DefaultConfigLoaderName = "json"

type JSONParser struct {
}

func (p *JSONParser) Parse(data []byte) (configloader.Part, error) {
	var m = interface{}(nil)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return configloader.NewMapPart(m), nil
}

var Parser = &JSONParser{}

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
