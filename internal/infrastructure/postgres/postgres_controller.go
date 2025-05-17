package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func PullSQL() {

	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("Запуск задержки")
	time.Sleep(10 * time.Second)

}

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Несколько попыток подключения
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Printf("Attempt %d: DB connection failed, retrying...", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %v", err)
}

func RunMigrations(db *sql.DB) error {
	// Проверяем существование таблицы books как индикатора применённых миграций
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_name = 'books'
    )`).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check books table: %v", err)
	}

	if exists {
		log.Println("Tables already exist, skipping migrations")
		return nil
	}

	log.Println("Applying database migrations...")
	files := []string{
		"migrations/000001_create_books_table.up.sql",
		"migrations/20231001_create_authors_table.sql",
		"migrations/000002_create_users_table.up.sql", // Добавляем новую миграцию
	}

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %v", file, err)
		}

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", file, err)
		}
		log.Printf("Applied migration: %s", file)
	}
	return nil
}

// InsertFakeBooks вставляет 100 фейковых книг в таблицу book
func InsertFakeBooks(db *sql.DB) {
	var authors []string
	for i := 0; i < 10; i++ {
		author := gofakeit.Name()
		authors = append(authors, author)
	}

	// Вставка книг в базу данных
	for i := 1; i < 101; i++ {
		block := false
		book := entities.Book{
			Index:     i,
			Title:     gofakeit.Sentence(1),                        // Генерация названия книги
			Author:    authors[gofakeit.Number(0, len(authors)-1)], // Случайный автор
			Block:     &block,                                      // Устанавливаем значение блокировки
			TakeCount: 0,                                           // Начальное значение take_count
		}

		_, err := db.Exec("INSERT INTO book (book, author, block) VALUES ($1, $2, $3)", book.Title, book.Author, book.Block)
		if err != nil {
			log.Fatalf("Error inserting book: %v", err)
		}
	}
}
