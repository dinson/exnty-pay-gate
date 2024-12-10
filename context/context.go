package context

import (
	"context"
	"payment-gateway/constant"
)

// GetUserID from context
func GetUserID(ctx context.Context) int {
	return ctx.Value(constant.UserID).(int)
}

// GetCountryID from context
func GetCountryID(ctx context.Context) int {
	return ctx.Value(constant.CountryID).(int)
}
