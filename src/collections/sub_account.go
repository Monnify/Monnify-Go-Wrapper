package collections

import (
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
