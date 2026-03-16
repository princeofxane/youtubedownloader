package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"youtube_downloader/pkg/config"
	"youtube_downloader/pkg/server"

	logr "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	cnf := config.ReadConfig()

	// run http server only in dev environment.
	if cnf.Environment != "e2" {
		go func() {
			r := server.NewHTTP(cnf)
			r.StartHTTPServer()
		}()
	}

	go func() {
		r := server.NewGRPC(cnf)
		r.StartGRPCServer(cnf)
	}()

	select {
	case <-ctx.Done():
		logr.Info("context done: shutting down")
	case s := <-interrupt:
		logr.WithField("signal", s).Info("server: received interrupt signal")

	}

}
