package internal

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"youtube_downloader/pkg/config"
	custerr "youtube_downloader/pkg/custom_error"

	logr "github.com/sirupsen/logrus"
)

func Downloader(videoUrl, videoQuality, ffmpegLocation string, conf *config.Config) error {
	path := conf.YTDLP.DownloadPath

	err := checkURL(videoUrl)
	if err != nil {
		return err
	}

	resolutionFlag := fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best", videoQuality)

	err = downloadVideo(resolutionFlag, ffmpegLocation, videoUrl, path)
	if err != nil {
		return err
	} else {
		logr.Info("download has been completed")
	}

	return nil
}

func checkURL(url string) error {
	// Run yt-dlp with --simulate to check the URL
	cmd := exec.Command("yt-dlp", "--simulate", url)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// yt-dlp returns an error if the URL is invalid or the video doesn't exist
		return fmt.Errorf("failed to validate URL: %s", string(output))
	}

	if len(output) == 0 {
		return errors.New("no response from yt-dlp")
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
