package configloader

import "fmt"

type Parser interface {
	Parse(data []byte) (Part, error)
}

type ConfigLoader struct {
	Assembler *Assembler
	parser    Parser
}

func (l *ConfigLoader) SetParser(p Parser) *ConfigLoader {
	l.parser = p
	return l
}
func (l *ConfigLoader) SetAssemblerConfig(c *Config) *ConfigLoader {
	l.Assembler = l.Assembler.WithConfig(c)
	return l
}
func (l *ConfigLoader) Load(data []byte, v interface{}) error {
	part, err := l.parser.Parse(data)
	if err != nil {
		return nil
	}
	_, err = l.Assembler.WithPart(part).Assemble(v)
	return err
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{
		Assembler: EmptyAssembler,
	}
}

var regsiteredConfigLoaders = map[string]*ConfigLoader{}

func RegisterConfigLoader(name string, c *ConfigLoader) {
	regsiteredConfigLoaders[name] = c
}

//LoadConfig load byte slice to data by given config loader.
//Return any error if raised
func LoadConfig(name string, data []byte, v interface{}) error {
	c := regsiteredConfigLoaders[name]
	if c == nil {
		return fmt.Errorf("configloader : %w (%s)", ErrConfigLoaderNotRegistered, name)
	}
	return c.Load(data, v)
}

// var parserConfigs = map[string]*Config{}

// //RegisterParser register config loader with given name.
// func RegisterParser(name string, p Parser) {
// 	registeredParsers[name] = p
// }

// func ConfigParser(name string, c *Config) {
// 	parserConfigs[name] = c
// }

// //UnregisterAllParsers unreister all config loaders.
// func UnregisterAllParsers() {
// 	registeredParsers = map[string]Parser{}
// }

// //LoadConfig load byte slice to data by given munarshaler.
// //Return any error if raised
// func LoadConfig(name string, data []byte, v interface{}) error {
// 	p := registeredParsers[name]
// 	if p == nil {
// 		return fmt.Errorf("configloader : %w (%s)", ErrParserNotRegistered, name)
// 	}
// 	c := parserConfigs[name]
// 	// if c==nil{
// 	// 	c=
// 	// }
// 	part, err := p.Parse(data)
// 	if err != nil {
// 		return err
// 	}
// 	return
// 	// return u.LoadConfig(data, v)
// }
