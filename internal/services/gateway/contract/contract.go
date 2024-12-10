package contract

type GetGatewayByCountryRequest struct {
	CountryID int
}

type Gateway struct {
	ID         int
	Name       string
	DataFormat string
}
