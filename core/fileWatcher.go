package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"time"
)

//
// FsWatcher watches for filesystem changes and notify it's subscribers
//
type FsWatcher struct {
	Subscribers map[int][]Runner
	watcher     *fsnotify.Watcher
}

//
// Subscribe subscribes an event listener to file watcher.
//
func (f *FsWatcher) Subscribe(runner Runner) {
	for _, ev := range runner.HandledEvents() {
		f.Subscribers[ev] = append(f.Subscribers[ev], runner)
	}
}

//
// notifySubscribers notifies all subscribed listeners when an event occurs.
//
func (f *FsWatcher) notifySubscribers(eventOp fsnotify.Op, event fsnotify.Event) {
	if runners, exist := f.Subscribers[int(eventOp)]; exist {
		for _, runner := range runners {
			fmt.Println("Calling run")
			runner.Run(event.Name)
		}
	}
}

//
//
//
func (f *FsWatcher) visitFileInfo(p string, fi os.FileInfo, err error) error {
	if fi.IsDir() {
		return f.watcher.Add(p)
	}

	return nil
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

	filepath.Walk(path, f.visitFileInfo)

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
		Subscribers: make(map[int][]Runner),
	}
}
