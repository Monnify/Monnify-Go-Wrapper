package monnify

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/cache"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/token"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type MonnifyOptions struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

func New(options *MonnifyOptions) string {
	baseUrl := utils.GetBaseUrl(options.IsProduction)
	cache := cache.NewCache()
	_ = token.NewToken(cache, baseUrl, options.ApiKey+":"+options.SecretKey)

	return "Hello"
}
