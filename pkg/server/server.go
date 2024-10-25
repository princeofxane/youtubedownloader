package server

import (
	"fmt"
	"net"
	"net/http"
	"youtube_downloader/pkg/config"

	"github.com/gorilla/mux"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	grpc_api "youtube_downloader/pkg/api/grpc_api"
	http_api "youtube_downloader/pkg/api/http_api"
	pb "youtube_downloader/pkg/grpc_gen"
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
	r = http_api.Handler(r, cnf)

	return &HTTPServer{
		endpoint: cnf.ServerCfg.HTTPEndPoint,
		port:     cnf.ServerCfg.HTTPPort,
		router:   r,
	}
}

type GRPCServer struct {
	endpoint string
	port     string
	options  []grpc.ServerOption
}

func (g *GRPCServer) StartGRPCServer(config config.YTDownloaderConfig) {
	logr.Infof("grpc has started at port: %s", g.port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", g.endpoint, g.port))
	if err != nil {
		logr.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(g.options...)
	pb.RegisterYTDownloadServiceServer(grpcServer, &grpc_api.API{
		Cnfg: config,
	})
	grpcServer.Serve(lis)

}

func NewGRPC(cnf config.YTDownloaderConfig) *GRPCServer {
	var opts []grpc.ServerOption

	return &GRPCServer{
		endpoint: cnf.ServerCfg.GRPCEndPoint,
		port:     cnf.ServerCfg.GRPCPort,
		options:  opts,
	}
}
