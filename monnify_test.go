package monnify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialization(t *testing.T) {
	result := New(&Options{
		ApiKey:    "",
		SecretKey: "",
	})

	assert.Equal(t, "Hello", result, "Should return 'Hello' ")
}
