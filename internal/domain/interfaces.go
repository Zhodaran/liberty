package domain

import (
	"context"
	"sync"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/infrastructure/postgres"
)

type AuthRepository interface {
	Register(user entities.UserAuth) error
	Login(username, password string) (entities.UserAuth, error)
}
type BookRepository interface {
	GetAllBooks() ([]entities.Book, error)
	TakeBook(id int) (entities.Book, error)
	AddBook(addaderBook postgres.AddaderBookRequest) error
	UpdateBook(index int, book entities.Book) error
	ReturnBook(index int) error
}

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	GetByID(ctx context.Context, id string) (entities.User, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]entities.User, error)
}

type AuthorRepository interface {
	AddAuthor(author entities.Author) error
	ListAuthors() ([]entities.Author, error)
	GetAuthors() ([]entities.Author, error)
}

type Responder interface {
	OutputJSON(responseData interface{}) error
	ErrorUnauthorized(err error) error
	ErrorBadRequest(err error) error
	ErrorForbidden(err error) error
	ErrorInternal(err error) error
}

type Library struct {
	Authors []string
	mu      sync.RWMutex
}
