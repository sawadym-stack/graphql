# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=direct && go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .
# Copy the .env file if available (useful for local testing the container)
COPY --from=builder /app/.env* ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
