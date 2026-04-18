# Space War Sim — Installation Guide

## Quick Start

### Linux (amd64)

```bash
# Download and extract
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/linux-x86/space-war-sim -o space-war-sim
chmod +x space-war-sim
sudo mv space-war-sim /usr/local/bin/

# Or install via .deb
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/space-war-sim_0.1.0-1_amd64.deb -o space-war-sim.deb
sudo dpkg -i space-war-sim.deb
```

### Linux (arm64)

```bash
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/linux-arm64/space-war-sim -o space-war-sim
chmod +x space-war-sim
sudo mv space-war-sim /usr/local/bin/
```

### macOS (Intel)

```bash
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/darwin-amd64/space-war-sim -o space-war-sim
chmod +x space-war-sim
sudo mv space-war-sim /usr/local/bin/
```

### macOS (Apple Silicon)

```bash
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/darwin-arm64/space-war-sim -o space-war-sim
chmod +x space-war-sim
sudo mv space-war-sim /usr/local/bin/
```

### Windows

```powershell
# Download
Invoke-WebRequest -Uri https://github.com/wezzels/forge-sims/raw/main/binaries/windows-amd64/space-war-sim.exe -OutFile space-war-sim.exe

# Add to PATH or run directly
.\space-war-sim.exe --help
```

## Build from Source

### Prerequisites

- Go 1.21 or later
- Git

### Steps

```bash
git clone https://idm.wezzel.com/crab-meat-repos/space-war-sim.git
cd space-war-sim

# Build
go build -o space-war-sim ./cmd/sim

# Run tests
go test ./... -v

# Install
go install ./cmd/sim
```

### Cross-compile

```bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -o dist/space-war-sim ./cmd/sim

# macOS arm64
GOOS=darwin GOARCH=arm64 go build -o dist/space-war-sim ./cmd/sim

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/space-war-sim.exe ./cmd/sim
```

## Verify Installation

```bash
space-war-sim --version
# space-war-sim v0.1.0 (commit: <hash>, built: <date>)

# Health check
space-war-sim --doctor

# Initialize sample configs
space-war-sim --init

# Run a scenario
space-war-sim --config configs/scenario.yaml --aar results.json
```

## Docker

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o space-war-sim ./cmd/sim

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/space-war-sim /usr/local/bin/
ENTRYPOINT ["space-war-sim"]
```

## Troubleshooting

| Issue | Solution |
|-------|----------|
| `Permission denied` | `chmod +x space-war-sim` |
| `command not found` | Add to PATH or use full path |
| Config not found | Run `space-war-sim --init` |
| Port conflicts | Check `configs/scenario.yaml` for port settings |