package config

import (
	"log"
	"os"
)

type configuration struct {
	Database *database
}

type database struct {
	Username string
	Password string
	DBName   string
	Host     string
	Port     string
}

var (
	config *configuration
)

func Get() *configuration {
	return config
}

func Load() {
	log.Println("loading configuration values...")

	c := &configuration{
		Database: &database{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}

	config = c

	log.Println("configuration loaded!")
}
