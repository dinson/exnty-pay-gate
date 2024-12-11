package stripe

import (
	"context"
	"log"
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
	log.Println("stripe deposit is not implemented")
	return nil
}

func (i impl) Withdraw(ctx context.Context, req *contract.WithdrawRequest) error {
	// TODO: implement exponential backoff
	log.Println("stripe withdrawal is not implemented")
	return nil
}
