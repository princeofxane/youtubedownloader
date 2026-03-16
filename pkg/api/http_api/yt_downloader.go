package api

import (
	"fmt"
	"net/http"
	"youtube_downloader/pkg/internal"

	custerr "youtube_downloader/pkg/custom_error"

	logr "github.com/sirupsen/logrus"
)

func (a *api) ytDownload(w http.ResponseWriter, r *http.Request) {
	videoUrl := r.Header.Get("video_url")
	videoQuality := r.Header.Get("video_quality")

	if videoUrl == "" || videoQuality == "" {
		logr.Errorln("needed parameter is not provided")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "please provide required parameters")
		return
	}
	ffmpegLocation := a.cnfg.YTDLPCfg.FFMPEGLocation

	err := internal.Downloader(videoUrl, videoQuality, ffmpegLocation, a.cnfg)
	if err != nil {
		cerr, ok := err.(*custerr.CustomError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}
		w.WriteHeader(cerr.StatusCode)
		fmt.Fprintln(w, cerr.Error())
		return

	}
}
