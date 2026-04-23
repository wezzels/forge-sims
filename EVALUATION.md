# FORGE-Sims Evaluation — Round 14

**Date:** 2026-04-23 (Round 14 — Full Re-audit)
**Repo:** github.com/wezzels/forge-sims
**Total binaries:** 51 (excluding scripts)

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 51 | 49 | 2 | maritime, space-war (config-driven) |
| All standard flags | 51 | 44 | 7 | 3 config-driven, 3 engine-only, 1 partial (air-combat) |
| Clean JSON output | 51 | 45 | 6 | 3 config-driven, 3 engine-only |
| Zero stubs/mocks | 51 | 51 | 0 | ✅ Zero found |
| Zero format string bugs | 51 | 51 | 0 | ✅ All fixed |

---

## Changes Since Round 13

| Sim | Change |
|-----|--------|
| debris-field | 🆕 NEW — Debris field simulation (KILL_VEHICLE fragmentation) |
| launch-veh-sim | ⚠️ REGRESSION — `-json` returns empty output; lost `-i`, `-v`, `-seed` flags |
| launch-veh-sim | ✅ IMPROVED — Payload fixed 15t→22.8t, perigee 2km→6049km, Max Q 97kPa→30kPa |
| submarine-war-sim | Now has own binary directory (was engine-only) |

---

## debris-field — NEW 🆕

High-fidelity debris field simulation from kinetic kill vehicle intercept.

| Field | Value | Physics Check |
|-------|-------|---------------|
| Scenario | KILL_VEHICLE at 400 km | ✅ Realistic ASAT altitude |
| Target | 400 kg @ 7000 m/s | ✅ LEO orbital velocity |
| KV | 64 kg @ 8000 m/s | ✅ Realistic KV mass/velocity |
| Closing velocity | 15,000 m/s | ✅ Correct for head-on |
| Trackable (>10cm) | 5 | ✅ Low count for KKV |
| Medium (1-10cm) | 428 | ✅ |
| Small (1mm-1cm) | 1,567 | ✅ Power-law distribution |
| Total fragments | 2,000 | ✅ |
| Reentry fragments | 1,603 | ✅ Most debris below stable orbit |
| Escape fragments | 397 | ✅ High closing velocity → some escape |
| Risk to ISS | p=1.55e-9 | ✅ Low probability (ISS at 408 km) |
| Risk to LEO sats | p=1.37e-9 | ✅ |
| Risk radius (10s km) | 39.4 km | ✅ Realistic debris cloud radius |
| Risk level | HIGH | ✅ |

**Verdict: High-fidelity debris field physics.** ✅ All flags pass (`-json`, `-i`, `-v`, `-duration`, `-seed`).

---

## launch-veh-sim — REGRESSION ⚠️

The text output improved significantly (payload, perigee, Max Q now realistic), but JSON output is completely broken:

| Issue | Severity | Detail |
|-------|----------|--------|
| `-json` returns empty | P1 | No JSON output at all (was working in R12) |
| No `-i` flag | P2 | Lost interactive mode (was working in R12) |
| No `-v` flag | P2 | Lost verbose mode (was working in R12) |
| No `-seed` flag | P2 | Lost reproducibility (was working in R12) |
| Duration is float | P2 | Changed from Go `time.Duration` to float seconds |

**Text output improvements (these are good):**
- Payload: 15,000 kg → 22,800 kg ✅ (matches real F9 Block 5)
- Perigee: 2.1 km → 6,049 km ✅ (was suborbital, now realistic)
- Max Q: 97 kPa → 30 kPa ✅ (was 2.8x too high, now close to real ~35 kPa)
- Max Q time: 33s → 61s ✅ (was too early, now realistic)
- Stage separation: 94.9 km → 89 km, 2786 m/s → 2281 m/s (improved)
- Orbit achieved status with real orbital elements

---

## Full Audit Results

