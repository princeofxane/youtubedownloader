# Step 1: Build the Go application
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

# Step 2: Create a smaller image to run the application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port that your application listens on (change as needed)
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]