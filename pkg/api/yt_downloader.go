package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func (a *api) ytDownload(w http.ResponseWriter, r *http.Request) {
	videoUrl := r.Header.Get("video_url")
	videoQuality := r.Header.Get("video_quality")

	if videoUrl == "" || videoQuality == "" {
		log.Println("needed parameter is not provided")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "please provide required parameters")
		return
	}

	path, err := storagePath()
	if err != nil {
		log.Println("failed to get storage path")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	ffmpegLocation := a.cnfg.YTDLPCfg.FFMPEGLocation
	resolutionFlag := fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best", videoQuality)

	err = downloadVideo(resolutionFlag, ffmpegLocation, videoUrl, path)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Println("download is completed")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "download is completed")
	}
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
		return fmt.Errorf("failed to download video: %w", err)
	}
	return nil
}

func storagePath() (string, error) {
	// check if env variable SNOWHIVE is present.
	// this is the path to mounted harddrive.
	path, exists := os.LookupEnv("SNOWHIVE")
	if !exists {
		log.Println("SNOWHIVE env is not present")
		return "", errors.New("storage location is not mounted")
	}

	// check if the directory is active
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println("SNOWHIVE path is not present")
		return "", errors.New("storage path is not present")
	}
	// Check if it's a directory
	log.Println(info.IsDir())

	return path, nil
}
