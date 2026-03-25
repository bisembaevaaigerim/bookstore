package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strings"
)

var categories = make(map[int]models.Category)
var categoryNextID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := []models.Category{}
	for _, c := range categories {
		list = append(list, c)
	}
	json.NewEncoder(w).Encode(list)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(category.Name) == "" {
		http.Error(w, "name is required", http.StatusUnprocessableEntity)
		return
	}
	category.ID = categoryNextID
	categoryNextID++
	categories[category.ID] = category
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
