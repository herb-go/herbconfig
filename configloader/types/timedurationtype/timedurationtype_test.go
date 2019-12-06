package timedurationtype

import (
	"testing"
	"time"

	"github.com/herb-go/herbconfig/configloader"
)

func TestTimeDurationType(t *testing.T) {
	defer func() {
		configloader.InitCommon()
	}()
	configloader.InitCommon()
	RegisterType()
	td := time.Duration(0)
	part := configloader.NewMapPart(time.Second.String())
	a := configloader.EmptyAssembler.WithConfig(configloader.NewCommonConfig()).WithPart(part)
	ok, err := a.Assemble(&td)
	if ok == false || err != nil {
		t.Fatal(ok, err)
	}
	if td != time.Second {
		t.Fatal(td)
	}
}
