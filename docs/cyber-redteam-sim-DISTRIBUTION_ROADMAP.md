# Cyber Red Team Simulator — Distribution Roadmap

## Overview

Package the simulator as installable artifacts for Windows, macOS, and Linux.
Single Go binary → polished installers with configs, docs, and integration.

---

## Phase 1: Build System & Cross-Compilation (1-2 days)

**Goal:** Automated build pipeline producing all platform binaries.

- [ ] Create `Makefile` with cross-compile targets
- [ ] Add `-ldflags -X main.version=... -X main.commit=... -X main.date=...` for version injection
- [ ] Add `version` command to CLI that prints version/commit/date
- [ ] Build matrix:
  - `cyber-redteam-sim-windows-amd64.exe`
  - `cyber-redteam-sim-darwin-amd64` (Intel Mac)
  - `cyber-redteam-sim-darwin-arm64` (Apple Silicon)
  - `cyber-redteam-sim-linux-amd64`
  - `cyber-redteam-sim-linux-arm64` (Raspberry Pi / ARM servers)
- [ ] Add `make dist` target that builds all + checksums (`sha256sum`)
- [ ] Test each binary starts and runs `sim version`
- [ ] Add `.goreleaser.yml` for CI-based releases (optional, Phase 4)

---

## Phase 2: Package Contents & Platform Integration (2-3 days)

**Goal:** Each package includes configs, docs, and platform-appropriate integration.

### Common to all platforms:
- [ ] `configs/` — sample scenarios (enterprise-ad.yaml, + 4 more)
- [ ] `docs/` — README, quickstart guide, scenario format reference
- [ ] `LICENSE`
- [ ] `CHANGELOG.md`

### Windows-specific:
- [ ] `sim.exe` — renamed binary
- [ ] `configs/` — scenarios with Windows-compatible paths (C:\ProgramData\cyber-redteam-sim\)
- [ ] `install.ps1` — PowerShell installer:
  - Copy to `C:\Program Files\CyberRedTeamSim\`
  - Add to PATH
  - Register uninstall info
- [ ] `uninstall.ps1` — clean removal
- [ ] Optional: Windows Service wrapper (`--service install` flag using `golang.org/x/sys/windows/svc/mgr`)

### macOS-specific:
- [ ] Universal binary or separate Intel/ARM
- [ ] `com.cyberredteamsim.plist` — launchd config for server mode
- [ ] `cyber-redteam-sim` — binary in `/usr/local/bin/` or app bundle
- [ ] Homebrew formula (`Formula/cyber-redteam-sim.rb`):
  - `homebrew tap crab-meat-repos/cyber`
  - Pulls from GitHub releases
  - `brew install cyber-redteam-sim`
- [ ] `.pkg` installer via `pkgbuild`/`productbuild`:
  - Welcome screen, license, install location
  - Post-install script to set permissions
- [ ] Code signing (requires Apple Developer cert — Phase 4)

### Linux-specific:
- [ ] `.deb` package (Debian/Ubuntu):
  - `debian/control`, `debian/rules`, `debian/changelog`
  - systemd unit: `cyber-redteam-sim.service`
  - Config in `/etc/cyber-redteam-sim/`
  - Data in `/var/lib/cyber-redteam-sim/`
  - Log to `/var/log/cyber-redteam-sim/`
  - `dpkg -i cyber-redteam-sim_0.2.0_amd64.deb`
- [ ] `.rpm` package (RHEL/Fedora):
  - `.spec` file with same structure as .deb
  - `rpm -i cyber-redteam-sim-0.2.0-1.x86_64.rpm`
- [ ] `.tar.gz` standalone (any Linux):
  - Binary + configs + README
  - Install script: `sudo ./install.sh`
  - Good for ARM, Alpine, NixOS, etc.

---

## Phase 3: Installers & User Experience (2-3 days)

**Goal:** Click-to-install on each platform with smooth UX.

### Windows:
- [ ] WiX Toolset `.msi` installer:
  - GUI wizard (Welcome → License → Install → Finish)
  - Start Menu shortcuts
  - Add to PATH automatically
  - Optional: "Run a sample scenario" checkbox on final page
- [ ] Alternative: Inno Setup `.exe` installer (simpler, no WiX learning curve)
- [ ] Right-click "Open Scenario" shell integration (optional)

### macOS:
- [ ] `.dmg` disk image:
  - Drag-to-Applications style (for GUI mode)
  - Or `.pkg` with standard Apple installer wizard
- [ ] Post-install: `sim doctor` — checks for common issues (port availability, etc.)

### Linux:
- [ ] `.deb` integrates with `apt`:
  - `sudo apt install ./cyber-redteam-sim_0.2.0_amd64.deb`
  - `cyber-redteam-sim --version` works immediately
- [ ] `.rpm` integrates with `dnf`/`yum`:
  - `sudo dnf install ./cyber-redteam-sim-0.2.0-1.x86_64.rpm`
- [ ] APT repository (optional, Phase 4):
  - `deb [signed-by=/usr/share/keyrings/cyber.gpg] https://apt.stsgym.com/ stable main`
  - `sudo apt update && sudo apt install cyber-redteam-sim`

