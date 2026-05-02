# FORGE-Sims TODO

**Last Updated:** 2026-05-02 (Round 26)

---

## ✅ Completed

- [x] bmd-sim-sbirs — SBIRS satellite simulator
- [x] bmd-sim-uewr — Upgraded Early Warning Radar
- [x] bmd-sim-gmd — Ground-Based Midcourse Defense
- [x] bmd-sim-sm3 — Standard Missile-3
- [x] bmd-sim-icbm — ICBM threat simulator
- [x] bmd-sim-hgv — Hypersonic Glide Vehicle
- [x] bmd-sim-c2bmc — C2BMC emulation
- [x] bmd-sim-lrdr — Long Range Discrimination Radar
- [x] bmd-sim-stss — Space Tracking & Surveillance System
- [x] bmd-sim-irbm — IRBM threat simulator
- [x] bmd-sim-sm6 — Standard Missile-6
- [x] bmd-sim-jamming — Electronic warfare / jamming
- [x] bmd-sim-link16 — Link 16 terminal emulator
- [x] bmd-sim-jreap — JREAP protocol simulator
- [x] bmd-sim-dsp — Defense Support Program
- [x] bmd-sim-cobra-judy — Sea-based X-band radar
- [x] bmd-sim-decoy — Countermeasure/decoy simulator
- [x] bmd-sim-slcm — Sea-Launched Cruise Missile
- [x] bmd-sim-gfcb — Ground-based Fire Control Battalion
- [x] bmd-sim-thaad-er — THAAD Extended Range
- [x] bmd-sim-ifxb — Integrated Fire X-band radar
- [x] bmd-sim-space-weather — Solar/weather effects
- [x] bmd-sim-atmospheric — Atmospheric propagation
- [x] All JSON output standardized (sim-cli library)
- [x] 65 binaries cross-compiled for 3 platforms
- [x] 58/59 produce valid JSON output
- [x] 20/59 have `data` sections with computed physics
- [x] 44/59 verified high-fidelity physics
- [x] 7 warfare sims (CBRN, cyber-kinetic, info-ops, land combat, logistics, nuclear effects, space data network)
- [x] Space Force Command Center (Cesium 3D globe)
- [x] NORAD Command Center (7 BMDS scenarios, 30+ sim endpoints)
- [x] **Round 26 Evaluation Complete** — 81 binaries audited
- [x] **space-data-network-sim documented** — about/space-data-network-sim.md created
- [x] **EVALUATION.md updated** — R26 status documented
- [x] **REMEDIATION-R26.md created** — Fix plan for 35 binaries

---

## Remaining Work (P1/P2 — Round 26 Remediation)

### space-data-network-sim (P1)
- [ ] Add `-json` flag for stdout JSON output
- [ ] Add `-i` interactive mode (live dashboard)
- [ ] Add `-v` verbose flag
- [ ] Maintain backward compatibility with `-output` flag
- [ ] Rebuild all platform binaries

### Flag Standardization (P2 — 20 binaries)
- [ ] bmd-sim-c2bmc: Add `-duration` alias for `-d`
- [ ] bmd-sim-link16: Add `-duration` alias for `-d`
- [ ] bmd-sim-jreap: Add `-duration` alias for `-d`
- [ ] bmd-sim-jrsc: Add `-duration` alias for `-d`
- [ ] bmd-sim-gfcb: Add `-duration` alias for `-d`
- [ ] bmd-sim-decoy: Add `-duration` alias for `-d`
- [ ] bmd-sim-ifxb: Add `-duration` alias for `-d`
- [ ] bmd-sim-ir-signature: Add `-duration` alias for `-d`
- [ ] bmd-sim-kill-assessment: Add `-duration` alias for `-d`
- [ ] bmd-sim-rcs: Add `-duration` alias for `-d`
- [ ] bmd-sim-satellite-weapon: Add `-duration` alias for `-d`
- [ ] bmd-sim-wta: Add `-duration` alias for `-d`
- [ ] boost-intercept: Add `-duration` alias for `-d`
- [ ] debris-field: Add `-duration` alias for `-d`
- [ ] electronic-war-sim: Add `-duration` alias for `-d`
- [ ] emp-effects: Add `-duration` alias for `-d`
- [ ] engagement-chain: Add `-duration` alias for `-d`
- [ ] hgv: Add `-duration` alias for `-d`
- [ ] kill-chain-rt: Add `-duration` alias for `-d`
- [ ] tactical-net: Add `-duration` alias for `-d`

### Duplicate Consolidation (P2 — 15 pairs)
- [ ] Remove `air-traffic` (keep `bmd-sim-air-traffic`)
- [ ] Remove `space-debris` (keep `bmd-sim-space-debris`)
- [ ] Remove `satellite-tracker` (keep `bmd-sim-satellite-tracker`)
- [ ] Remove `boost-intercept` (keep `bmd-sim-boost-intercept`)
- [ ] Remove `debris-field` (keep `bmd-sim-debris-field`)
- [ ] Remove `electronic-war-sim` (keep `bmd-sim-electronic-attack`)
- [ ] Remove `emp-effects` (keep `bmd-sim-emp-effects`)
- [ ] Remove `engagement-chain` (keep `bmd-sim-engagement-chain`)
- [ ] Remove `hgv` (keep `bmd-sim-hgv`)
- [ ] Remove `ir-signature` (keep `bmd-sim-ir-signature`)
- [ ] Remove `kill-assessment` (keep `bmd-sim-kill-assessment`)
- [ ] Remove `kill-chain-rt` (keep `bmd-sim-kill-chain-rt`)
- [ ] Remove `rcs` (keep `bmd-sim-rcs`)
- [ ] Remove `tactical-net` (keep `bmd-sim-tactical-net`)
- [ ] Remove `wta` (keep `bmd-sim-wta`)
- [ ] Update Makefile to skip non-prefixed builds
- [ ] Update API_REFERENCE.md
- [ ] Update README.md binary count

### Text Cleanup (P3 — 3 binaries)
- [ ] air-traffic: Suppress text before JSON output
- [ ] bmd-sim-air-traffic: Suppress text before JSON output
- [ ] satellite-tracker: Suppress text before JSON output

---

## Remaining Work (P3 — Legacy)

### launch-veh-sim
- [ ] Vulcan Centaur: Low-TWR S2 crashes (needs PEG guidance)
- [ ] Delta IV: Low-TWR S2 crashes (needs PEG guidance)
- [ ] Falcon Heavy: Orbit at 793 km (excess S2 delta-V for 26t payload)
- [ ] Atlas V: Orbit at 754 km (excess S2 delta-V)
- [ ] Ariane 6: Orbit at 1569 km (excess S2 delta-V)

### Documentation
- [ ] API_REFERENCE.md — needs update for 7 new warfare sims
- [ ] DOCS.md — needs update for new sims
- [ ] ACCURACY.md — needs update for air-combat-sim F-22 fuel fix

### Infrastructure
- [ ] Docker Compose for running all simulators together
- [ ] Grafana dashboards for each simulator
- [ ] ICD message standardization across all simulators

---

## Completed Phases

### Phase 1: Critical BMDS Coverage ✅
### Phase 2: Enhanced Capabilities ✅
### Phase 3: Full Spectrum ✅
### Phase 4: Advanced ✅
### Phase 5: Cross-Platform & Standards ✅
### Phase 6: Warfare Sims ✅
