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

	conf, err := config.LoadConfig()
	if err != nil {
		logr.Fatalf("failed to load config: %v", err)
	}

	r := server.NewServer(conf)
	go r.Start()

	select {
	case <-ctx.Done():
		logr.Info("context done: shutting down")
	case s := <-interrupt:
		logr.WithField("signal", s).Info("server: received interrupt signal")
	}
}

