package service

import (
	"time"

	"gopkg.in/go-playground/validator.v10"

	"gitlab.com/dpcat237/timer-api/model"
	"gitlab.com/dpcat237/timer-api/repository/db"
)

var validate *validator.Validate

// SessionService defines required methods for Session service
type SessionService interface {
	AddSession(se model.Session) model.Error
	GetSessions() (model.Sessions, model.Error)
}

// sessionService defines required services for Session service
type sessionService struct {
	sesRps db.SessionRepository
}

func init() {
	validate = validator.New()
}

// newSession initializes Session service
func newSession(sesRps db.SessionRepository) *sessionService {
	return &sessionService{sesRps: sesRps}
}

// AddSession adds new session to database
func (srv sessionService) AddSession(se model.Session) model.Error {
	if err := validate.Struct(se); err != nil {
		return model.NewErrorPrecondition(err.Error()).WithError(err)
	}

	se.CreatedAt = time.Now().UTC().Unix()
	ses, er := srv.GetSessions()
	if er.IsError() {
		return er
	}

	ses = append(ses, se)
	if err := srv.sesRps.SaveSessions(ses); err != nil {
		return model.NewErrorServer("Error saving sessions").WithError(err)
	}
	return model.NewErrorNil()
}

// GetSessions returns sessions from database
func (srv sessionService) GetSessions() (model.Sessions, model.Error) {
	ses, err := srv.sesRps.GetSessions()
	if err != nil {
		return ses, model.NewErrorServer("Error getting sessions").WithError(err)
	}
	return ses, model.NewErrorNil()
}
