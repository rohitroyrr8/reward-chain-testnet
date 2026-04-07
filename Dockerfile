FROM golang:1.25.7-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /out && go build -o /out/rewardchaind ./cmd/rewardchaind

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y ca-certificates curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root
COPY --from=builder /out/rewardchaind /usr/local/bin/rewardchaind

EXPOSE 26656 26657 1317 9090 9091

CMD ["rewardchaind", "start"]
