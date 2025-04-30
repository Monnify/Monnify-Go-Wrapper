package request

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getCredentials() string {
	credentials := utils.LoadConfig("../../..")
	return credentials.MonnifyAPIKey + ":" + credentials.MonnifySecretKey
}

func TestGenerateTokenFailed(t *testing.T) {
	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.generateToken()
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")
}

func TestGenerateToken(t *testing.T) {
	credentials := getCredentials()

	// Wrong credentials
	wReq := NewHttpRequest(utils.GetBaseUrl(false), "blah"+":"+"blah")
	_, err := wReq.generateToken()
	assert.NotNil(t, err)
	assert.Equal(t, false, err.Response.RequestSuccessful)

	// Valid credentials
	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)
	resp, _ := request.generateToken()
	assert.NotEmpty(t, resp)
	assert.IsType(t, string(""), resp, "value should be a string")
}

func TestGetTokenFailed(t *testing.T) {
	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.getToken()
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")
}

func TestGetToken(t *testing.T) {
	credentials := getCredentials()

	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)
	resp, _ := request.getToken()
	assert.NotEmpty(t, resp)
	assert.IsType(t, string(""), resp, "value should be a string")

	cacheResp, _ := request.getToken()
	assert.NotEmpty(t, cacheResp)
	assert.IsType(t, string(""), cacheResp, "value should be a string")
}

func TestCreateRequestFailed(t *testing.T) {
	request := NewHttpRequest("invalid+url", "")
	_, reqErr := request.CreateRequest(" ", "", nil)
	assert.Equal(t, "could not create request", reqErr.Message, "message should be 'could not create request'")
}

func TestPostRequest(t *testing.T) {
	credentials := getCredentials()

	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.Post(" ", nil)
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")

	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)

	// TEST FAILED
	_, fErr := request.Post("/", nil)
	assert.NotNil(t, fErr)
	assert.Equal(t, false, fErr.Response.RequestSuccessful)

	// TEST SUCCESS
	//resp, _ := request.Post("/", nil)
}

func TestGetRequest(t *testing.T) {
	credentials := getCredentials()

	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.Get(" ")
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")

	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)

	// TEST FAILED
	_, fErr := request.Get("/")
	assert.NotNil(t, fErr)
	assert.Equal(t, false, fErr.Response.RequestSuccessful)

	// TEST SUCCESS
	resp, _ := request.Get(constants.GetSubAccountEndpoint)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPutRequest(t *testing.T) {
	credentials := getCredentials()

	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.Put(" ", nil)
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")

	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)

	// TEST FAILED
	_, fErr := request.Put("/", nil)
	assert.NotNil(t, fErr)
	assert.Equal(t, false, fErr.Response.RequestSuccessful)

	// TEST SUCCESS
	//resp, _ := request.Post("/", nil)
}

func TestDeleteRequest(t *testing.T) {
	credentials := getCredentials()

	wReq := NewHttpRequest("invalid+url", "blah"+":"+"blah")
	_, err := wReq.Delete(" ")
	assert.NotNil(t, err) // Error not nil
	assert.Equal(t, "error making http request", err.Message, "message should be 'error making http request'")

	request := NewHttpRequest(utils.GetBaseUrl(false), credentials)

	// TEST FAILED
	_, fErr := request.Delete("/")
	assert.NotNil(t, fErr)
	assert.Equal(t, false, fErr.Response.RequestSuccessful)

	// TEST SUCCESS
	//resp, _ := request.Post("/", nil)
}
