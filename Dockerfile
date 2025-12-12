# Use a specific version of Go with Alpine for building
FROM golang:1.23.3-alpine AS builder

# Install build tools and dependencies (cgo + SQLite headers for go-sqlite3)
RUN apk add --no-cache git build-base sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set environment variables for Go (cgo enabled so go-sqlite3 works)
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# Build the Go application
RUN go build -o main ./cmd

# Use an Alpine image for runtime
FROM alpine:latest

# Install CA certificates for HTTPS and SQLite runtime libs (if linked dynamically)
RUN apk add --no-cache ca-certificates sqlite-libs

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080 8081

# Command to run the binary
ENTRYPOINT ["./main"]
