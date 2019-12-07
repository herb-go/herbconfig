package configfile

import (
	"html"
	"io/ioutil"
	"net/url"
	"os"
)

type File string

func (f File) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.AbsolutePath())
}

func (f File) WriteRaw(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.AbsolutePath(), data, perm)
}

func (f File) AbsolutePath() string {
	return string(f)
}
func (f File) ID() string {
	u := url.URL{
		Scheme: "file",
		Host:   "local",
		Path:   html.EscapeString(string(f)),
	}
	return u.String()
}
func (f File) Watcher() Watcher {
	return nil
}

func registerLocalFileCreator() {
	RegisterCreator("file", func(id string) (ConfigFile, error) {
		u, err := url.Parse(id)
		if err != nil {
			return nil, err
		}
		return File(u.Path), nil
	})
}

func init() {
	registerLocalFileCreator()
}
