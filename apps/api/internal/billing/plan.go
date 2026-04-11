package billing

type Plan struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    int64    `json:"price"`
	Currency string   `json:"currency"`
	Interval string   `json:"interval"`
	Features []string `json:"features"`
}

var Plans = []Plan{
	{ID: "free", Name: "Free", Price: 0, Currency: "USD", Interval: "month", Features: []string{"5 projects", "1 user"}},
	{ID: "starter", Name: "Starter", Price: 2900, Currency: "USD", Interval: "month", Features: []string{"25 projects", "5 users", "Priority support"}},
	{ID: "pro", Name: "Pro", Price: 9900, Currency: "USD", Interval: "month", Features: []string{"Unlimited projects", "25 users", "API access", "SSO"}},
	{ID: "enterprise", Name: "Enterprise", Price: 29900, Currency: "USD", Interval: "month", Features: []string{"Unlimited everything", "Dedicated support", "SLA", "Custom integrations"}},
}

func GetPlan(id string) *Plan {
	for _, p := range Plans {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
