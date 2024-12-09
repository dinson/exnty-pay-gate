package db

import (
	"context"
	"fmt"
	"time"
)

func (i impl) CreateCountry(ctx context.Context, country Country) error {
	query := `INSERT INTO countries (name, code, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := i.db.QueryRow(query, country.Name, country.Code, time.Now(), time.Now()).Scan(&country.ID)
	if err != nil {
		return fmt.Errorf("failed to insert country: %v", err)
	}
	return nil
}

func (i impl) GetCountries(ctx context.Context) ([]Country, error) {
	rows, err := i.db.Query(`SELECT id, name, code, created_at, updated_at FROM countries`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %v", err)
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		var country Country
		if err := rows.Scan(&country.ID, &country.Name, &country.Code, &country.CreatedAt, &country.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan country: %v", err)
		}
		countries = append(countries, country)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return countries, nil
}

func (i impl) GetSupportedCountriesByGateway(ctx context.Context, gatewayID int) ([]Country, error) {
	query := `
		SELECT c.id AS country_id, c.name AS country_name
		FROM countries c
		JOIN gateway_countries gc ON c.id = gc.country_id
		WHERE gc.gateway_id = $1
		ORDER BY c.name
	`

	rows, err := i.db.Query(query, gatewayID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries for gateway %d: %v", gatewayID, err)
	}
	defer rows.Close()

	var countries []Country
	for rows.Next() {
		var country Country
		if err := rows.Scan(&country.ID, &country.Name); err != nil {
			return nil, fmt.Errorf("failed to scan country: %v", err)
		}
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over rows: %v", err)
	}

	return countries, nil
}
