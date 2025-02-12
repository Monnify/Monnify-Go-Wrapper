package monnify

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/cache"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/token"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/Monnify/Monnify-Go-Wrapper/src/disbursement"
)

type MonnifyOptions struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

type Monnify struct {
	Disbursement *disbursement.Disbursement
}

func New(options *MonnifyOptions) *Monnify {
	baseUrl := utils.GetBaseUrl(options.IsProduction)
	cache := cache.NewCache()
	request := request.NewHttpRequest(baseUrl)
	token := token.NewToken(cache, baseUrl, options.ApiKey+":"+options.SecretKey)

	return &Monnify{
		Disbursement: disbursement.NewDisbursement(request, token),
	}
}
