# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy your Go service source code into the container
COPY . .

# Expose the gRPC service port
EXPOSE 3001

# Command to run the binary
CMD ["go", "run", "/app/server/server.go"]