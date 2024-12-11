package deposit

import (
	"encoding/json"
	"log"
	"net/http"
	"payment-gateway/context"
	"payment-gateway/enum"
	"payment-gateway/internal/api/handler/deposit/dto"
	"payment-gateway/internal/models"
	"payment-gateway/internal/services/gateway"
	"payment-gateway/internal/services/gateway/contract"
	"payment-gateway/internal/services/transaction"
	txnContract "payment-gateway/internal/services/transaction/contract"
)

type Handler struct {
	Gateway gateway.Gateway
	Txn     transaction.Transaction
}

// InitDeposit handles deposit requests
func (h Handler) InitDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := context.GetUserID(ctx)
	userCountryID := context.GetCountryID(ctx)
	var req dto.InitDepositRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// retrieve all payment gateways based on country and ascending order of priority
	gateways, err := h.Gateway.GetByCountry(ctx, &contract.GetGatewayByCountryRequest{
		CountryID: userCountryID,
	})
	if err != nil {
		log.Println("failed to retrieve payment gateways: ", err.Error())
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if len(gateways) == 0 {
		log.Println("no payment gateways configured for country: ", userCountryID)
		http.Error(w, "no gateway support your country", http.StatusForbidden)
		return
	}

	respData := &dto.InitDepositResponse{
		Success: false, // default
	}
	respStatusCode := http.StatusInternalServerError // default

	for _, g := range gateways {
		txnResp, err := h.Txn.Deposit(ctx, &txnContract.DepositRequest{
			UserID:          userID,
			GatewayID:       g.ID,
			GatewayProvider: enum.Provider(g.Name),
			Amount:          req.Amount,
			Currency:        req.Currency,
		})
		if err != nil {
			log.Println("failed to perform deposit: ", err.Error())
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if txnResp.Success {
			log.Println("deposit txn successful: ", txnResp.TransactionID)

			respData.Success = true
			respData.Gateway = g.Name
			respData.TransactionID = txnResp.TransactionID

			respStatusCode = http.StatusOK
			break
		}
	}

	w.WriteHeader(respStatusCode)

	response := models.APIResponse{
		StatusCode: respStatusCode,
		Message:    "success",
		Data:       respData,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
