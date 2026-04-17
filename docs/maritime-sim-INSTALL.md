# Maritime Sim — Installation Guide

## Overview

Maritime Sim is a high-fidelity maritime warfare simulation engine modeling surface combatants, submarines, sensors (active/passive sonar, radar, ESM), weapons (AShM, torpedoes), AI commanders, C2/ROE, and IEEE 1278.1 DIS federation.

**Binary:** `maritime-sim`  
**Version:** 0.1.0  
**License:** MIT  
**Platforms:** Linux (amd64/arm64), macOS (Intel/ARM), Windows (amd64)

---

## Quick Start (Any Platform)

```bash
# Download the binary for your platform from binaries/
# Linux/macOS:
chmod +x maritime-sim
./maritime-sim --version

# Initialize configs:
./maritime-sim --init

# Run a simulation:
./maritime-sim --config configs/scenario.yaml
```

---

## Installation by Platform

### Linux (Debian/Ubuntu) — .deb Package

```bash
# Download the .deb package
wget https://github.com/wezzels/forge-sims/raw/main/binaries/maritime-sim_0.1.0-1_amd64.deb

# Install
sudo dpkg -i maritime-sim_0.1.0-1_amd64.deb

# Verify
maritime-sim --version

# Run a simulation
maritime-sim --config /etc/maritime-sim/configs/scenario.yaml

# Run as a systemd service
sudo systemctl enable --now maritime-sim

# Uninstall
sudo dpkg -r maritime-sim
```

**What gets installed:**
- Binary: `/usr/local/bin/maritime-sim`
- Configs: `/etc/maritime-sim/configs/`
- Systemd unit: `/lib/systemd/system/maritime-sim.service`
- Man page: `/usr/share/man/man1/maritime-sim.1.gz`

### Linux — tar.gz

```bash
# Download and extract
tar xzf maritime-sim-0.1.0-linux-amd64.tar.gz
cd maritime-sim-0.1.0-linux-amd64

# Install using the included script
sudo ./install.sh

# Uninstall
sudo ./uninstall.sh
```

### macOS — tar.gz

```bash
# Download and extract (ARM for Apple Silicon, amd64 for Intel)
tar xzf maritime-sim-0.1.0-darwin-arm64.tar.gz
cd maritime-sim-0.1.0-darwin-arm64

# Install using the included script
sudo ./install.sh

# Remove quarantine attribute
sudo xattr -d com.apple.quarantine /usr/local/bin/maritime-sim

# Run as a launchd service
sudo launchctl load -w /Library/LaunchDaemons/com.maritimesim.plist

# Stop service
sudo launchctl unload /Library/LaunchDaemons/com.maritimesim.plist
```

### Windows — .zip

```powershell
# Download and extract the zip
Expand-Archive maritime-sim-0.1.0-windows-amd64.zip -DestinationPath C:\maritime-sim

# Option 1: Run directly
.\maritime-sim.exe --version
.\maritime-sim.exe --config configs\scenario.yaml

# Option 2: Install system-wide (run as Administrator)
.\install.bat

# Uninstall
.\uninstall.bat
```

### Linux ARM64 (Raspberry Pi, ARM servers)

```bash
tar xzf maritime-sim-0.1.0-linux-arm64.tar.gz
cd maritime-sim-0.1.0-linux-arm64
sudo ./install.sh
```

---

## CLI Reference

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

---

## Running Simulations

### Headless
```bash
maritime-sim --config configs/scenario.yaml
```

### With Web UI (Leaflet.js tactical display)
```bash
maritime-sim --config configs/scenario.yaml --web :8080
# Open http://localhost:8080 in browser
```

### Export After-Action Review
```bash
maritime-sim --config configs/scenario.yaml --aar results.json
```

### As a System Service (Linux)
```bash
sudo systemctl start maritime-sim
sudo systemctl status maritime-sim
journalctl -u maritime-sim -f
```

---

## Docker

```bash
docker pull stsgym/maritime-sim:latest
docker run -p 8080:8080 stsgym/maritime-sim:latest
```

---

## Health Check

```bash
maritime-sim --doctor

# Output:
# Version:     0.1.0 (commit: abc1234, built: 2026-04-17T12:00:00Z)
# OS/Arch:    linux/amd64
# Config:     configs/scenario.yaml ✓
# Scenarios:  2 scenario(s) found
# Scenario:    South China Sea ASW Scenario loaded ✓ (3 entities, 7200s)
# ✅ All checks passed. Ready to simulate.
```

---

## Removal

### Linux (.deb)
```bash
sudo dpkg -r maritime-sim
sudo rm -rf /etc/maritime-sim /var/lib/maritime-sim  # optional
```

### Linux/macOS (tar.gz)
```bash
sudo ./uninstall.sh
```

### Windows
```powershell
.\uninstall.bat    # Run as Administrator
```

---

## Troubleshooting

### "Config directory not found"
```bash
maritime-sim --init
```

### "Permission denied"
```bash
chmod +x maritime-sim
# macOS:
sudo xattr -d com.apple.quarantine maritime-sim
```

### Port 8080 in use
```bash
maritime-sim --config configs/scenario.yaml --web :9090
```