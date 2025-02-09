package token

import (
	"log"
	"os"
	"testing"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/cache"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestValidToken(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cache := cache.NewCache()
	baseUrl := utils.GetBaseUrl(false)
	apiKey := os.Getenv("MONNIFY_API_KEY")
	secretKey := os.Getenv("MONNIFY_SECRET_KEY")
	token := NewToken(cache, baseUrl, apiKey+":"+secretKey)

	// Get uncached token
	getToken, _ := token.GetToken()
	assert.Equal(t, true, getToken != "", "expected 'true' but got "+getToken != "")

	// Get cached token
	getTokenCached, _ := token.GetToken()
	assert.Equal(t, true, getTokenCached != "", "expected 'true' but got "+getTokenCached != "")
}
