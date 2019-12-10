package timedurationtype

import (
	"testing"
	"time"

	"github.com/herb-go/herbconfig/loader"
)

func TestTimeDurationType(t *testing.T) {
	defer func() {
		loader.InitCommon()
	}()
	loader.InitCommon()
	RegisterType()
	td := time.Duration(0)
	part := loader.NewMapPart(time.Second.String())
	a := loader.EmptyAssembler.WithConfig(loader.NewCommonConfig()).WithPart(part)
	ok, err := a.Assemble(&td)
	if ok == false || err != nil {
		t.Fatal(ok, err)
	}
	if td != time.Second {
		t.Fatal(td)
	}
}
