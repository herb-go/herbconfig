package json

import "testing"

var data = `
{
	"fieldint":15
}
`

type testStruct struct {
	FieldInt int
}

func TestJSON(t *testing.T) {
	var d = testStruct{}
	err := Unmarshaler([]byte(data), &d)
	if err != nil {
		t.Fatal(err)
	}
	if d.FieldInt != 15 {
		t.Fatal(d)
	}
}
