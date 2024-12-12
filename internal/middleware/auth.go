package middleware

import (
	"context"
	"net/http"
	"payment-gateway/constant"
)

func (i impl) VerifyAuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: resolve JWT token and extract user ID
		// TODO: get User country and email from DB and store to context

		ctx := context.WithValue(r.Context(), constant.UserID, 1)
		ctx = context.WithValue(r.Context(), constant.UserEmail, "")
		ctx = context.WithValue(r.Context(), constant.CountryID, 1)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
