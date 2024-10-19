# ReadMe

## Description
youtubedownloader is a service to download youtube video for a given resolution and store it in the home storage.
Target: RPI.

## Setup
### Configuration file.
This file defines the service and picks up the script needs to be executed.

* location: `/home/{user}/.config/yt_downloader/config.yaml`

```
server:
  port: 8080
yt-dlp:
  ffmpeg_location: "/usr/bin/ffmpeg"
  output_path: "${SNOWHIVE}"
  is_active: true
```

### Pre-Requisitines
Make sure you have these programs installed.
* yt-dlp `it used to download youtube videos.`
* ffmpeg `it used to merge audio and video file.`


todo: 1. Write about outputpath 2. Write about resolution. 

