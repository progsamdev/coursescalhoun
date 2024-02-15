package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            string
	Email         string
	PasswordHarsh string
}

type UserService struct {
	DB *sql.DB
}

func (u *UserService) Create(email, password_raw string) (*User, error) {
	email = strings.ToLower(email)
	password_hash, err := u.hashPassword(password_raw)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	newUser := User{
		Email:         email,
		PasswordHarsh: password_hash,
	}
	row := u.DB.QueryRow(
		`INSERT INTO users(email, password_hash)
		VALUES ($1, $2) RETURNING id`, newUser.Email, newUser.PasswordHarsh,
	)
	err = row.Scan((&newUser.ID))
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return &newUser, nil
}

func (u *UserService) hashPassword(password_raw string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password_raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (u *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{Email: email}

	row := u.DB.QueryRow(
		`SELECT id, password_hash 
					FROM users
					WHERE email = $1`, user.Email)
	err := row.Scan(&user.ID, &user.PasswordHarsh)

	if err != nil {
		return nil, fmt.Errorf("Authenticate failed: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHarsh), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Authenticate failed: %w", err)
	}
	return &user, nil
}
