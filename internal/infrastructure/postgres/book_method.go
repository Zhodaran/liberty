package postgres

import (
	"errors"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

// @Summary Get Geo Coordinates by Address
// @Description This endpoint allows you to get geo coordinates by address.
// @Tags User
// @Accept json
// @Produce json
// @Param index path int true "Book INDEX"
// @Param Authorization header string true "Bearer Token"
// @Param body body TakeBookRequest true "Request body"
// @Success 200 {object} Response "Успешное выполнение"
// @Failure 400 {object} mErrorResponse "Ошибка запроса"
// @Failure 500 {object} mErrorResponse "Ошибка подключения к серверу"
// @Security BearerAuth
// @Router /api/book/take/{index} [post]

// Реализация метода TakeBookHandler

// Реализация метода ReturnBook
func (r *PostgresBookRepository) GetAllBooks() ([]entities.Book, error) {
	rows, err := r.db.Query("SELECT index, book, author, block, take_count FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.Index, &book.Title, &book.Author, &book.Block, &book.TakeCount); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repo *PostgresBookRepository) TakeBook(id int) (entities.Book, error) {
	// Начинаем транзакцию
	tx, err := repo.db.Begin()
	if err != nil {
		return entities.Book{}, err
	}
	defer tx.Rollback()

	// Обновляем статус книги
	result, err := tx.Exec(
		"UPDATE books SET block = true, take_count = take_count + 1 WHERE index = $1 AND block = false",
		id,
	)
	if err != nil {
		return entities.Book{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entities.Book{}, errors.New("book not found or already taken")
	}

	// Получаем обновленную книгу
	var book entities.Book
	err = tx.QueryRow(
		"SELECT index, book, author, block, take_count FROM books WHERE id = $1",
		id,
	).Scan(&book.Index, &book.Title, &book.Author, &book.Block, &book.TakeCount)
	if err != nil {
		return entities.Book{}, err
	}

	// Фиксируем транзакцию
	if err := tx.Commit(); err != nil {
		return entities.Book{}, err
	}

	return book, nil
}

func (repo *PostgresBookRepository) AddBook(addaderBook AddaderBookRequest) error {
	_, err := repo.db.Exec(
		"INSERT INTO books (book, author, block, take_count) VALUES ($1, $2, false, 0)",
		addaderBook.Book,
		addaderBook.Author,
	)
	return err
}

func (r *PostgresBookRepository) UpdateBook(id int, book entities.Book) error {
	_, err := r.db.Exec(
		"UPDATE books SET book = $1, author = $2, block = $3 WHERE index = $4",
		book.Title,
		book.Author,
		book.Block,
		id,
	)
	return err
}

func (repo *PostgresBookRepository) ReturnBook(id int) error {
	result, err := repo.db.Exec(
		"UPDATE books SET block = false WHERE index = $1 AND block = true",
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found or already returned")
	}

	return nil
}
