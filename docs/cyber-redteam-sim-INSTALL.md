# Cyber Red Team Simulator — Installation Guide

## Overview

The Cyber Red Team Simulator is a high-fidelity network attack and defense simulation engine. It models 45+ MITRE ATT&CK techniques with probabilistic outcomes, defense modeling (maturity levels, IR playbooks, EDR, SIEM correlation), C2 infrastructure, and IEEE 1278.1 DIS federation.

**Binary:** `cyber-redteam-sim`  
**Version:** 0.2.0  
**License:** MIT  
**Platforms:** Linux (amd64/arm64), macOS (Intel/ARM), Windows (amd64)

---

## Quick Start (Any Platform)

```bash
# Download the binary for your platform from binaries/
# Linux/macOS:
chmod +x cyber-redteam-sim
./cyber-redteam-sim --version

# Initialize configs:
./cyber-redteam-sim --init

# Run a simulation:
./cyber-redteam-sim --config configs/enterprise-ad.yaml --no-server
```

---

## Installation by Platform

### Linux (Debian/Ubuntu) — .deb Package

The easiest way on Debian-based systems:

```bash
# Download the .deb package
wget https://github.com/wezzels/forge-sims/raw/main/binaries/cyber-redteam-sim_0.2.0-1_amd64.deb

# Install
sudo dpkg -i cyber-redteam-sim_0.2.0-1_amd64.deb

# Verify
cyber-redteam-sim --version

# Run a simulation
cyber-redteam-sim --config /etc/cyber-redteam-sim/configs/enterprise-ad.yaml --no-server

# Run as a systemd service
sudo systemctl enable --now cyber-redteam-sim

# Uninstall
sudo dpkg -r cyber-redteam-sim
```

**What gets installed:**
- Binary: `/usr/local/bin/cyber-redteam-sim`
- Configs: `/etc/cyber-redteam-sim/configs/`
- Data dir: `/var/lib/cyber-redteam-sim/`
- Systemd unit: `/lib/systemd/system/cyber-redteam-sim.service`
- Man page: `/usr/share/man/man1/cyber-redteam-sim.1`
- Bash completions: `/etc/bash_completion.d/cyber-redteam-sim`
- Docs: `/usr/local/share/doc/cyber-redteam-sim/`

### Linux — tar.gz

```bash
# Download and extract
tar xzf cyber-redteam-sim-0.2.0-linux-amd64.tar.gz
cd cyber-redteam-sim-0.2.0-linux-amd64

# Install using the included script
sudo ./install.sh

# Or install manually:
sudo cp usr/local/bin/cyber-redteam-sim /usr/local/bin/
sudo mkdir -p /etc/cyber-redteam-sim/configs
sudo cp etc/cyber-redteam-sim/configs/*.yaml /etc/cyber-redteam-sim/configs/

# Uninstall
sudo ./uninstall.sh
```

### macOS — tar.gz

```bash
# Download and extract (ARM for Apple Silicon, amd64 for Intel)
tar xzf cyber-redteam-sim-0.2.0-darwin-arm64.tar.gz
cd cyber-redteam-sim-0.2.0-darwin-arm64

# Install using the included script
sudo ./install.sh

# This installs:
# - Binary to /usr/local/bin/cyber-redteam-sim
# - Configs to /etc/cyber-redteam-sim/configs/
# - launchd plist to /Library/LaunchDaemons/com.cyberredteamsim.plist
# - Zsh and bash completions

# Remove quarantine attribute (macOS)
sudo xattr -d com.apple.quarantine /usr/local/bin/cyber-redteam-sim

# Run as a launchd service
sudo launchctl load -w /Library/LaunchDaemons/com.cyberredteamsim.plist

# Stop service
sudo launchctl unload /Library/LaunchDaemons/com.cyberredteamsim.plist
```

### Windows — .zip

```powershell
# Download and extract the zip
Expand-Archive cyber-redteam-sim-0.2.0-windows-amd64.zip -DestinationPath C:\cyber-redteam-sim

# Option 1: Run directly
.\cyber-redteam-sim.exe --version
.\cyber-redteam-sim.exe --config configs\enterprise-ad.yaml --no-server

# Option 2: Install system-wide (run as Administrator)
.\install.bat

# This installs:
# - Binary to C:\Program Files\CyberRedTeamSim\
# - Configs to C:\ProgramData\cyber-redteam-sim\configs\
# - Adds to system PATH
# - Registers in Add/Remove Programs

# Uninstall
.\uninstall.bat
```

### Linux ARM64 (Raspberry Pi, ARM servers)

```bash
tar xzf cyber-redteam-sim-0.2.0-linux-arm64.tar.gz
cd cyber-redteam-sim-0.2.0-linux-arm64
sudo ./install.sh
```

---

## CLI Reference

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

---

## Scenarios

### Built-in Scenarios

The `--config-wizard` flag offers 4 templates:

