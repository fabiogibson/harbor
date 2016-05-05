package main

import (
	"fmt"
	"github.com/fabiogibson/harbor/core"
	"github.com/fsnotify/fsnotify"
)

func main() {
	fmt.Println("Initializing watcher...")
	fsWatcher := core.NewFsWatcher()

	testRunner := func(ev fsnotify.Event) {
		go core.RunTests(ev.Name)
	}

	fsWatcher.Subscribe(16, testRunner)
	fsWatcher.Subscribe(2, testRunner)

	fmt.Println("Watching for filesystem changes...")
	fsWatcher.Init("./")
}
