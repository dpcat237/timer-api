package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func getBodyContent(r *http.Request, data interface{}) model.Error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, r.ContentLength))
	if err != nil {
		return model.NewErrorServer("Error reading request body").WithError(err)
	}

	if err := json.Unmarshal(body, data); err != nil {
		return model.NewErrorServer("Error parsing request body").WithError(err)
	}
	return model.NewErrorNil()
}

func returnCreatedNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

func returnFailed(w http.ResponseWriter, er model.Error) {
	w.WriteHeader(er.Status)
	if err := json.NewEncoder(w).Encode(er); err != nil {
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
