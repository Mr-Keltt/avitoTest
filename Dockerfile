# Dockerfile

# Use official Golang image as a base
FROM golang:1.23.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the container's working directory
COPY . .

# Download Go modules
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
