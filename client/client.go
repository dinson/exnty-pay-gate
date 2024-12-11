package client

import (
	"log"
	"payment-gateway/config"
	"payment-gateway/db"
)

type clients struct {
	DB db.DB
}

var (
	client = &clients{}
)

func Get() *clients {
	return client
}

func Init() {
	initPostgreSQLClient()
}

func initPostgreSQLClient() {
	log.Println("initializing postgresql client...")

	dbURL := "postgres://" + config.Get().Database.Username +
		":" + config.Get().Database.Password +
		"@" + config.Get().Database.Host +
		":" + config.Get().Database.Port +
		"/" + config.Get().Database.DBName +
		"?sslmode=disable"

	client.DB = db.Initialize(dbURL)

	log.Println("initialized postgresql client!")
}
