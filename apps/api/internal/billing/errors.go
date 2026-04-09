package billing

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrInvalidAmount       = errors.New("invalid payment amount")
	ErrSubscriptionExpired = errors.New("subscription expired")
	ErrPlanNotFound        = errors.New("plan not found")
	ErrPaymentFailed       = errors.New("payment processing failed")
	ErrRefundExceedsAmount = errors.New("refund exceeds payment amount")
	ErrDiscountExpired     = errors.New("discount code expired")
	ErrInvalidCurrency     = errors.New("unsupported currency")
)

func generateID() string {
	b := make([]byte, 12)
	rand.Read(b)
	return hex.EncodeToString(b)
}
