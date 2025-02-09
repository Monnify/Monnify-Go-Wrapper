package monnify

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type HttpRequest struct {
	baseUrl string
}

func NewHttpRequest(baseUrl string) *HttpRequest {
	return &HttpRequest{baseUrl}
}

func (h *HttpRequest) CreateRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.New("could not create request")
	}

	return req, nil
}

func (h *HttpRequest) Post(url, authorization string, body io.Reader) (*http.Response, error) {
	req, err := h.CreateRequest(http.MethodPost, h.baseUrl+url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("error making http request")
	}

	return res, nil
}
