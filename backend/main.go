package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/gggruja/PAC_SRB_2020/backend/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func healthChecking(rw http.ResponseWriter, r *http.Request) {
	log.Println("Healthy App - Running!")
	log.Println("Backend App is up and running - Goran Grujic")

	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Backend App - Goran Grujic!"))
}

func main() {

	// log instance
	l := hclog.Default()

	// load the configuration
	cnf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	l.Info("Loaded config: ", hclog.Fmt("%+v", cnf))

	sm := mux.NewRouter()

	sm.HandleFunc("/", healthChecking)
	sm.Handle("/metrics", promhttp.Handler())

	// create Server
	s := http.Server{
		Addr:         cnf.BindAddress,
		Handler:      sm,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 120,
	}

	go func() {
		l.Info("Starting server on " + s.Addr)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received.
	sig := <-c
	l.Info("Received signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	l.Info("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
