# Golang Monnify Package Documentation

A Golang package to effortlessly integrate the Monnify payment gateway API into your Golang projects.

## Installation

Install Go Package:

```bash
go get github.com/Monnify/Monnify-Go-Wrapper
```

## Quick Start

Here's how to quickly initialize a payment transaction:

```go
import (
  "log"

  "github.com/Monnify/Monnify-Go-Wrapper"
  "github.com/Monnify/Monnify-Go-Wrapper/src/collections"
)

monnifyPayment := monnify.New(&monnify.Options{
  ApiKey:    "api key",
  SecretKey: "secret key",
  IsProduction: true
})

resp, err := monnifyPayment.Transaction.InitializeTransaction(
  collections.InitializeTransactionModel{
    CustomerEmail:      "example@gmail.com",
    CustomerName:       "John Doe",
    Amount:             20.87,
    PaymentReference:   "payment reference",
    PaymentDescription: "description",
    ContractCode:       "contract code",
  },
)

if err != nil {
  log.Fatalln(err)
}

log.Println(resp)
```

# Available Services

This package provides the following services:

1. **Transaction Service**: Manage payments, authorizations, and statuses.
2. **Customer Reserved Account Service**: Create/manage virtual accounts.
3. **Disbursement Service**: Manage single and bulk fund transfers.
4. **Verification Service**: Perform account, BVN, NIN verifications.
5. **Sub Account Service**: Create/manage sub-accounts for split payments.
6. **Refund Service**: Handle payment refunds.

# Detailed Usage

## Transaction Service

The Transaction Service handles all payment-related operations.

### All Available Methods

```go
// Initialize a new transaction
monnifyPayment.Transaction.InitializeTransaction(body);

// Initialize bank transfer payment
monnifyPayment.Transaction.PayWithBankTransfer(body);

// Charge a card
monnifyPayment.Transaction.ChargeCard(body);

// Card tokenization
monnifyPayment.Transaction.CardTokenization(body);

/* Card Authorization */

// Authorize with OTP
monnifyPayment.Transaction.AuthorizeOTP(body);

// Authorize 3D secure card
monnifyPayment.Transaction.ThreeDsSecureAuthTransaction(body);

/* Transaction Information */

// Get transaction status v
monnifyPayment.Transaction.GetTransactionStatusv1(body);

// Get transaction status v2
monnifyPayment.Transaction.GetTransactionStatusv2(body);
```

## Transaction Initialization

```go
monnifyPayment.Transaction.InitializeTransaction(collections.InitializeTransactionModel{
    CustomerEmail:      "example@gmail.com",
    CustomerName:       "John Doe",
    Amount:             20.87,
    PaymentReference:   "payment reference",
    PaymentDescription: "description",
    ContractCode:       "contract code",
  },
);
```

**Required fields:** `CustomerEmail`, `CustomerName`, `Amount`, `PaymentReference`, `PaymentDescription`, `ContractCode`.

**Optional fields:** `PaymentMethods`, `IncomeSplitConfig`, `RedirectUrl`, `CurrencyCode`, `MetaData`.

### Pay with Bank Transfer

Initializes a bank transfer payment.

```go
monnifyPayment.Transaction.PayWithBankTransfer(collections.PayWithBankTransferModel{
		TransactionReference: "transactionReference",
		BankCode:             "058",
	});
```

### Charge Card

Process a card payment.

```go
monnifyPayment.Transaction.ChargeCard(collections.ChargeCardModel{
		TransactionReference: transactionReference,
		CollectionChannel:    "API_NOTIFICATION",
		Card: ChargeCard{
			Number:      "4111111111111111",
			ExpiryMonth: "12",
			ExpiryYear:  "2025",
			CVV:         "123",
			PIN:         "1234",
		},
		DeviceInformation: DeviceInformation{
			HttpBrowserLanguage:          "en-US",
			HttpBrowserJavaEnabled:       false,
			HttpBrowserJavaScriptEnabled: false,
			HttpBrowserColorDepth:        24,
			HttpBrowserScreenHeight:      1203,
			HttpBrowserScreenWidth:       2138,
			HttpBrowserTimeDifference:    "24",
			UserAgentBrowserValue:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	});
```

