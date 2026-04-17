# Cyber Red Team Simulator

High-fidelity network attack and defense simulation engine.

## Quick Install

### Linux (Debian/Ubuntu)
```bash
sudo dpkg -i cyber-redteam-sim_0.2.0-1_amd64.deb
cyber-redteam-sim --version
```

### macOS (Homebrew)
```bash
brew install wezzels/tap/cyber-redteam-sim
```

### macOS / Linux (tar.gz)
```bash
tar xzf cyber-redteam-sim-0.2.0-linux-amd64.tar.gz
sudo ./install.sh
```

### Windows
Extract the `.zip` and run `install.bat` as Administrator.

## Usage

```bash
# Run a simulation
cyber-redteam-sim --config configs/enterprise-ad.yaml --no-server

# Run with REST API
cyber-redteam-sim --config configs/enterprise-ad.yaml

# Health check
cyber-redteam-sim --doctor

# Initialize config directory
cyber-redteam-sim --init

# Validate a scenario config
cyber-redteam-sim --validate configs/enterprise-ad.yaml

# List available scenarios
cyber-redteam-sim --list-scenarios

# Interactive scenario wizard
cyber-redteam-sim --config-wizard

# Check for updates
cyber-redteam-sim --update-check

# Print version
cyber-redteam-sim --version
```

## Build from Source

```bash
make build          # Build for current platform
make test           # Run all tests
make dist            # Cross-compile all platforms + package + checksum
make deb             # Build .deb package
make docker          # Build Docker image
make cross GOOS=darwin GOARCH=arm64  # Single platform
make clean           # Remove built binary
make dist-clean      # Remove entire dist/ directory
```

## CLI Flags

| Flag | Description |
|------|-------------|
| `--config PATH` | Scenario config file (default: `configs/enterprise-ad.yaml`) |
| `--no-server` | Disable API servers (print AAR to stdout) |
| `--rest-port PORT` | REST API port (default: 8080) |
| `--version` | Print version, commit, build date, OS/arch |
| `--doctor` | Run health check diagnostics |
| `--init` | Create config directory with sample scenarios |
| `--validate PATH` | Validate a scenario config file |
| `--list-scenarios` | List available scenario files |
| `--config-wizard` | Interactive scenario generator |
| `--update-check` | Check for newer versions |

## Packages

| Package | Description |
|---------|-------------|
| `internal/engine` | Core simulation engine |
| `internal/attack` | 45+ MITRE ATT&CK techniques |
| `internal/defense` | Defense modeling, IR playbooks, EDR, SIEM |
| `internal/network` | Network topology, hosts, services, zones |
| `internal/attackgraph` | A* pathfinding, tradeoff analysis |
| `internal/ml` | Q-Learning and REINFORCE agents |
| `internal/stealth` | Cumulative detection probability |
| `internal/cloud` | Probabilistic cloud attack modeling |
| `internal/c2` | C2 infrastructure (redirectors, implants, dead drops) |
| `internal/aar` | After-action review, scoring, replay |
| `internal/team` | Purple team, observer mode, objectives |
| `internal/scenario` | YAML scenario loading and validation |
| `internal/editor` | Interactive scenario editor |
| `internal/federation` | IEEE 1278.1 DIS bridge |
| `internal/api` | REST API server |
| `internal/grpcapi` | JSON-RPC API server |
| `internal/webui` | Leaflet.js tactical display with SSE |
| `internal/replay` | Time-travel replay |
| `internal/benchmark` | Performance benchmarking |

## Distribution

| Platform | Format | Size |
|----------|--------|------|
| Linux amd64 | `.tar.gz` / `.deb` | ~2.8 MB |
| Linux arm64 | `.tar.gz` | ~2.5 MB |
| macOS Intel | `.tar.gz` | ~2.8 MB |
| macOS ARM | `.tar.gz` | ~2.6 MB |
| Windows amd64 | `.zip` | ~2.8 MB |
| Docker | `stsgym/cyber-redteam-sim:latest` | Alpine-based |

## License

MIT