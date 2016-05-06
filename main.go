package main

import (
	"fmt"
	"github.com/fabiogibson/harbor/core"
	"github.com/fabiogibson/harbor/runners"
)

func initWatcher() *core.FsWatcher {
	fsWatcher := core.NewFsWatcher()
	fsWatcher.Subscribe(runners.NewGoTestRunner())
	defer fmt.Println("Watching for filesystem changes...")
	return fsWatcher
}

func main() {
	fmt.Println("Initializing watcher...")
	initWatcher().Init("./")
}
