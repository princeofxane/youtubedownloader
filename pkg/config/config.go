package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YTDownloaderConfig struct {
	ServerCfg   ServerConfig `yaml:"server"`
	YTDLPCfg    YTDLPConfig  `yaml:"yt-dlp"`
	Environment string
}

type ServerConfig struct {
	HTTPEndPoint string `yaml:"http_endpoint"`
	HTTPPort     string `yaml:"http_port"`
	GRPCEndPoint string `yaml:"grpc_endpoint"`
	GRPCPort     string `yaml:"grpc_port"`
}

// YTDLP is the command line tool that is used to download yt videos.
type YTDLPConfig struct {
	FFMPEGLocation string `yaml:"ffmpeg_location"`
	OutputPath     string `yaml:"output_path"`
	IsActive       string `yaml:"is_active"`
}

func ReadConfig() YTDownloaderConfig {
	appENV := os.Getenv("APP_ENV")

	configFile := fmt.Sprintf("./config/%s_config.yaml", appENV)

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error reading config.yaml file: %v", err)
	}

	var config YTDownloaderConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing config.yaml file: %v", err)
	}

	config.Environment = appENV

	return config
}
