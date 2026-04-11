package handler

import (
	"encoding/json"
	"net/http"
)

func ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": "sub_001", "plan": "pro", "status": "active"},
		{"id": "sub_002", "plan": "starter", "status": "active"},
	})
}
