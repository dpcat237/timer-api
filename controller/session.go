package controller

import (
	"gitlab.com/dpcat237/timer-api/logger"
)

// SessionController defines required methods for Session controller
type SessionController interface {
}

// sessionController defines required services for Session controller
type sessionController struct {
	logg logger.Logger
}

// newSession initializes Session controller
func newSession(logg logger.Logger) *sessionController {
	return &sessionController{logg: logg}
}
