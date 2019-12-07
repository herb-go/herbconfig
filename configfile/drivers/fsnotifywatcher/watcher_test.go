package configfile

import "testing"
import "github.com/fsnotify/fsnotify"

func TestEvent(t *testing.T) {
	e := Event{
		e: &fsnotify.Event{
			Name: "testpath",
			Op:   fsnotify.Create,
		},
	}
	if e.Path() != "testpath" {
		t.Fatal(e)
	}
	if !e.IsCreate() {
		t.Fatal(e)
	}
	if e.IsChmod() {
		t.Fatal(e)
	}
	if e.IsReName() {
		t.Fatal(e)
	}
	if e.IsRemove() {
		t.Fatal(e)
	}
	if e.IsWrite() {
		t.Fatal(e)
	}
}
