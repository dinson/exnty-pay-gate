package api

import (
	"github.com/gorilla/mux"
	"net/http"
	callbackHandler "payment-gateway/internal/api/handler/callback"
	depositHandler "payment-gateway/internal/api/handler/deposit"
	"payment-gateway/internal/api/handler/withdrawal"
	"payment-gateway/internal/middleware"
	"payment-gateway/internal/services/gateway"
	"payment-gateway/internal/services/transaction"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	initProtectedRoutes(router)
	initPublicRoutes(router)

	return router
}

func initProtectedRoutes(router *mux.Router) {
	m := middleware.New()

	router.Use(m.VerifyAuthToken)

	deposit := depositHandler.Handler{
		Gateway: gateway.New(),
		Txn:     transaction.New(),
	}
	router.Handle("/deposit", http.HandlerFunc(deposit.InitDeposit)).Methods("POST")

	withdraw := withdrawal.Handler{
		Gateway: gateway.New(),
		Txn:     transaction.New(),
	}
	router.Handle("/withdrawal", http.HandlerFunc(withdraw.InitWithdrawal)).Methods("POST")
}

func initPublicRoutes(router *mux.Router) {
	publicRouter := router.PathPrefix("/public").Subrouter() // for routes that do not require authentication middleware

	callback := callbackHandler.Handler{
		Txn: transaction.New(),
	}
	publicRouter.Handle("/gateway/callback/deposit/success", http.HandlerFunc(callback.HandleDepositSuccess)).Methods("POST")
	publicRouter.Handle("/gateway/callback/deposit/failure", http.HandlerFunc(callback.HandleDepositFailure)).Methods("POST")
	publicRouter.Handle("/gateway/callback/withdrawal/success", http.HandlerFunc(callback.HandleWithdrawalSuccess)).Methods("POST")
	publicRouter.Handle("/gateway/callback/withdrawal/failure", http.HandlerFunc(callback.HandleWithdrawalFailure)).Methods("POST")
}
