package db

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"

	"gitlab.com/dpcat237/timer-api/model"
)

// SessionRepository defines required methods for Session repository
type SessionRepository interface {
	GetSessions() (model.Sessions, error)
}

// sessionRepository defines required services for Session repository
type sessionRepository struct {
	db *bolt.DB
}

// NewSession initializes Session repository
func NewSession(db *bolt.DB) *sessionRepository {
	return &sessionRepository{db: db}
}

// GetLastLocationByIP get last location details by IP address
func (rps sessionRepository) GetSessions() (model.Sessions, error) {
	var ses model.Sessions
	tx, err := rps.db.Begin(false)
	if err != nil {
		return ses, err
	}

	b := tx.Bucket([]byte(model.SessionBucket))
	if b == nil {
		return ses, nil
	}

	sesStr := b.Get([]byte(model.SessionsKey))
	if len(sesStr) == 0 {
		return ses, nil
	}

	if err := json.Unmarshal(sesStr, &ses); err != nil {
		return ses, err
	}
	return ses, nil
}
