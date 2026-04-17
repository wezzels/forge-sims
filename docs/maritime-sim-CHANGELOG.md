# Changelog

## v0.1.0 (2026-04-17)

### Initial Distribution Release

**New Features:**
- Build system with cross-compilation for 5 platforms (Linux amd64/arm64, macOS Intel/ARM, Windows amd64)
- `.deb` package with systemd unit, postinst/prerm lifecycle scripts
- Version command (`--version`) with build metadata injection
- `--doctor` flag: health check diagnostics
- `--init` flag: initialize config directory with sample scenarios
- `--validate` flag: validate scenario config files
- `--list-scenarios` flag: list available scenario files
- Distribution Makefile: `make dist`, `make deb`, `make docker`
- Platform installers: Linux install.sh/uninstall.sh, macOS install.sh + launchd, Windows install.bat/uninstall.bat
- Man page (`maritime-sim.1`)
- Docker image (Alpine 3.19, health check, port 8080)

**Core Simulation (pre-existing):**
- 21 packages, 90+ tests, all green
- Core engine with WGS84 kinematics, config loading, event bus
- Surface combatant (DDG/FFG/CVN) and submarine (SSN/SSK) models
- Acoustic sensors (Mackenzie sound speed, sonar equation, transmission loss)
- Radar (equation, horizon, detection)
- Sensor manager (active sonar scheduling, towed array)
- Track manager (alpha-beta filtering, classification, bearing-only)
- AI commander (threat assessment, decision tree, ROE)
- C2/ROE engine (weapons free/tight/hold)
- MSEL timeline engine
- AAR recorder (state snapshots, replay, scoring, JSON export)
- Weapon models (AShM, torpedo)
- Damage model
- EW (ECM chaff/jamming, ESM bearings)
- DIS (IEEE 1278) federation bridge
- Web tactical display (Leaflet.js, SSE)
- Hydrodynamics (Pierson-Moskowitz spectra)