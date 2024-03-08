package models

import "database/sql"

type Session struct {
	ID              string
	UserID          string
	NewSessionToken string
	TokenHash       string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID string) (*Session, error) {
	//1. Create the session token
	//2.

	return nil, nil

}

func (ss *SessionService) User(token string) (*User, error) {
	//1. Create the session token
	//2.

	return nil, nil

}