### Authorize with OTP

```go
monnifyPayment.Transaction.AuthorizeOTP(collections.AuthorizeOTPModel{
		TransactionReference: "MNFY|05|20230409198910|002725",
		CollectionChannel:    "API_NOTIFICATION",
		TokenId:              "20.87-b66bef0aa8e660863c4e1177a08fefba",
		Token:                "otpToken",
	});
```

### Authorize 3D secure card

```go
monnifyPayment.Transaction.ThreeDsSecureAuthTransaction(collections.ThreeDsSecureAuthTransactionModel{
  TransactionReference: "transactionReference",
  CollectionChannel: "CollectionChannel",
  Card: collections.ThreeDsSecureAuthTransactionCard{
    Number: "39837982098238",
    ExpiryMonth: "05",
    ExpiryYear: "2028",
    CVV: "983",
  },
  ApiKey: "apiKey",
});
```

### Get Transaction Status v1

```go
monnifyPayment.Transaction.GetTransactionStatusv1(collections.GetTransactionStatusv1Model{
  PaymentReference: "PaymentReference",
});
```

### Get Transaction Status v2

```go
monnifyPayment.Transaction.GetTransactionStatusv2(collections.GetTransactionStatusv2Model{
  TransactionReference: "TransactionReference",
});
```

### Card tokenization

```go
monnifyPayment.Transaction.CardTokenization(collections.CardTokenizationModel{
    CustomerName:       "John Doe",
		CustomerEmail:      "example@gmail.com",
		Amount:             25.7,
		PaymentDescription: "payment description",
		PaymentReference:   "paymentReference",
		CurrencyCode:       "NGN",
		ContractCode:       "contract code",
		CardToken:          "card token",
		ApiKey:             "api key",
});
```

## Customer Reserved Account Service

Manages reserved account operations.

### All Available Methods

```go
// Create reserved account
monnifyPayment.ReservedAccount.CreateReservedAccount(body);

// Add linked accounts
monnifyPayment.ReservedAccount.AddLinkedAccounts(body);

// Remove account
monnifyPayment.ReservedAccount.DeallocateReservedAccount(body);

// Account details
monnifyPayment.ReservedAccount.ReservedAccountDetails(body);

// Account transactions
monnifyPayment.ReservedAccount.ReservedAccountTransactions(body);

// Update account kyc
monnifyPayment.ReservedAccount.UpdateReservedAccountKycInfo(body);
```

### Create Reserved Account

Creates a new reserved account.

```go
monnifyPayment.ReservedAccount.CreateReservedAccount(collections.ReservedAccountSchema{
		CustomerName:         "John Doe",
		CustomerEmail:        "customer email",
		AccountName:          "John Doe",
		AccountReference:     "accRef",
		ContractCode:         "contract code",
		Bvn:                  "bvn",
		GetAllAvailableBanks: true,
		RestrictPaymentSource: false,
	});
```

### Add linked accounts

Add linked accounts.

```go
monnifyPayment.ReservedAccount.AddLinkedAccounts(collections.AddLinkedAccountSchema{
		AccountReference:     "accRef",
		GetAllAvailableBanks: true,
	});
```

### Get Account Details

Get the full reserved account detail.

```go
monnifyPayment.ReservedAccount.ReservedAccountDetails(collections.ReservedAccountDetailsSchema{
		AccountReference: "accRef",
	});
```

### Deallocate Account

```go
monnifyPayment.ReservedAccount.DeallocateReservedAccount(collections.DeallocateReservedAccountSchema{
		AccountReference: "accRef",
	})
```

### Update the KYC for a reserved account.

```go
monnifyPayment.ReservedAccount.UpdateReservedAccountKycInfo(collections.UpdateReservedAccountKycInfoSchema{
		AccountReference: accRef,
		Bvn:              "82782672868273",
	});
```

## Disbursement Service

