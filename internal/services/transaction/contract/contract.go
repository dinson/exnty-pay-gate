package contract

type DepositRequest struct {
	UserID    int
	Amount    float64
	Currency  string
	GatewayID int
	CountryID int
}

type DepositResponse struct {
	TransactionID int
	Success       bool
}
