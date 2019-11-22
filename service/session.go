package service

import (
	"gitlab.com/dpcat237/timer-api/model"
	"gitlab.com/dpcat237/timer-api/repository/db"
)

// SessionService defines required methods for Session service
type SessionService interface {
	GetSessions() (model.Sessions, model.Error)
}

// sessionService defines required services for Session service
type sessionService struct {
	sesRps db.SessionRepository
}

// newSession initializes Session service
func newSession(sesRps db.SessionRepository) *sessionService {
	return &sessionService{sesRps: sesRps}
}

// GetSessions returns sessions from database
func (srv sessionService) GetSessions() (model.Sessions, model.Error) {
	ses, err := srv.sesRps.GetSessions()
	if err != nil {
		return ses, model.NewErrorServer("Error getting sessions").WithError(err)
	}
	return ses, model.NewErrorNil()
}
