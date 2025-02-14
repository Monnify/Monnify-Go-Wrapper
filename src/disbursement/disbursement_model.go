package disbursement

type SingleTransfer struct {
	Amount                   float64 `validate:"required,number,min=20" json:"amount"`
	Reference                string  `validate:"required" json:"reference"`
	Narration                string  `validate:"required,min=3" json:"narration"`
	DestinationBankCode      string  `validate:"required,number,min=3" json:"destinationBankCode"`
	DestinationAccountNumber string  `validate:"required,number,len=10" json:"destinationAccountNumber"`
	CurrencyCode             string  `validate:"validateCurrency" json:"currencyCode"`
	SourceAccountNumber      string  `validate:"required,number,len=10" json:"sourceAccountNumber"`
}

func (s *SingleTransfer) SetDefault() {
	if s.CurrencyCode == "" {
		s.CurrencyCode = "NGN"
	}
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

type bulkTransferTransactionList struct {
	Narration                string  `validate:"required,min=3" json:"narration"`
	DestinationAccountNumber string  `validate:"number,len=10" json:"destinationAccountNumber"`
	Amount                   float64 `validate:"required,number,min=20" json:"amount"`
	DestinationBankCode      string  `validate:"required,number,min=3" json:"destinationBankCode"`
	Reference                string  `validate:"required" json:"reference"`
	CurrencyCode             string  `validate:"validateCurrency" json:"currencyCode"`
}

type BulkTransfer struct {
	Title                string                        `validate:"required,min=5" json:"title"`
	BatchReference       string                        `validate:"required" json:"batchReference"`
	OnValidationFailure  string                        `validate:"onValidationEnum" json:"onValidationFailure"`
	NotificationInterval int                           `validate:"number,min=10" json:"notificationInterval"`
	Narration            string                        `validate:"required,min=3" json:"narration"`
	SourceAccountNumber  string                        `validate:"number,len=10" json:"sourceAccountNumber"`
	TransactionList      []bulkTransferTransactionList `validate:"required" json:"transactionList"`
}

func (b *BulkTransfer) SetDefault() {
	if b.OnValidationFailure == "" {
		b.OnValidationFailure = "CONTINUE"
	}

	if b.NotificationInterval == 0 {
		b.NotificationInterval = 20
	}

	for i := range b.TransactionList {
		if b.TransactionList[i].CurrencyCode == "" {
			b.TransactionList[i].CurrencyCode = "NGN"
		}
	}
}

type BulkTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TotalAmount            int    `json:"totalAmount"`
		TotalFee               int    `json:"totalFee"`
		BatchReference         string `json:"batchReference"`
		BatchStatus            string `json:"batchStatus"`
		TotalTransactionsCount int    `json:"totalTransactionsCount"`
		DateCreated            string `json:"dateCreated"`
	} `json:"responseBody"`
}

type AuthorizeTransfer struct {
	Reference         string `validate:"required" json:"reference"`
	AuthorizationCode string `validate:"required,number" json:"authorizationCode"`
}

type AuthorizeTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TotalAmount            int    `json:"totalAmount"`
		TotalFee               int    `json:"totalFee"`
		BatchReference         string `json:"batchReference"`
		BatchStatus            string `json:"batchStatus"`
		TotalTransactionsCount int    `json:"totalTransactionsCount"`
		DateCreated            string `json:"dateCreated"`
	} `json:"responseBody"`
}
