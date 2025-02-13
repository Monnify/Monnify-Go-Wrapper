package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	request := NewHttpRequest("invalid+url", "")

	// Test against wrong method
	_, reqErr := request.CreateRequest(" ", "", nil)
	assert.Equal(t, "could not create request", reqErr.Error(), "message should be 'could not create request'")

	_, postErr := request.Post("", nil)
	assert.Equal(t, "error making http request", postErr.Error(), "message should be 'error making http request'")
}
