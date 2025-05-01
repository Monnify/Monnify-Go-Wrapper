package monnify

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/collections"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/Monnify/Monnify-Go-Wrapper/src/disbursement"
	"github.com/Monnify/Monnify-Go-Wrapper/src/verification"
)

type Options struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

type monnify struct {
	Disbursement    *disbursement.Disbursement
	ReservedAccount *collections.ReservedAccount
	SubAccount      *collections.SubAccount
	Transaction     *collections.Transaction
	Refund          *disbursement.Refund
	Verification    *verification.Verification
}

func New(options *Options) *monnify {
	baseUrl := utils.GetBaseUrl(options.IsProduction)
	httpRequest := request.NewHttpRequest(baseUrl, options.ApiKey+":"+options.SecretKey)

	return &monnify{
		Disbursement:    disbursement.NewDisbursement(httpRequest),
		ReservedAccount: collections.NewReservedAccount(httpRequest),
		SubAccount:      collections.NewSubAccount(httpRequest),
		Transaction:     collections.NewTransaction(httpRequest),
		Refund:          disbursement.NewRefund(httpRequest),
		Verification:    verification.NewVerification(httpRequest),
	}
}
