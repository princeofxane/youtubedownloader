# Build stage
FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ytdownloader

# Runtime stage
FROM alpine:latest

# Install dependencies
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3 \
    py3-pip \
    curl \
    unzip \
    nodejs

# Install yt-dlp and yt-dlp-ejs via pip (musl/Alpine compatible)
RUN pip3 install --no-cache-dir yt-dlp yt-dlp-ejs --break-system-packages

WORKDIR /app
COPY --from=build /app/ytdownloader .
COPY --from=build /app/config ./config

EXPOSE 8080
CMD ["./ytdownloader"]