package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Upmate/shipfast/apps/api/internal/model"
)

var users = []model.User{
	{ID: "usr_001", Email: "alice@example.com", Name: "Alice", Role: "admin", Active: true},
	{ID: "usr_002", Email: "bob@example.com", Name: "Bob", Role: "member", Active: true},
	{ID: "usr_003", Email: "carol@example.com", Name: "Carol", Role: "member", Active: false},
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, `{"error":"user not found"}`, http.StatusNotFound)
}
