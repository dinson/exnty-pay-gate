package db

import (
	"time"
)

type User struct {
	ID        int
	Username  string
	Email     string
	CountryID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Gateway struct {
	ID                  int
	Name                string
	DataFormatSupported string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Country struct {
	ID        int
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID        int
	Amount    float64
	Type      string // deposit or withdrawal
	Status    string // success, failed, cancelled, initialized
	UserID    int
	GatewayID int
	CountryID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GatewayPriorityConfig struct {
	ID        int
	GatewayID int
	CountryID int
	Priority  int // 1,2,3 and so on, lowest value means highest priority
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GatewaysForCountry struct {
	ID   int
	Name string
}
