package collections

type IncomeSplitConfig struct {
	SubAccountCode  string  `json:"subAccountCode" validate:"required"`
	SplitAmount     float64 `json:"splitAmount" validate:"omitempty,min=1"`
	SplitPercentage float64 `json:"splitPercentage" validate:"omitempty,min=1"`
	FeePercentage   float64 `json:"feePercentage" validate:"omitempty,min=0"`
	FeeBearer       bool    `json:"feeBearer" validate:"omitempty"`
}

type InitializeTransactionModel struct {
	CustomerEmail      string              `json:"customerEmail" validate:"required,email"`
	CustomerName       string              `json:"customerName" validate:"required,min=3"`
	Amount             float64             `validate:"required,number,min=20" json:"amount"`
	PaymentDescription string              `json:"paymentDescription" validate:"required,min=3"`
	PaymentReference   string              `json:"paymentReference" validate:"required,min=3"`
	CurrencyCode       string              `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	ContractCode       string              `json:"contractCode" validate:"required"`
	RedirectUrl        string              `json:"redirectUrl" validate:"omitempty,uri"`
	PaymentMethods     []string            `json:"paymentMethods" validate:"omitempty"`
	IncomeSplitConfig  []IncomeSplitConfig `json:"incomeSplitConfig" validate:"omitempty"`
	MetaData           interface{}         `json:"metaData" validate:"omitempty"`
}

type InitializeTransactionResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TransactionReference string   `json:"transactionReference"`
		PaymentReference     string   `json:"paymentReference"`
		MerchantName         string   `json:"merchantName"`
		ApiKey               string   `json:"apiKey"`
		EnabledPaymentMethod []string `json:"enabledPaymentMethod"`
		CheckoutUrl          string   `json:"checkoutUrl"`
	} `json:"responseBody"`
}
