package billing

import "time"

type Refund struct {
	ID        string    `json:"id"`
	PaymentID string    `json:"payment_id"`
	Amount    int64     `json:"amount"`
	Reason    string    `json:"reason"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewRefund(paymentID string, amount int64, reason string) *Refund {
	return &Refund{
		ID:        generateID(),
		PaymentID: paymentID,
		Amount:    amount,
		Reason:    reason,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
}

func (r *Refund) Approve() { r.Status = "approved" }
func (r *Refund) Reject()  { r.Status = "rejected" }
