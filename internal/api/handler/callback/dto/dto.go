package dto

type DepositSuccessCallbackRequest struct {
	TransactionID int `json:"transaction_id"`
}

type DepositFailureCallbackRequest struct {
	TransactionID int    `json:"transaction_id"`
	Reason        string `json:"reason"`
}
