package billing

import "time"

type Receipt struct {
	ID        string    `json:"id"`
	InvoiceID string    `json:"invoice_id"`
	PaymentID string    `json:"payment_id"`
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	IssuedAt  time.Time `json:"issued_at"`
}

func NewReceipt(invoiceID, paymentID string, amount int64, currency string) *Receipt {
	return &Receipt{
		ID:        generateID(),
		InvoiceID: invoiceID,
		PaymentID: paymentID,
		Amount:    amount,
		Currency:  currency,
		IssuedAt:  time.Now(),
	}
}
