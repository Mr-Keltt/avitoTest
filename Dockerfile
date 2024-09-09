# Go image based on Debian (default Go image is often Debian-based)
FROM golang:1.23.1 AS builder

# Update and install git
RUN apt-get update && apt-get install -y git

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN go build -o /go-tender-app main.go

# Use a minimal image for the final stage
FROM debian:bullseye-slim

# Set working directory for binary
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /go-tender-app .

# Expose the service port (matching SERVER_ADDRESS)
EXPOSE 8080

# Define entrypoint to start the app
CMD ["./go-tender-app"]
