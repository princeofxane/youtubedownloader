package api

import (
	"youtube_downloader/pkg/config"

	"github.com/gorilla/mux"
)

type api struct {
	conf *config.Config
}

func Handler(r *mux.Router, cnf *config.Config) *mux.Router {
	a := &api{
		conf: cnf,
	}
	r.HandleFunc("/health", a.healthCheck).Methods("GET")
	r.HandleFunc("/download", a.ytDownload).Methods("POST")

	return r
}
