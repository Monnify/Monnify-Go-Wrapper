package collections

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Monnify/Monnify-Go-Wrapper/src/common/request"
	"github.com/Monnify/Monnify-Go-Wrapper/src/common/utils"
)

func TestCreateReservedAccountFailed(t *testing.T) {
	credentials := utils.LoadConfig()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials.MonnifyAPIKey+":"+credentials.MonnifySecretKey)
	reservedAccount := NewReservedAccount(httpRequest)

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
	credentials := utils.LoadConfig()
	httpRequest := request.NewHttpRequest(utils.GetBaseUrl(false), credentials.MonnifyAPIKey+":"+credentials.MonnifySecretKey)
	reservedAccount := NewReservedAccount(httpRequest)

	body := ReservedAccountSchema{
		CustomerName:         "John Doe",
		CustomerEmail:        utils.GenerateRandomEmail(),
		AccountName:          "John Doe",
		AccountReference:     uuid.New().String(),
		CurrencyCode:         "NGN",
		ContractCode:         "0165673622",
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
}
