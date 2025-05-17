package controllers

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func GenerateUsers(count int) {
	for i := 0; i < count; i++ {
		username := gofakeit.Username()                                   // Генерация случайного имени пользователя
		password := gofakeit.Password(true, true, true, false, false, 10) // Генерация случайного пароля

		Users[username] = entities.UserAuth{
			Username: username,
			Password: password,
		}
		fmt.Printf("Created user: %s with password: %s\n", username, password)
	}
}

func (l *Library) AddBooks(books []entities.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()

}
