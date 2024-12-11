package dto

type InitDepositRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type InitDepositResponse struct {
	Success       bool   `json:"success"`
	Gateway       string `json:"gateway"`
	TransactionID int    `json:"transaction_id"`
}
