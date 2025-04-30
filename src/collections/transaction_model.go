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
	IncomeSplitConfig  []IncomeSplitConfig `json:"incomeSplitConfig,omitempty" validate:"omitempty"`
	MetaData           interface{}         `json:"metaData,omitempty" validate:"omitempty"`
}

func (s *InitializeTransactionModel) SetDefault() {
	if s.CurrencyCode == "" {
		s.CurrencyCode = "NGN"
	}
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

type GetTransactionStatusv2Model struct {
	TransactionReference string `json:"transactionReference" validate:"required"`
}

type GetTransactionStatusv2Response struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TransactionReference string `json:"transactionReference"`
		PaymentReference     string `json:"paymentReference"`
		AmountPaid           string `json:"amountPaid"`
		TotalPayable         string `json:"totalPayable"`
		SettlementAmount     string `json:"settlementAmount"`
		PaidOn               string `json:"paidOn"`
		PaymentStatus        string `json:"paymentStatus"`
		PaymentDescription   string `json:"paymentDescription"`
		Currency             string `json:"currency"`
		PaymentMethod        string `json:"paymentMethod"`
		Product              struct {
			Type      string `json:"type"`
			Reference string `json:"reference"`
		} `json:"product"`
		CardDetails struct {
			CardType             string      `json:"cardType"`
			Last4                string      `json:"last4"`
			ExpMonth             string      `json:"expMonth"`
			ExpYear              string      `json:"expYear"`
			Bin                  string      `json:"bin"`
			BankCode             interface{} `json:"bankCode"`
			BankName             string      `json:"bankName"`
			Reusable             bool        `json:"reusable"`
			CountryCode          string      `json:"countryCode"`
			CardToken            interface{} `json:"cardToken"`
			SupportsTokenization bool        `json:"supportsTokenization"`
			MaskedPan            string      `json:"maskedPan"`
		} `json:"cardDetails"`
		AccountDetails  interface{}   `json:"accountDetails"`
		AccountPayments []interface{} `json:"accountPayments"`
		Customer        struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"customer"`
		MetaData struct {
		} `json:"metaData"`
	} `json:"responseBody"`
}

type GetTransactionStatusv1Model struct {
	PaymentReference string `json:"paymentReference" validate:"required"`
}

type GetTransactionStatusv1Response struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TransactionReference string `json:"transactionReference"`
		PaymentReference     string `json:"paymentReference"`
		AmountPaid           string `json:"amountPaid"`
		TotalPayable         string `json:"totalPayable"`
		SettlementAmount     string `json:"settlementAmount"`
		PaidOn               string `json:"paidOn"`
		PaymentStatus        string `json:"paymentStatus"`
		PaymentDescription   string `json:"paymentDescription"`
		Currency             string `json:"currency"`
		PaymentMethod        string `json:"paymentMethod"`
		Product              struct {
			Type      string `json:"type"`
			Reference string `json:"reference"`
		} `json:"product"`
		CardDetails struct {
			CardType             string      `json:"cardType"`
			Last4                string      `json:"last4"`
			ExpMonth             string      `json:"expMonth"`
			ExpYear              string      `json:"expYear"`
			Bin                  string      `json:"bin"`
			BankCode             interface{} `json:"bankCode"`
			BankName             interface{} `json:"bankName"`
			Reusable             bool        `json:"reusable"`
			CountryCode          interface{} `json:"countryCode"`
			CardToken            interface{} `json:"cardToken"`
			SupportsTokenization bool        `json:"supportsTokenization"`
			MaskedPan            string      `json:"maskedPan"`
		} `json:"cardDetails"`
		AccountDetails  interface{}   `json:"accountDetails"`
		AccountPayments []interface{} `json:"accountPayments"`
		Customer        struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"customer"`
		MetaData struct {
		} `json:"metaData"`
	} `json:"responseBody"`
}

type PayWithUssdModel struct {
	TransactionReference string `json:"transactionReference" validate:"required"`
	BankUssdCode         string `json:"bankUssdCode" validate:"omitempty"`
}

type PayWithUssdResponse struct {
}

type PayWithBankTransferModel struct {
	TransactionReference string `json:"transactionReference" validate:"required"`
	BankCode             string `json:"bankCode,omitempty" validate:"omitempty"`
}

type PayWithBankTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		AccountNumber          string      `json:"accountNumber"`
		AccountName            string      `json:"accountName"`
		BankName               string      `json:"bankName"`
		BankCode               string      `json:"bankCode"`
		AccountDurationSeconds int         `json:"accountDurationSeconds"`
		UssdPayment            string      `json:"ussdPayment"`
		RequestTime            string      `json:"requestTime"`
		ExpiresOn              string      `json:"expiresOn"`
		TransactionReference   string      `json:"transactionReference"`
		PaymentReference       string      `json:"paymentReference"`
		Amount                 float64     `json:"amount"`
		Fee                    float64     `json:"fee"`
		TotalPayable           float64     `json:"totalPayable"`
		CollectionChannel      string      `json:"collectionChannel"`
		ProductInformation     interface{} `json:"productInformation"`
	} `json:"responseBody"`
}

type ChargeCard struct {
	Number      string `json:"number" validate:"required,number"`
	ExpiryMonth string `json:"expiryMonth" validate:"required,number,len=2"`
	ExpiryYear  string `json:"expiryYear" validate:"required,number,len=4"`
	CVV         string `json:"cvv" validate:"required,number,len=3"`
	PIN         string `json:"pin" validate:"required,number"`
}

