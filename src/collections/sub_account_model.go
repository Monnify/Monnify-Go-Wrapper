package collections

type CreateSubAccountModel struct {
	CurrencyCode           string  `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	AccountNumber          int     `json:"accountNumber" validate:"required,len=10"`
	BankCode               int     `json:"bankCode" validate:"required,min=3"`
	DefaultSplitPercentage float64 `json:"defaultSplitPercentage" validate:"required,min=0"`
	Email                  string  `json:"email" validate:"required,email"`
}

type CreateSubAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      []struct {
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
	} `json:"responseBody"`
}

type DeleteSubAccountModel struct {
	SubAccountCode string `json:"subAccountCode" validate:"required"`
}

type DeleteSubAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
}
