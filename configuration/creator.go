package configuration

import (
	"strings"
	"sync"
)

var registeredCreators = map[string]func(id string) (Configuration, error){}

var locker = sync.Mutex{}

func getFileTypeFromID(id string) string {
	index := strings.Index(id, "://")
	if index < 0 {
		return ""
	}
	return id[0:index]
}
func RegisterCreator(name string, creator func(id string) (Configuration, error)) {
	locker.Lock()
	defer locker.Unlock()
	registeredCreators[name] = creator
}

func New(id string) (Configuration, error) {
	locker.Lock()
	defer locker.Unlock()
	tp := getFileTypeFromID(id)
	creator := registeredCreators[tp]
	if tp == "" || creator == nil {
		return nil, NewFileObjectSchemeError(id)
	}
	return creator(id)
}
