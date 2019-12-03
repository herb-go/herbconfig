package toml

import "testing"

var data = `
fieldint=15
`

type testStruct struct {
	FieldInt int
}

func TestToml(t *testing.T) {
	var d = testStruct{}
	err := Unmarshaler([]byte(data), &d)
	if err != nil {
		t.Fatal(err)
	}
	if d.FieldInt != 15 {
		t.Fatal(d)
	}
}
