# Use a specific version of Go with Alpine for building
FROM golang:1.23.3-alpine AS builder

# Install build tools and dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the Go application
RUN go build -o main ./cmd

# Use an Alpine image for runtime
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["./main"]
