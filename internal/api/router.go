package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"payment-gateway/internal/api/handler"
	callbackHandler "payment-gateway/internal/api/handler/callback"
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

	publicRouter := router.PathPrefix("/public").Subrouter() // for routes that do not require authentication middleware

	callback := callbackHandler.Handler{
		Txn: transaction.New(),
	}
	publicRouter.Handle("/gateway/callback/deposit/success", http.HandlerFunc(callback.HandleDepositSuccess)).Methods("POST")
	publicRouter.Handle("/gateway/callback/deposit/failure", http.HandlerFunc(callback.HandleDepositFailure)).Methods("POST")
	//publicRouter.Handle("/gateway/callback/withdrawal/success", http.HandlerFunc()).Methods("POST")
	//publicRouter.Handle("/gateway/callback/withdrawal/failure", http.HandlerFunc()).Methods("POST")

	return router
}
