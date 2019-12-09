package fsnotifywatcher

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/herb-go/herbconfig/configuration"
)

// Op describes a set of file operations.
type Op uint32

var (
	OpCreate = Op(fsnotify.Create)
	OpWrite  = Op(fsnotify.Write)
	OpRemove = Op(fsnotify.Remove)
	OpRename = Op(fsnotify.Rename)
	OpChmod  = Op(fsnotify.Chmod)
)

type Event struct {
	e *fsnotify.Event
}

func (e *Event) Path() string {
	return e.e.Name
}

func (e *Event) IsCreate() bool {
	return e.e.Op&fsnotify.Create == fsnotify.Create
}

func (e *Event) IsWrite() bool {
	return e.e.Op&fsnotify.Write == fsnotify.Write
}

func (e *Event) IsRemove() bool {
	return e.e.Op&fsnotify.Remove == fsnotify.Remove
}

func (e *Event) IsReName() bool {
	return e.e.Op&fsnotify.Rename == fsnotify.Rename
}

func (e *Event) IsChmod() bool {
	return e.e.Op&fsnotify.Chmod == fsnotify.Chmod
}

type Watcher struct {
	locker          sync.Mutex
	fswatcher       *fsnotify.Watcher
	errc            chan error
	registeredFuncs map[string][]*func()
	c               chan int
}

func NewWatcher() *Watcher {
	return &Watcher{
		registeredFuncs: map[string][]*func(){},
	}
}
func (w *Watcher) Watch(cf configuration.Configuration, callback func()) (unwatcher func(), err error) {
	w.locker.Lock()
	defer w.locker.Unlock()
	path := cf.AbsolutePath()
	if path != "" {
		if w.registeredFuncs[path] == nil {
			w.registeredFuncs[path] = []*func(){&callback}
			err := w.fswatcher.Add(path)
			if err != nil {
				return nil, err
			}
		} else {
			w.registeredFuncs[path] = append(w.registeredFuncs[path], &callback)
		}
		return w.unwatch(path, &callback), nil
	}
	return nil, nil
}

func (w *Watcher) unwatch(path string, cb *func()) func() {
	return func() {
		w.locker.Lock()
		defer w.locker.Unlock()
		result := []*func(){}
		fns := w.registeredFuncs[path]
		for k := range fns {
			if fns[k] == cb {
				continue
			}
			result = append(result, fns[k])
		}
		w.registeredFuncs[path] = result
	}
}
func (w *Watcher) On(e *Event) {
	w.locker.Lock()
	defer w.locker.Unlock()
	if e.IsWrite() {
		fns := w.registeredFuncs[e.Path()]
		for k := range fns {
			(*fns[k])()
		}
	}

}
func (w *Watcher) StartWatching() error {
	w.c = make(chan int)
	w.errc = make(chan error)
	go func() {
		for {
			select {
			case event := <-w.fswatcher.Events:
				w.On(&Event{&event})
			case err, ok := <-w.fswatcher.Errors:
				if ok {
					w.errc <- err
				}
			case <-w.c:
				return
			}
		}
	}()
	return nil
}
func (w *Watcher) StopWatching() error {
	close(w.c)
	go func() {
		close(w.errc)
	}()
	return w.fswatcher.Close()
}
func (w *Watcher) ErrorChan() chan error {
	return w.errc
}

func (w *Watcher) Init() error {
	var err error
	w.locker.Lock()
	defer w.locker.Unlock()
	w.registeredFuncs = map[string][]*func(){}
	w.fswatcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	return nil
}
