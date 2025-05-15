# ── сборка ──
FROM golang:1.24-bullseye AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# копируем код
COPY cmd/statsbot ./cmd/statsbot
COPY internal ./internal
COPY config/config.toml ./config/config.toml

# собираем статический бинарь
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" \
    -o statsbot ./cmd/statsbot

# ── финальный образ ──
FROM ubuntu:24.04
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /build/statsbot .
COPY --from=builder /build/config/config.toml ./config.toml

EXPOSE 8080
ENTRYPOINT ["./statsbot", "-config", "config.toml"]
