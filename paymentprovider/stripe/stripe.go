package stripe

import (
	"context"
	"payment-gateway/paymentprovider/contract"
)

type Stripe interface {
	Deposit(ctx context.Context, req *contract.DepositRequest) error
	Withdraw(ctx context.Context, req *contract.WithdrawRequest) error
}

type impl struct{}

func New() Stripe {
	return &impl{}
}

func (i impl) Deposit(ctx context.Context, req *contract.DepositRequest) error {
	// TODO: implement exponential backoff
	panic("implement me")
}

func (i impl) Withdraw(ctx context.Context, req *contract.WithdrawRequest) error {
	// TODO: implement exponential backoff
	panic("implement me")
}
