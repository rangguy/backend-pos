# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copy dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy entire source code
COPY . .

# Build from root (main.go imports cmd/)
RUN go build -v -o backend-pos . && \
    chmod +x backend-pos

# Verify binary created and architecture
RUN ls -lh backend-pos && \
    file backend-pos

# Stage 2: Production image
FROM alpine:latest

RUN apk add --no-cache tzdata ca-certificates

ENV TZ=Asia/Jakarta
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app

RUN addgroup -g 1001 binarygroup && \
    adduser -D -u 1001 -G binarygroup userapp && \
    mkdir -p /app/uploads /app/temp /app/config && \
    chown -R userapp:binarygroup /app

# Copy binary from builder
COPY --from=builder --chown=userapp:binarygroup /app/backend-pos .

# Verify and set executable
RUN file /app/backend-pos && \
    chmod +x /app/backend-pos

USER userapp

EXPOSE 8085

CMD ["./backend-pos"]