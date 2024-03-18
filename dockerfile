FROM golang:1.21.5 AS builder

WORKDIR /app

RUN mkdir -p cmd
RUN mkdir -p internal/config
RUN mkdir -p internal/middleware
RUN mkdir -p internal/limiter
RUN mkdir -p internal/bd

COPY go.mod ./
COPY go.sum ./
COPY cmd/main.go ./cmd
COPY internal/config/config.go ./internal/config
COPY internal/config/config_test.go ./internal/config
COPY internal/middleware/limiter_middleware.go ./internal/middleware
COPY internal/middleware/limiter_middleware_test.go ./internal/middleware
COPY internal/limiter/limiter_interface.go ./internal/limiter
COPY internal/limiter/limiter.go ./internal/limiter
COPY internal/limiter/limiter_test.go ./internal/limiter
COPY internal/bd/redis_client.go ./internal/bd
COPY internal/bd/bd_interface.go ./internal/bd

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o go-rate-limit cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-rate-limit ./

CMD ["./go-rate-limit"]
