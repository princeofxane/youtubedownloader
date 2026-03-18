# Build stage
FROM golang:1.23.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ytdownloader

# Runtime stage
FROM alpine:latest

# Install ffmpeg + dependencies
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3 \
    curl

# Install yt-dlp (arch-aware)
ARG TARGETARCH
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux_aarch64 -o /usr/local/bin/yt-dlp; \
    else \
        curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp; \
    fi && \
    chmod +x /usr/local/bin/yt-dlp

WORKDIR /app
COPY --from=build /app/ytdownloader .
COPY --from=build /app/config ./config

EXPOSE 8080
CMD ["./ytdownloader"]