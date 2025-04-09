package disbursement

import (
	"fmt"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	mErr "github.com/Monnify/Monnify-Go-Wrapper/src/common/error"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type Refund struct {
	request *request.HttpRequest
}

func NewRefund(request *request.HttpRequest) *Refund {
	return &Refund{request}
}

func (r *Refund) InitiateRefund(body InitiateRefundModel) (*InitiateRefundResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()

	res, err := r.request.Post(constants.InitiateRefundEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[InitiateRefundResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *Refund) GetAllRefunds(body GetAllRefundsModel) (*GetAllRefundsResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.GetAllRefundsEndpoint, body.Page, body.Size)
	res, err := r.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetAllRefundsResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (r *Refund) GetRefundStatus(body GetRefundStatusModel) (*GetRefundStatusResponse, *mErr.Error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(constants.GetRefundStatusEndpoint, body.RefundReference)
	res, err := r.request.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetRefundStatusResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
