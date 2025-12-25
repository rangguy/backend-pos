# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o backend-pos ./cmd

# Stage 2: Create the production image
FROM alpine:latest

# Install dependencies
RUN apk add --no-cache tzdata ca-certificates

# Set timezone
ENV TZ=Asia/Jakarta
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

# Set working directory
WORKDIR /app

# Create user and group
RUN addgroup -g 1001 binarygroup && \
    adduser -D -u 1001 -G binarygroup userapp

# Create necessary directories
RUN mkdir -p /app/uploads /app/temp /app/config && \
    chown -R userapp:binarygroup /app

# Copy binary from builder
COPY --from=builder --chown=userapp:binarygroup /app/backend-pos .

# Switch to non-root user
USER userapp

# Expose port
EXPOSE 8085

# Health check (optional, sesuaikan endpoint)
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8085/health || exit 1

# Run application
CMD ["./backend-pos"]