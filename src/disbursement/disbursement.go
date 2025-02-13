package disbursement

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/constants"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

type Disbursement struct {
	request *request.HttpRequest
}

func NewDisbursement(request *request.HttpRequest) *Disbursement {
	return &Disbursement{request}
}

func (d *Disbursement) InitiateSingleTransfer(body SingleTransfer) (*SingleTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := d.request.Post(constants.DisbursementSingleEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[SingleTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
