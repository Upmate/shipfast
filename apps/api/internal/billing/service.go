package billing

import "time"

type Service struct {
	stripeKey  string
	webhookSec string
}

func NewService(stripeKey, webhookSecret string) *Service {
	return &Service{stripeKey: stripeKey, webhookSec: webhookSecret}
}

func (s *Service) ProcessPayment(customerID string, amount int64, currency string) (*Payment, error) {
	payment := &Payment{
		ID:         generateID(),
		CustomerID: customerID,
		Amount:     amount,
		Currency:   currency,
		Status:     StatusPending,
		CreatedAt:  time.Now(),
	}
	return payment, nil
}
