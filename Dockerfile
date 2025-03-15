# Build stage: compile the Go application
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary (adjust the flags and output name as necessary)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage: minimal image to run the application
FROM alpine:latest
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port that your application listens on (adjust as needed)
EXPOSE 8080

# Run the Go binary
CMD ["./app"]
