package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//
//
func TestShouldCreateNewFsWatcherInstance(t *testing.T) {
	// act
	fsWatcher := NewFsWatcher()

	// assert
	assert.NotNil(t, fsWatcher)
}

//
//
//
func TestShouldSubscribeAnEventListener(t *testing.T) {
	// mock
	fsWatcher := NewFsWatcher()

	// act
	fsWatcher.Subscribe(16, func(ev fsnotify.Event) { return })

	// assert
	assert.Len(t, fsWatcher.Subscribers, 1)
}

//
//
//
func TestShouldExecuteAnEventHandler(t *testing.T) {
	// mock
	fsWatcher := NewFsWatcher()
	invoked := false

	fsWatcher.Subscribe(16, func(e fsnotify.Event) { invoked = true })

	// act
	fsWatcher.notifySubscribers(16, fsnotify.Event{})

	// assert
	assert.True(t, invoked)
}
