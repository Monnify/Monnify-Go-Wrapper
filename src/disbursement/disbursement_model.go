package disbursement

type SingleTransfer struct {
	Amount                   float64 `validate:"required,number,min=20" json:"amount"`
	Reference                string  `validate:"required" json:"reference"`
	Narration                string  `validate:"required,min=3" json:"narration"`
	DestinationBankCode      string  `validate:"required,number,min=3" json:"destinationBankCode"`
	DestinationAccountNumber string  `validate:"required,number,len=10" json:"destinationAccountNumber"`
	Currency                 string  `validate:"required,validateCurrency" json:"currency"`
	SourceAccountNumber      string  `validate:"required,number,len=10" json:"sourceAccountNumber"`
}

type SingleTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Amount                   int    `json:"amount"`
		Reference                string `json:"reference"`
		Status                   string `json:"status"`
		DateCreated              string `json:"dateCreated"`
		TotalFee                 int    `json:"totalFee"`
		DestinationAccountName   string `json:"destinationAccountName"`
		DestinationBankName      string `json:"destinationBankName"`
		DestinationAccountNumber string `json:"destinationAccountNumber"`
		DestinationBankCode      string `json:"destinationBankCode"`
	} `json:"responseBody"`
}
