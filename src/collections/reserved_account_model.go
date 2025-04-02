package collections

type IncomeSplit struct {
	SubAccountCode            string  `json:"subAccountCode" validate:"required"`
	SplitPercentage           float64 `json:"splitPercentage" validate:"omitempty,min=0"`
	FeePercentage             float64 `json:"feePercentage" validate:"omitempty,min=0"`
	FeeBearer                 bool    `json:"feeBearer" validate:"omitempty"`
	ReservedAccountConfigCode bool    `json:"reservedAccountConfigCode" validate:"omitempty"`
}

type ReservedAccountSchema struct {
	CustomerName          string                 `json:"customerName" validate:"required,min=3"`
	CustomerEmail         string                 `json:"customerEmail" validate:"required,email"`
	AccountName           string                 `json:"accountName" validate:"required,min=3"`
	AccountReference      string                 `json:"accountReference" validate:"required"`
	CurrencyCode          string                 `json:"currencyCode,omitempty" validate:"omitempty,oneof=NGN USD EUR"`
	ContractCode          string                 `json:"contractCode" validate:"required"`
	Bvn                   string                 `json:"bvn,omitempty" validate:"omitempty,len=11,required_without=Nin"`
	Nin                   string                 `json:"nin,omitempty" validate:"omitempty,len=11"`
	GetAllAvailableBanks  bool                   `json:"getAllAvailableBanks,omitempty" validate:"omitempty"`
	PreferredBanks        []string               `json:"preferredBanks,omitempty" validate:"dive,required_if=GetAllAvailableBanks false"`
	IncomeSplitConfig     []IncomeSplit          `json:"incomeSplitConfig,omitempty" validate:"omitempty"`
	MetaData              map[string]interface{} `json:"metaData,omitempty"`
	RestrictPaymentSource bool                   `json:"restrictPaymentSource,omitempty" validate:"omitempty"`
	AllowedPaymentSources map[string]interface{} `json:"allowedPaymentSources,omitempty" validate:"required_if=RestrictPaymentSource true"`
}

func (s *ReservedAccountSchema) SetDefault() {
	if s.CurrencyCode == "" {
		s.CurrencyCode = "NGN"
	}
}

type ReservedAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		ContractCode          string        `json:"contractCode"`
		AccountReference      string        `json:"accountReference"`
		AccountName           string        `json:"accountName"`
		CurrencyCode          string        `json:"currencyCode"`
		CustomerEmail         string        `json:"customerEmail"`
		CustomerName          string        `json:"customerName"`
		AccountNumber         string        `json:"accountNumber"`
		BankName              string        `json:"bankName"`
		BankCode              string        `json:"bankCode"`
		CollectionChannel     string        `json:"collectionChannel"`
		ReservationReference  string        `json:"reservationReference"`
		ReservedAccountType   string        `json:"reservedAccountType"`
		Status                string        `json:"status"`
		CreatedOn             string        `json:"createdOn"`
		IncomeSplitConfig     []IncomeSplit `json:"incomeSplitConfig"`
		RestrictPaymentSource bool          `json:"restrictPaymentSource"`
	} `json:"responseBody"`
}

type AddLinkedAccountSchema struct {
	AccountReference     string   `json:"accountReference" validate:"required"`
	GetAllAvailableBanks bool     `json:"getAllAvailableBanks" validate:"omitempty"`
	PreferredBanks       []string `json:"preferredBanks" validate:"dive,required_if=GetAllAvailableBanks false"`
}

type AddLinkedAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		ContractCode     string `json:"contractCode"`
		AccountReference string `json:"accountReference"`
		AccountName      string `json:"accountName"`
		CurrencyCode     string `json:"currencyCode"`
		CustomerEmail    string `json:"customerEmail"`
		CustomerName     string `json:"customerName"`
		Accounts         []struct {
			AccountName   string `json:"accountName"`
			AccountNumber string `json:"accountNumber"`
			BankName      string `json:"bankName"`
			BankCode      string `json:"bankCode"`
		} `json:"accounts"`
		CollectionChannel     string        `json:"collectionChannel"`
		ReservationReference  string        `json:"reservationReference"`
		ReservedAccountType   string        `json:"reservedAccountType"`
		Status                string        `json:"status"`
		CreatedOn             string        `json:"createdOn"`
		IncomeSplitConfig     []IncomeSplit `json:"incomeSplitConfig"`
		BVN                   string        `json:"bvn"`
		RestrictPaymentSource bool          `json:"restrictPaymentSource"`
	} `json:"responseBody"`
}

type ReservedAccountDetailsSchema struct {
	AccountReference string `json:"accountReference" validate:"required"`
}

type ReservedAccountDetailsResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		ContractCode     string `json:"contractCode"`
		AccountReference string `json:"accountReference"`
		AccountName      string `json:"accountName"`
		CurrencyCode     string `json:"currencyCode"`
		CustomerEmail    string `json:"customerEmail"`
		CustomerName     string `json:"customerName"`
		Accounts         []struct {
			AccountName   string `json:"accountName"`
			AccountNumber string `json:"accountNumber"`
			BankName      string `json:"bankName"`
			BankCode      string `json:"bankCode"`
		} `json:"accounts"`
		CollectionChannel    string `json:"collectionChannel"`
		ReservationReference string `json:"reservationReference"`
		ReservedAccountType  string `json:"reservedAccountType"`
		Status               string `json:"status"`
		CreatedOn            string `json:"createdOn"`
		Contract             struct {
			Name        string `json:"name"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"contract"`
		TransactionCount      int           `json:"transactionCount"`
		IncomeSplitConfig     []IncomeSplit `json:"incomeSplitConfig"`
		BVN                   string        `json:"bvn"`
		RestrictPaymentSource bool          `json:"restrictPaymentSource"`
	} `json:"responseBody"`
}

type ReservedAccountTransactionsSchema struct {
	AccountReference string `json:"accountReference" validate:"required"`
	Page             int    `json:"page" validate:"omitempty,min=0"`
	Size             int    `json:"size" validate:"omitempty,min=1"`
}

func (s *ReservedAccountTransactionsSchema) SetDefault() {
	if s.Size == 0 {
		s.Size = 10
	}
}

type ReservedAccountTransactionsResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Content []struct {
			CustomerDTO struct {
				Email        string `json:"email"`
				Name         string `json:"name"`
				MerchantCode string `json:"merchantCode"`
			} `json:"customerDTO"`
			ProviderAmount       float64 `json:"providerAmount"`
			PaymentMethod        string  `json:"paymentMethod"`
			CreatedOn            string  `json:"createdOn"`
			Amount               float64 `json:"amount"`
			Flagged              bool    `json:"flagged"`
			ProviderCode         string  `json:"providerCode"`
			Fee                  float64 `json:"fee"`
			CurrencyCode         string  `json:"currencyCode"`
			CompletedOn          string  `json:"completedOn"`
			PaymentDescription   string  `json:"paymentDescription"`
			PaymentStatus        string  `json:"paymentStatus"`
			TransactionReference string  `json:"transactionReference"`
			PaymentReference     string  `json:"paymentReference"`
			MerchantCode         string  `json:"merchantCode"`
			MerchantName         string  `json:"merchantName"`
			PayableAmount        float64 `json:"payableAmount"`
			AmountPaid           float64 `json:"amountPaid"`
			Completed            bool    `json:"completed"`
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
			Paged      bool `json:"paged"`
			Unpaged    bool `json:"unpaged"`
		} `json:"pageable"`
		Last          bool `json:"last"`
		TotalPages    int  `json:"totalPages"`
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

type DeallocateReservedAccountSchema struct {
	AccountReference string `json:"accountReference" validate:"required"`
}
type DeallocateReservedAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		ContractCode          string        `json:"contractCode"`
		AccountReference      string        `json:"accountReference"`
		AccountName           string        `json:"accountName"`
		CurrencyCode          string        `json:"currencyCode"`
		CustomerEmail         string        `json:"customerEmail"`
		CustomerName          string        `json:"customerName"`
		AccountNumber         string        `json:"accountNumber"`
		BankName              string        `json:"bankName"`
		BankCode              string        `json:"bankCode"`
		CollectionChannel     string        `json:"collectionChannel"`
		ReservationReference  string        `json:"reservationReference"`
		ReservedAccountType   string        `json:"reservedAccountType"`
		Status                string        `json:"status"`
		CreatedOn             string        `json:"createdOn"`
		IncomeSplitConfig     []IncomeSplit `json:"incomeSplitConfig"`
		RestrictPaymentSource bool          `json:"restrictPaymentSource"`
	} `json:"responseBody"`
}

type UpdateReservedAccountKycInfoSchema struct {
	AccountReference string `json:"accountReference" validate:"required"`
	Bvn              string `json:"bvn" validate:"omitempty,len=11,required_without=Nin"`
	Nin              string `json:"nin" validate:"omitempty,len=11"`
}

type UpdateReservedAccountKycInfoResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		AccountReference string `json:"accountReference"`
		AccountName      string `json:"accountName"`
		CustomerEmail    string `json:"customerEmail"`
		CustomerName     string `json:"customerName"`
		Bvn              string `json:"bvn"`
	} `json:"responseBody"`
}
