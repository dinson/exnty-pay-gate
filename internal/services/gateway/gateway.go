package gateway

import (
	"context"
	"payment-gateway/client"
	"payment-gateway/db"
	"payment-gateway/internal/services/gateway/contract"
)

type Gateway interface {
	// GetByCountry
	// retrieve payment gateways for a country in the ascending order of their priority
	GetByCountry(ctx context.Context, req *contract.GetGatewayByCountryRequest) ([]*contract.Gateway, error)
}

type impl struct {
	db db.DB
}

func New() Gateway {
	return &impl{
		db: client.Get().DB,
	}
}
