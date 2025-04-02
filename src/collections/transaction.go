package collections

import (
	"fmt"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type Transaction struct {
	request *request.HttpRequest
}

func NewTransaction(request *request.HttpRequest) *Transaction {
	return &Transaction{request}
}

func (t *Transaction) InitializeTransaction(body InitializeTransactionModel) (*InitializeTransactionResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()

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

func (t *Transaction) GetTransactionStatusv2(body GetTransactionStatusv2Model) (*GetTransactionStatusv2Response, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.GetTransactionStatusv2Endpoint, body.TransactionReference)
	res, err := t.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetTransactionStatusv2Response](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (t *Transaction) GetTransactionStatusv1(body GetTransactionStatusv1Model) (*GetTransactionStatusv1Response, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.GetTransactionStatusv1Endpoint, body.PaymentReference)
	res, err := t.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetTransactionStatusv1Response](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

// TODO: can't find the endpoint for PayWithUssd
//func (t *Transaction) PayWithUssd(body PayWithUssdModel) (*PayWithUssdResponse, error) {
//
//}

func (t *Transaction) PayWithBankTransfer(body PayWithBankTransferModel) (*PayWithBankTransferResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := t.request.Post(constants.PayWithBankTransferEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[PayWithBankTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (t *Transaction) ChargeCard(body ChargeCardModel) (*ChargeCardResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := t.request.Post(constants.ChargeCardEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ChargeCardResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (t *Transaction) AuthorizeOTP(body AuthorizeOTPModel) (*AuthorizeOTPResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := t.request.Post(constants.AuthorizeOTPEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[AuthorizeOTPResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (t *Transaction) ThreeDsSecureAuthTransaction(body ThreeDsSecureAuthTransactionModel) (*ThreeDsSecureAuthTransactionResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := t.request.Post(constants.ThreeDsSecureAuthTransactionEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ThreeDsSecureAuthTransactionResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (t *Transaction) CardTokenization(body CardTokenizationModel) (*CardTokenizationResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()

	res, err := t.request.Post(constants.CardTokenizationEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[CardTokenizationResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
