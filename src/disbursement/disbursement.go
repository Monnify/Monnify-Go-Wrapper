package disbursement

import (
	"fmt"
	"net/url"

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

	body.SetDefault()

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

func (d *Disbursement) InitiateBulkTransfer(body BulkTransfer) (*BulkTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	body.SetDefault()

	res, err := d.request.Post(constants.BulkTransferEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[BulkTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) AuthorizeBulkTransfer(body AuthorizeTransfer) (*AuthorizeBulkTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := d.request.Post(constants.AuthorizeBulkTransferEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[AuthorizeBulkTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) AuthorizeSingleTransfer(body AuthorizeTransfer) (*AuthorizeSingleTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := d.request.Post(constants.AuthorizeSingleTransferEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[AuthorizeSingleTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) ResendTransferOTP(body ResendTransferOTP) (*ResendTransferOTPResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	res, err := d.request.Post(constants.ResendTransferOTPEndpoint, body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[ResendTransferOTPResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) GetSingleTransferStatus(body GetStatus) (*GetSingleTransferStatusResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	encodedReference := url.QueryEscape(body.Reference)
	newUrl := fmt.Sprintf(constants.GetSingleTransferStatusEndpoint, encodedReference)
	res, err := d.request.Get(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetSingleTransferStatusResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) GetBulkTransferStatus(body GetBulkStatus) (*GetBulkTransferStatusResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	newUrl := fmt.Sprintf(constants.GetBulkTransferStatusEndpoint, body.Reference, body.PageSize, body.PageNo)
	res, err := d.request.Get(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetBulkTransferStatusResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) GetAllSingleTransfer(body GetAllSingleTransfer) (*GetAllSingleTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	newUrl := fmt.Sprintf(constants.AllSingleTransferEndpoint, body.PageSize, body.PageNo)
	res, err := d.request.Get(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetAllSingleTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Disbursement) GetAllBulkTransfer(body GetAllBulkTransfer) (*GetAllBulkTransferResponse, error) {
	if err := utils.ValidateStruct(body); err != nil {
		return nil, err
	}

	newUrl := fmt.Sprintf(constants.AllBulkTransferEndpoint, body.PageSize, body.PageNo)
	res, err := d.request.Get(newUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := utils.ParseResponse[GetAllBulkTransferResponse](res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
