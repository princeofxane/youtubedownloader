package api

import (
	"youtube_downloader/pkg/grpc_gen"
	"youtube_downloader/pkg/config"
)

type API struct {
	Cnfg config.YTDownloaderConfig
	grpc_gen.UnimplementedYTDownloadServiceServer
}
