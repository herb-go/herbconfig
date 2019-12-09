package configuration

import (
	"html"
)

type Text string

func (f Text) ReadRaw() ([]byte, error) {
	return []byte(string(f)), nil
}

func (f Text) AbsolutePath() string {
	return ""
}
func (f Text) ID() string {
	return "text://" + html.EscapeString(string(f))
}

func registerFileObejctTextCreator() {
	RegisterCreator("text", func(id string) (Configuration, error) {
		return Text(id[7:]), nil
	})
}

func init() {
	registerFileObejctTextCreator()
}
