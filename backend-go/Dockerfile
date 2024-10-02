# Use the official Golang image as the base image
FROM golang:1.23-bookworm

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go source
COPY main.go ./

# Initialize the go module
RUN go mod init helloworld

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]