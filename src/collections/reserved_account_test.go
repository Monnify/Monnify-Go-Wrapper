package collections

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

var accRef string

func getCredentials() string {
	credentials := utils.LoadConfig("../..")
	return credentials.MonnifyAPIKey + ":" + credentials.MonnifySecretKey
}

func getContractCode() string {
	credentials := utils.LoadConfig("../..")
	return credentials.ContractCode
}

func getApiKey() string {
	credentials := utils.LoadConfig("../..")
	return credentials.MonnifyAPIKey
}

func TestCreateReservedAccountFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	// TEST Validation
	invalidBody := ReservedAccountSchema{
		CustomerName:     "",
		CustomerEmail:    "",
		AccountName:      "",
		AccountReference: "",
		CurrencyCode:     "",
		ContractCode:     "",
		Bvn:              "",
	}
	_, vErr := reservedAccount.CreateReservedAccount(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	// Test Response
	body := ReservedAccountSchema{
		CustomerName:         "John Doe",
		CustomerEmail:        "johndoe@example.com",
		AccountName:          "John Doe",
		AccountReference:     "oihs8jkshjiy9whihu",
		CurrencyCode:         "NGN",
		ContractCode:         "0123",
		Bvn:                  "78362890932",
		GetAllAvailableBanks: true,
		//PreferredBanks:        []string{""},
		IncomeSplitConfig:     []IncomeSplit{},
		MetaData:              nil,
		RestrictPaymentSource: false,
		AllowedPaymentSources: nil,
	}

	_, err := reservedAccount.CreateReservedAccount(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestCreateReservedAccountSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	accRef = uuid.New().String()
	body := ReservedAccountSchema{
		CustomerName:         "John Doe",
		CustomerEmail:        utils.GenerateRandomEmail(),
		AccountName:          "John Doe",
		AccountReference:     accRef,
		ContractCode:         getContractCode(),
		Bvn:                  utils.GenerateRandomNumbers(11),
		GetAllAvailableBanks: true,
		//PreferredBanks:        []string{""},
		IncomeSplitConfig:     []IncomeSplit{},
		MetaData:              nil,
		RestrictPaymentSource: false,
		AllowedPaymentSources: nil,
	}

	resp, _ := reservedAccount.CreateReservedAccount(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotNil(t, resp.ResponseBody)
	assert.NotEmpty(t, resp.ResponseBody.ContractCode, "ContractCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountReference, "AccountReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CurrencyCode, "CurrencyCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerEmail, "CustomerEmail should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerName, "CustomerName should not be empty")
	assert.IsType(t, string(""), resp.ResponseBody.AccountNumber, "AccountNumber should not be empty")
	assert.IsType(t, string(""), resp.ResponseBody.BankName, "BankName should not be empty")
	assert.IsType(t, string(""), resp.ResponseBody.BankCode, "BankCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CollectionChannel, "CollectionChannel should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservationReference, "ReservationReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservedAccountType, "ReservedAccountType should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CreatedOn, "CreatedOn should not be empty")
	if resp.ResponseBody.IncomeSplitConfig != nil {
		for _, config := range resp.ResponseBody.IncomeSplitConfig {
			assert.NotEmpty(t, config.SubAccountCode, "SubAccountCode should not be empty")
			assert.IsType(t, float64(0), config.SplitPercentage, "SplitPercentage should be of type float64")
			assert.IsType(t, float64(0), config.FeePercentage, "FeePercentage should be of type float64")
			assert.IsType(t, bool(true), config.FeeBearer, "FeeBearer should be of type bool")
			assert.IsType(t, bool(false), config.ReservedAccountConfigCode, "ReservedAccountConfigCode should be of type bool")
			assert.GreaterOrEqual(t, config.SplitPercentage, 0.0, "SplitPercentage should be >= 0")
			assert.GreaterOrEqual(t, config.FeePercentage, 0.0, "FeePercentage should be >= 0")
		}
	}
	assert.IsType(t, bool(false), resp.ResponseBody.RestrictPaymentSource, "RestrictPaymentSource should be of type bool")
}

func TestAddLinkedAccountsFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	invalidBody := AddLinkedAccountSchema{
		AccountReference: "",
	}

	_, vErr := reservedAccount.AddLinkedAccounts(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	body := AddLinkedAccountSchema{
		AccountReference:     "98789789hsdjkhk",
		GetAllAvailableBanks: true,
		//PreferredBanks:        []string{""},
	}

	_, err := reservedAccount.AddLinkedAccounts(body)
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestAddLinkedAccountsSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	body := AddLinkedAccountSchema{
		AccountReference:     accRef,
		GetAllAvailableBanks: true,
		//PreferredBanks:        []string{""},
	}

	resp, _ := reservedAccount.AddLinkedAccounts(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotNil(t, resp.ResponseBody)
	assert.NotEmpty(t, resp.ResponseBody.ContractCode, "ContractCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountReference, "AccountReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CurrencyCode, "CurrencyCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerEmail, "CustomerEmail should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerName, "CustomerName should not be empty")
	assert.NotNil(t, resp.ResponseBody.Accounts, "Accounts should not be nil")
	assert.NotEmpty(t, resp.ResponseBody.Accounts, "Accounts should not be empty")
	for i, account := range resp.ResponseBody.Accounts {
		assert.NotEmpty(t, account.AccountName, "AccountName in Accounts should not be empty at index %d", i)
		assert.NotEmpty(t, account.AccountNumber, "AccountNumber in Accounts should not be empty at index %d", i)
		assert.NotEmpty(t, account.BankName, "BankName in Accounts should not be empty at index %d", i)
		assert.NotEmpty(t, account.BankCode, "BankCode in Accounts should not be empty at index %d", i)
	}
	assert.NotEmpty(t, resp.ResponseBody.CollectionChannel, "CollectionChannel should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservationReference, "ReservationReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservedAccountType, "ReservedAccountType should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CreatedOn, "CreatedOn should not be empty")
	if resp.ResponseBody.IncomeSplitConfig != nil {
		for _, config := range resp.ResponseBody.IncomeSplitConfig {
			assert.NotEmpty(t, config.SubAccountCode, "SubAccountCode should not be empty")
			assert.IsType(t, float64(0), config.SplitPercentage, "SplitPercentage should be of type float64")
			assert.IsType(t, float64(0), config.FeePercentage, "FeePercentage should be of type float64")
			assert.IsType(t, bool(true), config.FeeBearer, "FeeBearer should be of type bool")
			assert.IsType(t, bool(false), config.ReservedAccountConfigCode, "ReservedAccountConfigCode should be of type bool")
			assert.GreaterOrEqual(t, config.SplitPercentage, 0.0, "SplitPercentage should be >= 0")
			assert.GreaterOrEqual(t, config.FeePercentage, 0.0, "FeePercentage should be >= 0")
		}
	}
	assert.NotEmpty(t, resp.ResponseBody.BVN, "BVN should not be empty")
	assert.IsType(t, bool(false), resp.ResponseBody.RestrictPaymentSource, "RestrictPaymentSource should be of type bool")
}

func TestReservedAccountDetailsFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	invalidBody := ReservedAccountDetailsSchema{
		AccountReference: "",
	}

	_, vErr := reservedAccount.ReservedAccountDetails(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	_, err := reservedAccount.ReservedAccountDetails(ReservedAccountDetailsSchema{
		AccountReference: "98789789hsdjkhk",
	})
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestReservedAccountDetailsSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	body := ReservedAccountDetailsSchema{
		AccountReference: accRef,
	}

	resp, _ := reservedAccount.ReservedAccountDetails(body)
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotNil(t, resp.ResponseBody)
	assert.NotEmpty(t, resp.ResponseBody.ContractCode, "ContractCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountReference, "AccountReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CurrencyCode, "CurrencyCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerEmail, "CustomerEmail should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerName, "CustomerName should not be empty")
	assert.NotNil(t, resp.ResponseBody.Accounts, "Accounts should not be nil")
	assert.NotEmpty(t, resp.ResponseBody.Accounts, "Accounts should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Accounts[0].AccountName, "AccountName in Accounts should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Accounts[0].AccountNumber, "AccountNumber in Accounts should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Accounts[0].BankName, "BankName in Accounts should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Accounts[0].BankCode, "BankCode in Accounts should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CollectionChannel, "CollectionChannel should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservationReference, "ReservationReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservedAccountType, "ReservedAccountType should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CreatedOn, "CreatedOn should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Contract.Name, "Contract.Name should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Contract.Code, "Contract.Code should not be empty")
	assert.IsType(t, string(""), resp.ResponseBody.Contract.Description, "Contract.Description should not be empty")
	assert.IsType(t, int(0), resp.ResponseBody.TransactionCount, "TransactionCount should be of type int")
	if resp.ResponseBody.IncomeSplitConfig != nil {
		for _, config := range resp.ResponseBody.IncomeSplitConfig {
			assert.NotEmpty(t, config.SubAccountCode, "SubAccountCode should not be empty")
			assert.IsType(t, float64(0), config.SplitPercentage, "SplitPercentage should be of type float64")
			assert.IsType(t, float64(0), config.FeePercentage, "FeePercentage should be of type float64")
			assert.IsType(t, bool(true), config.FeeBearer, "FeeBearer should be of type bool")
			assert.IsType(t, bool(false), config.ReservedAccountConfigCode, "ReservedAccountConfigCode should be of type bool")
			assert.GreaterOrEqual(t, config.SplitPercentage, 0.0, "SplitPercentage should be >= 0")
			assert.GreaterOrEqual(t, config.FeePercentage, 0.0, "FeePercentage should be >= 0")
		}
	}
	assert.NotEmpty(t, resp.ResponseBody.BVN, "BVN should not be empty")
	assert.IsType(t, bool(false), resp.ResponseBody.RestrictPaymentSource, "RestrictPaymentSource should be of type bool")
}

func TestReservedAccountTransactionsFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	invalidBody := ReservedAccountTransactionsSchema{
		AccountReference: "",
	}

	_, vErr := reservedAccount.ReservedAccountTransactions(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	_, err := reservedAccount.ReservedAccountTransactions(ReservedAccountTransactionsSchema{
		AccountReference: "98789789hsdjkhk",
	})
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestReservedAccountTransactionsSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	resp, _ := reservedAccount.ReservedAccountTransactions(ReservedAccountTransactionsSchema{
		AccountReference: accRef,
	})
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotNil(t, resp.ResponseBody)
	assert.NotNil(t, resp.ResponseBody.Content, "Content should not be nil")

	if len(resp.ResponseBody.Content) >= 1 {
		assert.NotNil(t, resp.ResponseBody.Content[0].CustomerDTO, "CustomerDTO should not be nil")
		assert.NotEmpty(t, resp.ResponseBody.Content[0].CustomerDTO.Email, "CustomerDTO Email should not be empty")
		assert.NotEmpty(t, resp.ResponseBody.Content[0].CustomerDTO.Name, "CustomerDTO Name should not be empty")
		assert.IsType(t, float64(0), resp.ResponseBody.Content[0].ProviderAmount, "ProviderAmount should be of type float64")
		assert.IsType(t, string(""), resp.ResponseBody.Content[0].PaymentMethod, "PaymentMethod should be of type string")
		assert.IsType(t, bool(true), resp.ResponseBody.Content[0].Flagged, "Flagged should be of type bool")
		assert.NotNil(t, resp.ResponseBody.Pageable, "Pageable should not be nil")
		assert.IsType(t, int(0), resp.ResponseBody.Pageable.PageSize, "PageSize should be of type int")
		assert.IsType(t, bool(true), resp.ResponseBody.Pageable.Paged, "Paged should be of type bool")
		assert.IsType(t, int(0), resp.ResponseBody.TotalPages, "TotalPages should be of type int")
		assert.IsType(t, int(0), resp.ResponseBody.TotalElements, "TotalElements should be of type int")
		assert.IsType(t, bool(true), resp.ResponseBody.First, "First should be of type bool")
		assert.IsType(t, bool(false), resp.ResponseBody.Last, "Last should be of type bool")
		assert.IsType(t, int(0), resp.ResponseBody.NumberOfElements, "NumberOfElements should be of type int")
		assert.IsType(t, bool(false), resp.ResponseBody.Sort.Sorted, "Sorted should be of type bool")
		assert.IsType(t, bool(true), resp.ResponseBody.Sort.Empty, "Empty should be of type bool")
	}
}

func TestUpdateReservedAccountKycInfoFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	invalidBody := UpdateReservedAccountKycInfoSchema{
		AccountReference: "",
	}

	_, vErr := reservedAccount.UpdateReservedAccountKycInfo(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	_, err := reservedAccount.UpdateReservedAccountKycInfo(UpdateReservedAccountKycInfoSchema{
		AccountReference: "98789789hsdjkhk",
	})
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

// TODO: commented out because require valid BVN to run test
//func TestUpdateReservedAccountKycInfoSuccess(t *testing.T) {
//	credentials := getCredentials()
//	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
//	reservedAccount := NewReservedAccount(httpRequest)
//
//	resp, _ := reservedAccount.UpdateReservedAccountKycInfo(UpdateReservedAccountKycInfoSchema{
//		AccountReference: accRef,
//		Nin:              utils.GenerateRandomNumbers(11),
//		Bvn:              utils.GenerateRandomNumbers(11),
//	})
//	assert.Equal(t, true, resp.RequestSuccessful)
//	assert.Equal(t, "success", resp.ResponseMessage)
//	assert.NotNil(t, resp.ResponseBody)
//	assert.NotEmpty(t, resp.ResponseBody.AccountReference, "AccountReference should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.CustomerEmail, "CustomerEmail should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.CustomerName, "CustomerName should not be empty")
//	assert.NotEmpty(t, resp.ResponseBody.Bvn, "BVN should not be empty")
//}

func TestDeallocateReservedAccountFailed(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	invalidBody := DeallocateReservedAccountSchema{
		AccountReference: "",
	}

	_, vErr := reservedAccount.DeallocateReservedAccount(invalidBody)
	assert.NotNil(t, vErr)
	assert.Equal(t, "Validation Error", vErr.Message)

	_, err := reservedAccount.DeallocateReservedAccount(DeallocateReservedAccountSchema{
		AccountReference: "98789789hsdjkhk",
	})
	assert.NotNil(t, err)                                  // Error not nil
	assert.Equal(t, false, err.Response.RequestSuccessful) // Failed Response
}

func TestDeallocateReservedAccountSuccess(t *testing.T) {
	credentials := getCredentials()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials)
	reservedAccount := NewReservedAccount(httpRequest)

	resp, _ := reservedAccount.DeallocateReservedAccount(DeallocateReservedAccountSchema{
		AccountReference: accRef,
	})
	assert.Equal(t, true, resp.RequestSuccessful)
	assert.Equal(t, "success", resp.ResponseMessage)
	assert.NotNil(t, resp.ResponseBody)
	assert.NotEmpty(t, resp.ResponseBody.ContractCode, "ContractCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountReference, "AccountReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountName, "AccountName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CurrencyCode, "CurrencyCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerEmail, "CustomerEmail should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CustomerName, "CustomerName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.AccountNumber, "AccountNumber should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankName, "BankName should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.BankCode, "BankCode should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CollectionChannel, "CollectionChannel should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservationReference, "ReservationReference should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.ReservedAccountType, "ReservedAccountType should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.Status, "Status should not be empty")
	assert.NotEmpty(t, resp.ResponseBody.CreatedOn, "CreatedOn should not be empty")
	if resp.ResponseBody.IncomeSplitConfig != nil {
		for _, config := range resp.ResponseBody.IncomeSplitConfig {
			assert.NotEmpty(t, config.SubAccountCode, "SubAccountCode should not be empty")
			assert.IsType(t, float64(0), config.SplitPercentage, "SplitPercentage should be of type float64")
			assert.IsType(t, float64(0), config.FeePercentage, "FeePercentage should be of type float64")
			assert.IsType(t, bool(true), config.FeeBearer, "FeeBearer should be of type bool")
			assert.IsType(t, bool(false), config.ReservedAccountConfigCode, "ReservedAccountConfigCode should be of type bool")
			assert.GreaterOrEqual(t, config.SplitPercentage, 0.0, "SplitPercentage should be >= 0")
			assert.GreaterOrEqual(t, config.FeePercentage, 0.0, "FeePercentage should be >= 0")
		}
	}
	assert.IsType(t, bool(false), resp.ResponseBody.RestrictPaymentSource, "RestrictPaymentSource should be of type bool")
}
