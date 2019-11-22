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

// getBodyContent extract body from request
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

// getQueryVal returns URL query variable
func getQueryVal(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// returnCreatedNoContent sets status of response to 201
func returnCreatedNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

// returnCreatedNoContent sets status of response to 500 and returns an error message
func returnFailed(w http.ResponseWriter, er model.Error) {
	w.WriteHeader(er.Status)
	if err := json.NewEncoder(w).Encode(er); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}

// returnCreatedNoContent returns marshaled array to JSON and empty array if array is empty
func returnJsonArray(w http.ResponseWriter, v interface{}) {
	if zero.IsZero(v) {
		v = make([]string, 0)
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, model.ErrorServer, http.StatusInternalServerError)
		return
	}
}
