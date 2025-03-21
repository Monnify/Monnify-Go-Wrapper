package disbursement

type InitiateRefundModel struct {
	RefundReason             string  `json:"refundReason" validate:"required,max=64"`
	DestinationAccountNumber string  `json:"destinationAccountNumber" validate:"omitempty,number,len=10"`
	RefundAmount             float64 `json:"refundAmount" validate:"required,number,min=20"`
	DestinationBankCode      string  `json:"destinationBankCode" validate:"omitempty,number,min=3"`
	RefundReference          string  `json:"refundReference" validate:"required,min=3"`
	CurrencyCode             string  `json:"currencyCode" validate:"omitempty,oneof=NGN USD EUR"`
	TransactionReference     string  `json:"transactionReference" validate:"required"`
	CustomerNote             string  `json:"customerNote" validate:"required,max=16"`
}

func (s *InitiateRefundModel) SetDefault() {
	if s.CurrencyCode == "" {
		s.CurrencyCode = "NGN"
	}
}

type InitiateRefundResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		RefundReference      string `json:"refundReference"`
		TransactionReference string `json:"transactionReference"`
		RefundReason         string `json:"refundReason"`
		CustomerNote         string `json:"customerNote"`
		RefundAmount         int    `json:"refundAmount"`
		RefundType           string `json:"refundType"`
		RefundStatus         string `json:"refundStatus"`
		RefundStrategy       string `json:"refundStrategy"`
		Comment              string `json:"comment"`
		CompletedOn          string `json:"completedOn"`
		CreatedOn            string `json:"createdOn"`
	} `json:"responseBody"`
}

type GetAllRefundsModel struct {
	Page int `json:"page" validate:"required,number,min=0"`
	Size int `json:"size" validate:"required,number,min=1"`
}

type GetAllRefundsResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Content []struct {
			RefundReference      string `json:"refundReference"`
			TransactionReference string `json:"transactionReference"`
			RefundReason         string `json:"refundReason"`
			CustomerNote         string `json:"customerNote"`
			RefundAmount         int    `json:"refundAmount"`
			RefundType           string `json:"refundType"`
			RefundStatus         string `json:"refundStatus"`
			RefundStrategy       string `json:"refundStrategy"`
			Comment              string `json:"comment"`
			CompletedOn          string `json:"completedOn"`
			CreatedOn            string `json:"createdOn"`
		} `json:"content"`
		Pageable         interface{} `json:"pageable"`
		Last             bool        `json:"last"`
		TotalElements    int         `json:"totalElements"`
		TotalPages       int         `json:"totalPages"`
		Sort             interface{} `json:"sort"`
		First            bool        `json:"first"`
		NumberOfElements int         `json:"numberOfElements"`
		Size             int         `json:"size"`
		Number           int         `json:"number"`
		Empty            bool        `json:"empty"`
	} `json:"responseBody"`
}

type GetRefundStatusModel struct {
	RefundReference string `json:"refundReference" validate:"required"`
}

type GetRefundStatusResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		RefundReference      string `json:"refundReference"`
		TransactionReference string `json:"transactionReference"`
		RefundReason         string `json:"refundReason"`
		CustomerNote         string `json:"customerNote"`
		RefundAmount         int    `json:"refundAmount"`
		RefundType           string `json:"refundType"`
		RefundStatus         string `json:"refundStatus"`
		RefundStrategy       string `json:"refundStrategy"`
		Comment              string `json:"comment"`
		CompletedOn          string `json:"completedOn"`
		CreatedOn            string `json:"createdOn"`
	} `json:"responseBody"`
}
