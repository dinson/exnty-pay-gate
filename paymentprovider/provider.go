package paymentprovider

import (
	"context"
	"payment-gateway/paymentprovider/contract"
	"payment-gateway/paymentprovider/link"
	"payment-gateway/paymentprovider/stripe"
)

type PaymentProvider interface {
	Deposit(ctx context.Context, req *contract.DepositRequest) error
	Withdraw(ctx context.Context, req *contract.WithdrawRequest) error
}

type impl struct {
	stripe stripe.Stripe
	link   link.Link
}

func New() PaymentProvider {
	return &impl{
		stripe: stripe.New(),
		link:   link.New(),
	}
}
