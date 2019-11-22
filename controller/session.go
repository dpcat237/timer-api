package controller

import (
	"net/http"
	"time"

	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/model"
	"gitlab.com/dpcat237/timer-api/service"
)

// SessionController defines required methods for Session controller
type SessionController interface {
	AddSession(w http.ResponseWriter, r *http.Request)
	GetSessions(w http.ResponseWriter, r *http.Request)
}

// sessionController defines required services for Session controller
type sessionController struct {
	logg   logger.Logger
	sesSrv service.SessionService
}

// newSession initializes Session controller
func newSession(logg logger.Logger, sesSrv service.SessionService) *sessionController {
	return &sessionController{logg: logg, sesSrv: sesSrv}
}

// AddSession adds session
func (cnt *sessionController) AddSession(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	act := "session.add"
	errMsg := ""
	defer cnt.logg.RequestEnd(time.Now(), act, &status, &errMsg)

	var se model.Session
	if er := getBodyContent(r, &se); er.IsError() {
		status = er.Status
		errMsg = er.Message
		returnFailed(w, er)
		return
	}

	if er := cnt.sesSrv.AddSession(se); er.IsError() {
		status = er.Status
		errMsg = er.Message
		returnFailed(w, er)
		return
	}
	returnCreatedNoContent(w)
}

// GetSessions return existent sessions on database
func (cnt *sessionController) GetSessions(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	act := "sessions.list"
	errMsg := ""
	defer cnt.logg.RequestEnd(time.Now(), act, &status, &errMsg)

	dto, er := cnt.sesSrv.GetSessions()
	if er.IsError() {
		status = er.Status
		errMsg = er.Message
		returnFailed(w, er)
		return
	}
	returnJsonArray(w, dto)
}
