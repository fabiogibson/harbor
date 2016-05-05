package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"harbor/runner"
	"harbor/watcher"
)

func main() {
	fmt.Println("Initializing watcher...")
	fsWatcher := watcher.NewFsWatcher()

	testRunner := func(ev fsnotify.Event) {
		go runner.RunTests(ev.Name)
	}

	fsWatcher.Subscribe(16, testRunner)
	fsWatcher.Subscribe(2, testRunner)

	fmt.Println("Watching for filesystem changes...")
	fsWatcher.Init("./")
}