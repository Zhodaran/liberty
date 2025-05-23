package controllers

import (
	"encoding/json"
	"net/http"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/facades"
)

// controllers/auth_controller.go
func Login(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userAuth entities.UserAuth

		if err := json.NewDecoder(r.Body).Decode(&userAuth); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		user, err := facade.Login(userAuth.Username, userAuth.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Здесь можно сгенерировать JWT токен или установить сессию
		// Пока просто возвращаем успешный ответ
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login successful",
			"user":    user,
		})
	}
}

func Register(facade *facades.LibraryFacade) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userAuth entities.UserAuth
		if err := json.NewDecoder(r.Body).Decode(&userAuth); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := facade.Register(userAuth); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User registered successfully",
		})
	}
}
