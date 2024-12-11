package paymentprovider

import (
	"context"
	"payment-gateway/enum"
	"payment-gateway/errors"
	"payment-gateway/paymentprovider/contract"
)

func (i impl) Deposit(ctx context.Context, req *contract.DepositRequest) error {
	switch req.GatewayProvider {
	case enum.ProviderStripe:
		return i.stripe.Deposit(ctx, req)
	case enum.ProviderLink:
		return i.link.Deposit(ctx, req)
	default:
		return errors.ErrInvalidProvider
	}
}
