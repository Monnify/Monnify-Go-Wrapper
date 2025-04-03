package collections

import (
	"fmt"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type ReservedAccount struct {
	request *request.HttpRequest
}

func NewReservedAccount(request *request.HttpRequest) *ReservedAccount {
	return &ReservedAccount{request}
}

func (r *ReservedAccount) CreateReservedAccount(body ReservedAccountSchema) (*ReservedAccountResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()

	res, err := r.request.Post(constants.CreateReservedAccountEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ReservedAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *ReservedAccount) AddLinkedAccounts(body AddLinkedAccountSchema) (*AddLinkedAccountResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.AddLinkedAccountEndpoint, body.AccountReference)
	res, err := r.request.Put(url, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[AddLinkedAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *ReservedAccount) ReservedAccountDetails(body ReservedAccountDetailsSchema) (*ReservedAccountDetailsResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.ReservedAccountDetailsEndpoint, body.AccountReference)
	res, err := r.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ReservedAccountDetailsResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *ReservedAccount) ReservedAccountTransactions(body ReservedAccountTransactionsSchema) (*ReservedAccountTransactionsResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()
	url := fmt.Sprintf(constants.ReservedAccountTransactionsEndpoint, body.AccountReference, body.Page, body.Size)
	res, err := r.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ReservedAccountTransactionsResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *ReservedAccount) DeallocateReservedAccount(body DeallocateReservedAccountSchema) (*DeallocateReservedAccountResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.DeallocateReservedAccountEndpoint, body.AccountReference)
	res, err := r.request.Delete(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[DeallocateReservedAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *ReservedAccount) UpdateReservedAccountKycInfo(body UpdateReservedAccountKycInfoSchema) (*UpdateReservedAccountKycInfoResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.UpdateReservedAccountKycInfoEndpoint, body.AccountReference)
	res, err := r.request.Put(url, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[UpdateReservedAccountKycInfoResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
