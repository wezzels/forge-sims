# Space War Sim

High-fidelity space domain warfare simulation engine.

## Quick Install

### Linux (Debian/Ubuntu)
```bash
sudo dpkg -i space-war-sim_0.1.0-1_amd64.deb
space-war-sim --version
```

### macOS / Linux (tar.gz)
```bash
tar xzf space-war-sim-0.1.0-linux-amd64.tar.gz
sudo ./install.sh
```

### Windows
Extract the `.zip` and run `install.bat` as Administrator.

## Usage

```bash
# Run a simulation
space-war-sim --config configs/scenario.yaml

# With web UI
space-war-sim --config configs/scenario.yaml --web :8080

# Export AAR
space-war-sim --config configs/scenario.yaml --aar results.json

# Health check
space-war-sim --doctor

# Initialize configs
space-war-sim --init

# Validate scenario
space-war-sim --validate configs/scenario.yaml

# List scenarios
space-war-sim --list-scenarios

# Print version
space-war-sim --version
```

## CLI Reference

| Flag | Description |
|------|-------------|
| `--config PATH` | Scenario config file (default: `configs/scenario.yaml`) |
| `--aar PATH` | AAR JSON export path |
| `--web ADDR` | Web UI address (e.g. `:8080`) |
| `--version` | Print version, commit, build date, OS/arch |
| `--doctor` | Run health check diagnostics |
| `--init` | Create config directory with sample scenarios |
| `--validate PATH` | Validate a scenario config file |
| `--list-scenarios` | List available scenario files |

## Packages

| Package | Description |
|---------|-------------|
| `internal/core` | Engine, event bus, config loading, orbital state |
| `internal/orbital` | Keplerian mechanics, J2 perturbation, Hohmann transfers, conjunction analysis |
| `internal/ssa` | Space situational awareness, tracking, sensor models |
| `internal/asat` | ASAT weapons (DA-ASAT, co-orbital, laser, EMP), engagement, debris generation |
| `internal/debris` | Debris field modeling, fragmentation events, Kessler syndrome |
| `internal/ew` | Electronic warfare (jamming, spoofing, link disruption) |
| `internal/c2` | C2/ROE engine, space tasking orders |
| `internal/groundstation` | Ground station tracking, pass windows, link budgets |
| `internal/satellite` | Satellite platforms (LEO, MEO, GEO, constellations) |
| `internal/scenario` | YAML scenario loading and validation |
| `internal/aar` | After-action review, recording, JSON export |

## Build from Source

```bash
make build          # Build for current platform
make test           # Run all tests
make dist           # Cross-compile all platforms + package + checksum
make deb            # Build .deb package
make docker          # Build Docker image
make clean           # Remove built binary
make dist-clean      # Remove entire dist/ directory
```

## Scenarios

| Scenario | Entities | Duration | Description |
|----------|----------|----------|-------------|
| `scenario.yaml` | 8 | 7200s | South China Sea space domain (SSA, ASAT, EW, constellations) |
| `engagement.yaml` | 4 | 3600s | DA-ASAT engagement with debris modeling |

## License

MIT