# Use a minimal base image - no Go needed, binary is pre-built in CI
FROM alpine:latest

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from the CI runner
COPY server .

# Copy the .env file if available (useful for local testing)
COPY .env* ./

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./server"]
