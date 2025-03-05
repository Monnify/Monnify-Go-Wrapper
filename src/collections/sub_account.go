package collections

import "github.com/Monnify/Monnify-Go-Wrapper/src/common/request"

type SubAccount struct {
	request *request.HttpRequest
}

func NewSubAccount(request *request.HttpRequest) *SubAccount {
	return &SubAccount{request}
}
