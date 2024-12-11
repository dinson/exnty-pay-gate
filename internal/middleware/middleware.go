package middleware

import (
	"net/http"
	"payment-gateway/client"
	"payment-gateway/db"
)

type Middleware interface {
	VerifyAuthToken(next http.Handler) http.Handler
}

type impl struct {
	db db.DB
}

func New() Middleware {
	return &impl{
		db: client.Get().DB,
	}
}
