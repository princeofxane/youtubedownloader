package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v2"
)

type YTDownloaderConfig struct {
	ServerCfg ServerConfig `yaml:"server"`
	YTDLPCfg  YTDLPConfig  `yaml:"yt-dlp"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

// YTDLP is the command line tool that is used to download yt videos.
type YTDLPConfig struct {
	FFMPEGLocation string `yaml:"ffmpeg_location"`
	OutputPath     string `yaml:"output_path"`
	IsActive       string `yaml:"is_active"`
}

func ReadConfig() YTDownloaderConfig {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("rrror getting hostname: %v\n", err)
	}

	username := currentUser.Username
	if strings.Contains(username, " ") {
		// Take the first part if there's a space
		username = strings.Split(username, " ")[0]
	}

	path := fmt.Sprintf("/home/%s/.config/yt_downloader/config.yaml", username)

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config.yaml file: %v", err)
	}

	var config YTDownloaderConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing config.yaml file: %v", err)
	}

	return config
}
