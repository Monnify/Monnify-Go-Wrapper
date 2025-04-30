package disbursement

type SingleTransfer struct {
	Amount                   float64 `validate:"required,number,min=20" json:"amount"`
	Reference                string  `validate:"required" json:"reference"`
	Narration                string  `validate:"required,min=3" json:"narration"`
	DestinationBankCode      string  `validate:"required,number,min=3" json:"destinationBankCode"`
	DestinationAccountNumber string  `validate:"required,number,len=10" json:"destinationAccountNumber"`
	Currency                 string  `validate:"omitempty,oneof=NGN USD EUR" json:"currency"`
	SourceAccountNumber      string  `validate:"required,number,len=10" json:"sourceAccountNumber"`
}

func (s *SingleTransfer) SetDefault() {
	if s.Currency == "" {
		s.Currency = "NGN"
	}
}

type SingleTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Amount                   float64 `json:"amount"`
		Reference                string  `json:"reference"`
		Status                   string  `json:"status"`
		DateCreated              string  `json:"dateCreated"`
		TotalFee                 float64 `json:"totalFee"`
		DestinationAccountName   string  `json:"destinationAccountName"`
		DestinationBankName      string  `json:"destinationBankName"`
		DestinationAccountNumber string  `json:"destinationAccountNumber"`
		DestinationBankCode      string  `json:"destinationBankCode"`
	} `json:"responseBody"`
}

type bulkTransferTransactionList struct {
	Narration                string  `validate:"required,min=3" json:"narration"`
	DestinationAccountNumber string  `validate:"required,number,len=10" json:"destinationAccountNumber"`
	Amount                   float64 `validate:"required,number,min=20" json:"amount"`
	DestinationBankCode      string  `validate:"required,number,min=3" json:"destinationBankCode"`
	Reference                string  `validate:"required" json:"reference"`
	Currency                 string  `validate:"omitempty,oneof=NGN" json:"currency"`
}

type BulkTransfer struct {
	Title                string                        `validate:"required,min=5" json:"title"`
	BatchReference       string                        `validate:"required" json:"batchReference"`
	OnValidationFailure  string                        `validate:"required,oneof=BREAK CONTINUE" json:"onValidationFailure"`
	NotificationInterval int                           `validate:"required,number,oneof=25 50 75 100" json:"notificationInterval"`
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
		if b.TransactionList[i].Currency == "" {
			b.TransactionList[i].Currency = "NGN"
		}
	}
}

type BulkTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TotalAmount            float64 `json:"totalAmount"`
		TotalFee               float64 `json:"totalFee"`
		BatchReference         string  `json:"batchReference"`
		BatchStatus            string  `json:"batchStatus"`
		TotalTransactionsCount int     `json:"totalTransactionsCount"`
		DateCreated            string  `json:"dateCreated"`
	} `json:"responseBody"`
}

type AuthorizeTransfer struct {
	Reference         string `validate:"required" json:"reference"`
	AuthorizationCode string `validate:"required,number" json:"authorizationCode"`
}

type AuthorizeBulkTransferResponse struct {
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

type AuthorizeSingleTransferResponse struct {
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

type ResendTransferOTP struct {
	Reference string `validate:"required" json:"reference"`
}

type ResendTransferOTPResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Message string `json:"message"`
	} `json:"responseBody"`
}

type GetStatus struct {
	Reference string `validate:"required" json:"reference"`
}

type GetBulkStatus struct {
	Reference string `validate:"required" json:"reference"`
	PageNo    int    `validate:"required,number,min=0" json:"pageNo"`
	PageSize  int    `validate:"required,number,min=1" json:"pageSize"`
}

