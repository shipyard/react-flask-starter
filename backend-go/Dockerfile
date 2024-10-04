# Use the official Golang image as the base image
FROM golang:1.23-bookworm

# Set the Current Working Directory inside the container
WORKDIR /app

# Initialize the go module
RUN go mod init main
# Get 3rd party modules
RUN go get github.com/jackc/pgx/v5 github.com/stoewer/go-strcase golang.org/x/sync/errgroup

# Copy go source
COPY *.go ./

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]