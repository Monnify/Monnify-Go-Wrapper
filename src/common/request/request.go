package request

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/cache"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type HttpRequest struct {
	baseUrl     string
	cache       *cache.Cache
	credentials string
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

func NewHttpRequest(baseUrl string, credentials string) *HttpRequest {
	return &HttpRequest{baseUrl: baseUrl, cache: cache.NewCache(), credentials: credentials}
}

func (h *HttpRequest) generateToken() (string, error) {
	base64Str := base64.StdEncoding.EncodeToString([]byte(h.credentials))
	req, err := h.CreateRequest(http.MethodPost, constants.LoginEndpoint, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64Str)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("error making http request")
	}

	defer res.Body.Close()

	body, err := utils.ParseResponse[Login](res.Body)
	if err != nil {
		return "", err
	}

	h.cache.Set(constants.AuthentionKey, body.ResponseBody.AccessToken, body.ResponseBody.ExpiresIn*time.Second)
	return body.ResponseBody.AccessToken, nil
}

func (h *HttpRequest) getToken() (string, error) {
	value, ok := h.cache.Get(constants.AuthentionKey)
	if ok {
		return "Bearer " + value, nil
	}

	token, err := h.generateToken()
	if err != nil {
		return "", err
	}

	return "Bearer " + token, nil
}

func (h *HttpRequest) CreateRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, h.baseUrl+url, body)
	if err != nil {
		return nil, errors.New("could not create request")
	}

	return req, nil
}

func (h *HttpRequest) Post(url string, body any) (*http.Response, error) {
	token, err := h.getToken()
	if err != nil {
		return nil, err
	}

	parsedBody, err := utils.ParseBody(body)
	if err != nil {
		return nil, err
	}

	req, err := h.CreateRequest(http.MethodPost, url, parsedBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error making http request")
	}

	return res, nil
}

func (h *HttpRequest) Get(url string) (*http.Response, error) {
	token, err := h.getToken()
	if err != nil {
		return nil, err
	}

	req, err := h.CreateRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error making http request")
	}

	return res, nil
}
