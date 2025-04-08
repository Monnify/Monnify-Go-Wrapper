package collections

type CreateSubAccountModel struct {
	CurrencyCode           string  `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	AccountNumber          string  `json:"accountNumber" validate:"required,number,len=10"`
	BankCode               string  `json:"bankCode" validate:"required,number,min=3"`
	DefaultSplitPercentage float64 `json:"defaultSplitPercentage" validate:"required,min=0"`
	Email                  string  `json:"email" validate:"required,email"`
}

type SubAccountBody struct {
	SubAccountCode         string   `json:"subAccountCode"`
	AccountName            string   `json:"accountName"`
	CurrencyCode           string   `json:"currencyCode"`
	AccountNumber          string   `json:"accountNumber"`
	BankName               string   `json:"bankName"`
	BankCode               string   `json:"bankCode"`
	Email                  string   `json:"email"`
	DefaultSplitPercentage float64  `json:"defaultSplitPercentage"`
	SettlementProfileCode  string   `json:"settlementProfileCode"`
	SettlementReportEmails []string `json:"settlementReportEmails"`
}

type CreateSubAccountResponse struct {
	RequestSuccessful bool             `json:"requestSuccessful"`
	ResponseMessage   string           `json:"responseMessage"`
	ResponseCode      string           `json:"responseCode"`
	ResponseBody      []SubAccountBody `json:"responseBody"`
}

type DeleteSubAccountModel struct {
	SubAccountCode string `json:"subAccountCode" validate:"required"`
}

type DeleteSubAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
}

type GetSubAccountsResponse struct {
	RequestSuccessful bool             `json:"requestSuccessful"`
	ResponseMessage   string           `json:"responseMessage"`
	ResponseCode      string           `json:"responseCode"`
	ResponseBody      []SubAccountBody `json:"responseBody"`
}

type UpdateSubAccountModel struct {
	CurrencyCode           string  `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	AccountNumber          string  `json:"accountNumber" validate:"required,number,len=10"`
	BankCode               string  `json:"bankCode" validate:"required,number,min=3"`
	DefaultSplitPercentage float64 `json:"defaultSplitPercentage" validate:"required,min=0"`
	Email                  string  `json:"email" validate:"required,email"`
	SubAccountCode         string  `json:"subAccountCode" validate:"required"`
}

type UpdateSubAccountResponse struct {
	RequestSuccessful bool           `json:"requestSuccessful"`
	ResponseMessage   string         `json:"responseMessage"`
	ResponseCode      string         `json:"responseCode"`
	ResponseBody      SubAccountBody `json:"responseBody"`
}
