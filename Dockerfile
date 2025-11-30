# Build stage
FROM golang:1.24.1-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o /app/bin/bikes-compass \
    ./cmd/api/main.go

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/bin/bikes-compass .

# Copy any necessary files (if needed)
# COPY --from=builder /app/queries ./queries

# Change ownership
RUN chown -R appuser:appuser /home/appuser

# Switch to non-root user
USER appuser

# Expose port (default 8081 based on swagger docs)
EXPOSE 8081

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://${HOST}:${PORT}/v1/health || exit 1

# Run the application
CMD ["./bikes-compass"]
