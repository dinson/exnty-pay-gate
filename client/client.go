package client

import (
	"database/sql"
	"payment-gateway/config"
	"payment-gateway/db"
)

type clients struct {
	Database *database
}

type database struct {
	Client *sql.DB
}

var (
	client = &clients{}
)

func Get() *clients {
	return client
}

func Init() {
	initPostgresQLClient()
}

func initPostgresQLClient() {
	dbURL := "postgres://" + config.Get().Database.Username +
		":" + config.Get().Database.Password +
		"@" + config.Get().Database.Host +
		":" + config.Get().Database.Port +
		"/" + config.Get().Database.DBName +
		"?sslmode=disable"
	dbClient := db.InitializeDB(dbURL)

	client.Database = &database{
		Client: dbClient,
	}
}
