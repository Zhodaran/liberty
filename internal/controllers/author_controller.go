package controllers

import (
	"encoding/json"
	"net/http"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func (c *AuthorController) AddAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author entities.Author
		if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := c.facade.AddAuthor(author); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Author added successfully"})
	}
}

func (c *AuthorController) ListAuthors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := c.facade.ListAuthors()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(authors)
	}
}
