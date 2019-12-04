package configloader_test

import "testing"
import "github.com/herb-go/herbconfig/configloader"
import "github.com/herb-go/herbconfig/configloader/example"

func TestExample(t *testing.T) {
	c := configloader.NewCommonConfig()
	a := configloader.EmptyAssembler.WithConfig(c).WithPart(configloader.NewMapPart(example.ExampleData))
	v := &example.ExampleStruct{}
	ok, err := a.Assemble(v)
	if ok == false {
		t.Fatal(ok)
	}
	if err != nil {
		t.Fatal(err)
	}
}
