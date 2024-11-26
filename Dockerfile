FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p ./bin && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/migrate ./cmd/migrate/main.go && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/image_service ./cmd/image_service/main.go

FROM --platform=linux/arm64 debian:bullseye-slim

WORKDIR /app

RUN apt-get update && apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin /app/bin

COPY ./config /app/config
COPY ./migrations /app/migrations

EXPOSE 50051

CMD ["/app/bin/image_service", "--config=/app/config/local.yaml"]
