package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/controllers"
)

func NewRouter(auth *controllers.AuthController, user *controllers.UserController, book *controllers.BookController, author *controllers.AuthorController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		// Аутентификация
		r.Post("/auth/login", auth.Login())
		r.Post("/auth/register", auth.Register())
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		// Пользователи
		r.Post("/users", user.CreateUser())
		r.Get("/users/{id}", user.GetUser())
		r.Put("/users", user.UpdateUser())
		r.Delete("/users/{id}", user.DeleteUser())
		r.Get("/users", user.ListUsers())

		// Книги
		r.Post("/api/book/take/{index}", book.TakeBook())
		r.Delete("/api/book/return/{index}", book.ReturnBook())
		r.Post("/api/book", book.AddBook())
		r.Get("/api/books", book.GetAllBooksHandler())
		r.Put("/api/books/{index}", book.UpdateBook())

		// Авторы
		r.Post("/api/authors", author.AddAuthor()) // Без скобок!
		r.Get("/api/authors", author.ListAuthors())
	})

	return r
}
