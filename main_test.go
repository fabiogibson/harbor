package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//
//
func TestShouldInitNewWatcher(t *testing.T) {
	// act
	fsWatcher := initWatcher()

	// assert
	assert.NotNil(t, fsWatcher)
	assert.Len(t, fsWatcher.Subscribers[2], 1)
	assert.Len(t, fsWatcher.Subscribers[16], 1)
}
