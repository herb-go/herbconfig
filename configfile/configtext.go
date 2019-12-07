package configfile

import (
	"html"
	"os"
)

type ConfigText string

func (f ConfigText) ReadRaw() ([]byte, error) {
	return []byte(string(f)), nil
}

func (f ConfigText) WriteRaw(data []byte, perm os.FileMode) error {
	return NewFileObjectNotWriteableError(f.ID())
}

func (f ConfigText) AbsolutePath() string {
	return ""
}
func (f ConfigText) ID() string {
	return "text://" + html.EscapeString(string(f))
}
func (f ConfigText) Watcher() FileWatcher {
	return nil
}

func registerFileObejctTextCreator() {
	RegisterCreator("text", func(id string) (ConfigFile, error) {
		return ConfigText(id[7:]), nil
	})
}

func init() {
	registerFileObejctTextCreator()
}
