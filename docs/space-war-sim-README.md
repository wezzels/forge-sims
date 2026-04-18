# Space War Sim

High-fidelity orbital warfare simulation engine with SGP4 propagation, Monte Carlo ASAT engagement, NASA breakup debris modeling, covariance conjunction analysis, and constellation degradation simulation.

## Quick Install

```bash
# Linux amd64
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/linux-x86/space-war-sim -o space-war-sim
chmod +x space-war-sim && sudo mv space-war-sim /usr/local/bin/

# Or .deb package
curl -L https://github.com/wezzels/forge-sims/raw/main/binaries/space-war-sim_0.1.0-1_amd64.deb -o space-war-sim.deb
sudo dpkg -i space-war-sim.deb
```

See [INSTALL.md](docs/space-war-sim-INSTALL.md) for macOS, Windows, and build-from-source.

## Documentation

- [INSTALL.md](docs/space-war-sim-INSTALL.md) — Installation for all platforms
- [USAGE.md](docs/space-war-sim-USAGE.md) — CLI reference, Go API examples, all packages
- [CHANGELOG.md](docs/space-war-sim-CHANGELOG.md) — Version history

## Capabilities

| Package | Description | Tests |
|---------|-------------|-------|
| `sgp4` | SGP4/SDP4 orbit propagation from TLE data | 6 |
| `propagator` | RK4 numerical integrator (J2/J3/J4 + drag) | 6 |
| `atm` | NRLMSISE-00 atmospheric density (solar flux, Kp) | 11 |
| `asat` | Monte Carlo ASAT engagement (5 weapon types) | 12 |
| `debris` | NASA Standard Breakup Model + Kessler syndrome | 10 |
| `ssa` | Covariance conjunction analysis (Akella-Alfriend) | 11 |
| `maneuver` | Orbital maneuvers + collision avoidance | 11 |
| `satellite` | Constellation degradation (Starlink/GPS/Iridium) | 10 |
| `orbital` | Keplerian mechanics, Hohmann, J2, phasing | 6 |
| `core` | Engine, event bus, scenario loading | 9 |
| **Total** | **11 packages** | **104** |

## Binaries

| Platform | Binary | Size |
|----------|--------|------|
| Linux amd64 | `binaries/linux-x86/space-war-sim` | 3.5 MB |
| Linux arm64 | `binaries/linux-arm64/space-war-sim` | 3.4 MB |
| macOS Intel | `binaries/darwin-amd64/space-war-sim` | 3.5 MB |
| macOS ARM | `binaries/darwin-arm64/space-war-sim` | 3.4 MB |
| Windows | `binaries/windows-amd64/space-war-sim.exe` | 3.6 MB |
| Debian/Ubuntu | `binaries/space-war-sim_0.1.0-1_amd64.deb` | 1.9 MB |

## Web UI

Open `web/tactical.html` in a browser for the dark-theme Leaflet.js tactical display with real-time tracking, constellation status, conjunction alerts, and debris field visualization.