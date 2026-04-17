# Changelog

## v0.1.0 (2026-04-17)

### Initial Release

**Core Packages (11):**
- `internal/core` — Engine, event bus, config loading, orbital state management
- `internal/orbital` — Keplerian mechanics, J2 perturbation, Hohmann transfers, phasing maneuvers, conjunction analysis, atmospheric drag
- `internal/ssa` — Space situational awareness, tracking, sensor models (ground radar, optical, space-based)
- `internal/asat` — ASAT weapons (DA-ASAT, co-orbital, laser, EMP), engagement modeling, debris generation, Kessler syndrome
- `internal/debris` — Orbital debris field management, fragmentation events, collision probability
- `internal/ew` — Electronic warfare (uplink/downlink jamming, GPS spoofing, JNR calculation)
- `internal/c2` — C2/ROE engine, space tasking orders
- `internal/groundstation` — Ground station tracking, pass windows, link budgets
- `internal/satellite` — LEO/MEO/GEO platforms, constellation management
- `internal/scenario` — YAML scenario loading and validation
- `internal/aar` — After-action review, recording, JSON export

**CLI:**
- `--version` — Build metadata injection
- `--doctor` — Health check diagnostics
- `--init` — Initialize config directory
- `--validate` — Validate scenario configs
- `--list-scenarios` — List available scenarios

**Distribution:**
- Cross-compilation for 5 platforms (Linux amd64/arm64, macOS Intel/ARM, Windows amd64)
- `.deb` package with systemd unit
- Platform installers (Linux, macOS, Windows)
- Docker image (Alpine 3.19)
- GoReleaser, GitHub Actions CI/CD, GitLab CI
- Shell completions (bash, zsh, PowerShell)

**Tests:**
- Orbital mechanics (period, Hohmann, J2, velocity, drag, conjunction)
- All 11 packages compile clean