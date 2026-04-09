package billing

import "time"

type Invoice struct {
	ID        string     `json:"id"`
	CustomerID string    `json:"customer_id"`
	Amount    int64      `json:"amount"`
	Currency  string     `json:"currency"`
	Status    string     `json:"status"`
	DueDate   time.Time  `json:"due_date"`
	PaidAt    time.Time  `json:"paid_at,omitempty"`
	LineItems []LineItem `json:"line_items"`
	CreatedAt time.Time  `json:"created_at"`
}

type LineItem struct {
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	UnitPrice   int64  `json:"unit_price"`
	Total       int64  `json:"total"`
}

func NewInvoice(customerID string, items []LineItem) *Invoice {
	var total int64
	for _, item := range items {
		total += item.Total
	}
	return &Invoice{
		ID:         generateID(),
		CustomerID: customerID,
		Amount:     total,
		Currency:   "USD",
		Status:     "draft",
		LineItems:  items,
		CreatedAt:  time.Now(),
	}
}
