package constants

const (
	LoginEndpoint                   = "/api/v1/auth/login"
	DisbursementSingleEndpoint      = "/api/v2/disbursements/single"
	BulkTransferEndpoint            = "/api/v2/disbursements/batch"
	AuthorizeBulkTransferEndpoint   = "/api/v2/disbursements/batch/validate-otp"
	AuthorizeSingleTransferEndpoint = "/api/v2/disbursements/single/validate-otp"
	ResendTransferOTPEndpoint       = "/api/v2/disbursements/single/resend-otp"
	GetSingleTransferStatusEndpoint = "/api/v2/disbursements/single/summary?reference=%s"
	GetBulkTransferStatusEndpoint   = "/api/v2/disbursements/bulk/%s/transactions?pageSize=%d&pageNo=%d"
	AllSingleTransferEndpoint       = "/api/v2/disbursements/single/transactions?pageSize=%d&pageNo=%d"
	AllBulkTransferEndpoint         = "/api/v2/disbursements/bulk/transactions?pageSize=%d&pageNo=%d"
	CreateReservedAccountEndpoint   = "/api/v2/bank-transfer/reserved-accounts"
	AddLinkedAccountEndpoint        = "/api/v1/bank-transfer/reserved-accounts/add-linked-accounts/%s"

	AuthentionKey = "authentication"
)
