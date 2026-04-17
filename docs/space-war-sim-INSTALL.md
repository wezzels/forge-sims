# Space War Sim — Installation Guide

## Overview

Space War Sim is a high-fidelity space domain warfare simulation engine modeling orbital mechanics, SSA tracking, ASAT weapons, debris fields, electronic warfare, C2/ROE, and ground station link budgets.

**Binary:** `space-war-sim`
**Version:** 0.1.0
**License:** MIT
**Platforms:** Linux (amd64/arm64), macOS (Intel/ARM), Windows (amd64)

---

## Quick Start

```bash
chmod +x space-war-sim
./space-war-sim --version
./space-war-sim --init
./space-war-sim --config configs/scenario.yaml
```

---

## Installation by Platform

### Linux (.deb)
```bash
sudo dpkg -i space-war-sim_0.1.0-1_amd64.deb
space-war-sim --doctor
```

### Linux/macOS (tar.gz)
```bash
tar xzf space-war-sim-0.1.0-linux-amd64.tar.gz
sudo ./install.sh
```

### macOS (launchd)
```bash
sudo ./install.sh
sudo launchctl load -w /Library/LaunchDaemons/com.spacewarsim.plist
```

### Windows
```powershell
.\install.bat    # Run as Administrator
```

---

## CLI Reference

| Flag | Description |
|------|-------------|
| `--config PATH` | Scenario config file (default: `configs/scenario.yaml`) |
| `--aar PATH` | AAR JSON export path |
| `--web ADDR` | Web UI address |
| `--version` | Print version, commit, build date, OS/arch |
| `--doctor` | Health check diagnostics |
| `--init` | Create config directory with sample scenarios |
| `--validate PATH` | Validate a scenario config file |
| `--list-scenarios` | List available scenario files |

---

## Removal

### Linux (.deb)
```bash
sudo dpkg -r space-war-sim
sudo rm -rf /etc/space-war-sim /var/lib/space-war-sim  # optional
```

### Linux/macOS
```bash
sudo ./uninstall.sh
```

### Windows
```powershell
.\uninstall.bat    # Run as Administrator
```

---

## Troubleshooting

```bash
space-war-sim --init          # Create configs if missing
space-war-sim --doctor        # Check installation
space-war-sim --validate configs/scenario.yaml  # Validate scenario
```