package dto

type InitDepositRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type InitDepositResponse struct {
	Success       bool   `json:"success"`
	Gateway       string `json:"gateway"`
	TransactionID int    `json:"transactionID"`
	Currency      string `json:"currency"`
	UserID        int    `json:"userID"`
}
