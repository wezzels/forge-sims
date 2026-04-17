# Maritime Sim

High-fidelity maritime warfare simulation engine.

## Quick Install

### Linux (Debian/Ubuntu)
```bash
sudo dpkg -i maritime-sim_0.1.0-1_amd64.deb
maritime-sim --version
```

### macOS / Linux (tar.gz)
```bash
tar xzf maritime-sim-0.1.0-linux-amd64.tar.gz
sudo ./install.sh
```

### Windows
Extract the `.zip` and run `install.bat` as Administrator.

## Usage

```bash
# Run a simulation
maritime-sim --config configs/scenario.yaml

# Run with web UI (Leaflet.js tactical display)
maritime-sim --config configs/scenario.yaml --web :8080

# Export After-Action Review
maritime-sim --config configs/scenario.yaml --aar results.json

# Health check
maritime-sim --doctor

# Initialize config directory
maritime-sim --init

# Validate a scenario config
maritime-sim --validate configs/scenario.yaml

# List available scenarios
maritime-sim --list-scenarios

# Print version
maritime-sim --version
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
| `--config PATH` | Scenario config file (default: `configs/scenario.yaml`) |
| `--aar PATH` | AAR JSON export path (empty = no export) |
| `--web ADDR` | Web UI address (e.g. `:8080`, empty = no web UI) |
| `--version` | Print version, commit, build date, OS/arch |
| `--doctor` | Run health check diagnostics |
| `--init` | Create config directory with sample scenarios |
| `--validate PATH` | Validate a scenario config file |
| `--list-scenarios` | List available scenario files |

## Packages

| Package | Description |
|---------|-------------|
| `internal/core` | Simulation engine, event bus, config loading |
| `internal/platform/surface` | Surface combatant models (DDG/FFG/CVN) |
| `internal/platform/submarine` | Submarine models (SSN/SSK) |
| `internal/sensor/acoustic` | Mackenzie sound speed, sonar equation, transmission loss |
| `internal/sensor/radar` | Radar equation, horizon, detection |
| `internal/sensor/sensormgr` | Active sonar scheduling, towed array |
| `internal/sensor/trackmgr` | Alpha-beta filtering, classification, bearing-only TMA |
| `internal/sensor/signature` | Aspect-dependent target strength |
| `internal/ai` | Commander AI, threat assessment, decision tree |
| `internal/comms/c2` | C2/ROE engine, weapons free/tight/hold |
| `internal/scenario/msel` | MSEL timeline engine |
| `internal/scenario/aar` | AAR recorder, replay, scoring, JSON export |
| `internal/weapon/ashm` | Anti-ship missile flyout model |
| `internal/weapon/torpedo` | Torpedo runout/search/homing |
| `internal/weapon/damage` | Damage model (compartments, flooding) |
| `internal/weapon/engagement` | Engagement manager, ROE check |
| `internal/ew/ecm` | ECM chaff/jamming |
| `internal/ew/esm` | ESM bearings |
| `internal/hydro` | Pierson-Moskowitz spectra, seakeeping |
| `internal/fed/dis` | IEEE 1278.1 DIS federation bridge |
| `internal/webui` | Leaflet.js tactical display with SSE |

## Distribution

| Platform | Format | Size |
|----------|--------|------|
| Linux amd64 | `.tar.gz` / `.deb` | ~4.6 MB |
| Linux arm64 | `.tar.gz` | ~4.3 MB |
| macOS Intel | `.tar.gz` | ~4.6 MB |
| macOS ARM | `.tar.gz` | ~4.4 MB |
| Windows amd64 | `.zip` | ~4.7 MB |
| Docker | `stsgym/maritime-sim:latest` | Alpine-based |

## License

MIT