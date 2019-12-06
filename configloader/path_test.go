package configloader

import "testing"

func TestPath(t *testing.T) {
	is := NewInterfaceStep("interfacestep")
	if is.String() != "interfacestep" {
		t.Fatal(is)
	}
}
