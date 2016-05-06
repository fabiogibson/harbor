package runners

import (
	"github.com/fabiogibson/harbor/core"
	"github.com/fatih/color"
	"os/exec"
	"path/filepath"
	"strings"
)

//
//
//
type TestRunner interface {
	ExecuteTest(filename string) ([]byte, error)
}

//
//
//
type AbstractTestRunner struct {
	core.Runner
	SupportedExtensions []string
	TestRunner
}

//
//
//
func (t *AbstractTestRunner) Run(filename string) error {
	if !t.HandleExtension(filepath.Ext(filename)) {
		return nil
	}

	t.printHeader(filename)

	testOutput, err := t.ExecuteTest(filename)

	if err != nil {
		t.printFail(testOutput)
	} else {
		t.printSuccess(testOutput)
	}

	return err
}

//
//
//
func (t *AbstractTestRunner) ExecuteTest(filename string) ([]byte, error) {
	return exec.Command("go", "test", "./"+filepath.Dir(filename)).CombinedOutput()
}

//
//
//
func (t *AbstractTestRunner) printHeader(filename string) {
	color.Blue("===> Running tests for %s", filename)
}

//
//
//
func (t *AbstractTestRunner) printFail(testOutput []byte) {
	color.Red("--- Tests Failed")
	color.Red("%s", testOutput)
	color.Set(color.BgRed)
	color.White("Go back to work!")
}

//
//
//
func (t *AbstractTestRunner) printSuccess(testOutput []byte) {
	color.Green("Test Result: %s", testOutput)
	color.Set(color.BgGreen)
	color.White("Good job! Keep going...")
}

//
//
//
func (t *AbstractTestRunner) HandleExtension(extension string) bool {
	for _, val := range t.SupportedExtensions {
		if val == strings.Replace(extension, ".", "", 1) {
			return true
		}
	}

	return false
}

//
//
//
func (t *AbstractTestRunner) HandledEvents() []int {
	return []int{16, 2}
}
