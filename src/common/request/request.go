package request

import (
	"encoding/base64"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
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

func (h *HttpRequest) generateToken() (string, *mErr.Error) {
	base64Str := base64.StdEncoding.EncodeToString([]byte(h.credentials))
	req, reqErr := h.CreateRequest(http.MethodPost, constants.LoginEndpoint, nil)
	if reqErr != nil {
		return "", reqErr
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64Str)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, resErr := client.Do(req)
	if resErr != nil {
		return "", mErr.ErrorHandler("error making http request", resErr, nil)
	}

	defer res.Body.Close()

	body, parseErr := utils.ParseResponse[Login](res.Body)
	if parseErr != nil {
		return "", parseErr
	}

	h.cache.Set(constants.AuthentionKey, body.ResponseBody.AccessToken, body.ResponseBody.ExpiresIn*time.Second)
	return body.ResponseBody.AccessToken, nil
}

func (h *HttpRequest) getToken() (string, *mErr.Error) {
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

func (h *HttpRequest) CreateRequest(method, url string, body io.Reader) (*http.Request, *mErr.Error) {
	req, err := http.NewRequest(method, h.baseUrl+url, body)
	if err != nil {
		return nil, mErr.ErrorHandler("could not create request", err, nil)
	}

	return req, nil
}

func (h *HttpRequest) Post(url string, body any) (*http.Response, *mErr.Error) {
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

	res, resError := client.Do(req)
	if resError != nil {
		return nil, mErr.ErrorHandler("error making http request", resError, nil)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	errResp, err := utils.ParseResponse[mErr.ErrResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return nil, mErr.ErrorHandler("", nil, errResp)
}

func (h *HttpRequest) Get(url string) (*http.Response, *mErr.Error) {
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

	res, resError := client.Do(req)
	if resError != nil {
		return nil, mErr.ErrorHandler("error making http request", resError, nil)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	errResp, err := utils.ParseResponse[mErr.ErrResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return nil, mErr.ErrorHandler("", nil, errResp)
}

func (h *HttpRequest) Put(url string, body any) (*http.Response, *mErr.Error) {
	token, err := h.getToken()
	if err != nil {
		return nil, err
	}

	parsedBody, err := utils.ParseBody(body)
	if err != nil {
		return nil, err
	}

	req, err := h.CreateRequest(http.MethodPut, url, parsedBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, resError := client.Do(req)
	if resError != nil {
		return nil, mErr.ErrorHandler("error making http request", resError, nil)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	errResp, err := utils.ParseResponse[mErr.ErrResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return nil, mErr.ErrorHandler("", nil, errResp)
}

func (h *HttpRequest) Delete(url string) (*http.Response, *mErr.Error) {
	token, err := h.getToken()
	if err != nil {
		return nil, err
	}

	req, err := h.CreateRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, resError := client.Do(req)
	if resError != nil {
		return nil, mErr.ErrorHandler("error making http request", resError, nil)
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return res, nil
	}

	errResp, err := utils.ParseResponse[mErr.ErrResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return nil, mErr.ErrorHandler("", nil, errResp)
}
