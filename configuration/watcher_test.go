package configuration

import (
	"errors"
	"strings"
	"testing"
	"time"
)

type testWatcher struct {
	errc    chan error
	watched []*func()
}

func newTestWatcher() *testWatcher {
	return &testWatcher{
		errc:    make(chan error),
		watched: []*func(){},
	}
}
func (w *testWatcher) Watch(cf Configuration, callback func()) (unwatcher func(), err error) {
	if !strings.HasPrefix(cf.ID(), "text://") {
		return nil, nil
	}
	w.watched = append(w.watched, &callback)
	return func() {
		filtered := []*func(){}
		for _, v := range w.watched {
			if v == &callback {
				continue
			}
			filtered = append(filtered, v)
		}
		w.watched = filtered
	}, nil
}
func (w *testWatcher) Init() error {
	return nil
}
func (w *testWatcher) StartWatching() error {
	return nil
}
func (w *testWatcher) StopWatching() error {
	return nil
}
func (w *testWatcher) ErrorChan() chan error {
	return w.errc
}

func TestWatcher(t *testing.T) {
	var result = []interface{}{}
	wm := NewWatchManager()
	tm := newTestWatcher()
	m := NewSchemeWatcher("text", tm)
	wm.RegisterWatcher(m)
	text, err := New("text://test")
	if err != nil {
		t.Fatal(err)
	}
	uw, err := wm.Watch(text, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw == nil {
		t.Fatal(wm)
	}
	if len(tm.watched) != 1 {
		t.Fatal(tm)
	}
	uw2, err := wm.Watch(text, func() {
		result = append(result, text)
	})
	if len(tm.watched) != 2 {
		t.Fatal(tm)
	}
	file, err := New("file://tmp/test")
	if err != nil {
		t.Fatal(err)
	}
	uw3, err := wm.Watch(file, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw3 != nil {
		t.Fatal(wm)
	}
	if len(tm.watched) != 2 {
		t.Fatal(tm)
	}
	uw()
	if len(tm.watched) != 1 {
		t.Fatal(tm)
	}
	uw2()
	if len(tm.watched) != 0 {
		t.Fatal(tm)
	}
}

func TestUnwatch(t *testing.T) {
	var result = []interface{}{}
	wm := NewWatchManager()
	tm := newTestWatcher()
	m := NewSchemeWatcher("text", tm)
	wm.RegisterWatcher(m)
	text, err := New("text://test")
	if err != nil {
		t.Fatal(err)
	}
	uw, err := wm.Watch(text, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw == nil {
		t.Fatal(wm)
	}
	uw, err = wm.Watch(text, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw == nil {
		t.Fatal(wm)
	}
	uw, err = wm.Watch(text, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw == nil {
		t.Fatal(wm)
	}
	if len(tm.watched) != 3 {
		t.Fatal(tm)
	}
	wm.Unwatch()
	if len(tm.watched) != 0 {
		t.Fatal(tm)
	}
}

func TestStartAndStop(t *testing.T) {
	var result = []interface{}{}
	wm := NewWatchManager()
	tm := newTestWatcher()
	m := NewSchemeWatcher("text", tm)
	wm.RegisterWatcher(m)
	text, err := New("text://test")
	if err != nil {
		t.Fatal(err)
	}
	err = wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	uw, err := wm.Watch(text, func() {
		result = append(result, text)
	})
	if err != nil {
		t.Fatal(err)
	}
	if uw == nil {
		t.Fatal(wm)
	}
	if len(tm.watched) != 1 {
		t.Fatal(tm)
	}
	err = wm.Stop()
	if err != nil {
		t.Fatal(err)
	}
	if len(tm.watched) != 0 {
		t.Fatal(tm)
	}
	_, ok := <-wm.C()
	if ok {
		t.Fatal(wm)
	}
}

func TestErrorPanicHandler(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(t)
		}
		err := r.(error)
		if err == nil {
			t.Fatal(err)
		}
	}()
	PanicErrorHandler(errors.New("err"))
}

func TestErrorHandler(t *testing.T) {
	var result = []interface{}{}
	wm := NewWatchManager()
	tm := newTestWatcher()
	m := NewSchemeWatcher("text", tm)
	wm.RegisterWatcher(m)
	wm.ErrorHandler = func(err error) {
		result = append(result, err)
	}
	err := wm.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer wm.Stop()
	if len(result) != 0 {
		t.Fatal(result)
	}
	tm.errc <- errors.New("err")
	time.Sleep(time.Microsecond)
	if len(result) != 1 {
		t.Fatal(result)
	}
}
