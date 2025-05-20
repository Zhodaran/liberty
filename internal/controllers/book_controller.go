package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/facades"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/infrastructure/postgres"
)

func GetAllBooksHandler(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := facade.GetAllBooks()
		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(books)
	}
}

func UpdateBook(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем индекс книги из URL-параметров
		indexStr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "недопустимый индекс", http.StatusBadRequest)
			return
		}

		// Декодируем тело запроса в структуру updatedBook
		var updatedBook entities.Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, "недопустимый формат данных", http.StatusBadRequest)
			return
		}

		// Вызов метода обновления книги в репозитории
		if err := facade.UpdateBook(index, updatedBook); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Возвращаем обновленную книгу
		w.WriteHeader(http.StatusOK) // Устанавливаем статус 200 OK
		json.NewEncoder(w).Encode(updatedBook)
	}
}

func AddBook(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var addaderBook postgres.AddaderBookRequest
		if err := json.NewDecoder(r.Body).Decode(&addaderBook); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		// Вызов метода добавления книги в фасаде
		if err := facade.AddBook(addaderBook); err != nil {
			if err.Error() == "book already exists" {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "Failed to add book", http.StatusInternalServerError)
			}
			return
		}

		// Возвращаем сообщение об успешном добавлении книги
		json.NewEncoder(w).Encode(map[string]string{"message": "Book added successfully"})
	}
}

func TakeBook(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexStr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "invalid index", http.StatusBadRequest)
			return
		}

		var requestBody TakeBookRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if requestBody.Username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}

		book, err := facade.TakeBook(index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Book taken successfully",
			"book":    book,
		})
	}
}

func ReturnBook(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexStr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "invalid index", http.StatusBadRequest)
			return
		}

		var requestBody TakeBookRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if requestBody.Username == "" {
			http.Error(w, "Username is required", http.StatusBadRequest)
			return
		}

		if err := facade.ReturnBook(index); err != nil {
			if err.Error() == "book not found or already returned" {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "Failed to return book", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"message": "Book returned successfully"})
	}
}
