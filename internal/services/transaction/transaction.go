package transaction

import (
	"context"
	"payment-gateway/client"
	"payment-gateway/db"
	"payment-gateway/internal/services/transaction/contract"
)

type Transaction interface {
	Deposit(ctx context.Context, req *contract.DepositRequest) (*contract.DepositResponse, error)
	UpdateStatus(ctx context.Context, req *contract.UpdateStatusRequest) error
	Withdraw(ctx context.Context, req *contract.WithdrawRequest) (*contract.WithdrawResponse, error)
}

type impl struct {
	db db.DB
}

func New() Transaction {
	return &impl{
		db: client.Get().DB,
	}
}
