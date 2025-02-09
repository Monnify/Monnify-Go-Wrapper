package monnify

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestValidToken(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cache := NewCache()
	baseUrl := GetBaseUrl(false)
	apiKey := os.Getenv("MONNIFY_API_KEY")
	secretKey := os.Getenv("MONNIFY_SECRET_KEY")
	token := NewToken(cache, baseUrl, apiKey+":"+secretKey)

	// Get uncached token
	getToken, _ := token.GetToken()
	assert.Equal(t, true, getToken != "", "Token is valid")

	// Get cached token
	getTokenCached, _ := token.GetToken()
	assert.Equal(t, true, getTokenCached != "", "Token is valid")
}
