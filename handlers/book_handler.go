package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var books = make(map[int]models.Book)
var bookNextID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	q := r.URL.Query()
	categoryFilter := strings.ToLower(q.Get("category"))
	authorIDFilter, _ := strconv.Atoi(q.Get("author_id"))

	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	list := []models.Book{}
	for _, b := range books {
		if authorIDFilter > 0 && b.AuthorID != authorIDFilter {
			continue
		}
		if categoryFilter != "" {
			cat, ok := categories[b.CategoryID]
			if !ok || strings.ToLower(cat.Name) != categoryFilter {
				continue
			}
		}
		list = append(list, b)
	}
	total := len(list)
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	json.NewEncoder(w).Encode(map[string]any{
		"page":  page,
		"books": list[start:end],
	})
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(book.Title) == "" {
		http.Error(w, "title is required", http.StatusUnprocessableEntity)
		return
	}
	if _, ok := authors[book.AuthorID]; !ok {
		http.Error(w, "author not found", http.StatusUnprocessableEntity)
		return
	}
	if _, ok := categories[book.CategoryID]; !ok {
		http.Error(w, "category not found", http.StatusUnprocessableEntity)
		return
	}
	if book.Price <= 0 {
		http.Error(w, "price must be greater than 0", http.StatusUnprocessableEntity)
		return
	}
	book.ID = bookNextID
	bookNextID++
	books[book.ID] = book
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}
	book, ok := books[id]
	if !ok {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}
	if _, ok := books[id]; !ok {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	var input models.Book
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(input.Title) == "" {
		http.Error(w, "title is required", http.StatusUnprocessableEntity)
		return
	}
	if _, ok := authors[input.AuthorID]; !ok {
		http.Error(w, "author not found", http.StatusUnprocessableEntity)
		return
	}
	if _, ok := categories[input.CategoryID]; !ok {
		http.Error(w, "category not found", http.StatusUnprocessableEntity)
		return
	}
	if input.Price <= 0 {
		http.Error(w, "price must be greater than 0", http.StatusUnprocessableEntity)
		return
	}
	input.ID = id
	books[id] = input
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books[id])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}
	if _, ok := books[id]; !ok {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	delete(books, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "book deleted successfully"})
}
