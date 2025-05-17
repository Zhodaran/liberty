package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

// controllers/auth_controller.go
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	user, err := c.facade.Login(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	claims := map[string]interface{}{
		"user_id": user.Username,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	_, tokenString, err := TokenAuth.Encode(claims)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+tokenString)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user entities.UserAuth
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := c.facade.Register(user); err != nil {
		if err.Error() == "user already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, "Registration failed", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}
