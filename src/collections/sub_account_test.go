package collections

import (
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var subAccountCode string

func TestCreateSubAccountFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	// TEST Validation
	invalidBody := []CreateSubAccountModel{
		{
			CurrencyCode:           "",
			AccountNumber:          "",
			BankCode:               "",
			Email:                  "",
			DefaultSplitPercentage: 20.87,
		},
	}
	_, vErr := subAccount.CreateSubAccount(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := []CreateSubAccountModel{
		{
			CurrencyCode:           "NGN",
			AccountNumber:          "1853790782",
			BankCode:               "097",
			Email:                  utils.GenerateRandomEmail(),
			DefaultSplitPercentage: 20.87,
		},
	}

	_, err := subAccount.CreateSubAccount(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestCreateSubAccountSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	body := []CreateSubAccountModel{
		{
			CurrencyCode:           "NGN",
			AccountNumber:          "0211319282",
			BankCode:               "058",
			Email:                  utils.GenerateRandomEmail(),
			DefaultSplitPercentage: 20.87,
		},
	}

	resp, _ := subAccount.CreateSubAccount(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	for _, res := range resp.ResponseBody {
		subAccountCode = res.SubAccountCode
		assert.NotEmpty(t, res.SubAccountCode, "SubAccountCode should not be empty")
		assert.NotEmpty(t, res.AccountName, "AccountName should not be empty")
		assert.NotEmpty(t, res.CurrencyCode, "CurrencyCode should not be empty")
		assert.NotEmpty(t, res.AccountNumber, "AccountNumber should not be empty")
		assert.NotEmpty(t, res.BankName, "BankName should not be empty")
		assert.NotEmpty(t, res.BankCode, "BankCode should not be empty")
		assert.NotEmpty(t, res.Email, "Email should not be empty")
		assert.NotZero(t, res.DefaultSplitPercentage, "DefaultSplitPercentage should not be zero")
		assert.NotEmpty(t, res.SettlementProfileCode, "SettlementProfileCode should not be empty")
		if len(res.SettlementReportEmails) > 0 {
			assert.NotEmpty(t, res.SettlementReportEmails, "SettlementReportEmails should not be empty when provided")
		}
	}
}

func TestGetSubAccounts(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	resp, _ := subAccount.GetSubAccounts()
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	for _, res := range resp.ResponseBody {
		assert.NotEmpty(t, res.SubAccountCode, "SubAccountCode should not be empty")
		assert.NotEmpty(t, res.AccountName, "AccountName should not be empty")
		assert.NotEmpty(t, res.CurrencyCode, "CurrencyCode should not be empty")
		assert.NotEmpty(t, res.AccountNumber, "AccountNumber should not be empty")
		assert.NotEmpty(t, res.BankName, "BankName should not be empty")
		assert.NotEmpty(t, res.BankCode, "BankCode should not be empty")
		assert.NotEmpty(t, res.Email, "Email should not be empty")
		assert.NotZero(t, res.DefaultSplitPercentage, "DefaultSplitPercentage should not be zero")
		assert.NotEmpty(t, res.SettlementProfileCode, "SettlementProfileCode should not be empty")
		if len(res.SettlementReportEmails) > 0 {
			assert.NotEmpty(t, res.SettlementReportEmails, "SettlementReportEmails should not be empty when provided")
		}
	}
}

func TestUpdateSubAccountFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	// TEST Validation
	invalidBody := UpdateSubAccountModel{
		CurrencyCode:           "",
		AccountNumber:          "",
		BankCode:               "",
		Email:                  "",
		DefaultSplitPercentage: 20.87,
		SubAccountCode:         "",
	}
	_, vErr := subAccount.UpdateSubAccount(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := UpdateSubAccountModel{
		CurrencyCode:           "NGN",
		AccountNumber:          "1853790782",
		BankCode:               "097",
		Email:                  utils.GenerateRandomEmail(),
		DefaultSplitPercentage: 20.87,
		SubAccountCode:         "MFY_SUB_426724844716",
	}

	_, err := subAccount.UpdateSubAccount(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestUpdateSubAccountSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	body := UpdateSubAccountModel{
		CurrencyCode:           "NGN",
		AccountNumber:          "0211319282",
		BankCode:               "058",
		Email:                  utils.GenerateRandomEmail(),
		DefaultSplitPercentage: 20.87,
		SubAccountCode:         subAccountCode,
	}
	resp, _ := subAccount.UpdateSubAccount(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.SubAccountCode, "SubAccountCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CurrencyCode, "CurrencyCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountNumber, "AccountNumber should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankName, "BankName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankCode, "BankCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Email, "Email should not be empty")
	assert.NotZero(t, resp.ResponseBody.DefaultSplitPercentage, "DefaultSplitPercentage should not be zero")
	assert.NotEmpty(t, resp.ResponseBody.SettlementProfileCode, "SettlementProfileCode should not be empty")
	if len(resp.ResponseBody.SettlementReportEmails) > 0 {
		assert.NotEmpty(t, resp.ResponseBody.SettlementReportEmails, "SettlementReportEmails should not be empty when provided")
	}
}

func TestDeleteSubAccountFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	// TEST Validation
	invalidBody := DeleteSubAccountModel{
		SubAccountCode: "",
	}
	_, vErr := subAccount.DeleteSubAccount(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := DeleteSubAccountModel{
		SubAccountCode: "MFY_SUB_426724533",
	}

	_, err := subAccount.DeleteSubAccount(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestDeleteSubAccountSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	subAccount := NewSubAccount(httpRequest)

	body := DeleteSubAccountModel{
		SubAccountCode: subAccountCode,
	}

	resp, _ := subAccount.DeleteSubAccount(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
}
