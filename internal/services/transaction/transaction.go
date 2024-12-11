package transaction

import (
	"context"
	"payment-gateway/client"
	"payment-gateway/db"
	"payment-gateway/internal/services/transaction/contract"
	"payment-gateway/paymentprovider"
)

type Transaction interface {
	Deposit(ctx context.Context, req *contract.DepositRequest) (*contract.DepositResponse, error)
	UpdateStatus(ctx context.Context, req *contract.UpdateStatusRequest) error
	Withdraw(ctx context.Context, req *contract.WithdrawRequest) (*contract.WithdrawResponse, error)
}

type impl struct {
	db              db.DB
	paymentProvider paymentprovider.PaymentProvider
}

func New() Transaction {
	return &impl{
		db:              client.Get().DB,
		paymentProvider: paymentprovider.New(),
	}
}
