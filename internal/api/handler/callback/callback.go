package callback

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"payment-gateway/enum"
	"payment-gateway/internal/api/handler/callback/dto"
	"payment-gateway/internal/models"
	"payment-gateway/internal/services/transaction"
	"payment-gateway/internal/services/transaction/contract"
)

type Handler struct {
	Txn transaction.Transaction
}

// HandleDepositSuccess handles deposit success callback
func (h Handler) HandleDepositSuccess(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.DepositSuccessCallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	respStatusCode := http.StatusOK

	if err = h.Txn.UpdateStatus(ctx, &contract.UpdateStatusRequest{
		TransactionID: req.TransactionID,
		Status:        enum.TxnStatusSuccess,
	}); err != nil {
		log.Println(fmt.Sprintf("failed to update txn status: %v", err))
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(respStatusCode)

	response := models.APIResponse{
		StatusCode: respStatusCode,
		Message:    "success",
		Data:       nil,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println(fmt.Sprintf("error encoding json: %v", err))
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// HandleDepositFailure handles deposit failure callback
func (h Handler) HandleDepositFailure(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.DepositFailureCallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	respStatusCode := http.StatusOK

	if err = h.Txn.UpdateStatus(ctx, &contract.UpdateStatusRequest{
		TransactionID: req.TransactionID,
		Status:        enum.TxnStatusFailed,
	}); err != nil {
		log.Println(fmt.Sprintf("failed to update txn status: %v", err))
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(respStatusCode)

	response := models.APIResponse{
		StatusCode: respStatusCode,
		Message:    "success",
		Data:       nil,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println(fmt.Sprintf("error encoding json: %v", err))
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// HandleWithdrawalSuccess handles withdrawal success callback
func (h Handler) HandleWithdrawalSuccess(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.WithdrawalSuccessCallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	respStatusCode := http.StatusOK

	if err = h.Txn.UpdateStatus(ctx, &contract.UpdateStatusRequest{
		TransactionID: req.TransactionID,
		Status:        enum.TxnStatusSuccess,
	}); err != nil {
		log.Println(fmt.Sprintf("failed to update txn status: %v", err))
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(respStatusCode)

	response := models.APIResponse{
		StatusCode: respStatusCode,
		Message:    "success",
		Data:       nil,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println(fmt.Sprintf("error encoding json: %v", err))
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// HandleWithdrawalFailure handles withdrawal failure callback
func (h Handler) HandleWithdrawalFailure(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.WithdrawalFailureCallbackRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("invalid payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	respStatusCode := http.StatusOK

	if err = h.Txn.UpdateStatus(ctx, &contract.UpdateStatusRequest{
		TransactionID: req.TransactionID,
		Status:        enum.TxnStatusFailed,
	}); err != nil {
		log.Println(fmt.Sprintf("failed to update txn status: %v", err))
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(respStatusCode)

	response := models.APIResponse{
		StatusCode: respStatusCode,
		Message:    "success",
		Data:       nil,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Println(fmt.Sprintf("error encoding json: %v", err))
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
