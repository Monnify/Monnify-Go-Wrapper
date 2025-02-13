package constants

const (
	LoginEndpoint              = "/api/v1/auth/login"
	DisbursementSingleEndpoint = "/api/v2/disbursements/single"
	BulkTransferEndpoint       = "/api/v2/disbursements/batch"

	AuthentionKey = "authentication"
)

var (
	SupportedCurrency = []string{"NGN"}
	ValidationEnum    = []string{"BREAK", "CONTINUE"}
)
