# Deployment and Local Operations

## Overview

This document covers the current local/container workflow for Reward Chain Testnet.

## Local development

### Start with Ignite

```bash
ignite chain serve
```

### Build the binary

```bash
make build
```

### Install locally

```bash
make install
```

## Docker

### Build the image

```bash
docker build -t rewardchain:local .
```

### Run the binary in a container

```bash
docker run --rm -p 26656:26656 -p 26657:26657 -p 1317:1317 -p 9090:9090 -p 9091:9091 rewardchain:local
```

## Docker Compose

### Start the containerized node

```bash
docker compose up --build
```

### Stop it

```bash
docker compose down
```

## Ports

- `26656` — P2P
- `26657` — Tendermint/CometBFT RPC
- `1317` — REST API
- `9090` — gRPC
- `9091` — gRPC-web / auxiliary endpoint

## Notes

This is a foundation deployment setup for local and testnet-oriented development. Production hardening, persistent infra design, observability, secrets handling, and validator operations should be documented separately as the chain evolves.
