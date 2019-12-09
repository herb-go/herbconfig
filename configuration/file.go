package configuration

import (
	"html"
	"io/ioutil"
	"net/url"
)

type File string

func (f File) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.AbsolutePath())
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

func registerLocalFileCreator() {
	RegisterCreator("file", func(id string) (Configuration, error) {
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
