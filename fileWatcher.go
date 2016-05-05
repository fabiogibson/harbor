package harbor

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

//
// FsWatcher watches for filesystem changes and notify it's subscribers
//
type FsWatcher struct {
	Subscribers map[fsnotify.Op][]func(ev fsnotify.Event)
	watcher     *fsnotify.Watcher
}

//
// Subscribe subscribes an event listener to file watcher.
//
func (f *FsWatcher) Subscribe(eventOp fsnotify.Op, handler func(ev fsnotify.Event)) {
	f.Subscribers[eventOp] = append(f.Subscribers[eventOp], handler)
}

//
// notifySubscribers notifies all subscribed listeners when an event occurs.
//
func (f *FsWatcher) notifySubscribers(eventOp fsnotify.Op, event fsnotify.Event) {
	if val, ok := f.Subscribers[eventOp]; ok {
		for _, handler := range val {
			handler(event)
		}
	}
}

//
// loadDirectories adds all sub directories in rootPath to fsNotify Watcher.
//
func (f *FsWatcher) loadDirectories(rootPath string) {
	visit := func(p string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			err = f.watcher.Add(path.Join(rootPath, p))
		}

		return nil
	}

	filepath.Walk(rootPath, visit)
}

//
// init initializes new fsNotify Watcher in path.
//
func (f *FsWatcher) Init(path string) {
	watcher, err := fsnotify.NewWatcher()
	f.watcher = watcher

	if err != nil {
		log.Fatal(err)
	}

	defer f.watcher.Close()
	done := make(chan bool)
	locked := false

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if !locked {
					locked = true
					f.notifySubscribers(event.Op, event)
					time.AfterFunc(1*time.Second, func() { locked = false })
				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	f.loadDirectories(path)

	if err != nil {
		log.Fatal(err)
	}

	<-done
}

//
// NewFsWatcher creates a new FsWatcher instance
//
func NewFsWatcher() *FsWatcher {
	return &FsWatcher{
		Subscribers: make(map[fsnotify.Op][]func(ev fsnotify.Event)),
	}
}
