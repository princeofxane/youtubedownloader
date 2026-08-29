#!/bin/sh
set -e

echo "Upgrading yt-dlp and yt-dlp-ejs..."
pip3 install -U --no-cache-dir yt-dlp yt-dlp-ejs --break-system-packages ||
  echo "WARNING: yt-dlp upgrade failed, continuing with existing version"

echo "Starting application..."
exec "$@"
