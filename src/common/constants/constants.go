package constants

const (
	LoginEndpoint                 = "/api/v1/auth/login"
	DisbursementSingleEndpoint    = "/api/v2/disbursements/single"
	BulkTransferEndpoint          = "/api/v2/disbursements/batch"
	AuthorizeBulkTransferEndpoint = "/api/v2/disbursements/batch/validate-otp"

	AuthentionKey = "authentication"
)

var (
	SupportedCurrency = []string{"NGN"}
	OnValidationEnum  = []string{"BREAK", "CONTINUE"}
)
