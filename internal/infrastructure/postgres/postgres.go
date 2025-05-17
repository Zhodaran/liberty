package postgres

import (
	"database/sql"
	"sync"

	"github.com/go-chi/jwtauth"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{Db: db}
}

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
	return &PostgresBookRepository{db: db}
}

type PostgresUserRepository struct {
	Db *sql.DB
}

func NewPostgresAuthorRepository(db *sql.DB) *PostgresAuthorRepository {
	return &PostgresAuthorRepository{db: db}
}

type PostgresAuthorRepository struct {
	db *sql.DB
}

func NewPostgresAuthRepository(db *sql.DB) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

type PostgresAuthRepository struct {
	db *sql.DB
}

type TokenResponse struct {
	Token string `json:"token"`
}

var (
	TokenAuth = jwtauth.New("HS256", []byte("your_secret_key"), nil)
	mu        sync.Mutex
)

type AuthorRequest struct {
	Name string `json:"name"`
}

type TakeBookRequest struct {
	Username string `json:"username"` // Поле для имени пользователя
}

type AddaderBookRequest struct {
	Book   string `json:"book"`
	Author string `json:"author"`
}

type AddBooksRequest struct {
	Books []AddaderBookRequest `json:"books"` // Массив книг, который мы ожидаем в запросе
}

type CreateResponse struct {
	Message string          `json:"message"`
	Books   []entities.Book `json:"books"` // Добавляем поле для списка книг
}

type BookController struct {
	repo *PostgresBookRepository
}

func NewBookController(repo *PostgresBookRepository) *BookController {
	return &BookController{repo: repo}
}
