# Reward Chain Testnet

Reward Chain Testnet is a Cosmos SDK-based blockchain scaffolded with Ignite CLI and prepared for iterative development around reward operations, partner management, liquidity flows, and future IBC/EVM compatibility work.

## Current status

This repository currently contains the base chain scaffold and local development setup.

### Chain identity
- **Chain name:** `rewardchain`
- **Binary:** `rewardchaind`
- **Address prefix:** `reward`
- **Default denom:** `stake`

## Prerequisites

Make sure these tools are installed:
- Go `1.25.7+`
- Ignite CLI
- Git

Optional but recommended:
- Docker
- Docker Compose

## Local development

### Start the chain with Ignite

```bash
ignite chain serve
```

This command installs dependencies, builds the chain, initializes local state, and starts a development node.

### Install the chain binary locally

```bash
make install
```

### Build the binary

```bash
make build
```

### Run tests

```bash
make test
```

### Run unit tests only

```bash
make test-unit
```

### Generate protobuf files

```bash
make proto-gen
```

## Container workflow

### Build the Docker image

```bash
docker build -t rewardchain:local .
```

### Start with Docker Compose

```bash
docker compose up --build
```

See `docs/deployment.md` for local deployment and container usage notes.

## Configuration

Local chain configuration lives in `config.yml`.

Current scaffold defaults include:
- validator accounts
- faucet account
- default stake balances
- OpenAPI generation output path

## Project goals

The roadmap for this repository includes:
- reward-focused chain configuration and docs
- partner management flows
- liquidity operations
- reward issuance and burn flows
- reward swap logic
- IBC compatibility
- EVM compatibility planning and implementation
- deployment and operations documentation

## Repository structure

- `app/` — application wiring and chain configuration
- `cmd/rewardchaind/` — node binary entrypoints and CLI commands
- `proto/` — protobuf configuration
- `docs/` — generated and supporting documentation
- `config.yml` — Ignite development chain config
- `Makefile` — common developer commands

## Notes

This project started from an Ignite scaffold and is being customized incrementally. Commits should remain focused, descriptive, and tied to meaningful project progress.
