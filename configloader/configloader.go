package configloader

import (
	"fmt"
)

//ConfigLoader config lodaer interface
type ConfigLoader interface {
	//LoadConfig load data to giver interface.
	//Return any error if raised.
	LoadConfig(data []byte, v interface{}) error
}

//configloaders all registered config loaders
var configloaders = map[string]ConfigLoader{}

//RegisterConfigLoader register config loader with given name.
func RegisterConfigLoader(name string, u ConfigLoader) {
	configloaders[name] = u
}

//UnregisterAllConfigLoader unreister all config loaders.
func UnregisterAllConfigLoader() {
	configloaders = map[string]ConfigLoader{}
}

//LoadConfig load byte slice to data by given munarshaler.
//Return any error if raised
func LoadConfig(name string, data []byte, v interface{}) error {
	u := configloaders[name]
	if u == nil {
		return fmt.Errorf("configloader : %w (%s)", ErrConfigLoaderNotRegistered, name)
	}
	return u.LoadConfig(data, v)
}
