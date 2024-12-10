package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"payment-gateway/internal/api/handler"
	depositHandler "payment-gateway/internal/api/handler/deposit"
	"payment-gateway/internal/api/middleware"
	"payment-gateway/internal/services/gateway"
	"payment-gateway/internal/services/transaction"
)

func SetupRouter() *mux.Router {
	m := middleware.New()

	router := mux.NewRouter()
	router.Use(m.VerifyAuthToken)

	deposit := depositHandler.Handler{
		Gateway: gateway.New(),
		Txn:     transaction.New(),
	}

	router.Handle("/deposit", http.HandlerFunc(deposit.InitDeposit)).Methods("POST")
	router.Handle("/withdrawal", http.HandlerFunc(handler.WithdrawalHandler)).Methods("POST")

	// gateway callback routes
	//router.Handle("/gateway/callback/deposit/success", http.HandlerFunc()).Methods("POST")
	//router.Handle("/gateway/callback/deposit/failure", http.HandlerFunc()).Methods("POST")
	//router.Handle("/gateway/callback/withdrawal/success", http.HandlerFunc()).Methods("POST")
	//router.Handle("/gateway/callback/withdrawal/failure", http.HandlerFunc()).Methods("POST")

	return router
}
