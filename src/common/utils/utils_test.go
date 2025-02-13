package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseUrl(t *testing.T) {
	baseUrl := GetBaseUrl(false)
	assert.Equal(t, "https://sandbox.monnify.com", baseUrl, "expected 'https://sandbox.monnify.com' but got "+baseUrl)

	baseProdUrl := GetBaseUrl(true)
	assert.Equal(t, "https://api.monnify.com", baseProdUrl, "expected 'https://api.monnify.com' but got "+baseProdUrl)
}
