package server

import (
	"youtube_downloader/config"
	"youtube_downloader/pkg/api"

	"github.com/gorilla/mux"
)

func New(cnf config.YTDownloaderConfig) *mux.Router {
	r := mux.NewRouter()
	r = api.Handler(r, cnf)
	return r
}
