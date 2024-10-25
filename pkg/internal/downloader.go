package internal

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	custerr "youtube_downloader/pkg/custom_error"

	logr "github.com/sirupsen/logrus"
)

func Downloader(videoUrl, videoQuality, ffmpegLocation string) error {
	path, err := storagePath()
	if err != nil {
		return err
	}

	resolutionFlag := fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best", videoQuality)

	err = downloadVideo(resolutionFlag, ffmpegLocation, videoUrl, path)
	if err != nil {
		return err
	} else {
		logr.Info("download is completed")
	}

	return nil
}

func downloadVideo(resolutionFlag, ffmpegLocation, videoUrl, path string) error {
	outputTemplate := fmt.Sprintf("%s/%s.%%(ext)s", path, "%(title)s")

	cmd := exec.Command("yt-dlp", "-f", resolutionFlag, "--ffmpeg-location", ffmpegLocation, "-o", outputTemplate, videoUrl)

	// Set the output to the terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		msg := fmt.Sprintf("failed to download video: %v", err)

		logr.Errorln(msg)
		return custerr.CreateErr(msg, http.StatusInternalServerError)
	}
	return nil
}

func storagePath() (string, error) {
	// check if env variable SNOWHIVE is present.
	// this is the path to mounted harddrive.
	path, exists := os.LookupEnv("SNOWHIVE")
	if !exists {
		msg := "SNOWHIVE env is not present"

		logr.Errorln(msg)
		return "", custerr.CreateErr(msg, http.StatusInternalServerError)
	}

	// check if the directory is active
	_, err := os.Stat(path)
	if os.IsNotExist(err) {

		msg := fmt.Sprintf("SNOWHIVE path is not present, err: %v", err)

		logr.Errorln(msg)
		return "", custerr.CreateErr(msg, http.StatusInternalServerError)
	}

	return path, nil
}
