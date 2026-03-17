package api

import (
	"youtube_downloader/pkg/config"

	"github.com/gorilla/mux"
)

type api struct {
	cnfg config.YTDownloaderConfig
}

func Handler(r *mux.Router, cnf config.YTDownloaderConfig) *mux.Router {
	a := &api{
		cnfg: cnf,
	}
	r.HandleFunc("/health", a.healthCheck).Methods("GET")
	r.HandleFunc("/download", a.ytDownload).Methods("POST")

	return r
}
