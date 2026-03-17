package config

import (
	"fmt"
	"os"

	u "youtube_downloader/pkg/util"

	"gopkg.in/yaml.v2"
)

const (
	AppEnv = "APP_ENV"
	K3s    = "k3s"
	Dev    = "e0"
)

type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Env     string
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	YTDLP struct {
		FFMPEGLocation string `yaml:"ffmpeg_location"`
		DownloadPath   string `yaml:"download_path"`
	} `yaml:"yt-dlp"`
}

func LoadConfig() (*Config, error) {
	env, ok := u.GetEnv(AppEnv)
	if !ok {
		return nil, fmt.Errorf("environment variable %s is not set", AppEnv)
	}

	// validate environment variable
	if env != K3s && env != Dev {
		return nil, fmt.Errorf("invalid environment: %s. Must be either %s or %s", env, K3s, Dev)
	}

	return parseConfig(env)
}

func parseConfig(env string) (*Config, error) {
	var config Config

	config.App.Env = env

	configFile := fmt.Sprintf("./config/%s.yaml", env)

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config.yaml file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config.yaml file: %v", err)
	}

	return &config, nil
}
