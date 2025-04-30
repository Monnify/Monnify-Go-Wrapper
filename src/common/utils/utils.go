package utils

import (
	"bytes"
	"encoding/json"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	MonnifyAPIKey    string `mapstructure:"MONNIFY_API_KEY"`
	MonnifySecretKey string `mapstructure:"MONNIFY_SECRET_KEY"`
	ContractCode     string `mapstructure:"CONTRACT_CODE"`
	SourceAccount    string `mapstructure:"SOURCE_ACCOUNT"`
}

func LoadConfig(path string) (config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("unable to read config", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unable to unmarshal config", err)
	}

	return
}

func GetBaseUrl(isProduction bool) string {
	if isProduction {
		return "https://api.monnify.com"
	}

	return "https://sandbox.monnify.com"
}

func ValidateStruct(data any) *mErr.Error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(data); err != nil {
		return mErr.ErrorHandler("Validation Error", err, nil)
	}

	return nil
}

func ParseBody(body any) (*bytes.Reader, *mErr.Error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, mErr.ErrorHandler(err.Error(), err, nil)
	}

	reader := bytes.NewReader(jsonData)
	return reader, nil
}

func ParseResponse[G any](body io.ReadCloser) (*G, *mErr.Error) {
	resBody, err := io.ReadAll(body)
	if err != nil {
		return nil, mErr.ErrorHandler("failed to read response body", err, nil)
	}

	var resp G
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, mErr.ErrorHandler("failed to unmarshal response body", err, nil)
	}

	return &resp, nil
}

func GenerateRandomNumbers(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(r.Intn(10)))
	}

	return sb.String()
}

func GenerateRandomEmail() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "example.com"}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	usernameLength := r.Intn(7) + 6
	var username strings.Builder
	for i := 0; i < usernameLength; i++ {
		username.WriteByte(letters[r.Intn(len(letters))])
	}

	domain := domains[r.Intn(len(domains))]
	email := username.String() + "@" + domain
	return email
}
