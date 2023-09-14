package cumulocity


import (
    "testing"
	"github.com/stretchr/testify/assert"
)

func TestCleanupFile(t *testing.T) {
	var dirToRemove string
	osRemoveAll = func (name string) error {
		dirToRemove = name
		return nil
	}
	m := Microservice{}

	m.CleanupFile("/the/good/directory/is/this/but_not_this")

	assert.Equal(t, "/the/good/directory/is/this", dirToRemove)
} 