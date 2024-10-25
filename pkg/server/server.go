package server

import (
	"youtube_downloader/pkg/api"
	"youtube_downloader/pkg/config"

	"github.com/gorilla/mux"
)

func New(cnf config.YTDownloaderConfig) *mux.Router {
	r := mux.NewRouter()
	r = api.Handler(r, cnf)
	return r
}
