package middleware

import (
	"context"
	"net/http"
	"payment-gateway/constant"
)

type Middleware interface {
	VerifyAuthToken(next http.Handler) http.Handler
}

type impl struct {
}

func New() Middleware {
	return &impl{}
}

func (i impl) VerifyAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with a value
		// TODO: resolve JWT token and extract user ID
		// Extract User country and store to context
		ctx := context.WithValue(r.Context(), constant.UserID, 1)
		ctx = context.WithValue(r.Context(), constant.CountryID, 1)

		// Pass the updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
