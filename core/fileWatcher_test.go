package core

import (
	"github.com/bouk/monkey"
	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"reflect"
	"testing"
)

//
//
//
type MockFileInfo struct {
	mock.Mock
	os.FileInfo
}

//
//
//
func (i *MockFileInfo) IsDir() bool {
	args := i.Called()
	return args.Bool(0)
}

//
//
//
type MockedWatcher struct {
	mock.Mock
}

//
//
//
type MockedRunner struct {
	mock.Mock
	Runner
}

//
//
//
func (m *MockedRunner) Run(filename string) error {
	m.Called(filename)
	return nil
}

//
//
//
func (m *MockedRunner) HandledEvents() []int {
	return []int{16, 2}
}

//
//
//
func PatchAdd() *bool {
	var w *fsnotify.Watcher
	var called bool

	monkey.PatchInstanceMethod(reflect.TypeOf(w), "Add", func(watcher *fsnotify.Watcher, path string) error {
		called = true
		return nil
	})

	return &called
}

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
	watcher := NewFsWatcher()
	runner := &MockedRunner{}

	// act
	watcher.Subscribe(runner)

	// assert
	assert.Len(t, watcher.Subscribers[16], 1)
	assert.Len(t, watcher.Subscribers[2], 1)
}

//
//
//
func TestShouldInvokeRunnersRunMethod(t *testing.T) {
	// mock
	watcher := NewFsWatcher()
	runner := &MockedRunner{}
	runner.On("Run", "file.go").Return(nil)

	// act
	watcher.Subscribers[16] = []Runner{runner}
	watcher.notifySubscribers(16, fsnotify.Event{Name: "file.go"})

	// assert
	runner.AssertCalled(t, "Run", "file.go")
}

//
//
//
func TestVisitFileInfoShouldNotAddFilesToWatcher(t *testing.T) {
	// mock
	mockFileInfo := &MockFileInfo{}
	mockFileInfo.On("IsDir").Return(false)
	fw := NewFsWatcher()
	fw.watcher, _ = fsnotify.NewWatcher()
	invoked := PatchAdd()

	// act
	fw.visitFileInfo("path", mockFileInfo, nil)

	// assert
	assert.False(t, *invoked)

	// cleanup
	monkey.UnpatchAll()
}

//
//
//
func TestVisitFileInfoShouldAddDirectoryToWatcher(t *testing.T) {
	// mock
	mockFileInfo := &MockFileInfo{}
	mockFileInfo.On("IsDir").Return(true)
	fw := NewFsWatcher()
	fw.watcher, _ = fsnotify.NewWatcher()
	invoked := PatchAdd()

	// act
	fw.visitFileInfo("path", mockFileInfo, nil)

	// assert
	assert.True(t, *invoked)

	// cleanup
	monkey.UnpatchAll()
}
