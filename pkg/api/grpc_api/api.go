package api

import (
	"youtube_downloader/pkg/config"
	"youtube_downloader/pkg/grpc_gen"
)

type API struct {
	Cnfg config.YTDownloaderConfig
	grpc_gen.UnimplementedYTDownloadServiceServer
}
