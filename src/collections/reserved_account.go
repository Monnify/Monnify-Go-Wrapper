package collections

import (
	"fmt"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type ReservedAccount struct {
	request *request.HttpRequest
}

func NewReservedAccount(request *request.HttpRequest) *ReservedAccount {
	return &ReservedAccount{request}
}

func (r *ReservedAccount) CreateReservedAccount(body ReservedAccountSchema) (*ReservedAccountResponse, error) {
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

func (r *ReservedAccount) AddLinkedAccounts(body AddLinkedAccountSchema) (*AddLinkedAccountResponse, error) {
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
