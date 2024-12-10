package contract

import (
	"payment-gateway/enum"
)

type DepositRequest struct {
	UserID    int
	Amount    float64
	Currency  string
	GatewayID int
	CountryID int
}

type DepositResponse struct {
	TransactionID int
	Success       bool
}

type UpdateStatusRequest struct {
	TransactionID int
	Status        enum.TxnStatus
}

type WithdrawRequest struct {
	UserID    int
	Amount    float64
	Currency  string
	GatewayID int
	CountryID int
}

type WithdrawResponse struct {
	TransactionID int
	Success       bool
}