Handles money transfers and disbursements.

### All Available Methods

```go
// Single Transfers
monnifyPayment.Disbursement.InitiateSingleTransfer(body)

// Authorize single transfer
monnifyPayment.Disbursement.AuthorizeSingleTransfer(body)

// Get single transfer status
monnifyPayment.Disbursement.GetSingleTransferStatus(body)

// Bulk Transfer
monnifyPayment.Disbursement.InitiateBulkTransfer(body)

// Authorize bulk transfer
monnifyPayment.Disbursement.AuthorizeBulkTransfer(body)

// Get bulk status
monnifyPayment.Disbursement.GetBulkTransferStatus(body)

// Get All Single transactions
monnifyPayment.Disbursement.GetAllSingleTransfer(body)

// Get All Bulk transactions
monnifyPayment.Disbursement.GetAllBulkTransfer(body)

/** Other Operations **/

// Resend OTP
monnifyPayment.Disbursement.ResendTransferOTP(body)
```

### Single Transfer

Process a single money transfer.

```go
monnifyPayment.Disbursement.InitiateSingleTransfer(disbursement.SingleTransfer{
		Amount:                   20.5,
		Reference:                "879867856545687",
		Narration:                "narration",
		DestinationBankCode:      "058",
		DestinationAccountNumber: "89786798678879",
		Currency:                 "NGN",
		SourceAccountNumber:      "098787867675,
	})
```

### Bulk Transfer

Process multiple transfers at once.

```go
monnifyPayment.Disbursement.InitiateBulkTransfer(disbursement.BulkTransfer{
		Title:                "Bulk Transfer",
		BatchReference:       "878675656454565",
		SourceAccountNumber:  "8786754564345",
		Narration:            "Bulk Transfer Narration",
		OnValidationFailure:  "CONTINUE",
		NotificationInterval: 25,
		TransactionList: []bulkTransferTransactionList{
			{
				Amount:                   200.5,
				Reference:                "65654534534",
				Narration:                "narration",
				DestinationBankCode:      "058",
				DestinationAccountNumber: "54675675456",
				Currency:                 "NGN",
			},
		},
	})
```

### Authorize a single transfer with OTP.

```go
monnifyPayment.Disbursement.AuthorizeSingleTransfer(disbursement.AuthorizeTransfer{
		Reference:         "reference-98678578769756745",
		AuthorizationCode: "491763",
	})
```

### Authorize a bulk transfer with OTP.

```go
monnifyPayment.Disbursement.AuthorizeBulkTransfer(disbursement.AuthorizeTransfer{
		Reference:         "reference-98678578769756745",
		AuthorizationCode: "491763",
	})
```

### Check Single Transfer Status

```go
monnifyPayment.Disbursement.GetSingleTransferStatus(disbursement.GetStatus{
		Reference:         "reference-98678578769756745",
	})
```

### Check Bulk Transfer Status

```go
monnifyPayment.Disbursement.GetBulkTransferStatus(disbursement.GetBulkStatus{
		Reference:         "reference-98678578769756745",
    PageNo: 1,
    PageSize: 10
	})
```

### Resend Transfer OTP

```go
monnifyPayment.Disbursement.ResendTransferOTP(disbursement.ResendTransferOTP{
		Reference:         "reference-98678578769756745",
	})
```

### Get All Single Transfer

```go
monnifyPayment.Disbursement.GetAllSingleTransfer(disbursement.GetAllSingleTransfer{
		PageNo: 1,
    PageSize: 10
	})
```

### Get All Bulk Transfer

```go
monnifyPayment.Disbursement.GetAllBulkTransfer(disbursement.GetAllBulkTransfer{
		PageNo: 1,
    PageSize: 10
	})
```

## Verification Service

### All Available Methods

```go
// Verify account
monnifyPayment.verification.ValidateBankAccount(body);
// Verify BVN
monnifyPayment.verification.VerifyBvnInformation(body);
// Match BVN
monnifyPayment.verification.MatchBvnAndAccountName(body);
```

### Verify Bank Account

