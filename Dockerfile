# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend-pos .

RUN ls -lh backend-pos

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

COPY --from=builder --chown=userapp:binarygroup /app/backend-pos .

RUN chmod +x backend-pos

USER userapp

EXPOSE 8085

CMD ["./backend-pos"]