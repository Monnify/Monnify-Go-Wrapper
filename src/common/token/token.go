package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/cache"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
)

type Token struct {
	cache       *cache.Cache
	credentials string
	request     *request.HttpRequest
}

type Login struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		AccessToken string        `json:"accessToken"`
		ExpiresIn   time.Duration `json:"expiresIn"`
	} `json:"responseBody"`
}

func NewToken(cache *cache.Cache, baseUrl, credentials string) *Token {
	return &Token{cache: cache, credentials: credentials, request: request.NewHttpRequest(baseUrl)}
}

func (t *Token) GenerateToken() (string, error) {
	base64Str := base64.StdEncoding.EncodeToString([]byte(t.credentials))
	res, err := t.request.Post(constants.LoginEndpoint, "Basic "+base64Str, nil)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	var data Login
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", errors.New("failed to unmarshal response body")
	}

	t.cache.Set(constants.AuthentionKey, data.ResponseBody.AccessToken, data.ResponseBody.ExpiresIn*time.Second)
	return data.ResponseBody.AccessToken, nil
}

func (t *Token) GetToken() (string, error) {
	value, ok := t.cache.Get(constants.AuthentionKey)
	if ok {
		return value, nil
	}

	return t.GenerateToken()
}