type TransferStatus struct {
	Amount                   int    `json:"amount"`
	Reference                string `json:"reference"`
	Narration                string `json:"narration"`
	Currency                 string `json:"currency"`
	Fee                      int    `json:"fee"`
	TwoFaEnabled             bool   `json:"twoFaEnabled"`
	Status                   string `json:"status"`
	TransactionDescription   string `json:"transactionDescription"`
	TransactionReference     string `json:"transactionReference"`
	CreatedOn                string `json:"createdOn"`
	SourceAccountNumber      string `json:"sourceAccountNumber"`
	DestinationAccountName   string `json:"destinationAccountName"`
	DestinationBankName      string `json:"destinationBankName"`
	DestinationAccountNumber string `json:"destinationAccountNumber"`
	DestinationBankCode      string `json:"destinationBankCode"`
}

type Pageable struct {
	Sort struct {
		Sorted   bool `json:"sorted"`
		Unsorted bool `json:"unsorted"`
		Empty    bool `json:"empty"`
	} `json:"sort"`
	PageSize   int  `json:"pageSize"`
	PageNumber int  `json:"pageNumber"`
	Offset     int  `json:"offset"`
	Paged      bool `json:"paged"`
	Unpaged    bool `json:"unpaged"`
}

type GetSingleTransferStatusResponse struct {
	RequestSuccessful bool           `json:"requestSuccessful"`
	ResponseMessage   string         `json:"responseMessage"`
	ResponseCode      string         `json:"responseCode"`
	ResponseBody      TransferStatus `json:"responseBody"`
}

type GetBulkTransferStatusResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Content       []TransferStatus `json:"content"`
		Pageable      Pageable         `json:"pageable"`
		Last          bool             `json:"last"`
		TotalPages    int              `json:"totalPages"`
		TotalElements int              `json:"totalElements"`
		Sort          struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		First            bool `json:"first"`
		NumberOfElements int  `json:"numberOfElements"`
		Size             int  `json:"size"`
		Number           int  `json:"number"`
		Empty            bool `json:"empty"`
	} `json:"responseBody"`
}

type GetAllSingleTransfer struct {
	PageNo   int `validate:"required,number,min=0" json:"pageNo"`
	PageSize int `validate:"required,number,min=1" json:"pageSize"`
}

type GetAllSingleTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Content       []TransferStatus `json:"content"`
		Pageable      Pageable         `json:"pageable"`
		Last          bool             `json:"last"`
		TotalPages    int              `json:"totalPages"`
		TotalElements int              `json:"totalElements"`
		Sort          struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		First            bool `json:"first"`
		NumberOfElements int  `json:"numberOfElements"`
		Size             int  `json:"size"`
		Number           int  `json:"number"`
		Empty            bool `json:"empty"`
	} `json:"responseBody"`
}

type GetAllBulkTransfer struct {
	PageNo   int `validate:"required,number,min=0" json:"pageNo"`
	PageSize int `validate:"required,number,min=1" json:"pageSize"`
}

type GetAllBulkTransferResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Content []struct {
			TotalAmount            int    `json:"totalAmount"`
			TotalFee               int    `json:"totalFee"`
			BatchReference         string `json:"batchReference"`
			BatchStatus            string `json:"batchStatus"`
			TotalTransactionsCount int    `json:"totalTransactionsCount"`
			DateCreated            string `json:"dateCreated"`
		} `json:"content"`
		Pageable struct {
			Sort struct {
				Sorted   bool `json:"sorted"`
				Unsorted bool `json:"unsorted"`
				Empty    bool `json:"empty"`
			} `json:"sort"`
			PageSize   int  `json:"pageSize"`
			PageNumber int  `json:"pageNumber"`
			Offset     int  `json:"offset"`
			Unpaged    bool `json:"unpaged"`
			Paged      bool `json:"paged"`
		} `json:"pageable"`
		TotalPages    int  `json:"totalPages"`
		Last          bool `json:"last"`
		TotalElements int  `json:"totalElements"`
		Sort          struct {
			Sorted   bool `json:"sorted"`
			Unsorted bool `json:"unsorted"`
			Empty    bool `json:"empty"`
		} `json:"sort"`
		First            bool `json:"first"`
		NumberOfElements int  `json:"numberOfElements"`
		Size             int  `json:"size"`
		Number           int  `json:"number"`
		Empty            bool `json:"empty"`
	} `json:"responseBody"`
}
