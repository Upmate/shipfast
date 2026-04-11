package billing

import "time"

type Customer struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	StripeID  string    `json:"stripe_id,omitempty"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCustomer(email, name string) *Customer {
	return &Customer{
		ID:        generateID(),
		Email:     email,
		Name:      name,
		Plan:      "free",
		CreatedAt: time.Now(),
	}
}

func (c *Customer) UpgradePlan(planID string) {
	c.Plan = planID
}
