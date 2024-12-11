package dto

type DepositSuccessCallbackRequest struct {
	TransactionID int `json:"transaction_id" binding:"required"`
}

type DepositFailureCallbackRequest struct {
	TransactionID int    `json:"transaction_id" binding:"required"`
	Reason        string `json:"reason"`
}

type WithdrawalSuccessCallbackRequest struct {
	TransactionID int `json:"transaction_id" binding:"required"`
}

type WithdrawalFailureCallbackRequest struct {
	TransactionID int    `json:"transaction_id" binding:"required"`
	Reason        string `json:"reason"`
}
