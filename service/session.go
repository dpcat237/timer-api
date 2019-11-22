package service

import (
	"time"

	"gopkg.in/go-playground/validator.v10"

	"gitlab.com/dpcat237/timer-api/model"
	"gitlab.com/dpcat237/timer-api/repository/db"
)

const (
	FilterQuery = "filter"

	filterDay   = "day"
	filterWeek  = "week"
	filterMonth = "month"
)

var validate *validator.Validate

// SessionService defines required methods for Session service
type SessionService interface {
	AddSession(se model.Session) model.Error
	GetSessions(flt string) (model.Sessions, model.Error)
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
	ses, er := srv.GetSessions("")
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
func (srv sessionService) GetSessions(flt string) (model.Sessions, model.Error) {
	ses, err := srv.sesRps.GetSessions()
	if err != nil {
		return ses, model.NewErrorServer("Error getting sessions").WithError(err)
	}

	if flt == "" {
		return ses, model.NewErrorNil()
	}
	return srv.filterSessions(ses, flt), model.NewErrorNil()
}

// filterSessions filters sessions by creation date
func (srv sessionService) filterSessions(ses model.Sessions, flt string) model.Sessions {
	var rst model.Sessions
	var after int64
	switch flt {
	case filterDay:
		after = srv.getTodayTimestamp()
	case filterWeek:
		after = srv.getWeekTimestamp()
	case filterMonth:
		after = srv.getMonthTimestamp()
	default:
		return ses
	}

	for _, se := range ses {
		if se.CreatedAt >= after {
			rst = append(rst, se)
		}
	}
	return rst
}

// getTodayTimestamp returns timestamp of beginning of today
func (srv sessionService) getTodayTimestamp() int64 {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
}

// getTodayTimestamp returns timestamp of beginning last 7 days
func (srv sessionService) getWeekTimestamp() int64 {
	now := time.Now()
	return srv.getDayTimestamp(now.AddDate(0, 0, -7))
}

// getTodayTimestamp returns timestamp of beginning last month
func (srv sessionService) getMonthTimestamp() int64 {
	now := time.Now()
	return srv.getDayTimestamp(now.AddDate(0, -1, 0))
}

// getTodayTimestamp returns timestamp of date without hours, minutes, and seconds
func (srv sessionService) getDayTimestamp(tm time.Time) int64 {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).Unix()
}
