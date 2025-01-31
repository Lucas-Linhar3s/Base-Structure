# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . /app

# Swag init to generate docs
# RUN swag init  -g ./cmd/server/main.go -o ./server/docs --parseDependency

# Build the Go application
RUN go build -ldflags="-s -w" -o ./cmd/bin/server ./cmd/server/main.go

# Expose the port that the application will run on
EXPOSE 8000

# Command to run the executable
CMD ["./cmd/bin/server"]