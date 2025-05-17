package postgres

import (
	"errors"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func (repo *PostgresAuthorRepository) GetAuthors() ([]entities.Author, error) {
	rows, err := repo.db.Query("SELECT id, name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

// Реализация метода ListAuthorsHandler
func (repo *PostgresAuthorRepository) ListAuthors() ([]entities.Author, error) {
	rows, err := repo.db.Query("SELECT DISTINCT id, name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []entities.Author
	for rows.Next() {
		var author entities.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

// Реализация метода AddAuthorHandler
func (repo *PostgresAuthorRepository) AddAuthor(author entities.Author) error {
	if author.Name == "" {
		return errors.New("author name is required")
	}

	_, err := repo.db.Exec("INSERT INTO authors (name) VALUES ($1)", author.Name)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: authors.name" {
			return errors.New("author already exists")
		}
		return err
	}

	return nil
}
