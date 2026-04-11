package handler

import (
	"encoding/json"
	"net/http"
)

func ListInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": "inv_001", "amount": 9900, "status": "paid"},
		{"id": "inv_002", "amount": 29900, "status": "pending"},
	})
}
