package monnify

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/collections"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/Monnify/Monnify-Go-Wrapper/src/disbursement"
)

type MonnifyOptions struct {
	ApiKey       string
	SecretKey    string
	IsProduction bool
}

type Monnify struct {
	Disbursement    *disbursement.Disbursement
	ReservedAccount *collections.ReservedAccount
	SubAccount      *collections.SubAccount
}

func New(options *MonnifyOptions) *Monnify {
	baseUrl := utils.GetBaseUrl(options.IsProduction)
	request := request.NewHttpRequest(baseUrl, options.ApiKey+":"+options.SecretKey)

	return &Monnify{
		Disbursement:    disbursement.NewDisbursement(request),
		ReservedAccount: collections.NewReservedAccount(request),
		SubAccount:      collections.NewSubAccount(request),
	}
}
