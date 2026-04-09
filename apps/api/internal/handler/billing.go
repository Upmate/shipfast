package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Upmate/shipfast/apps/api/internal/billing"
)

func ListPlans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(billing.Plans)
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := billing.Metrics{
		TotalRevenue:     1250000,
		MonthlyRecurring: 104166,
		ActiveSubs:       47,
		ChurnRate:        2.3,
		ARPU:             9900,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