```go
monnifyPayment.verification.ValidateBankAccount(verification.ValidateBankAccountModel{
  AccountNumber: "8789678675764",
  BankCode: "058"
})
```

### Verify BVN Information

```go
monnifyPayment.verification.VerifyBvnInformation(verification.VerifyBvnInformationModel{
  BVN: "8789678675764",
  DateOfBirth: "05-05-2025"
  MobileNo: "09878676567"
  Name: "John Doe"
})
```

### Match BVN with Bank Account

```go
monnifyPayment.verification.MatchBvnAndAccountName(verification.MatchBvnAndAccountNameModel{
  BVN: "8789678675764",
  AccountNumber: "8789678675764",
  BankCode: "058"
})
```

## Sub Account Service

Manages sub-accounts for split payments.

### All Available Methods

```go
// Create sub account
monnifyPayment.collections.CreateSubAccount(body)
// Get all sub accounts
monnifyPayment.collections.GetSubAccounts()
// Update sub account
monnifyPayment.collections.UpdateSubAccount(body)
// Delete sub account
monnifyPayment.collections.DeleteSubAccount(body)
```

### Create Sub Account

Creates a new sub-account for split payments.

```go
monnifyPayment.collections.CreateSubAccount([]collections.CreateSubAccountModel{
		{
			CurrencyCode:           "NGN",
			AccountNumber:          "0211319282",
			BankCode:               "058",
			Email:                  "example@gmail.com",
			DefaultSplitPercentage: 20.87,
		},
	})
```

### Get All Sub Accounts

Retrieves all sub-accounts associated with your contract.

```go
monnifyPayment.collections.GetSubAccounts()
```

### Update Sub Account

Updates an existing sub-account's details.

```go
monnifyPayment.collections.UpdateSubAccount(collections.UpdateSubAccountModel{
  CurrencyCode:           "NGN",
  AccountNumber:          "0211319282",
  BankCode:               "058",
  Email:                  "example@gmail.com",
  DefaultSplitPercentage: 20.87,
})
```

### Delete Sub Account

Removes a sub-account from your contract.

```go
monnifyPayment.collections.DeleteSubAccount(collections.DeleteSubAccountModel{
  SubAccountCode: "subAccountCode",
})
```

## Refund Service

### All Available Methods

```go
// Initialize a refund
monnifyPayment.disbursement.InitiateRefund(body)
// Get all refunds
monnifyPayment.disbursement.GetAllRefunds(body)
// Check refund status
monnifyPayment.disbursement.GetRefundStatus(body)
```

### Initialize Refund

Creates a new refund request.

```go
monnifyPayment.disbursement.InitiateRefund(disbursement.InitiateRefundModel{
  RefundReason: "reason",
  RefundAmount: 20.8,
  RefundReference: "reference-87687564535487",
  TransactionReference: "transaction-ref-978656745",
  CustomerNote: "customer note"
})
```

### Get Refund Status

Check the status of a specific refund.

```go
monnifyPayment.disbursement.GetRefundStatus(disbursement.GetRefundStatusModel{
  RefundReference: "reference-87687564535487",
})
```

### Get All Refunds

Retrieves all refunds with pagination.

```go
monnifyPayment.disbursement.GetAllRefunds(disbursement.GetAllRefundsModel{
  Page: 1,
  Size: 10
})
```

## Error Handling

> This package returns error in an organized structure. Below is the error structure

```go
  type Error Struct {
    Message  string
    Error    error
    Response struct {
      RequestSuccessful bool
      ResponseMessage   string
      ResponseCode      string
    }
  }
```

> Response can be nil when an error is thrown on code level but available when an error is thrown from API

## Testing

Run package tests with:

```bash
make test
```

## Contributing

- Fork repository
- Create feature/fix branch
- Submit Pull Request

## Credits

- [Franklin Isaiah](https://github.com/Cavdy)

## License

This package is licensed under the [MIT License](LICENSE.md).

## Support

For any support or security issues, please contact [integration-support@monnify.com](mailto:integration-support@monnify.com).
