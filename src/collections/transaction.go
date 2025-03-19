package collections

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type Transaction struct {
	request *request.HttpRequest
}

func NewTransaction(request *request.HttpRequest) *Transaction {
	return &Transaction{request}
}

func (t *Transaction) InitializeTransaction(body InitializeTransactionModel) (*InitializeTransactionResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := t.request.Post(constants.InitTransactionEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[InitializeTransactionResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
