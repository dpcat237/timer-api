package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cydev/zero"

	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/model"
	"gitlab.com/dpcat237/timer-api/service"
)

// Collector defines controllers
type Collector struct {
	SesCnt SessionController
	logg   logger.Logger
}

// InitCollector initializes collector of controllers for gRPC
func InitCollector(logg logger.Logger, sCll service.Collector) Collector {
	return Collector{
		SesCnt: newSession(logg, sCll.SesSrv),
		logg:   logg,
	}
}

// HealthCheck checks service health
func (srv *Collector) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`{"success": true}`)); err != nil {
		srv.logg.WithError(err).Error("Error health check")
	}
}

func returnFailed(w http.ResponseWriter, er model.Error) {
	w.WriteHeader(er.Status)
	if err := json.NewEncoder(w).Encode(er); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}

func returnJson(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}

func returnJsonArray(w http.ResponseWriter, v interface{}) {
	if zero.IsZero(v) {
		v = make([]string, 0)
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}
