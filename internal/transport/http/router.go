package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/controllers"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/facades"
)

func NewRouter(facade *facades.LibraryFacade) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		// Аутентификация
		r.Post("/api/login", controllers.Login(facade))
		r.Post("/api/register", controllers.Register(facade))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)

		// Пользователи
		r.Post("/users", controllers.CreateUser(facade))
		r.Get("/users/{id}", controllers.GetUser(facade))
		r.Put("/users", controllers.UpdateUser(facade))
		r.Delete("/users/{id}", controllers.DeleteUser(facade))
		r.Get("/users", controllers.ListUsers(facade))

		// Книги
		r.Post("/api/book/take/{index}", controllers.TakeBook(facade))
		r.Delete("/api/book/return/{index}", controllers.ReturnBook(facade))
		r.Post("/api/book", controllers.AddBook(facade))
		r.Get("/api/books", controllers.GetAllBooksHandler(facade))
		r.Put("/api/books/{index}", controllers.UpdateBook(facade))

		// Авторы
		r.Post("/api/authors", controllers.AddAuthor(facade)) // Без скобок!
		r.Get("/api/authors", controllers.ListAuthors(facade))
	})

	return r
}
