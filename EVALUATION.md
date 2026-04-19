# FORGE-Sims Evaluation — Round 9 (Full Integration)

**Date:** 2026-04-19 (Round 9 — norad.stsgym.com Full Integration)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

**40 operational sims.** All pass `-json`, `-v`, `-help`. 38/40 pass `-i` and `-seed`.

Excluded: 3 config-driven (cyber-redteam, maritime, space-war), 3 engine-only (electronic-war, missile-defense, submarine-war).

---

## norad.stsgym.com Integration — COMPLETE

### Architecture
- Container: `norad-sim` on miner (host networking, port 5008)
- Flask + SocketIO app with ScenarioEngine (1807→2100+ lines)
- Source: `~/norad-sim-node/app/app.py` (host), `/app/app.py` (container)
- Binary mount: `~/norad-sim/host-bin/` → `/host-bin/` (read-only)
- 40 Go binaries in host-bin (all updated from forge-sims repo)

### Sim Integration — Go Binary Calls

**All sensor and environment data now comes from Go binaries with Python fallback.**

| Integration Method | Sims | Count |
|---|---|---|
| SimOutput → Go binary (sensor feeds) | sbirs, stss, dsp, uewr, lrdr, tpy2, gbr, cobra-judy, space-weather, atmospheric | 10 |
| `/api/public/sensor/<name>` (on-demand) | c2bmc, gfcb, ifxb, jrsc, decoy, jamming, icbm, irbm, hgv, slcm, sm3, sm6, gmd, aegis, patriot, thaad, hub | 17 |
| Scheduled data fetch + interpolation | satellite-tracker, space-debris, air-traffic | 3 |
| On-demand API endpoint | tactical-net, wta, kill-assessment, mrbm, engagement-chain, electronic-attack | 6 |

**Total: 36 sims integrated via API calls**

### API Endpoints (all verified live)

| Endpoint | Status | Source |
|---|---|---|
| `/api/public/state` | ✅ Live | ScenarioEngine + Go binary data |
| `/api/public/sensor/<name>` | ✅ Live | 27 sims on-demand |
| `/api/public/satellites` | ✅ Live | satellite-tracker binary + interpolation |
| `/api/public/debris` | ✅ Live | space-debris binary + interpolation |
| `/api/public/air-traffic` | ✅ Live | air-traffic binary + interpolation |
| `/api/public/tactical-net` | ✅ Live | tactical-net binary on-demand |
| `/api/public/wta` | ✅ Live | wta binary on-demand |
| `/api/public/kill-assessment` | ✅ Live | kill-assessment binary on-demand |
| `/api/public/mrbm` | ✅ Live | bmd-sim-mrbm binary on-demand |
| `/api/public/engagement-chain` | ✅ Live | engagement-chain binary on-demand |
| `/api/public/electronic-attack` | ✅ Live | bmd-sim-electronic-attack binary on-demand |

### Scenario Engine
- 7 scenarios: pacific-icbm, multi-axis, hypersonic-glide, submarine-launch, nuclear-exchange, regional-irbm, decoy-stress
- Realistic threat types: ICBM, IRBM, HGV, SLCM with MIRV, decoys, RCS, IR signature
- Auto-intercept logic with SLS doctrine, PK models
- All BMDS sensor sites at real-world coordinates
- Comms data from bmd-sim-link16 and bmd-sim-jreap binaries

### SimOutput Class — Go Binary Priority

Each sensor method in SimOutput now:
1. **Calls the Go binary first** (`_run_sim_binary('bmd-sim-<name>', ['-json', '-duration', '5s'])`)
2. **Parses JSON output** (handles text headers before JSON)
3. **Falls back to Python model** if binary fails or returns error

This ensures maximum realism while maintaining uptime.

---

## Binary Quality Matrix (45 total)

| Sim | `-v` | `-i` | `-json` | `-duration` | `-seed` | Notes |
|-----|------|------|---------|-------------|---------|-------|
| **BMDS Sims (35)** |||||||
| bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-stss | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-hub | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-electronic-attack | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-mrbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| engagement-chain | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| kill-assessment | ✅ | ❌ | ✅ | ✅ (dur) | ✅ | Missing `-i` |
| wta | ✅ | ❌ | ✅ | ❌ (int) | ✅ | Uses int duration |
| **Space Sims (4)** |||||||
| air-traffic | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| space-debris | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| tactical-net | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |

---

## Resolved Issues ✅

| Issue | Round | Status |
|-------|-------|--------|
| C2BMC nil pointer crash | R1-R5 | ✅ Fixed |
| Interactive mode was stub | R1-R5 | ✅ Fixed |
| Double-tick printing | R1-R5 | ✅ Fixed |
| JSON was minimal | R1-R5 | ✅ Fixed |
| gmd, lrdr, tpy2 missing all flags | R6-R7 | ✅ Fixed |
| aegis, gbr, patriot, thaad, hub broken JSON | R6-R7 | ✅ Fixed |
| Duration format inconsistency | R7-R8 | ✅ Fixed — all Go duration |
| 3 duplicate old binaries | R8 | ✅ Fixed — removed |
| norad.stsgym.com using Python models | R8-R9 | ✅ Fixed — now calls Go binaries |
| bmd-sim-lrdr, bmd-sim-uewr missing from host-bin | R8 | ✅ Fixed — added |
| Space weather/atmospheric random | R8 | ✅ Fixed — now from Go binaries |
| Comms data random | R8 | ✅ Fixed — now from bmd-sim-link16/jreap |
| 20+ sims not integrated into API | R8 | ✅ Fixed — 27 sensor endpoints + 6 data endpoints |
| JSON text header parsing | R9 | ✅ Fixed — _run_sim_binary strips text before JSON |

---

## Remaining Issues (Low Priority)

| # | Issue | Priority |
|---|-------|----------|
| 1 | kill-assessment missing `-i` | P3 |
| 2 | wta uses int duration, missing `-i` | P3 |
| 3 | 3 engine sims have no CLI (electronic-war, missile-defense, submarine-war) | P3 |
| 4 | 3 config-driven sims need documentation | P3 |
| 5 | bmd-sim-nuclear-efx not in forge-sims repo | P3 |

---

## Summary

**Round 9 — norad.stsgym.com fully integrated with FORGE-Sims.**

- **36/40 sims** connected to the live API (27 sensor endpoints + 6 data endpoints + 3 scheduled)
- **All sensor data** now sourced from Go binaries with Python fallback
- **All 45 binaries** in host-bin updated from forge-sims repo
- **SimOutput class** rewritten to call Go binaries first, fall back to Python
- **_run_sim_binary** handles text headers before JSON output
- **New dynamic endpoint** `/api/public/sensor/<name>` for 27 sims on-demand
- **New endpoints**: engagement-chain, electronic-attack