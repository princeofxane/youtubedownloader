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

func Downloader(videoUrl, videoQuality, ffmpegLocation string, cnf config.YTDownloaderConfig) error {
	var path string

	if cnf.Environment == "e0" {
		path = cnf.YTDLPCfg.OutputPath
	} else {
		xpath, err := storagePath()
		if err != nil {
			logr.Errorf("failed to get storage path: %v", err)
			return err
		}
		path = xpath
	}


	fmt.Println("-------------------2-------------------------------")
	fmt.Println("video url: ", videoUrl)
	fmt.Println("--------------------------------------------------")

	err := checkURL(videoUrl)
	if err != nil {
		return err
	}

	fmt.Println("------------------3--------------------------------")
	fmt.Println("video url: ", videoUrl)
	fmt.Println("--------------------------------------------------")
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
		logr.Errorf("failed to validate URL: %s", err)
		return fmt.Errorf("failed to validate URL: %s", string(output))
	}

	if len(output) == 0 {
		logr.Errorln("no response from yt-dlp")
		return errors.New("no response from yt-dlp")
	}

	return nil
}

func downloadVideo(resolutionFlag, ffmpegLocation, videoUrl, path string) error {
	outputTemplate := fmt.Sprintf("%s/%s.%%(ext)s", path, "%(title)s")


	fmt.Println("-----------------3---------------------------------")
	fmt.Println("video url: ", videoUrl)
	fmt.Println("--------------------------------------------------")

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
