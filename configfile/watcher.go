package configfile

type Watcher func(callback func()) (unwatcher func())

var registeredWatcher = map[string]Watcher{}

func RegisterWatcher(name string, creator func(id string) (ConfigFile, error)) {
	lock.Lock()
	defer lock.Unlock()
	registeredCreators[name] = creator
}