### All platforms:
- [ ] `sim doctor` command — health check:
  - Binary runs
  - Config directory writable
  - Port available (if --server)
  - Sample scenario loads
- [ ] `sim init` command — first-run setup:
  - Create config directory
  - Copy sample scenarios
  - Set default config
- [ ] Man page (`man cyber-redteam-sim` on Linux/macOS)
- [ ] Shell completions (bash, zsh, fish, PowerShell)

---

## Phase 4: CI/CD & Release Automation (2-3 days)

**Goal:** Tag a release → all packages built and published automatically.

- [ ] GitHub Actions / GitLab CI pipeline:
  - On tag `v*`: build all 5 binaries
  - Run tests on each platform
  - Generate `.deb`, `.rpm`, `.msi`, `.pkg`, `.tar.gz`
  - Compute SHA256 checksums
  - Upload to GitHub/GitLab Releases
- [ ] GoReleaser integration:
  - `.goreleaser.yml` handles cross-compile, checksums, Docker images, Homebrew tap
  - Single `goreleaser release` command
- [ ] Homebrew tap auto-update:
  - Push formula to `homebrew-tap` repo on release
- [ ] APT repo hosting:
  - `reprepro` on miner for `apt.stsgym.com`
  - Signed with GPG key
- [ ] Code signing:
  - Windows: Authenticode sig (`signtool`)
  - macOS: Apple Developer ID (`codesign` + `notarytool`)
  - Linux: GPG sign `.deb`/`.rpm`
- [ ] Docker image:
  - `docker pull stsgym/cyber-redteam-sim:latest`
  - Multi-arch (amd64 + arm64)
  - Entrypoint: `sim --config /etc/cyber-redteam-sim/scenario.yaml --server`

---

## Phase 5: Polish & Extras (1-2 days)

- [ ] `sim config` interactive wizard — generate scenario YAML
- [ ] `sim scenario list` — list built-in scenarios
- [ ] `sim scenario validate <file>` — validate a scenario config
- [ ] Auto-update mechanism (`sim update check`)
- [ ] Telemetry opt-in (anonymous usage stats)
- [ ] Sample output gallery in docs

---

## Timeline Summary

| Phase | Duration | Deliverable |
|-------|----------|-------------|
| 1. Build System | 1-2 days | `make dist` produces 5 binaries + checksums |
| 2. Package Contents | 2-3 days | Platform-specific files, services, configs |
| 3. Installers | 2-3 days | .msi, .pkg/.dmg, .deb/.rpm that install cleanly |
| 4. CI/CD | 2-3 days | Tag → automated release with all artifacts |
| 5. Polish | 1-2 days | Doctor, init, completions, Docker |

**Total: ~8-13 days** (working incrementally, Phase 1 is usable standalone)

---

## Quick Start (Phase 1 Only — 1 day)

```bash
# Build everything
make dist

# Output:
# dist/cyber-redteam-sim-v0.2.0-windows-amd64.exe
# dist/cyber-redteam-sim-v0.2.0-darwin-amd64
# dist/cyber-redteam-sim-v0.2.0-darwin-arm64
# dist/cyber-redteam-sim-v0.2.0-linux-amd64
# dist/cyber-redteam-sim-v0.2.0-linux-arm64
# dist/checksums.txt
```