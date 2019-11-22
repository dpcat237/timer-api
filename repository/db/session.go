package db

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"

	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/model"
)

// SessionRepository defines required methods for Session repository
type SessionRepository interface {
	GetSessions() (model.Sessions, error)
	SaveSessions(ses model.Sessions) error
}

// sessionRepository defines required services for Session repository
type sessionRepository struct {
	db   *bolt.DB
	logg logger.Logger
}

// NewSession initializes Session repository
func NewSession(db *bolt.DB, logg logger.Logger) *sessionRepository {
	return &sessionRepository{db: db, logg: logg}
}

// GetLastLocationByIP get last location details by IP address
func (rps sessionRepository) GetSessions() (model.Sessions, error) {
	var ses model.Sessions
	tx, err := rps.db.Begin(false)
	if err != nil {
		return ses, err
	}
	defer logRollback(tx, rps.logg)

	b := tx.Bucket([]byte(model.SessionBucket))
	if b == nil {
		return ses, nil
	}

	sesBt := b.Get([]byte(model.SessionsKey))
	if len(sesBt) == 0 {
		return ses, nil
	}

	if err := json.Unmarshal(sesBt, &ses); err != nil {
		return ses, err
	}
	return ses, nil
}

func (rps sessionRepository) SaveSessions(ses model.Sessions) error {
	tx, err := rps.db.Begin(true)
	if err != nil {
		return err
	}

	b := tx.Bucket([]byte(model.SessionBucket))
	if b == nil {
		newB, err := tx.CreateBucket([]byte(model.SessionBucket))
		if err != nil {
			return err
		}
		b = newB
	}

	sesBt, err := json.Marshal(ses)
	if err != nil {
		return err
	}

	if err := b.Put([]byte(model.SessionsKey), sesBt); err != nil {
		return err
	}
	return tx.Commit()
}
