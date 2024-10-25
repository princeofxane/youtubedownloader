package api

import (
	"context"
	"fmt"

	logr "github.com/sirupsen/logrus"

	pb "youtube_downloader/pkg/grpc_gen"

	"youtube_downloader/pkg/internal"
)

func (s *API) YTDownload(_ context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	videoUrl := in.GetVideoUrl()
	videoQuality := in.GetVideQuality()

	if videoUrl == "" || videoQuality == "" {
		logr.Errorln("needed parameter is not provided")

		//todo: return an error or some status.
		return &pb.DownloadResponse{Message: "please provide required parameters"}, nil
	}

	ffmpegLocation := s.Cnfg.YTDLPCfg.FFMPEGLocation

	err := internal.Downloader(videoUrl, videoQuality, ffmpegLocation)
	if err != nil {
		return &pb.DownloadResponse{Message: fmt.Sprintf("an error occured: %v", err)}, nil
	}

	return &pb.DownloadResponse{Message: "video has been downloaded"}, nil

}
