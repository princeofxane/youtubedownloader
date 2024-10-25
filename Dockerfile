# Step 1: Build the Go application
FROM golang:1.23.2-bookworm as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Copy the source code into the container
COPY . .

# Download the Go modules
RUN go mod download

# Build the Go application
# RUN go build -o myapp .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

# Step 2: Create a smaller image to run the application
FROM alpine:3.20.3

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port that your application listens on (change as needed)
EXPOSE 8050

# Command to run the executable
CMD ["./myapp"]