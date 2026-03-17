package server

import (
	"fmt"
	"net/http"
	"youtube_downloader/pkg/config"

	"github.com/gorilla/mux"
	logr "github.com/sirupsen/logrus"

	"youtube_downloader/pkg/api"
)

type HTTPServer struct {
	endpoint string
	port     string
	router   *mux.Router
}

func NewServer(conf *config.Config) *HTTPServer {
	r := mux.NewRouter()
	r = api.Handler(r, conf)

	return &HTTPServer{
		endpoint: conf.Server.Host,
		port:     conf.Server.Port,
		router:   r,
	}
}

func (h *HTTPServer) Start() {
	logr.Infof("http server has started at port: %s", h.port)
	logr.Errorln(http.ListenAndServe(fmt.Sprintf("%s:%s", h.endpoint, h.port), h.router))
}
