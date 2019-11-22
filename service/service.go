package service

import (
	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/repository/db"
)

// Collector defines services
type Collector struct {
	SesSrv SessionService
}

// Init initializes services and required repositories
func Init(dbCl db.DatabaseCollector, logg logger.Logger) Collector {
	// Initialize repositories
	sesRps := db.NewSession(dbCl.GetDatabase(), logg)

	// Initialize services
	sesSrv := newSession(sesRps)

	return Collector{
		SesSrv: sesSrv,
	}
}
