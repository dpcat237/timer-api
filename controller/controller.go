package controller

import (
	"net/http"

	"gitlab.com/dpcat237/timer-api/logger"
)

// Collector defines controllers
type Collector struct {
	SesCnt SessionController
	logg   logger.Logger
}

// InitCollector initializes collector of controllers for gRPC
func InitCollector(logg logger.Logger) Collector {
	return Collector{
		SesCnt: newSession(logg),
		logg:   logg,
	}
}

// HealthCheck checks service health
func (srv *Collector) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`{"success": true}`)); err != nil {
		srv.logg.WithError(err).Error("Error health check")
	}
}
