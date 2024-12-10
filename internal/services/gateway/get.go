package gateway

import (
	"context"
	"fmt"
	"log"
	"payment-gateway/db"
	"payment-gateway/internal/services/gateway/contract"
)

// GetByCountry
// retrieve payment gateways for a country in the ascending order of their priority
func (i impl) GetByCountry(ctx context.Context, req *contract.GetGatewayByCountryRequest) ([]*contract.Gateway, error) {
	res, err := i.db.ListCountryGatewaysByPriority(ctx, req.CountryID)
	if err != nil {
		log.Println(fmt.Sprintf("failed to retrieve country specific gateways: %v countryID: %d\n", err, req.CountryID))
		return nil, err
	}

	return mapToContractGateway(res), nil
}

func mapToContractGateway(res []*db.GatewaysForCountry) []*contract.Gateway {
	var gs []*contract.Gateway

	for _, g := range res {
		gs = append(gs, &contract.Gateway{
			ID:   g.ID,
			Name: g.Name,
		})
	}

	return gs
}
