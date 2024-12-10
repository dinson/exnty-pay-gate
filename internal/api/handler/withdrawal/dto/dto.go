package dto

type InitWithdrawalRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type InitWithdrawalResponse struct {
	Success bool `json:"success"`
}
