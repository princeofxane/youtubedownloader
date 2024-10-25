package api

import (
	"context"
	pb "youtube_downloader/pkg/grpc_gen"

	logr "github.com/sirupsen/logrus"
)

func (s *API) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	logr.Infof("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}
