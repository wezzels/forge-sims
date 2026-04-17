# Changelog

## v0.2.0 (2026-04-17)

### New Features
- Probabilistic cloud attack modeling (`cloud` package)
- gRPC/JSON-RPC engine adapter (`grpcapi/adapter.go`)
- Build system with cross-compilation for 5 platforms (Windows, macOS Intel/ARM, Linux amd64/arm64)
- `.deb` package with systemd unit, postinst/prerm lifecycle scripts
- Version command (`--version`) with build metadata injection
- Distribution Makefile: `make dist`, `make deb`, `make docker`, `make cross`
- `--doctor` flag: health check diagnostics (version, config, scenario validation, port)
- `--init` flag: initialize config directory with sample scenarios
- `--validate` flag: validate scenario config files (hosts, zones, services, vulns)
- `--list-scenarios` flag: list available scenario files with metadata
- `--config-wizard` flag: interactive scenario generator with 4 templates
  - enterprise-ad: 5 hosts, domain admin objective
  - dmz-breach: 3 zones, web app to internal
  - cloud-kill-chain: AWS/Azure multi-cloud
  - apt-persistence: long-duration APT, persistence/C2 focus
- `--update-check` flag: check for newer versions
- Shell completions (bash, zsh, PowerShell)
- Man page (`cyber-redteam-sim.1`)
- Platform installers: Linux (install.sh/uninstall.sh), macOS (install.sh + launchd), Windows (install.bat/uninstall.bat)
- GoReleaser config for automated releases
- GitHub Actions CI/CD (test + cross-compile on push, full release on tag)
- GitLab CI for IDM mirror (.deb pipeline)
- Docker image (Alpine 3.19, health check, port 8080)
- Homebrew formula (`brew install wezzels/tap/cyber-redteam-sim`)
- Snap package config
- Video demo scripts (mock + composite live-vs-mock comparison)
- Distribution roadmap (docs/DISTRIBUTION_ROADMAP.md)

### Improvements
- Cloud `CloudExfil` and `CloudC2` now use `*CloudEnvironment` instead of `(Provider, string)`
- All 65+ tests pass across 20 packages
- Config nested directory fix in packages (`configs/configs/` → `configs/`)

### Bug Fixes
- Fixed `benchmark_test` and `integration_test` callers for new cloud API signature

## v0.1.0 (2026-04-09)

### Initial Release
- Core engine with WGS84 kinematics, event bus, config loading
- 45+ MITRE ATT&CK techniques
- Defense modeling (maturity levels, IR playbooks, EDR, SIEM correlation)
- Network simulation (hosts, services, zones, connections)
- Attack graph (A* pathfinding, tradeoff analysis)
- AI commander (Q-Learning, threat assessment)
- C2 infrastructure (redirectors, CDN fronts, dead drops, implants)
- IEEE 1278.1 DIS federation bridge
- Web UI with Leaflet.js and SSE real-time push
- After-action review with scoring
- Replay with time-travel
- Scenario editor with validation