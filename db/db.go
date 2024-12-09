package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"payment-gateway/internal/services"
)

type DB interface {
	CreateUser(ctx context.Context, user User) error
	GetUsers(ctx context.Context) ([]User, error)
	CreateGateway(ctx context.Context, gateway Gateway) error
	GetGateways(ctx context.Context) ([]Gateway, error)
	CreateCountry(ctx context.Context, country Country) error
	GetCountries(ctx context.Context) ([]Country, error)
	CreateTransaction(ctx context.Context, transaction Transaction) error
	GetTransactions(ctx context.Context) ([]Transaction, error)
	GetSupportedCountriesByGateway(ctx context.Context, gatewayID int) ([]Country, error)
}

type impl struct {
	db *sql.DB
}

// Initialize initializes the database connection
func Initialize(dbURL string) DB {
	var db *sql.DB
	var err error

	err = services.RetryOperation(func() error {
		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			return err
		}

		return db.Ping()
	}, 5)

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	return &impl{
		db: db,
	}
}
