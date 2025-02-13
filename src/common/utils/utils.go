package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"slices"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/go-playground/validator/v10"
)

func GetBaseUrl(isProduction bool) string {
	if isProduction {
		return "https://api.monnify.com"
	}

	return "https://sandbox.monnify.com"
}

func validateCurrency(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if !slices.Contains(constants.SupportedCurrency, field) {
		return false
	}

	return true
}

func validationEnum(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if !slices.Contains(constants.ValidationEnum, field) {
		return false
	}

	return true
}

func ValidateStruct(data any) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("validateCurrency", validateCurrency)
	validate.RegisterValidation("validationEnum", validationEnum)
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

func ParseResponse[G comparable](body io.ReadCloser) (*G, error) {
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
