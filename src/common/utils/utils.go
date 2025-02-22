package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
)

func GetBaseUrl(isProduction bool) string {
	if isProduction {
		return "https://api.monnify.com"
	}

	return "https://sandbox.monnify.com"
}

func ValidateStruct(data any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(data)
}

func ParseBody(body any) (*bytes.Reader, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonData)
	return reader, nil
}

func ParseResponse[G any](body io.ReadCloser) (*G, error) {
	resBody, err := io.ReadAll(body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	var resp G
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	return &resp, nil
}
