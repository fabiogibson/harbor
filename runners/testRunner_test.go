package runners

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//
//
func TestShouldCreateNewTestRunnerInstance(t *testing.T) {
	// act
	testRunner := NewGoTestRunner()

	// assert
	assert.NotNil(t, testRunner)
}

//
//
//
func TestShouldReturnSuportedExtensions(t *testing.T) {
	// act
	testRunner := NewGoTestRunner()

	// assert
	assert.Contains(t, testRunner.SupportedExtensions, "go")
}

//
//
//
func TestShouldShouldReturnTrueForGoExtension(t *testing.T) {
	// act
	testRunner := NewGoTestRunner()

	// assert
	assert.True(t, testRunner.HandleExtension("go"))
}

//
//
//
func TestShouldShouldReturnHandledEvents(t *testing.T) {
	// act
	testRunner := NewGoTestRunner()

	// assert
	assert.Contains(t, testRunner.HandledEvents(), 16)
	assert.Contains(t, testRunner.HandledEvents(), 2)
}
