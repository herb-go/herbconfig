package jsonconfig

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/herb-go/herbconfig/loader"
)

const DefaultConfigLoaderName = "json"

type JSONParser struct {
}

func (p *JSONParser) Parse(data []byte) (loader.Part, error) {
	var err error
	r := bytes.NewBuffer(data)
	var bytes = []byte{}
	var line string
	for err != io.EOF {
		line, err = r.ReadString(10)
		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if len(line) > 2 && line[0:2] == "//" {
			continue
		}
		bytes = append(bytes, []byte(line)...)
	}
	var m = interface{}(nil)
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return loader.NewMapPart(m), nil
}

var Parser = &JSONParser{}

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
