package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Upmate/shipfast/apps/api/internal/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, `{"error":"email and password required"}`, http.StatusBadRequest)
		return
	}

	user := model.User{
		ID:        "usr_" + time.Now().Format("20060102150405"),
		Email:     req.Email,
		Name:      req.Name,
		Role:      "member",
		Active:    true,
		CreatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
