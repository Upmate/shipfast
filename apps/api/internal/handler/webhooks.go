package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Upmate/shipfast/apps/api/internal/billing"
)

func HandleStripeWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"failed to read body"}`, http.StatusBadRequest)
		return
	}

	event := billing.NewWebhookEvent("stripe", string(body))
	event.MarkProcessed()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
