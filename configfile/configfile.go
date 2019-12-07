package configfile

import "os"

func IsSame(src ConfigFile, dst ConfigFile) bool {
	return src.ID() == dst.ID()
}

type ConfigFile interface {
	ReadRaw() ([]byte, error)
	WriteRaw([]byte, os.FileMode) error
	AbsolutePath() string
	ID() string
	Watcher() FileWatcher
}

func ReadFile(file ConfigFile) ([]byte, error) {
	return file.ReadRaw()
}

func WriteFile(file ConfigFile, data []byte, mode os.FileMode) error {
	return file.WriteRaw(data, mode)
}

func New(id string) (ConfigFile, error) {
	tp := getFileTypeFromID(id)
	creator := registeredCreators[tp]
	if tp == "" || creator == nil {
		return nil, NewFileObjectSchemeError(id)
	}
	return creator(id)
}
