package verification

import (
	"fmt"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type Verification struct {
	request *request.HttpRequest
}

func NewVerification(request *request.HttpRequest) *Verification {
	return &Verification{request}
}

func (v *Verification) ValidateBankAccount(body ValidateBankAccountModel) (*ValidateBankAccountResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	newUrl := fmt.Sprintf(constants.ValidateBankAccountEndpoint, body.AccountNumber, body.BankCode)
	res, err := v.request.Get(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ValidateBankAccountResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (v *Verification) VerifyBvnInformation(body VerifyBvnInformationModel) (*VerifyBvnInformationResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := v.request.Post(constants.VerifyBvnInformationEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[VerifyBvnInformationResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (v *Verification) MatchBvnAndAccountName(body MatchBvnAndAccountNameModel) (*MatchBvnAndAccountNameResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := v.request.Post(constants.MatchBvnAndAccountNameEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[MatchBvnAndAccountNameResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
