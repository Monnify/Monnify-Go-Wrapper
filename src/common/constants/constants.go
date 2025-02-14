package constants

const (
	LoginEndpoint                   = "/api/v1/auth/login"
	DisbursementSingleEndpoint      = "/api/v2/disbursements/single"
	BulkTransferEndpoint            = "/api/v2/disbursements/batch"
	AuthorizeBulkTransferEndpoint   = "/api/v2/disbursements/batch/validate-otp"
	AuthorizeSingleTransferEndpoint = "/api/v2/disbursements/single/validate-otp"
	ResendTransferOTPEndpoint       = "/api/v2/disbursements/single/resend-otp"
	GetSingleTransferStatusEndpoint = "/api/v2/disbursements/single/summary?reference=%s"

	AuthentionKey = "authentication"
)

var (
	SupportedCurrency = []string{"NGN"}
	OnValidationEnum  = []string{"BREAK", "CONTINUE"}
)
