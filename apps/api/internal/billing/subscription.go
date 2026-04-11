package billing

import "time"

type Subscription struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	PlanID     string    `json:"plan_id"`
	Status     string    `json:"status"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewSubscription(customerID, planID string) *Subscription {
	return &Subscription{
		ID:         generateID(),
		CustomerID: customerID,
		PlanID:     planID,
		Status:     "active",
		StartDate:  time.Now(),
		CreatedAt:  time.Now(),
	}
}

func (s *Subscription) Cancel() {
	s.Status = "canceled"
	s.EndDate = time.Now()
}

func (s *Subscription) IsActive() bool {
	return s.Status == "active"
}
