package api

import (
	"youtube_downloader/config"

	"github.com/gorilla/mux"
)

type api struct {
	cnfg config.YTDownloaderConfig
}

func Handler(r *mux.Router, cnf config.YTDownloaderConfig) *mux.Router {
	a := &api{
		cnfg: cnf,
	}
	r.HandleFunc("/healthcheck", a.healthCheck).Methods("GET")
	r.HandleFunc("/downloadvideo", a.ytDownload).Methods("POST")

	return r
}
