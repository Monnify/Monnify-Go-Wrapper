package monnify

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/collections"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/Monnify/Monnify-Go-Wrapper/src/disbursement"
)

type Options struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

type Monnify struct {
	Disbursement    *disbursement.Disbursement
	ReservedAccount *collections.ReservedAccount
	SubAccount      *collections.SubAccount
	Transaction     *collections.Transaction
}

func New(options *Options) *Monnify {
	baseUrl := utils.GetBaseUrl(options.IsProduction)
	httpRequest := request.NewHttpRequest(baseUrl, options.ApiKey+":"+options.SecretKey)

	return &Monnify{
		Disbursement:    disbursement.NewDisbursement(httpRequest),
		ReservedAccount: collections.NewReservedAccount(httpRequest),
		SubAccount:      collections.NewSubAccount(httpRequest),
		Transaction:     collections.NewTransaction(httpRequest),
	}
}
