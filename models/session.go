package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/progsamdev/coursescalhoun/rand"
)

const (
	//The minimum number of bytes to be used for each session token.
	MinBytesPerToken = 32
)

type Session struct {
	ID              string
	UserID          string
	NewSessionToken string
	TokenHash       string
}

type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

func (ss *SessionService) Create(userID string) (*Session, error) {
	if ss.BytesPerToken < MinBytesPerToken {
		ss.BytesPerToken = MinBytesPerToken
	}
	token, err := rand.Strings(ss.BytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	session := Session{
		UserID:          userID,
		NewSessionToken: token,
		TokenHash:       ss.hash(token),
	}

	row := ss.DB.QueryRow(`
		UPDATE sessions
		SET token_hash = $2
		WHERE user_id = $1
    RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows {
		// If no session exists, we will get ErrNoRows. That means we need to
		// create a session object for that user.
		row = ss.DB.QueryRow(`
			INSERT INTO sessions (user_id, token_hash)
			VALUES ($1, $2)
			RETURNING id;`, session.UserID, session.TokenHash)
		// The error will be overwritten with either a new error, or nil
		err = row.Scan(&session.ID)
	}

	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	//1. Create the session token
	//2.

	return nil, nil

}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}