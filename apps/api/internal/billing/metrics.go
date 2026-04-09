package billing

type Metrics struct {
	TotalRevenue     int64   `json:"total_revenue"`
	MonthlyRecurring int64   `json:"monthly_recurring"`
	ActiveSubs       int     `json:"active_subscriptions"`
	ChurnRate        float64 `json:"churn_rate"`
	ARPU             int64   `json:"arpu"`
}

func CalculateMetrics(customers []Customer, subscriptions []Subscription, payments []Payment) *Metrics {
	var totalRevenue int64
	var activeSubs int

	for _, p := range payments {
		if p.Status == StatusCompleted {
			totalRevenue += p.Amount
		}
	}

	for _, s := range subscriptions {
		if s.IsActive() {
			activeSubs++
		}
	}

	var arpu int64
	if len(customers) > 0 {
		arpu = totalRevenue / int64(len(customers))
	}

	return &Metrics{
		TotalRevenue:     totalRevenue,
		MonthlyRecurring: totalRevenue / 12,
		ActiveSubs:       activeSubs,
		ARPU:             arpu,
	}
}
