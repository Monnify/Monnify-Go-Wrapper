package collections

import (
	"fmt"
	"net/url"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type SubAccount struct {
	request *request.HttpRequest
}

func NewSubAccount(request *request.HttpRequest) *SubAccount {
	return &SubAccount{request}
}

func (s *SubAccount) CreateSubAccount(body CreateSubAccountModel) (*CreateSubAccountResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := s.request.Post(constants.CreateSubAccountEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[CreateSubAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (s *SubAccount) DeleteSubAccount(body DeleteSubAccountModel) (*DeleteSubAccountResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	encodedSubAccountCode := url.QueryEscape(body.SubAccountCode)
	newUrl := fmt.Sprintf(constants.DeleteSubAccountEndpoint, encodedSubAccountCode)
	res, err := s.request.Delete(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[DeleteSubAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (s *SubAccount) GetSubAccounts() (*GetSubAccountsResponse, error) {
	res, err := s.request.Get(constants.GetSubAccountEndpoint)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetSubAccountsResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (s *SubAccount) UpdateSubAccount(body UpdateSubAccountModel) (*UpdateSubAccountResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := s.request.Put(constants.UpdateSubAccountEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[UpdateSubAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
