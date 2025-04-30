package collections

import (
	"fmt"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var transactionReference string
var paymentReference string
var otpTokenId string
var otpToken = "123456"

func TestInitializeTransactionFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	// TEST Validation
	invalidBody := InitializeTransactionModel{
		CustomerEmail:      "",
		CustomerName:       "",
		Amount:             20.87,
		PaymentReference:   "",
		PaymentDescription: "",
		ContractCode:       "",
	}

	_, vErr := transaction.InitializeTransaction(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := InitializeTransactionModel{
		CustomerEmail:      utils.GenerateRandomEmail(),
		CustomerName:       "John Doe",
		Amount:             20.87,
		PaymentReference:   "reference",
		PaymentDescription: "description",
		ContractCode:       "9289784",
	}

	_, err := transaction.InitializeTransaction(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestInitializeTransactionSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := InitializeTransactionModel{
		CustomerEmail:      utils.GenerateRandomEmail(),
		CustomerName:       "John Doe",
		Amount:             20.87,
		PaymentReference:   uuid.New().String(),
		PaymentDescription: "description",
		ContractCode:       getContractCode(),
	}

	resp, _ := transaction.InitializeTransaction(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	transactionReference = resp.ResponseBody.TransactionReference
	paymentReference = resp.ResponseBody.PaymentReference
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.MerchantName, "MerchantName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ApiKey, "ApiKey should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.EnabledPaymentMethod, "EnabledPaymentMethod should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CheckoutUrl, "CheckoutUrl should not be empty")
	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.MerchantName, "MerchantName should be a string")
	assert.IsType(t, "", resp.ResponseBody.ApiKey, "ApiKey should be a string")
	assert.IsType(t, []string{}, resp.ResponseBody.EnabledPaymentMethod, "EnabledPaymentMethod should be a slice of strings")
	assert.IsType(t, "", resp.ResponseBody.CheckoutUrl, "CheckoutUrl should be a string")
}

func TestGetTransactionStatusv2Failed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	invalidBody := GetTransactionStatusv2Model{
		TransactionReference: "",
	}
	_, vErr := transaction.GetTransactionStatusv2(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := GetTransactionStatusv2Model{
		TransactionReference: "oisuiosuiuijniojij",
	}
	_, err := transaction.GetTransactionStatusv2(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestGetTransactionStatusv2Success(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := GetTransactionStatusv2Model{
		TransactionReference: transactionReference,
	}
	resp, _ := transaction.GetTransactionStatusv2(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AmountPaid, "AmountPaid should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TotalPayable, "TotalPayable should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.SettlementAmount, "SettlementAmount should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.PaidOn, "PaidOn should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentStatus, "PaymentStatus should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentDescription, "PaymentDescription should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Currency, "Currency should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.PaymentMethod, "PaymentMethod should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.Product.Type, "Product Type should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.Product.Reference, "Product Reference should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.CardType, "CardType should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.Last4, "Last4 should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.ExpMonth, "ExpMonth should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.ExpYear, "ExpYear should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.Bin, "Bin should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.BankName, "BankName should not be empty")
	assert.IsType(t, true, resp.ResponseBody.CardDetails.Reusable, "Reusable should be a bool")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.CountryCode, "CountryCode should not be empty")
	assert.IsType(t, nil, resp.ResponseBody.CardDetails.BankCode, "BankCode should be nil")
	assert.IsType(t, nil, resp.ResponseBody.CardDetails.CardToken, "CardToken should be nil")
	assert.IsType(t, "", resp.ResponseBody.CardDetails.MaskedPan, "MaskedPan should be a string")
	assert.NotEmpty(t, resp.ResponseBody.Customer.Email, "Customer Email should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Customer.Name, "Customer Name should not be empty")

}

func TestGetTransactionStatusv1Failed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	invalidBody := GetTransactionStatusv1Model{
		PaymentReference: "",
	}
	_, vErr := transaction.GetTransactionStatusv1(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := GetTransactionStatusv1Model{
		PaymentReference: "oisuiosuiuijniojij",
	}
	_, err := transaction.GetTransactionStatusv1(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestGetTransactionStatusv1Success(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := GetTransactionStatusv1Model{
		PaymentReference: paymentReference,
	}
	resp, _ := transaction.GetTransactionStatusv1(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AmountPaid, "AmountPaid should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TotalPayable, "TotalPayable should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.SettlementAmount, "SettlementAmount should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.PaidOn, "PaidOn should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentStatus, "PaymentStatus should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentDescription, "PaymentDescription should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Currency, "Currency should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.PaymentMethod, "PaymentMethod should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.Product.Type, "Product Type should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.Product.Reference, "Product Reference should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.CardType, "CardType should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.Last4, "Last4 should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.ExpMonth, "ExpMonth should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.ExpYear, "ExpYear should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.Bin, "Bin should not be empty")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.BankName, "BankName should not be empty")
	assert.IsType(t, true, resp.ResponseBody.CardDetails.Reusable, "Reusable should be a bool")
	//assert.NotEmpty(t, resp.ResponseBody.CardDetails.CountryCode, "CountryCode should not be empty")
	assert.IsType(t, nil, resp.ResponseBody.CardDetails.BankCode, "BankCode should be nil")
	assert.IsType(t, nil, resp.ResponseBody.CardDetails.CardToken, "CardToken should be nil")
	assert.IsType(t, "", resp.ResponseBody.CardDetails.MaskedPan, "MaskedPan should be a string")
	assert.NotEmpty(t, resp.ResponseBody.Customer.Email, "Customer Email should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Customer.Name, "Customer Name should not be empty")

}

func TestPayWithBankTransferFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	invalidBody := PayWithBankTransferModel{
		TransactionReference: "",
		BankCode:             "",
	}
	_, vErr := transaction.PayWithBankTransfer(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := PayWithBankTransferModel{
		TransactionReference: "oisuiosuiuijniojij",
		BankCode:             "098",
	}
	_, err := transaction.PayWithBankTransfer(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestPayWithBankTransferSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := PayWithBankTransferModel{
		TransactionReference: transactionReference,
		BankCode:             "058",
	}
	resp, _ := transaction.PayWithBankTransfer(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.AccountNumber, "AccountNumber should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankName, "BankName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankCode, "BankCode should not be empty")
	assert.NotZero(t, resp.ResponseBody.AccountDurationSeconds, "AccountDurationSeconds should not be zero")
	assert.NotEmpty(t, resp.ResponseBody.UssdPayment, "UssdPayment should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.RequestTime, "RequestTime should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ExpiresOn, "ExpiresOn should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	assert.NotZero(t, resp.ResponseBody.Amount, "Amount should not be zero")
	//assert.NotZero(t, resp.ResponseBody.Fee, "Fee should not be zero")
	assert.NotZero(t, resp.ResponseBody.TotalPayable, "TotalPayable should not be zero")
	assert.NotEmpty(t, resp.ResponseBody.CollectionChannel, "CollectionChannel should not be empty")
	assert.Nil(t, resp.ResponseBody.ProductInformation, "ProductInformation should be nil")
	assert.IsType(t, "", resp.ResponseBody.AccountNumber, "AccountNumber should be a string")
	assert.IsType(t, "", resp.ResponseBody.AccountName, "AccountName should be a string")
	assert.IsType(t, "", resp.ResponseBody.BankName, "BankName should be a string")
	assert.IsType(t, "", resp.ResponseBody.BankCode, "BankCode should be a string")
	assert.IsType(t, 0, resp.ResponseBody.AccountDurationSeconds, "AccountDurationSeconds should be an int")
	assert.IsType(t, "", resp.ResponseBody.UssdPayment, "UssdPayment should be a string")
	assert.IsType(t, "", resp.ResponseBody.RequestTime, "RequestTime should be a string")
	assert.IsType(t, "", resp.ResponseBody.ExpiresOn, "ExpiresOn should be a string")
	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
	assert.IsType(t, 0.00, resp.ResponseBody.Amount, "Amount should be an int")
	assert.IsType(t, 0.00, resp.ResponseBody.Fee, "Fee should be an int")
	assert.IsType(t, 0.00, resp.ResponseBody.TotalPayable, "TotalPayable should be an int")
	assert.IsType(t, "", resp.ResponseBody.CollectionChannel, "CollectionChannel should be a string")
	assert.IsType(t, nil, resp.ResponseBody.ProductInformation, "ProductInformation should be nil")
}

func TestChargeCardFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	invalidBody := ChargeCardModel{
		TransactionReference: "",
		CollectionChannel:    "",
		Card: ChargeCard{
			Number:      "",
			ExpiryMonth: "",
			ExpiryYear:  "",
			CVV:         "",
			PIN:         "",
		},
		DeviceInformation: DeviceInformation{
			HttpBrowserLanguage:          "",
			HttpBrowserJavaEnabled:       false,
			HttpBrowserJavaScriptEnabled: false,
			HttpBrowserColorDepth:        0,
			HttpBrowserScreenHeight:      0,
			HttpBrowserScreenWidth:       0,
			HttpBrowserTimeDifference:    "",
			UserAgentBrowserValue:        "",
		},
	}
	_, vErr := transaction.ChargeCard(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := ChargeCardModel{
		TransactionReference: "kjskjsoijoisjois",
		CollectionChannel:    "API_NOTIFICATION",
		Card: ChargeCard{
			Number:      "4111111111111111",
			ExpiryMonth: "12",
			ExpiryYear:  "2025",
			CVV:         "123",
			PIN:         "1234",
		},
		DeviceInformation: DeviceInformation{
			HttpBrowserLanguage:          "en-US",
			HttpBrowserJavaEnabled:       false,
			HttpBrowserJavaScriptEnabled: false,
			HttpBrowserColorDepth:        24,
			HttpBrowserScreenHeight:      1203,
			HttpBrowserScreenWidth:       2138,
			HttpBrowserTimeDifference:    "24",
			UserAgentBrowserValue:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	}
	_, err := transaction.ChargeCard(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestChargeCardWithoutOTPSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := ChargeCardModel{
		TransactionReference: transactionReference,
		CollectionChannel:    "API_NOTIFICATION",
		Card: ChargeCard{
			Number:      "4111111111111111",
			ExpiryMonth: "12",
			ExpiryYear:  "2025",
			CVV:         "123",
			PIN:         "1234",
		},
		DeviceInformation: DeviceInformation{
			HttpBrowserLanguage:          "en-US",
			HttpBrowserJavaEnabled:       false,
			HttpBrowserJavaScriptEnabled: false,
			HttpBrowserColorDepth:        24,
			HttpBrowserScreenHeight:      1203,
			HttpBrowserScreenWidth:       2138,
			HttpBrowserTimeDifference:    "24",
			UserAgentBrowserValue:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	}
	resp, _ := transaction.ChargeCard(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Message, "Message should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	//assert.Greater(t, resp.ResponseBody.AuthorizedAmount, 0, "AuthorizedAmount should be greater than 0")
	assert.IsType(t, "", resp.ResponseBody.Status, "Status should be a string")
	assert.IsType(t, "", resp.ResponseBody.Message, "Message should be a string")
	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
	assert.IsType(t, 0.00, resp.ResponseBody.AuthorizedAmount, "AuthorizedAmount should be an float")

}

func TestChargeCardWithOTPSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := ChargeCardModel{
		TransactionReference: "MNFY|05|20250409193710|002465",
		CollectionChannel:    "API_NOTIFICATION",
		Card: ChargeCard{
			Number:      "5060995994247093",
			ExpiryMonth: "12",
			ExpiryYear:  "2025",
			CVV:         "123",
			PIN:         "1234",
		},
		DeviceInformation: DeviceInformation{
			HttpBrowserLanguage:          "en-US",
			HttpBrowserJavaEnabled:       false,
			HttpBrowserJavaScriptEnabled: false,
			HttpBrowserColorDepth:        24,
			HttpBrowserScreenHeight:      1203,
			HttpBrowserScreenWidth:       2138,
			HttpBrowserTimeDifference:    "24",
			UserAgentBrowserValue:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	}
	resp, _ := transaction.ChargeCard(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Message, "Message should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.OtpData.Id, "OTP ID should not be empty")
	otpTokenId = resp.ResponseBody.OtpData.Id
	fmt.Println("OTP TOKEN", otpTokenId)
	assert.NotEmpty(t, resp.ResponseBody.OtpData.Message, "OTP Message should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.OtpData.AuthData, "OTP AuthData should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")

	assert.IsType(t, "", resp.ResponseBody.Status, "Status should be a string")
	assert.IsType(t, "", resp.ResponseBody.Message, "Message should be a string")
	assert.IsType(t, "", resp.ResponseBody.OtpData.Id, "OTP ID should be a string")
	assert.IsType(t, "", resp.ResponseBody.OtpData.Message, "OTP Message should be a string")
	assert.IsType(t, "", resp.ResponseBody.OtpData.AuthData, "OTP AuthData should be a string")
	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
}

//func TestAuthorizeOTPSuccess(t *testing.T) {
//	credentials := getCredentials()
//	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
//	transaction := NewTransaction(httpRequest)
//
//	body := AuthorizeOTPModel{
//		TransactionReference: "MNFY|05|20250409193710|002465",
//		CollectionChannel:    "API_NOTIFICATION",
//		TokenId:              "20.87-b66bef0aa8e660863c4e1177a08fefba",
//		Token:                otpToken,
//	}
//
//	resp, err := transaction.AuthorizeOTP(body)
//	fmt.Println(err.Response)
//	assert.Equal(t, true, resp.RequestSuccessful)
//	assert.Equal(t, "success", resp.ResponseMessage)
//	assert.NotEmpty(t, resp.ResponseBody.PaymentStatus, "PaymentStatus should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.PaymentDescription, "PaymentDescription should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
//	assert.IsType(t, "", resp.ResponseBody.PaymentStatus, "PaymentStatus should be a string")
//	assert.IsType(t, "", resp.ResponseBody.PaymentDescription, "PaymentDescription should be a string")
//	assert.IsType(t, 0.00, resp.ResponseBody.AmountPaid, "AmountPaid should be a string")
//	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
//	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
//	assert.IsType(t, "", resp.ResponseBody.CurrencyPaid, "CurrencyPaid should be a string")
//}

//func TestThreeDsSecureAuthTransactionSuccess(t *testing.T) {
//	credentials := getCredentials()
//	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
//	transaction := NewTransaction(httpRequest)
//
//	body := ThreeDsSecureAuthTransactionModel{
//		TransactionReference: "MNFY|97|20250409221804|002133",
//		CollectionChannel:    "API_NOTIFICATION",
//		ApiKey:               getApiKey(),
//		Card: ThreeDsSecureAuthTransactionCard{
//			Number:      "4000000000000",
//			ExpiryMonth: "12",
//			ExpiryYear:  "2025",
//			CVV:         "123",
//		},
//	}
//
//	resp, err := transaction.ThreeDsSecureAuthTransaction(body)
//	fmt.Println(err.Response)
//	assert.Equal(t, true, resp.RequestSuccessful)
//	assert.Equal(t, "success", resp.ResponseMessage)
//	assert.NotEmpty(t, resp.ResponseBody.PaymentStatus, "PaymentStatus should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.PaymentDescription, "PaymentDescription should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
//	assert.IsType(t, "", resp.ResponseBody.PaymentStatus, "PaymentStatus should be a string")
//	assert.IsType(t, "", resp.ResponseBody.PaymentDescription, "PaymentDescription should be a string")
//	assert.IsType(t, 0.00, resp.ResponseBody.AmountPaid, "AmountPaid should be a string")
//	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
//	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
//	assert.IsType(t, "", resp.ResponseBody.CurrencyPaid, "CurrencyPaid should be a string")
//}

func TestCardTokenizationSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	transaction := NewTransaction(httpRequest)

	body := CardTokenizationModel{
		CustomerName:       "John Doe",
		CustomerEmail:      utils.GenerateRandomEmail(),
		Amount:             25.7,
		PaymentDescription: "payment description",
		PaymentReference:   paymentReference,
		CurrencyCode:       "NGN",
		ContractCode:       getContractCode(),
		CardToken:          "",
		ApiKey:             getApiKey(),
	}

	resp, err := transaction.CardTokenization(body)
	fmt.Println(err.Response)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotEmpty(t, resp.ResponseBody.PaymentStatus, "PaymentStatus should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentDescription, "PaymentDescription should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.TransactionReference, "TransactionReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.PaymentReference, "PaymentReference should not be empty")
	assert.IsType(t, "", resp.ResponseBody.PaymentStatus, "PaymentStatus should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentDescription, "PaymentDescription should be a string")
	assert.IsType(t, 0.00, resp.ResponseBody.AmountPaid, "AmountPaid should be a string")
	assert.IsType(t, "", resp.ResponseBody.TransactionReference, "TransactionReference should be a string")
	assert.IsType(t, "", resp.ResponseBody.PaymentReference, "PaymentReference should be a string")
	//assert.IsType(t, "", resp.ResponseBody.CurrencyPaid, "CurrencyPaid should be a string")
}
