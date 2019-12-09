package fsnotifywatcher

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

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
	result := []interface{}{}
	wm := configuration.NewWatchManager()
	m := NewWatcher()
	wm.RegisterWatcher(configuration.NewSchemeWatcher("file", m))
	err := wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer wm.Stop()
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	fn := tmpfile.Name()
	if fn == "" {
		t.Fatal(fn)
	}
	defer os.Remove(fn)
	cf, err := configuration.New("file://" + fn)
	if err != nil {
		t.Fatal(err)
	}
	wm.Watch(cf, func() {
		result = append(result, cf)
	})
	time.Sleep(time.Millisecond)

	err = ioutil.WriteFile(fn, []byte("t"), 700)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)
	if len(result) == 0 {
		t.Fatal(result)
	}
}

func TestWatcher2(t *testing.T) {
	result := []interface{}{}
	wm := configuration.NewWatchManager()
	m := NewWatcher()
	wm.RegisterWatcher(configuration.NewSchemeWatcher("file", m))
	err := wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer wm.Stop()
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	fn := tmpfile.Name()
	if fn == "" {
		t.Fatal(fn)
	}
	e := fsnotify.Event{
		Name: fn,
		Op:   fsnotify.Write,
	}
	defer os.Remove(fn)
	cf, err := configuration.New("file://" + fn)
	if err != nil {
		t.Fatal(err)
	}
	wm.Watch(cf, func() {
		result = append(result, cf)
	})
	time.Sleep(time.Millisecond)
	m.fswatcher.Events <- e
	time.Sleep(time.Second)
	if len(result) != 1 {
		t.Fatal(result)
	}
	result = []interface{}{}
	wm.Watch(cf, func() {
		result = append(result, cf)
	})
	if len(m.registeredFuncs[fn]) != 2 {
		t.Fatal(m)
	}
	time.Sleep(time.Millisecond)
	m.fswatcher.Events <- e
	time.Sleep(time.Second)
	if len(result) != 2 {
		t.Fatal(result)
	}

}

func TestErr(t *testing.T) {
	result := []interface{}{}
	wm := configuration.NewWatchManager()
	wm.ErrorHandler = func(err error) {
		result = append(result, err)
	}
	m := NewWatcher()
	wm.RegisterWatcher(configuration.NewSchemeWatcher("file", m))
	err := wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer wm.Stop()
	time.Sleep(time.Millisecond)
	go func() {
		m.fswatcher.Errors <- errors.New("error")
	}()
	time.Sleep(time.Millisecond)
	if len(result) == 0 || result[0].(error).Error() != "error" {
		t.Fatal(result)
	}
}
