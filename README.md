# Uptime Sentinel

A lightweight, modern infrastructure monitoring tool written in Go.

## Features

- Real-time HTTP health checks
- In-memory status storage
- JSON API for integration with dashboards
- Concurrency-first design

## Getting Started

### Prerequisites

- Go 1.16+

### Installation

1. Clone the repository
2. Run `go mod tidy` to install dependencies

### Configuration (IMPORTANT)

The application requires a configuration file to know which services to monitor.

**You MUST create a `targets.json` file in the `configs/` directory before running the application.**

A template is provided:

```bash
cp configs/targets.json.example configs/targets.json
```

If you skip this step, the application will fail to start with a configuration error.

### Running the Server

```bash
go run cmd/server/main.go
```

The server will start on port `8080`.

### API

- `GET /status`: Returns the current status of all monitored targets.

## Development

This project is a work in progress.
TODO:

- [ ] Add support for a custom interval
- [ ] Implement persistent storage (Redis/Postgres)
- [ ] Fix race conditions in the store map
- [ ] Add Dockerfile
