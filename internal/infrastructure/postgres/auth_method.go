package postgres

import (
	"database/sql"
	"errors"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
)

func (r *PostgresAuthRepository) Login(username, password string) (entities.UserAuth, error) {
	var user entities.UserAuth
	err := r.db.QueryRow(
		"SELECT username, password FROM users WHERE username = $1",
		username,
	).Scan(&user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.UserAuth{}, errors.New("user not found")
		}
		return entities.UserAuth{}, err
	}

	if user.Password != password {
		return entities.UserAuth{}, errors.New("invalid password")
	}

	return user, nil
}
func (r *PostgresAuthRepository) Register(user entities.UserAuth) error {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)",
		user.Username,
	).Scan(&exists)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	_, err = r.db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username, user.Password,
	)
	return err
}
