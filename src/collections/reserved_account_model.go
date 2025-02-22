package collections

type IncomeSplit struct {
	SubAccountCode  string  `json:"subAccountCode" validate:"required"`
	SplitPercentage float64 `json:"splitPercentage" validate:"omitempty,min=0"`
	FeePercentage   float64 `json:"feePercentage" validate:"omitempty,min=0"`
	FeeBearer       bool    `json:"feeBearer" validate:"omitempty"`
}

type ReservedAccountSchema struct {
	CustomerName          string                 `json:"customerName" validate:"required,min=3"`
	CustomerEmail         string                 `json:"customerEmail" validate:"required,email"`
	AccountName           string                 `json:"accountName" validate:"required,min=3"`
	AccountReference      string                 `json:"accountReference" validate:"required"`
	CurrencyCode          string                 `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	ContractCode          string                 `json:"contractCode" validate:"required"`
	Bvn                   string                 `json:"bvn" validate:"omitempty,len=11,required_without=Nin"`
	Nin                   string                 `json:"nin" validate:"omitempty,len=11"`
	GetAllAvailableBanks  bool                   `json:"getAllAvailableBanks" validate:"omitempty"`
	PreferredBanks        []string               `json:"preferredBanks" validate:"dive,required_if=GetAllAvailableBanks,false"`
	IncomeSplitConfig     []IncomeSplit          `json:"incomeSplitConfig" validate:"omitempty"`
	MetaData              map[string]interface{} `json:"metaData,omitempty"`
	RestrictPaymentSource bool                   `json:"restrictPaymentSource" validate:"omitempty"`
	AllowedPaymentSources map[string]interface{} `json:"allowedPaymentSources,omitempty" validate:"required_if=RestrictPaymentSource,true"`
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
