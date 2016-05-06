package runners

import ()

//
//
//
type GoTestRunner struct {
	AbstractTestRunner
}

//
//
//
func NewGoTestRunner() *GoTestRunner {
	runner := &GoTestRunner{}
	runner.SupportedExtensions = []string{"go"}
	return runner
}
