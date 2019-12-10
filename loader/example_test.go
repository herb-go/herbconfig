package loader_test

import (
	"testing"

	"github.com/herb-go/herbconfig/loader"
	"github.com/herb-go/herbconfig/loader/example"
)

func TestExample(t *testing.T) {
	c := loader.NewCommonConfig()
	a := loader.EmptyAssembler.WithConfig(c).WithPart(loader.NewMapPart(example.ExampleData))
	v := &example.ExampleStruct{}
	ok, err := a.Assemble(v)
	if ok == false {
		t.Fatal(ok)
	}
	if err != nil {
		t.Fatal(err)
	}
}
