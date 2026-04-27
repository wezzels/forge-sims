# Forge-Sims

**Ballistic Missile Defense System simulation suite — 57 binaries, physically accurate, JSON-native.**

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

# List scenarios
./binaries/linux-x86/launch-veh-sim -scenario list
```

## Binary Count: 57

| Category | Count | Description |
|----------|-------|-------------|
| **BMDS Sensors** | 10 | Radars and space-based IR sensors |
| **BMDS Threats** | 9 | ICBM, IRBM, MRBM, HGV, SLCM, decoys, EW, nuclear effects, UAP |
| **BMDS Interceptors** | 6 | GMD/GBI, SM-3, SM-6, THAAD, THAAD-ER, Patriot |
| **C2/Infrastructure** | 8 | C2BMC, GFCB, comms links, engagement chain, kill assessment, WTA |
| **Support** | 5 | Air traffic, space weather, satellite tracker, space debris, launch vehicles |
| **Assessment** | 6 | RCS, IR signature, EMP effects, debris field, kill assessment, population impact |
| **Engine Sims** | 3 | Electronic warfare, missile defense, submarine warfare (Go libraries) |
| **Standalone** | 6 | Launch vehicle, satellite weapon, air combat, boost intercept, HGV, tactical net |
| **Config-driven** | 3 | Cyber red team, maritime, space warfare (YAML scenario files) |
| **Intelligence** | 1 | Intel collection, counterintelligence, covert action (Go library) |

## About Each Simulator

Detailed physics, fidelity assessment, and improvement recommendations for every binary:

| Binary | About | Accuracy | One-line |
|--------|-------|----------|----------|
| bmd-sim-sbirs | [about/bmd-sim-sbirs.md](about/bmd-sim-sbirs.md) | Medium | SBIRS IR early warning constellation |
| bmd-sim-stss | [about/bmd-sim-stss.md](about/bmd-sim-stss.md) | Low-Med | STSS midcourse tracking constellation |
| bmd-sim-dsp | [about/bmd-sim-dsp.md](about/bmd-sim-dsp.md) | Low-Med | DSP legacy missile warning |
| bmd-sim-uewr | [about/bmd-sim-uewr.md](about/bmd-sim-uewr.md) | High | UEWR UHF phased-array radar |
| bmd-sim-tpy2 | [about/bmd-sim-tpy2.md](about/bmd-sim-tpy2.md) | High | AN/TPY-2 X-band radar |
| bmd-sim-gbr | [about/bmd-sim-gbr.md](about/bmd-sim-gbr.md) | High | Ground-Based Radar (XBR) |
| bmd-sim-cobra-judy | [about/bmd-sim-cobra-judy.md](about/bmd-sim-cobra-judy.md) | Medium | Cobra Judy S-band collection radar |
| bmd-sim-patriot | [about/bmd-sim-patriot.md](about/bmd-sim-patriot.md) | High | PATRIOT air defense system |
| bmd-sim-lrdr | [about/bmd-sim-lrdr.md](about/bmd-sim-lrdr.md) | High | Long-Range Discrimination Radar |
| bmd-sim-aegis | [about/bmd-sim-aegis.md](about/bmd-sim-aegis.md) | High | Aegis combat system (SPY-1D + SM-3/6) |
| bmd-sim-icbm | [about/bmd-sim-icbm.md](about/bmd-sim-icbm.md) | High | ICBM threat trajectories |
| bmd-sim-irbm | [about/bmd-sim-irbm.md](about/bmd-sim-irbm.md) | Med-High | IRBM threat trajectories |
| bmd-sim-mrbm | [about/bmd-sim-mrbm.md](about/bmd-sim-mrbm.md) | Medium | MRBM threat trajectories |
| bmd-sim-slcm | [about/bmd-sim-slcm.md](about/bmd-sim-slcm.md) | Medium | Submarine-launched cruise missiles |
| bmd-sim-hgv | [about/bmd-sim-hgv.md](about/bmd-sim-hgv.md) | High | Hypersonic glide vehicles |
| bmd-sim-decoy | [about/bmd-sim-decoy.md](about/bmd-sim-decoy.md) | Medium | Decoys & penetration aids |
| bmd-sim-electronic-attack | [about/bmd-sim-electronic-attack.md](about/bmd-sim-electronic-attack.md) | Med-High | ECM vs BMDS radars |
| bmd-sim-nuclear-efx | [about/bmd-sim-nuclear-efx.md](about/bmd-sim-nuclear-efx.md) | High | Nuclear detonation effects |
| bmd-sim-ufo | [about/bmd-sim-ufo.md](about/bmd-sim-ufo.md) | N/A | Unknown Anomalous Phenomena |
| bmd-sim-gmd | [about/bmd-sim-gmd.md](about/bmd-sim-gmd.md) | High | Ground-Based Midcourse Defense |
| bmd-sim-sm3 | [about/bmd-sim-sm3.md](about/bmd-sim-sm3.md) | High | Standard Missile 3 |
| bmd-sim-sm6 | [about/bmd-sim-sm6.md](about/bmd-sim-sm6.md) | Medium | Standard Missile 6 |
| bmd-sim-thaad | [about/bmd-sim-thaad.md](about/bmd-sim-thaad.md) | High | THAAD terminal defense |
| bmd-sim-thaad-er | [about/bmd-sim-thaad-er.md](about/bmd-sim-thaad-er.md) | Medium | THAAD Extended Range |
| boost-intercept | [about/boost-intercept.md](about/boost-intercept.md) | Medium | Boost-phase intercept |
| bmd-sim-c2bmc | [about/bmd-sim-c2bmc.md](about/bmd-sim-c2bmc.md) | Medium | C2 Battle Management |
| bmd-sim-gfcb | [about/bmd-sim-gfcb.md](about/bmd-sim-gfcb.md) | Medium | Global Force Coordination |
| bmd-sim-ifxb | [about/bmd-sim-ifxb.md](about/bmd-sim-ifxb.md) | Medium | Integrated Fire Control |
| bmd-sim-jrsc | [about/bmd-sim-jrsc.md](about/bmd-sim-jrsc.md) | Low-Med | Joint Regional Security Center |
| bmd-sim-link16 | [about/bmd-sim-link16.md](about/bmd-sim-link16.md) | Med-High | Link-16 tactical data link |
| bmd-sim-jreap | [about/bmd-sim-jreap.md](about/bmd-sim-jreap.md) | Low-Med | JREAP range extension |
| bmd-sim-hub | [about/bmd-sim-hub.md](about/bmd-sim-hub.md) | Low-Med | BMDS integration hub |
| engagement-chain | [about/engagement-chain.md](about/engagement-chain.md) | Medium | Full kill chain |
| kill-assessment | [about/kill-assessment.md](about/kill-assessment.md) | Med-High | Post-intercept damage assessment |
| kill-chain-rt | [about/kill-chain-rt.md](about/kill-chain-rt.md) | High | Real-time kill chain |
| tactical-net | [about/tactical-net.md](about/tactical-net.md) | Med-High | BMDS RF mesh network |
| wta | [about/wta.md](about/wta.md) | High | Weapon-Target Assignment |
| bmd-sim-atmospheric | [about/bmd-sim-atmospheric.md](about/bmd-sim-atmospheric.md) | Med-High | Atmospheric effects on BMD |
| bmd-sim-space-weather | [about/bmd-sim-space-weather.md](about/bmd-sim-space-weather.md) | Medium | Space weather effects |
| bmd-sim-air-traffic | [about/bmd-sim-air-traffic.md](about/bmd-sim-air-traffic.md) | Low-Med | Air traffic discrimination |
| satellite-tracker | [about/satellite-tracker.md](about/satellite-tracker.md) | High | NORAD/Celestrak orbital tracker |
| bmd-sim-space-debris | [about/bmd-sim-space-debris.md](about/bmd-sim-space-debris.md) | Medium | Space debris risk |
| debris-field | [about/debris-field.md](about/debris-field.md) | Medium | Post-intercept debris |
| launch-veh-sim | [about/launch-veh-sim.md](about/launch-veh-sim.md) | High | Launch vehicle ascent physics |
| population-impact-sim | [about/population-impact-sim.md](about/population-impact-sim.md) | Medium | Nuclear strike impact |
| electronic-war-sim | [about/electronic-war-sim.md](about/electronic-war-sim.md) | Med-High | Electronic warfare |
| missile-defense-sim | [about/missile-defense-sim.md](about/missile-defense-sim.md) | Medium | Layered missile defense |
| submarine-war-sim | [about/submarine-war-sim.md](about/submarine-war-sim.md) | Med-High | Submarine warfare |
| bmd-sim-satellite-weapon | [about/bmd-sim-satellite-weapon.md](about/bmd-sim-satellite-weapon.md) | Medium | ASAT weapons |
| rcs | [about/rcs.md](about/rcs.md) | Med-High | Radar cross-section modeling |
| ir-signature | [about/ir-signature.md](about/ir-signature.md) | Med-High | IR signature modeling |
| emp-effects | [about/emp-effects.md](about/emp-effects.md) | Med-High | EMP effects on BMDS |
| hgv | [about/hgv.md](about/hgv.md) | High | Hypersonic glide vehicle |
| air-combat-sim | [about/air-combat-sim.md](about/air-combat-sim.md) | Low-Med | Air combat engagement |
| cyber-redteam-sim | [about/cyber-redteam-sim.md](about/cyber-redteam-sim.md) | Unknown | Cyber red team (no source) |
| maritime-sim | [about/maritime-sim.md](about/maritime-sim.md) | Unknown | Maritime warfare (config-driven) |
| space-war-sim | [about/space-war-sim.md](about/space-war-sim.md) | Unknown | Space warfare (config-driven) |
| intel-collection-sim | [docs/intel-collection-sim-README.md](docs/intel-collection-sim-README.md) | High | Intelligence collection & counterintelligence (Go library) |

## NORAD Integration

The **satellite-tracker** integrates with NORAD/Celestrak for live orbital data:

```bash
./satellite-tracker -group gps-ops -json     # GPS constellation
./satellite-tracker -group starlink -json     # Starlink
./satellite-tracker -group active -json       # All active satellites
```

Supported groups: `active`, `stations`, `starlink`, `gps-ops`, `weather`, `geo`, `military`, `intelsat`, `resource`, `sarsat`, `argos`, `planet`, `dmc`, `orbcomm`, `x-comm`

## Physics Accuracy

All sensor and interceptor parameters are derived from open-source technical data:

- **Radar ranges**: Energy-on-target equation with PeakPower × PulseWidth × NPulses
- **System losses**: 10-14 dB per radar type (environment-corrected)
- **Interceptor PK**: Derived from closing velocity, divert capability, seeker range, lethal radius
- **Threat trajectories**: Keplerian mechanics with proper boost/midcourse/terminal phases
- **Nuclear effects**: HEMP E1/E2/E3, thermal blast, ionospheric scintillation
- **EW modeling**: J/S ratios, burnthrough ranges, DRFM, worst-case sensor degradation
- **Launch vehicles**: Spherical gravity, atmosphere-relative drag, Max Q validated against F9 (~30.7 kPa)

See [about/](about/) for detailed physics, fidelity assessment, and improvement recommendations for each simulator.

## Common Flags

| Flag | Description |
|------|-------------|
| `-json` | Machine-readable JSON output |
| `-i` | Interactive mode (live terminal dashboard) |
| `-scenario NAME` | Run specific scenario (use `-scenario list` to see options) |
| `-duration DURATION` | Run duration (e.g., `30s`, `5m`) |
| `-seed int` | Random seed (0 = non-deterministic) |
| `-v` | Verbose output (selected binaries) |
| `-help` | Show help and flag descriptions |

## Documentation

| File | Description |
|------|-------------|
| [about/](about/) | Detailed physics, fidelity, and improvement notes for every binary |
| [SOURCE.md](SOURCE.md) | Complete binary → source repository mapping |
| [API_REFERENCE.md](API_REFERENCE.md) | JSON output schema for all binaries |
| [ACCURACY.md](ACCURACY.md) | Accuracy assessment and limitations |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Build process, testing, and PR guidelines |
| [NORAD-INTEGRATION.md](NORAD-INTEGRATION.md) | Live orbital data integration |

## Repository Structure

```
forge-sims/
├── about/                      # Detailed about pages for each simulator
│   ├── bmd-sim-sbirs.md        # SBIRS constellation
│   ├── bmd-sim-uewr.md         # UEWR radar
│   ├── launch-veh-sim.md       # Launch vehicle physics
│   └── ...                     # 56 total about pages
├── binaries/linux-x86/         # 56 compiled Linux x86_64 binaries
├── configs/                    # YAML scenario configs for config-driven sims
├── docs/                       # Additional documentation
├── scripts/                    # Build and deploy scripts
├── SOURCE.md                   # Binary → source repository map
├── API_REFERENCE.md            # Complete JSON schema reference
├── ACCURACY.md                 # Accuracy assessment and limitations
├── README.md                   # This file
└── EVALUATION-*.md             # Evaluation history
```

## Source Code

This repository contains binaries and documentation only. Source code lives on IDM GitLab:

```
git@idm.wezzel.com:crab-meat-repos/<sim-name>.git
```

See [SOURCE.md](SOURCE.md) for the complete binary → source mapping and [CONTRIBUTING.md](CONTRIBUTING.md) for the build process.