package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/controllers/app"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	// Сборка контейнера зависимостей
	container, err := app.BuildContainer(logger)
	if err != nil {
		logger.Fatal("Failed to build container", zap.Error(err))
	}

	// Настройка сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      container.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запуск сервера
	go func() {
		logger.Info("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	// Ожидание graceful shutdown
	gracefulShutdown(server, logger)
}

func gracefulShutdown(server *http.Server, logger *zap.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Graceful shutdown failed", zap.Error(err))
	} else {
		logger.Info("Server stopped gracefully")
	}
}
