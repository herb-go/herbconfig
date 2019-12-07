package configfile

import (
	"strings"
	"sync"
)

var registeredCreators = map[string]func(id string) (ConfigFile, error){}

var lock = sync.Mutex{}

func getFileTypeFromID(id string) string {
	index := strings.Index(id, "://")
	if index < 0 {
		return ""
	}
	return id[0:index]
}
func RegisterCreator(name string, creator func(id string) (ConfigFile, error)) {
	lock.Lock()
	defer lock.Unlock()
	registeredCreators[name] = creator
}
