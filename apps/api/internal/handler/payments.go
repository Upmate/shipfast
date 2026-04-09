package handler

import (
	"encoding/json"
	"net/http"
)

func ListPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": "pay_001", "amount": 9900, "status": "completed"},
		{"id": "pay_002", "amount": 29900, "status": "completed"},
		{"id": "pay_003", "amount": 2900, "status": "pending"},
	})
}
