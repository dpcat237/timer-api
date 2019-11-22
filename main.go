package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/dpcat237/timer-api/config"
	"gitlab.com/dpcat237/timer-api/controller"
	"gitlab.com/dpcat237/timer-api/logger"
	"gitlab.com/dpcat237/timer-api/repository/db"
	"gitlab.com/dpcat237/timer-api/router"
	"gitlab.com/dpcat237/timer-api/service"
)

func main() {
	cfg := config.LoadConfigData()
	logg := logger.New()
	dbCl, er := db.InitDbCollector(cfg.DbName)
	if er.IsError() {
		logg.Error(er)
	}

	srvCll := service.Init(dbCl)
	cntCll := controller.InitCollector(logg, srvCll)
	rtrMng := router.NewManager(cntCll)
	rtrMng.LunchRouter(cfg.HTTPport)
	logg.Infof("Router started at on port %s", cfg.HTTPport)

	gracefulStop(cfg, logg, dbCl, rtrMng)
}

// gracefulStop stops router after receiving system or key interruption
func gracefulStop(cfg config.Config, logg logger.Logger, dbCl db.DatabaseCollector, rtrMng router.Manager) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	close(c)

	rtrMng.Shutdown(cfg.HTTPport)
	dbCl.GracefulStop()
	logg.Info("Service stopped")
}
