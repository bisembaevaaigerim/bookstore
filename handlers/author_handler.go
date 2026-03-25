package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strings"
)

var authors = make(map[int]models.Author)
var authorNextID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorsList := []models.Author{}
	for _, author := range authors {
		authorsList = append(authorsList, author)
	}
	json.NewEncoder(w).Encode(authorsList)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(author.Name) == "" {
		http.Error(w, "name is required", http.StatusUnprocessableEntity)
		return
	}
	author.ID = authorNextID
	authorNextID++
	authors[author.ID] = author
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
