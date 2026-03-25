package main

import (
	"bookstore/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")

	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "Bookstore API is running"}`)
	}).Methods("GET")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
