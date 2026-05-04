# Stage 1: Build the binary inside a Go container
FROM golang:1.23 AS builder

WORKDIR /app

# Copy dependency files and the vendor folder
COPY go.mod go.sum ./
COPY vendor ./vendor

# Copy the rest of the source code
COPY . .

# Build the binary using the vendor folder
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o server server.go

# Stage 2: Create the final tiny image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/server .
COPY .env* ./

EXPOSE 8080
CMD ["./server"]
