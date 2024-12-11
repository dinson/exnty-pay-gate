package link

import (
	"context"
	"log"
	"payment-gateway/paymentprovider/contract"
)

type Link interface {
	Deposit(ctx context.Context, req *contract.DepositRequest) error
	Withdraw(ctx context.Context, req *contract.WithdrawRequest) error
}

type impl struct{}

func New() Link {
	return &impl{}
}

func (i impl) Deposit(ctx context.Context, req *contract.DepositRequest) error {
	// TODO: implement exponential backoff
	log.Println("link deposit is not implemented")
	return nil
}

func (i impl) Withdraw(ctx context.Context, req *contract.WithdrawRequest) error {
	// TODO: implement exponential backoff
	log.Println("link withdrawal is not implemented")
	return nil
}
