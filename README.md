# Forge-Sims

**Ballistic Missile Defense System simulation suite — 47 binaries, physically accurate, JSON-native.**

## Overview

Forge-Sims is a comprehensive simulation suite for the US Ballistic Missile Defense System (BMDS). It covers every layer of the kill chain: satellite-based early warning, ground/sea radars, C2 networks, interceptor weapons, threat modeling, and post-intercept assessment.

All simulations output machine-readable JSON for integration with higher-level wargaming frameworks, AI planners, and visualization systems.

### Quick Start

```bash
# Run any sim with JSON output
./binaries/linux-x86/bmd-sim-gmd -json -duration 1s

# Interactive mode (live dashboard)
./binaries/linux-x86/bmd-sim-thaad -i

# With seed for reproducibility
./binaries/linux-x86/bmd-sim-icbm -json -seed 42
```

## Binary Count: 47

| Category | Count | Description |
|----------|-------|-------------|
| **BMDS Sensors** | 10 | Radars and space-based IR sensors |
| **BMDS Threats** | 9 | ICBM, IRBM, MRBM, HGV, SLCM, decoys, EW, nuclear effects, UAP |
| **BMDS Interceptors** | 6 | GMD/GBI, SM-3, SM-6, THAAD, THAAD-ER, Patriot |
| **C2/Infrastructure** | 8 | C2BMC, GFCB, comms links, engagement chain, kill assessment, WTA |
| **Support** | 5 | Air traffic, space weather, satellite tracker, space debris, launch vehicles |
| **Config-driven** | 3 | Cyber red team, maritime, space warfare (YAML scenario files) |
| **Engine** | 3 | Electronic warfare, missile defense, submarine warfare (Go libraries) |

## New in This Release

### 7 New Simulation Binaries
- **ufo** — Unknown Anomalous Phenomena threat simulator (physics-violating detection/engagement failures)
- **bmd-sim-nuclear-efx** — Nuclear detonation effects (EMP, radar degradation, satellite kill, ionospheric blackout)
- **bmd-sim-electronic-attack** — ECM vs BMDS radars (DRFM, RGPO/VGPO, chaff, barrage/spot jamming)
- **tactical-net** — BMDS RF mesh network simulation with propagation and EW effects
- **engagement-chain** — Full kill chain from boost detection through intercept
- **kill-assessment** — Post-intercept damage assessment (warhead damage, structural, fuzing, guidance)
- **wta** — Weapon-Target Assignment optimizer with doctrine and ROE modeling

### 3 Config-Driven Sims
- **cyber-redteam-sim** — Enterprise network cyber attack/defense scenarios
- **maritime-sim** — Naval surface/subsurface warfare scenarios
- **space-war-sim** — Orbital warfare and ASAT scenarios

These use YAML config files and export AAR JSON via `-aar` flag rather than streaming `-json`.

## NORAD Integration

The **satellite-tracker** binary integrates with NORAD/Celestrak for live orbital data:

```bash
# Track GPS constellation from live TLEs
./satellite-tracker -group gps-ops -json

# Track Starlink
./satellite-tracker -group starlink -json

# Track all active satellites
./satellite-tracker -group active -json
```

Supported groups: `active`, `stations`, `starlink`, `gps-ops`, `weather`, `geo`, `military`, `intelsat`, `resource`, `sarsat`, `argos`, `planet`, `dmc`, `orbcomm`, `x-comm`

Orbital propagation uses SGP4/SDP4 models with live TLE fetches from Celestrak.

## Physics Accuracy

All sensor and interceptor parameters are derived from open-source technical data:

- **Radar ranges**: Calculated from NEFD, aperture, frequency, RCS, and system losses using the radar equation
- **Interceptor PK**: Derived from closing velocity, divert capability, seeker range, and lethal radius
- **Threat trajectories**: Keplerian mechanics with proper boost/midcourse/terminal phase modeling
- **Nuclear effects**: Validated EMP, thermal blast, ionospheric scintillation, and satellite kill models
- **EW modeling**: J/S ratios, burnthrough ranges, DRFM delay, and ECCM technique effects

See [API_REFERENCE.md](API_REFERENCE.md) for complete JSON schemas and example output for every binary.

## Documentation

| File | Description |
|------|-------------|
| [API_REFERENCE.md](API_REFERENCE.md) | JSON output schema for all 47 binaries with examples |
| [BINARY_STATUS.md](BINARY_STATUS.md) | Flag support and clean JSON status table |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Build process, testing, and PR guidelines |
| [NORAD-INTEGRATION.md](NORAD-INTEGRATION.md) | Live orbital data integration details |
| [EVALUATION-2026-04-21.md](EVALUATION-2026-04-21.md) | Round 11 evaluation results |

## Common Flags

Most `bmd-sim-*` binaries support:

| Flag | Description |
|------|-------------|
| `-json` | Machine-readable JSON output |
| `-i` | Interactive mode (live terminal dashboard) |
| `-duration DURATION` | Run duration (e.g., `30s`, `5m`) |
| `-seed int` | Random seed (0 = non-deterministic) |
| `-v` | Verbose output (selected binaries) |
| `-help` | Show help and flag descriptions |

## Repository Structure

```
forge-sims/
├── binaries/linux-x86/     # 47 compiled Linux x86_64 binaries
├── configs/                # YAML scenario configs for config-driven sims
├── docs/                   # Additional documentation
├── scripts/                # Build and deploy scripts
├── API_REFERENCE.md        # Complete JSON schema reference
├── BINARY_STATUS.md        # Binary flag/status table
├── CONTRIBUTING.md         # Contributing guide
├── README.md               # This file
└── EVALUATION-*.md         # Evaluation history
```

## Source Code

This repository contains binaries and documentation only. Source code lives on IDM GitLab:

```
git@idm.wezzel.com:crab-meat-repos/<sim-name>.git
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for the build process.