type DeviceInformation struct {
	HttpBrowserLanguage          string `json:"httpBrowserLanguage" validate:"required"`
	HttpBrowserJavaEnabled       bool   `json:"httpBrowserJavaEnabled"`
	HttpBrowserJavaScriptEnabled bool   `json:"httpBrowserJavaScriptEnabled"`
	HttpBrowserColorDepth        int    `json:"httpBrowserColorDepth" validate:"required"`
	HttpBrowserScreenHeight      int    `json:"httpBrowserScreenHeight" validate:"required"`
	HttpBrowserScreenWidth       int    `json:"httpBrowserScreenWidth" validate:"required"`
	HttpBrowserTimeDifference    string `json:"httpBrowserTimeDifference" validate:"required"`
	UserAgentBrowserValue        string `json:"userAgentBrowserValue" validate:"required"`
}

type ChargeCardModel struct {
	TransactionReference string            `json:"transactionReference" validate:"required"`
	CollectionChannel    string            `json:"collectionChannel" validate:"required"`
	Card                 ChargeCard        `json:"card" validate:"required"`
	DeviceInformation    DeviceInformation `json:"deviceInformation" validate:"required"`
}

type ChargeCardResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Status               string  `json:"status"`
		Message              string  `json:"message"`
		TransactionReference string  `json:"transactionReference"`
		PaymentReference     string  `json:"paymentReference"`
		AuthorizedAmount     float64 `json:"authorizedAmount"`
		OtpData              struct {
			Id       string `json:"id"`
			Message  string `json:"message"`
			AuthData string `json:"authData"`
		} `json:"otpData,omitempty"`
	} `json:"responseBody"`
}

type AuthorizeOTPModel struct {
	TransactionReference string `json:"transactionReference" validate:"required"`
	CollectionChannel    string `json:"collectionChannel" validate:"required"`
	TokenId              string `json:"tokenId" validate:"required"`
	Token                string `json:"token" validate:"required,number"`
}

type AuthorizeOTPResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		PaymentStatus        string  `json:"paymentStatus"`
		PaymentDescription   string  `json:"paymentDescription"`
		TransactionReference string  `json:"transactionReference"`
		PaymentReference     string  `json:"paymentReference"`
		AmountPaid           float64 `json:"amountPaid"`
		CurrencyPaid         string  `json:"currencyPaid"`
	} `json:"responseBody"`
}

type ThreeDsSecureAuthTransactionCard struct {
	Number      string `json:"number" validate:"required,number"`
	ExpiryMonth string `json:"expiryMonth" validate:"required,number,len=2"`
	ExpiryYear  string `json:"expiryYear" validate:"required,number,len=4"`
	CVV         string `json:"cvv" validate:"required,number,len=3"`
}

type ThreeDsSecureAuthTransactionModel struct {
	TransactionReference string                           `json:"transactionReference" validate:"required"`
	CollectionChannel    string                           `json:"collectionChannel" validate:"required"`
	Card                 ThreeDsSecureAuthTransactionCard `json:"card" validate:"required"`
	ApiKey               string                           `json:"apiKey" validate:"required"`
}

type ThreeDsSecureAuthTransactionResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		PaymentStatus        string `json:"paymentStatus"`
		PaymentDescription   string `json:"paymentDescription"`
		TransactionReference string `json:"transactionReference"`
		PaymentReference     string `json:"paymentReference"`
		AmountPaid           int    `json:"amountPaid"`
		CurrencyPaid         string `json:"currencyPaid"`
	} `json:"responseBody"`
}

type CardTokenizationModel struct {
	CustomerName       string  `json:"customerName" validate:"required,min=3"`
	CustomerEmail      string  `json:"customerEmail" validate:"required,email"`
	Amount             float64 `validate:"required,number,min=20" json:"amount"`
	PaymentDescription string  `json:"paymentDescription" validate:"required,min=3"`
	PaymentReference   string  `json:"paymentReference" validate:"required,min=3"`
	CurrencyCode       string  `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	ContractCode       string  `json:"contractCode" validate:"required"`
	CardToken          string  `json:"cardToken" validate:"required"`
	ApiKey             string  `json:"apiKey" validate:"required"`
}

func (s *CardTokenizationModel) SetDefault() {
	if s.CurrencyCode == "" {
		s.CurrencyCode = "NGN"
	}
}

type CardTokenizationResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TransactionReference string `json:"transactionReference"`
		PaymentReference     string `json:"paymentReference"`
		AmountPaid           string `json:"amountPaid"`
		TotalPayable         string `json:"totalPayable"`
		SettlementAmount     string `json:"settlementAmount"`
		PaidOn               string `json:"paidOn"`
		PaymentStatus        string `json:"paymentStatus"`
		PaymentDescription   string `json:"paymentDescription"`
		Currency             string `json:"currency"`
		PaymentMethod        string `json:"paymentMethod"`
		Product              struct {
			Type      string `json:"type"`
			Reference string `json:"reference"`
		} `json:"product"`
		CardDetails struct {
			CardType             string `json:"cardType"`
			Last4                string `json:"last4"`
			ExpMonth             string `json:"expMonth"`
			ExpYear              string `json:"expYear"`
			Bin                  string `json:"bin"`
			BankCode             string `json:"bankCode"`
			BankName             string `json:"bankName"`
			Reusable             bool   `json:"reusable"`
			CountryCode          string `json:"countryCode"`
			CardToken            string `json:"cardToken"`
			SupportsTokenization bool   `json:"supportsTokenization"`
			MaskedPan            string `json:"maskedPan"`
		} `json:"cardDetails"`
		AccountDetails struct {
		} `json:"accountDetails"`
		AccountPayments []struct {
		} `json:"accountPayments"`
		Customer struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"customer"`
		MetaData struct {
			IpAddress  string `json:"ipAddress"`
			DeviceType string `json:"deviceType"`
		} `json:"metaData"`
	} `json:"responseBody"`
}
