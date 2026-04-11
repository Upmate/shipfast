package handler

import (
	"encoding/json"
	"net/http"
)

func ListRefunds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": "ref_001", "payment_id": "pay_001", "amount": 9900, "status": "approved"},
	})
}
