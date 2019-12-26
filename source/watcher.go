package source

import (
	"strings"
)

type Watcher interface {
	Watch(s Source, callback func()) (unwatcher func(), err error)
	Init() error
	StartWatching() error
	StopWatching() error
	ErrorChan() chan error
}

type SchemeWatcher struct {
	Scheme string
	Watcher
}

func (w *SchemeWatcher) Watch(s Source, callback func()) (unwatcher func(), err error) {
	if strings.HasPrefix(s.ID(), w.Scheme+"://") {
		return w.Watcher.Watch(s, callback)
	}
	return nil, nil
}

func NewSchemeWatcher(scheme string, watcher Watcher) *SchemeWatcher {
	return &SchemeWatcher{
		Scheme:  scheme,
		Watcher: watcher,
	}
}

type Unwatcher struct {
	Hanlder func()
	Source  Source
}

func PanicErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

type WatchManager struct {
	registeredWatchers []Watcher
	Unwatchers         []*Unwatcher
	c                  chan int
	ErrorHandler       func(err error)
}

func (w *WatchManager) C() chan int {
	return w.c
}
func (w *WatchManager) AddUnwatcher(s Source, h func()) {
	w.Unwatchers = append(w.Unwatchers, &Unwatcher{
		Hanlder: h,
		Source:  s,
	})
}

func (w *WatchManager) Start() error {
	w.c = make(chan int)
	for _, v := range w.registeredWatchers {
		err := v.StartWatching()
		if err != nil {
			return err
		}
		go func() {
			for {
				select {
				case e := <-v.ErrorChan():
					w.ErrorHandler(e)
				case <-w.c:
					return
				}
			}
		}()
	}

	return nil
}

func (w *WatchManager) Stop() error {
	w.Unwatch()
	for _, v := range w.registeredWatchers {
		err := v.StopWatching()
		if err != nil {
			return err
		}
	}
	close(w.c)
	return nil
}
func (w *WatchManager) Reset() error {
	for _, v := range w.registeredWatchers {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
func (w *WatchManager) Unwatch() {
	for _, v := range w.Unwatchers {
		v.Hanlder()
	}
	w.Unwatchers = []*Unwatcher{}
}
func (w *WatchManager) Watch(s Source, callback func()) (unwatcher func(), err error) {
	for _, v := range w.registeredWatchers {
		uw, err := v.Watch(s, callback)
		if err != nil {
			return nil, err
		}
		if uw != nil {
			w.AddUnwatcher(s, uw)
			return uw, nil
		}
	}
	return nil, nil
}
func (w *WatchManager) RegisterWatcher(watcher Watcher) error {
	err := watcher.Init()
	if err != nil {
		return err
	}
	w.registeredWatchers = append(w.registeredWatchers, watcher)
	return nil
}

func NewWatchManager() *WatchManager {
	return &WatchManager{
		registeredWatchers: []Watcher{},
		Unwatchers:         []*Unwatcher{},
		ErrorHandler:       PanicErrorHandler,
	}
}