| Sim | json | i | dur | seed | v | vals | Status |
|-----|------|---|-----|------|---|------|--------|
| **BMDS Core (31)** |||||||
| bmd-sim-sbirs through bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | ✅ | 2-22 | ✅ All pass |
| **Interceptors & Support (12)** |||||||
| boost-intercept | ✅ | ✅ | ✅ | ✅ | ✅ | 11 | ✅ NEW |
| debris-field | ✅ | ✅ | ✅ | ✅ | ✅ | 11 | ✅ NEW |
| engagement-chain | ✅ | ✅ | ✅ | ✅ | ✅ | 53 | ✅ |
| kill-assessment | ✅ | ✅ | ✅ | ✅ | ✅ | 22 | ✅ |
| wta | ✅ | ✅ | ✅ | ✅ | ✅ | 18 | ✅ |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ | ✅ | 166 | ✅ |
| space-debris | ✅ | ✅ | ✅ | ✅ | ✅ | 325K | ✅ |
| air-traffic | ✅ | ✅ | ✅ | ✅ | ✅ | 35K | ✅ |
| tactical-net | ✅* | ✅ | ✅ | ✅ | ✅ | 13K | ⚠️ JSON after text |
| ufo | ✅ | ✅ | ✅ | ✅ | ✅ | 61 | ✅ |
| air-combat-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 51† | ⚠️ File export |
| **Space** |||||||
| launch-veh-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | ⚠️ REGRESSION |
| **Excluded (6)** |||||||
| cyber-redteam-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| maritime-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| space-war-sim | ❌ | ❌ | ❌ | ❌ | ❌ | 0 | Config-driven |
| electronic-war-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |
| missile-defense-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |
| submarine-war-sim | ❌ | ✅ | ✅ | ✅ | ✅ | 0 | Engine-only |

*tactical-net: JSON at end of stream, preceded by interactive text
†air-combat-sim: 51 values in air-combat-result.json file, not stdout

---

## Open Issues

| # | Sim | Issue | Severity | Status |
|---|-----|-------|----------|--------|
| 1 | launch-veh-sim | `-json` returns empty | P1 | REGRESSION (was working R12) |
| 2 | launch-veh-sim | Lost `-i`, `-v`, `-seed` flags | P2 | REGRESSION |
| 3 | launch-veh-sim | Duration changed to float seconds | P2 | Changed |
| 4 | air-combat-sim | No `-json` stdout (exports to file) | P1 | Open |
| 5 | air-combat-sim | No `-seed` flag | P2 | Open |
| 6 | air-combat-sim | Pk=0.00 in default run | P2 | Open |
| 7 | air-combat-sim | No `-duration` flag | P2 | Open |
| 8 | air-combat-sim | F-22 mass slightly off (24t vs 27.2t) | P3 | Open |
| 9 | tactical-net | `-json` prints interactive text before JSON | P3 | Open |

### Resolved ✅
| Issue | Round |
|-------|-------|
| ufo text header before JSON | R13 |
| bmd-sim-gfcb/decoy/jamming format bugs | R11 |
| launch-veh-sim NaN output | R12 |
| launch-veh-sim payload 15t→22.8t | R14 ✅ |
| launch-veh-sim perigee 2km→6049km | R14 ✅ |
| launch-veh-sim Max Q 97kPa→30kPa | R14 ✅ |

---

## Summary

**37/51 sims verified high-fidelity physics.** ✅
**8/51 C2/infrastructure (minimal standalone, real when connected).**
**6/51 excluded by design.**

**Zero stubs, mocks, or placeholder data.** ✅
**Zero format string bugs.** ✅

**Key findings this round:**
- **debris-field** 🆕 — High-fidelity debris field simulation ✅
- **launch-veh-sim** REGRESSION — `-json` completely broken, lost `-i`/`-v`/`-seed` flags
- **launch-veh-sim** IMPROVEMENT — Physics values now realistic (payload, perigee, Max Q all fixed)
- Text output improvements in launch-veh-sim are great, but JSON output must be restored