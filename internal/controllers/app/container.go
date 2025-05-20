package app

import (
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/facades"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/infrastructure/postgres"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/transport/http"
)

type Container struct {
	Router *chi.Mux
}

func BuildContainer(logger *zap.Logger) (*Container, error) {
	// Добавляем задержку для инициализации БД в Docker
	time.Sleep(5 * time.Second)

	// Инициализация БД
	db, err := postgres.InitDB()
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return nil, err
	}

	// Проверка соединения с БД
	if err := db.Ping(); err != nil {
		logger.Error("Database connection failed", zap.Error(err))
		return nil, err
	}

	// Применение миграций
	postgres.RunMigrations(db)

	// Инициализация репозиториев
	authRepo := postgres.NewPostgresAuthRepository(db)
	bookRepo := postgres.NewPostgresBookRepository(db)
	authorRepo := postgres.NewPostgresAuthorRepository(db)
	userRepo := postgres.NewPostgresUserRepository(db)

	// Создание фасада
	library := facades.NewLibraryFacade(authRepo, bookRepo, authorRepo, userRepo)

	// Инициализация контроллеров

	// Создание роутера
	r := http.NewRouter(library)

	return &Container{Router: r}, nil
}
