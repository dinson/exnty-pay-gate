package db

import (
	"context"
	"fmt"
	"time"
)

func (i impl) CreateGateway(ctx context.Context, gateway Gateway) error {
	query := `INSERT INTO gateways (name, data_format_supported, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := i.db.QueryRow(query, gateway.Name, gateway.DataFormatSupported, time.Now(), time.Now()).Scan(&gateway.ID)
	if err != nil {
		return fmt.Errorf("failed to insert gateway: %v", err)
	}
	return nil
}

func (i impl) GetGateways(ctx context.Context) ([]Gateway, error) {
	rows, err := i.db.Query(`SELECT id, name, data_format_supported, created_at, updated_at FROM gateways`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gateways: %v", err)
	}
	defer rows.Close()

	var gateways []Gateway
	for rows.Next() {
		var gateway Gateway
		if err := rows.Scan(&gateway.ID, &gateway.Name, &gateway.DataFormatSupported, &gateway.CreatedAt, &gateway.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan gateway: %v", err)
		}
		gateways = append(gateways, gateway)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return gateways, nil
}

func (i impl) ListCountryGatewaysByPriority(ctx context.Context, countryID int) ([]*GatewaysForCountry, error) {
	query := `SELECT * FROM gateway_priority where country_id = $1 ORDER BY priority`

	rows, err := i.db.QueryContext(ctx, query, countryID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gateways for country %d: %v", countryID, err)
	}
	defer rows.Close()

	var gateways []*GatewaysForCountry
	for rows.Next() {
		var gateway GatewaysForCountry
		if err := rows.Scan(&gateway.ID, &gateway.Name); err != nil {
			return nil, fmt.Errorf("failed to scan gateway_priority: %v for country: %d", err, countryID)
		}
		gateways = append(gateways, &gateway)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v for country: %d", err, countryID)
	}

	return gateways, nil
}
