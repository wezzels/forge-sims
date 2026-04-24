# FORGE-Sims Evaluation — Round 15

**Date:** 2026-04-23 (Round 15 — Full Re-audit)
**Repo:** github.com/wezzels/forge-sims
**Total binaries:** 52 (excluding scripts)

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 52 | 50 | 2 | maritime, space-war (config-driven) |
| All standard flags | 52 | 44 | 8 | 3 config-driven, 3 engine-only, 2 partial |
| Clean JSON output | 52 | 46 | 6 | 3 config-driven, 2 engine-only, 1 needs-scenario |
| Zero stubs/mocks | 52 | 52 | 0 | ✅ Zero found |
| Zero format string bugs | 52 | 52 | 0 | ✅ All fixed |

---

## Changes Since Round 14

| Sim | Change |
|-----|--------|
| air-combat-sim | ✅ **NOW HAS `-json` stdout!** 46 values, all flags pass |
| air-combat-sim | ✅ `-seed`, `-duration`, `-v`, `-i` all working |
| kill-chain-rt | 🆕 NEW — Kill chain real-time simulator (6 scenarios) |
| launch-veh-sim | ✅ `-json` FIXED (was returning empty) |
| launch-veh-sim | ⚠️ Still missing `-i`, `-v`, `-seed` flags |
| launch-veh-sim | ⚠️ `-duration` is float seconds (not Go Duration) |
| launch-veh-sim | ⚠️ Max Q still 94 kPa (real F9 ~35 kPa) |
| tactical-net | ⚠️ `-json` still prints interactive text before JSON |

---

## Full Audit Results

| Sim | json | i | dur | seed | v | vals | Status |
|-----|------|---|-----|------|---|------|--------|
| **BMDS Core (31)** |||||||
| bmd-sim-sbirs through bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | ✅ | 2-22 | ✅ All pass |
| **Interceptors & Support (14)** |||||||
| boost-intercept | ✅ | ✅ | ✅ | ✅ | ✅ | 11 | ✅ |
| debris-field | ✅ | ✅ | ✅ | ✅ | ✅ | 11 | ✅ |
| kill-chain-rt | ✅* | ✅ | ✅ | ✅ | ✅ | 5† | ✅ Needs -scenario |
| engagement-chain | ✅ | ✅ | ✅ | ✅ | ✅ | 53 | ✅ |
| kill-assessment | ✅ | ✅ | ✅ | ✅ | ✅ | 22 | ✅ |
| wta | ✅ | ✅ | ✅ | ✅ | ✅ | 18 | ✅ |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ | ✅ | 166 | ✅ |
| space-debris | ✅ | ✅ | ✅ | ✅ | ✅ | 325K | ✅ |
| air-traffic | ✅ | ✅ | ✅ | ✅ | ✅ | 35K | ✅ |
| tactical-net | ✅‡ | ✅ | ✅ | ✅ | ✅ | 13K | ⚠️ JSON after text |
| ufo | ✅ | ✅ | ✅ | ✅ | ✅ | 61 | ✅ |
| air-combat-sim | ✅ | ✅ | ✅ | ✅ | ✅ | 46 | ✅ FIXED |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | ✅ |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | ✅ | 3 | ✅ |
| **Space** |||||||
| launch-veh-sim | ✅ | ❌ | ❌ | ❌ | ❌ | 31 | ⚠️ JSON fixed, still missing flags |
| **Excluded (6)** |||||||
| cyber-redteam-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| maritime-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| space-war-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| electronic-war-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |
| missile-defense-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |
| submarine-war-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |

*kill-chain-rt: requires `-scenario <name>` flag for JSON output
†kill-chain-rt: 5 top-level values, 19 total in parameters
‡tactical-net: JSON at end of stream, preceded by interactive text

---

## kill-chain-rt — NEW 🆕

High-fidelity kill chain simulator with 6 scenarios:

| Scenario | Interceptor | PK | Shots | Total Time |
|----------|-------------|-----|-------|------------|
| gmd-icbm-midcourse | GBI | 0.55 | 2 | 194s |
| aegis-irbm-pacific | SM-3 | 0.65 | 2 | 142s |
| thaad-terminal-korea | THAAD | 0.70 | 1 | 100s |
| patriot-terminal | PAC-3 | 0.75 | 4 | 72s |
| sm6-hgv-terminal | SM-6 | 0.10 | 2 | 97s |
| kei-boost-phase | KEI | 0.68 | 2 | 135s |