| Template | Hosts | Duration | Objective | Description |
|----------|-------|----------|-----------|-------------|
| `enterprise-ad` | 5 | 3600s | domain-admin | Corporate AD compromise |
| `dmz-breach` | 4 | 1800s | data-exfiltration | DMZ web app to internal |
| `cloud-kill-chain` | 4 | 7200s | data-exfiltration | AWS/Azure multi-cloud |
| `apt-persistence` | 4 | 14400s | data-exfiltration | Long-duration APT |

### Validating Scenarios

```bash
# Validate a scenario config
cyber-redteam-sim --validate configs/enterprise-ad.yaml

# Output:
# ✓ Name:          Enterprise AD Compromise
# ✓ Duration:      3600s
# ✓ Tick interval: 60s
# ✓ Maturity:      1
# ✓ Objective:     domain-admin
# ✓ Hosts:         5
# ✓ Zones:         2 (map[dmz:1 internal:4])
# ✓ Services:      12
# ✓ Vulnerabilities: 2 (exploitable)
# ✅ Scenario valid — ready to simulate
```

### Listing Scenarios

```bash
cyber-redteam-sim --list-scenarios

# Output:
# ✓ enterprise-ad.yaml              Enterprise AD Compromise
#                                   5 hosts, 3600s, maturity 1
#                                   objective: domain-admin
```

---

## Running Simulations

### Headless (No Server)

```bash
# Run simulation, print AAR to stdout
cyber-redteam-sim --config configs/enterprise-ad.yaml --no-server
```

### With REST API

```bash
# Start with REST API on port 8080
cyber-redteam-sim --config configs/enterprise-ad.yaml

# Query the API
curl http://localhost:8080/api/v1/status
```

### As a System Service (Linux)

```bash
# Using systemd (installed via .deb)
sudo systemctl start cyber-redteam-sim
sudo systemctl status cyber-redteam-sim

# View logs
journalctl -u cyber-redteam-sim -f
```

### As a launchd Service (macOS)

```bash
# Start
sudo launchctl load -w /Library/LaunchDaemons/com.cyberredteamsim.plist

# Stop
sudo launchctl unload /Library/LaunchDaemons/com.cyberredteamsim.plist

# View logs
tail -f /var/log/cyber-redteam-sim.log
```

---

## Docker

```bash
# Pull and run
docker pull stsgym/cyber-redteam-sim:latest
docker run -p 8080:8080 stsgym/cyber-redteam-sim:latest

# With custom config
docker run -p 8080:8080 -v ./configs:/etc/cyber-redteam-sim/configs stsgym/cyber-redteam-sim:latest

# Headless
docker run stsgym/cyber-redteam-sim:latest --config /etc/cyber-redteam-sim/configs/enterprise-ad.yaml --no-server
```

---

## Health Check

```bash
cyber-redteam-sim --doctor

# Output:
# Version:     0.2.0 (commit: abc1234, built: 2026-04-17T12:00:00Z)
# OS/Arch:     linux/amd64
# Config:       configs ✓
# Scenarios:    1 scenario(s) found
# Scenario:    Enterprise AD Compromise loaded ✓ (5 hosts, 12 services)
# REST API:    Port 8080 (will bind on start)
# ✅ All checks passed. Ready to simulate.
```

---

## Removal

### Linux (.deb)

```bash
sudo dpkg -r cyber-redteam-sim
# Configs preserved at /etc/cyber-redteam-sim/
# To remove everything:
sudo rm -rf /etc/cyber-redteam-sim /var/lib/cyber-redteam-sim
```

### Linux/macOS (tar.gz)

```bash
sudo ./uninstall.sh    # Linux
# Or manually:
sudo rm /usr/local/bin/cyber-redteam-sim
sudo rm -rf /etc/cyber-redteam-sim
sudo rm -rf /var/lib/cyber-redteam-sim
sudo rm /etc/bash_completion.d/cyber-redteam-sim      # Linux
sudo rm /usr/local/share/zsh/site-functions/_cyber-redteam-sim  # macOS
sudo launchctl unload /Library/LaunchDaemons/com.cyberredteamsim.plist  # macOS
```

### Windows

```powershell
.\uninstall.bat    # Run as Administrator
# Removes binary, configs, PATH entry, and registry entries
```

---

## Troubleshooting

### "Config directory not found"
```bash
cyber-redteam-sim --init    # Creates configs/ with sample scenarios
```

### "Permission denied" on Linux/macOS
```bash
chmod +x cyber-redteam-sim
# On macOS, also:
sudo xattr -d com.apple.quarantine cyber-redteam-sim
```

### Port 8080 already in use
```bash
cyber-redteam-sim --config configs/enterprise-ad.yaml --rest-port 9090
```

### View simulation results
```bash
# Headless mode prints full After-Action Review to stdout
cyber-redteam-sim --config configs/enterprise-ad.yaml --no-server 2>&1 | tee results.txt
```