# FORGE-Sims Binary Status

## Linux x86-64 Binaries

| Binary | Source Repo | Status |
|--------|-------------|--------|
| bmd-sim-aegis | bmd-sim-aegis | ✅ Pre-existing |
| bmd-sim-atmospheric | bmd-sim-atmospheric | ✅ New |
| bmd-sim-c2bmc | bmd-sim-c2bmc | ✅ New |
| bmd-sim-cobra-judy | bmd-sim-cobra-judy | ✅ New |
| bmd-sim-decoy | bmd-sim-decoy | ✅ New |
| bmd-sim-dsp | bmd-sim-dsp | ✅ New |
| bmd-sim-gbr | bmd-sim-gbr | ✅ Pre-existing |
| bmd-sim-gfcb | bmd-sim-gfcb | ✅ New |
| bmd-sim-gmd | bmd-sim-gmd | ✅ New |
| bmd-sim-hgv | bmd-sim-hgv | ✅ New |
| bmd-sim-hub | bmd-sim-hub | ✅ Pre-existing |
| bmd-sim-icbm | bmd-sim-icbm | ✅ New |
| bmd-sim-ifxb | bmd-sim-ifxb | ✅ New |
| bmd-sim-irbm | bmd-sim-irbm | ✅ New |
| bmd-sim-jamming | bmd-sim-jamming | ✅ New |
| bmd-sim-jreap | bmd-sim-jreap | ✅ New |
| bmd-sim-jrsc | bmd-sim-jrsc | ✅ New |
| bmd-sim-link16 | bmd-sim-link16 | ✅ New |
| bmd-sim-lrdr | bmd-sim-lrdr | ✅ New |
| bmd-sim-patriot | bmd-sim-patriot | ✅ Pre-existing |
| bmd-sim-sbirs | bmd-sim-sbirs | ✅ New |
| bmd-sim-slcm | bmd-sim-slcm | ✅ New |
| bmd-sim-sm3 | bmd-sim-sm3 | ✅ New |
| bmd-sim-sm6 | bmd-sim-sm6 | ✅ New |
| bmd-sim-space-weather | bmd-sim-space-weather | ✅ New |
| bmd-sim-stss | bmd-sim-stss | ✅ New |
| bmd-sim-thaad | bmd-sim-thaad | ✅ Pre-existing |
| bmd-sim-thaad-er | bmd-sim-thaad-er | ✅ New |
| bmd-sim-tpy2 | bmd-sim-tpy2 | ✅ Pre-existing |
| bmd-sim-uewr | bmd-sim-uewr | ✅ New |

**Total: 30 binaries (7 pre-existing + 23 new)**
---

## Cyber Red Team Simulator

| Binary | Source Repo | Platforms | Status |
|--------|-------------|----------|--------|
| cyber-redteam-sim | [cyber-redteam-sim](https://github.com/wezzels/cyber-redteam-sim) | linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64 | ✅ v0.2.0 |

**Additional files:**
- `cyber-redteam-sim_0.2.0-1_amd64.deb` — Debian/Ubuntu package with systemd unit
- `checksums.txt` — SHA256 checksums for all platform binaries
- `configs/cyber-redteam-sim/*.yaml` — Scenario configs (enterprise-ad, dmz-breach, cloud-kill-chain, apt-persistence)

**Install guide:** [docs/cyber-redteam-sim-INSTALL.md](docs/cyber-redteam-sim-INSTALL.md)

---

## Maritime Sim

| Binary | Source Repo | Platforms | Status |
|--------|-------------|----------|--------|
| maritime-sim | [maritime-sim](https://idm.wezzel.com/crab-meat-repos/maritime-sim) | linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64 | ✅ v0.1.0 |

**Additional files:**
- `maritime-sim_0.1.0-1_amd64.deb` — Debian/Ubuntu package with systemd unit
- `configs/maritime-sim/` — Scenario configs (scenario.yaml, engagement.yaml)

**Install guide:** [docs/maritime-sim-INSTALL.md](docs/maritime-sim-INSTALL.md)

---

## Space War Sim

| Binary | Source Repo | Platforms | Status |
|--------|-------------|----------|--------|
| space-war-sim | [space-war-sim](https://idm.wezzel.com/crab-meat-repos/space-war-sim) | linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64 | ✅ v0.1.0 |

**Additional files:**
- `space-war-sim_0.1.0-1_amd64.deb` — Debian/Ubuntu package with systemd unit
- `configs/space-war-sim/` — Scenario configs (scenario.yaml, engagement.yaml)

**Install guide:** [docs/space-war-sim-INSTALL.md](docs/space-war-sim-INSTALL.md)
boost-intercept | 2.4M | 5 ICBM profiles, 4 interceptor variants, 5 scenarios, PK model, detection timeline | PASS
