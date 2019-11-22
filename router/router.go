package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"gitlab.com/dpcat237/timer-api/controller"
	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/model"
)

// apiVersion defines current API version
const apiVersion = "v1"

// Manager is router's manager
type Manager struct {
	logg logger.Logger
	rtr  *mux.Router
	srv  *http.Server
}

// NewManager initializes router manager
func NewManager(cntCll controller.Collector) Manager {
	var mng Manager
	mng.rtr = mux.NewRouter().StrictSlash(true)
	mng.addRoutes(mng.rtr, mng.getSysRoutes(), true)
	mng.rtr.Handle("/debug/vars", http.DefaultServeMux)
	mng.addRoutes(mng.rtr.PathPrefix("/"+apiVersion).Subrouter(), mng.getV1Routes(cntCll), true)
	return mng
}

// LunchRouter runs router listener
func (mng *Manager) LunchRouter(port string) {
	mng.srv = &http.Server{Addr: ":" + port, Handler: mng.rtr}
	go func() {
		err := mng.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			mng.logg.Errorf("Error starting http service: %s", err)
		}
	}()
}

// Shutdown close router's connection
func (mng *Manager) Shutdown(port string) {
	if err := mng.srv.Shutdown(context.Background()); err != nil {
		mng.logg.Errorf("Error stopping http service %s", err)
	}
}

// addRoutes set route for router
func (mng *Manager) addRoutes(r *mux.Router, routes []model.Route, useStatusCode bool) {
	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

// getSysRoutes sets system routes
func (mng *Manager) getSysRoutes() []model.Route {
	return []model.Route{
		model.NewRoute(http.MethodGet, "/services/health", mng.healthCheck, "Check service health"),
	}
}

// getV1Routes sets version 1 routes
func (mng *Manager) getV1Routes(cntCll controller.Collector) []model.Route {
	return []model.Route{
		/** Session **/
		model.NewRoute(http.MethodGet, "/sessions", cntCll.SesCnt.GetSessions, "Get sessions"),
		model.NewRoute(http.MethodPost, "/sessions", cntCll.SesCnt.AddSession, "Add session"),
	}
}

// healthCheck checks service health
func (mng *Manager) healthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(`{"success": true}`)); err != nil {
		mng.logg.Errorf("Error returning health check %s", err)
	}
}
