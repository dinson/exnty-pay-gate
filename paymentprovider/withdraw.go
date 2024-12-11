package paymentprovider

import (
	"context"
	"payment-gateway/enum"
	"payment-gateway/errors"
	"payment-gateway/paymentprovider/contract"
)

func (i impl) Withdraw(ctx context.Context, req *contract.WithdrawRequest) error {
	switch req.GatewayProvider {
	case enum.ProviderStripe:
		return i.stripe.Withdraw(ctx, req)
	case enum.ProviderLink:
		return i.link.Withdraw(ctx, req)
	default:
		return errors.ErrInvalidProvider
	}
}
