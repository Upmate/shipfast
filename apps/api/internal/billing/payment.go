package billing

import "time"

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusRefunded  = "refunded"
)

type Payment struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	Amount      int64     `json:"amount"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	Method      string    `json:"method"`
	ExternalID  string    `json:"external_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

func (p *Payment) Complete(externalID string) {
	p.Status = StatusCompleted
	p.ExternalID = externalID
	p.CompletedAt = time.Now()
}

func (p *Payment) Fail() {
	p.Status = StatusFailed
}
