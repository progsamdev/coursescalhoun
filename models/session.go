package models

import (
	"database/sql"
	"fmt"

	"github.com/progsamdev/coursescalhoun/rand"
)

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
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	//TODO: hash de session toekn
	session := Session{
		UserID:          userID,
		NewSessionToken: token,
	}
	//store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	//1. Create the session token
	//2.

	return nil, nil

}