**Physics verified:**
- Kill chain timeline: detect → track → fire order → launch → intercept
- SM-6 vs HGV PK=0.10 ✅ (correct — terminal defense vs hypersonic is very hard)
- Patriot PAC-3 PK=0.75 ✅ (4 shots for layered defense)
- GMD PK=0.55 per shot, effective 0.80 with 2 shots ✅
- All flags pass (`-json`, `-i`, `-v`, `-duration`, `-seed`)

**Requires `-scenario` flag for meaningful output** — without it, just prints available scenarios.

---

## air-combat-sim — FIXED ✅

All previously identified CLI issues resolved:
- ✅ `-json` now outputs to stdout (was file-only)
- ✅ `-seed` flag added
- ✅ `-duration` flag working
- ✅ `-i` interactive mode
- ✅ `-v` verbose mode

**Remaining issues:**

| Issue | Severity | Detail |
|-------|----------|--------|
| Pk=0.00 in default run | P2 | No weapons employment in BVR scenario |
| Events list empty | P2 | No radar detections, no missile launches recorded |
| F-22 fuel=6100 kg | P3 | Real internal fuel ~8200 kg (was 5900, improved) |
| F-22 no mass field | P3 | JSON has no mass field (fuel_kg only) |
| Su-57 fuel=7800 kg | P3 | Real internal fuel ~10300 kg |

---

## launch-veh-sim — PARTIALLY FIXED

**Fixed:**
- ✅ `-json` now works (was completely broken in R14)
- ✅ Physics improved: payload 22,800 kg, perigee 574 km, orbit achieved

**Still broken/missing:**

| Issue | Severity | Detail |
|-------|----------|--------|
| No `-i` flag | P2 | Interactive mode removed |
| No `-v` flag | P2 | Verbose mode removed |
| No `-seed` flag | P2 | No reproducibility |
| `-duration` is float | P2 | Not Go time.Duration format |
| Max Q = 94 kPa | P2 | Real F9 ~35 kPa (2.7x too high) |
| Inconsistent JSON keys | P3 | Mixed case (Apogee/Perigee vs circ_alt_km) |

---

## Open Issues Summary

| # | Sim | Issue | Severity | Status |
|---|-----|-------|----------|--------|
| 1 | air-combat-sim | Pk=0.00, no weapons employment | P2 | Open |
| 2 | air-combat-sim | Events list always empty | P2 | Open |
| 3 | air-combat-sim | F-22 fuel 6100 kg (real ~8200) | P3 | Open |
| 4 | launch-veh-sim | No `-i`, `-v`, `-seed` flags | P2 | Open |
| 5 | launch-veh-sim | Max Q 94 kPa (real ~35 kPa) | P2 | Open |
| 6 | launch-veh-sim | `-duration` is float not Go Duration | P2 | Open |
| 7 | launch-veh-sim | Inconsistent JSON key casing | P3 | Open |
| 8 | tactical-net | `-json` prints text before JSON | P3 | Open |
| 9 | kill-chain-rt | Requires `-scenario` for JSON output | P3 | By design |

### Resolved ✅ (since R14)

| Issue | Round |
|-------|-------|
| air-combat-sim no `-json` stdout | R15 ✅ |
| air-combat-sim no `-seed`/`-duration` | R15 ✅ |
| launch-veh-sim `-json` empty | R15 ✅ |
| launch-veh-sim perigee 2.1 km | R14 ✅ |
| launch-veh-sim payload 15t | R14 ✅ (now 22.8t) |

---

## Summary

**38/52 sims verified high-fidelity physics.** ✅
**8/52 C2/infrastructure (minimal standalone, real when connected).**
**6/52 excluded by design (3 config-driven, 3 engine-only).**

**Zero stubs, mocks, or placeholder data.** ✅
**Zero format string bugs.** ✅

**Major improvements this round:**
- air-combat-sim: `-json` stdout FIXED, all flags now pass ✅
- launch-veh-sim: `-json` FIXED ✅
- kill-chain-rt: NEW high-fidelity kill chain simulator ✅
- Total passing sims: 44/52 (up from 42/51 last round)