package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Upmate/shipfast/apps/api/internal/model"
)

var products = []model.Product{
	{ID: "1", Name: "Starter Plan", Description: "For small teams", Price: 2900, Currency: "USD"},
	{ID: "2", Name: "Pro Plan", Description: "For growing companies", Price: 9900, Currency: "USD"},
	{ID: "3", Name: "Enterprise", Description: "Custom solutions", Price: 29900, Currency: "USD"},
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, `{"error":"product not found"}`, http.StatusNotFound)
}
