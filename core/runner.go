package core

//
//
//
type Runner interface {
	Run(filename string) error
	HandleExtension(extension string) bool
	HandledEvents() []int
}
