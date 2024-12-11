package contract

import (
	"payment-gateway/enum"
)

type DepositRequest struct {
	Email           string
	Amount          float64
	Currency        string
	GatewayProvider enum.Provider
}

type WithdrawRequest struct {
	Email           string
	Amount          float64
	Currency        string
	GatewayProvider enum.Provider
}
