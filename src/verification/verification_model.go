package verification

type ValidateBankAccountModel struct {
	AccountNumber string `json:"accountNumber" validate:"required,number,len=10"`
	BankCode      string `json:"bankCode" validate:"required,number,min=3"`
}

type ValidateBankAccountResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		AccountNumber string `json:"accountNumber"`
		AccountName   string `json:"accountName"`
		BankCode      string `json:"bankCode"`
	} `json:"responseBody"`
}

type VerifyBvnInformationModel struct {
	BVN         string `json:"bvn" validate:"required,number,min=11"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
	MobileNo    string `json:"mobileNo" validate:"required,min=11"`
	Name        string `json:"name" validate:"omitempty"`
}

type VerifyBvnInformationResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Bvn  string `json:"bvn"`
		Name struct {
			MatchStatus     string `json:"matchStatus"`
			MatchPercentage int    `json:"matchPercentage"`
		} `json:"name"`
		DateOfBirth string `json:"dateOfBirth"`
		MobileNo    string `json:"mobileNo"`
	} `json:"responseBody"`
}

type MatchBvnAndAccountNameModel struct {
	BVN           string `json:"bvn" validate:"required,number,min=11"`
	AccountNumber string `json:"accountNumber" validate:"required,number,len=10"`
	BankCode      string `json:"bankCode" validate:"required,number,min=3"`
}

type MatchBvnAndAccountNameResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		Bvn             string `json:"bvn"`
		AccountNumber   string `json:"accountNumber"`
		AccountName     string `json:"accountName"`
		MatchStatus     string `json:"matchStatus"`
		MatchPercentage int    `json:"matchPercentage"`
	} `json:"responseBody"`
}
