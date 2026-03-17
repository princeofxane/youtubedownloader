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

func (h *HTTPServer) StartHTTPServer() {
	logr.Infof("http server has started at port: %s", h.port)
	logr.Errorln(http.ListenAndServe(fmt.Sprintf("%s:%s", h.endpoint, h.port), h.router))
}

func NewHTTP(cnf config.YTDownloaderConfig) *HTTPServer {
	r := mux.NewRouter()
	r = api.Handler(r, cnf)

	return &HTTPServer{
		endpoint: cnf.ServerCfg.HTTPEndPoint,
		port:     cnf.ServerCfg.HTTPPort,
		router:   r,
	}
}
