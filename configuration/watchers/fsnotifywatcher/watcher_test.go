package fsnotidywatcher

import (
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/herb-go/herbconfig/configuration"
)

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

func TestWatcher(t *testing.T) {
	wm := configuration.NewWatchManager()
	wm.RegisterWatcher(configuration.NewSchemeWatcher("file", NewWatcher()))
	err := wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer wm.Stop()
}
