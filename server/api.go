package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/mahakamcloud/netd/appcontext"
	"github.com/mahakamcloud/netd/config"
	"github.com/mahakamcloud/netd/service"
	nrgorilla "github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
	log "github.com/sirupsen/logrus"
)

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Info("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	log.Info("API server shutdown complete")
}

func StartAPIServer() {
	log.Info("Starting Service API")
	newRelicApp := appcontext.NewrelicApp()
	muxRouter := Router(service.NewServices())
	router := nrgorilla.InstrumentRoutes(muxRouter, newRelicApp)
	handlerFunc := router.ServeHTTP

	n := negroni.New(negroni.NewRecovery())
	n.Use(httpStatLogger())
	n.UseHandlerFunc(handlerFunc)
	portInfo := ":" + strconv.Itoa(config.Port())
	server := &http.Server{Addr: portInfo, Handler: n}
	go listenServer(server)
	waitForShutdown(server)
}

func httpStatLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		next(rw, r)
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000

		if r.URL.Path != "/ping" {
			log.WithFields(log.Fields{
				"RequestTime":   startTime.Format(time.RFC3339),
				"ResponseTime":  responseTime.Format(time.RFC3339),
				"DeltaTime":     deltaTime,
				"RequestUrl":    r.URL.Path,
				"RequestMethod": r.Method,
				"RequestProxy":  r.RemoteAddr,
				"RequestSource": r.Header.Get("X-FORWARDED-FOR"),
			}).Debug("Http Logs")
		}
	})
}
