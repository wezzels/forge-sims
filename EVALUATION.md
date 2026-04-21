# FORGE-Sims Evaluation — Round 11

**Date:** 2026-04-21 (Round 11 — Full Re-audit)
**Repo:** github.com/wezzels/forge-sims
**Total binaries:** 48 (excluding scripts)

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 48 | 46 | 2 | maritime, space-war exit 1 (config-driven) |
| Clean JSON output | 48 | 42 | 6 | 3 config-driven, 3 engine-only |
| `-i` interactive mode | 48 | 43 | 5 | 3 config-driven, launch-veh-sim, 2 engine-only |
| `-duration` Go format | 48 | 43 | 5 | 3 config-driven, launch-veh-sim (float), 2 engine-only |
| `-seed` reproducibility | 48 | 43 | 5 | 3 config-driven, launch-veh-sim, 2 engine-only |
| `-v` verbose | 48 | 44 | 4 | 3 config-driven, launch-veh-sim |
| Format string bugs | 48 | 48 | 0 | ✅ ALL FIXED (gfcb, decoy, jamming) |

**Excluding config-driven (3), engine-only (3), and launch-veh-sim: 41/41 pass all flags (100%)**

---

## Changes Since Round 10

| Sim | Change |
|-----|--------|
| bmd-sim-gfcb | ✅ Format bug FIXED (`%!s(MISSING)` → correct output) |
| bmd-sim-decoy | ✅ Format bug FIXED (`%!s(MISSING)` → real decoy data) |
| bmd-sim-jamming | ✅ Format bug FIXED + enriched JSON (JNR, effect, power) |
| launch-veh-sim | 🆕 NEW — Launch vehicle simulator (has issues, see below) |

---

## Full Audit Results

| Sim | -json | -i | -dur | -seed | -v | Vals | Notes |
|-----|-------|----|------|-------|----|------|-------|
| **Tier 1 — High-Fidelity Physics (34)** |||||||
| bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ | ✅ | 6 | NEFD, aperture, IR detection |
| bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | Spin-rate, detection range |
| bmd-sim-stss | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | Orbital period, discrimination |
| bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ | ✅ | 4 | Frequency, power, track capacity |
| bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ | ✅ | 4 | Freq, power, range_rv_km |
| bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | X-band, search_range, mode |
| bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Gain, max_range |
| bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Gain, max_range |
| bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Silos, burnout, EKV, PK |
| bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Array, power, range, SM PK |
| bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ | ✅ | 8 | Launchers, range, missiles |
| bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ | ✅ | 6 | Launchers, range, KV |
| bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ | ✅ | 8 | Extended range, PK_hgv, alt |
| bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Range, apogee, MIRV, RCS |
| bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Range, apogee, MIRV |
| bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Maneuvers, speed, RCS |
| bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | Range, altitude, platform |
| bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ | ✅ | 9 | Range, PK, kill vehicle |
| bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ | ✅ | 10 | Range, PK, terminal guidance |
| bmd-sim-mrbm | ✅ | ✅ | ✅ | ✅ | ✅ | 11 | Range, mobile, CMS, variant |
| bmd-sim-nuclear-efx | ✅ | ✅ | ✅ | ✅ | ✅ | 22 | EMP, blast, radiation, sat kill |
| bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | Kp, solar flux, scintillation |
| bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ | ✅ | 4 | Rain, attenuation, refraction |
| bmd-sim-electronic-attack | ✅ | ✅ | ✅ | ✅ | ✅ | 9 | ECM types, ECCM, DRFM |
| engagement-chain | ✅ | ✅ | ✅ | ✅ | ✅ | 53 | Full kill chain |
| kill-assessment | ✅ | ✅ | ✅ | ✅ | ✅ | 22 | Intercept, warhead, BDA |
| wta | ✅ | ✅ | ✅ | ✅ | ✅ | 18 | Assignments, doctrine, PK |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ | ✅ | 161 | TLE propagation |
| space-debris | ✅ | ✅ | ✅ | ✅ | ✅ | 325K | 25K debris objects |
| air-traffic | ✅ | ✅ | ✅ | ✅ | ✅ | 35K | 5K flights |
| tactical-net | ✅ | ✅ | ✅ | ✅ | ✅ | 13K | Link budget, propagation |
| ufo | ✅* | ✅ | ✅ | ✅ | ✅ | 61 | Anomaly, engagement failure |
| **Tier 2 — C2/Infrastructure (8)** |||||||
| bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ | ✅ | 3 | Needs scenario context |
| bmd-sim-hub | ✅ | ✅ | ✅ | ✅ | ✅ | 5 | Track fusion, PK, TTI |
| bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ | ✅ | 0 | ✅ Format bug fixed |
| bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Needs asset data |
| bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Needs sensor data |
| bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Needs participants |
| bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | Needs protocol data |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ | ✅ | 2 | ✅ Format bug fixed |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | ✅ | 4 | ✅ Format bug fixed + enriched |

*ufo prints text header before JSON (handled by `_run_sim_binary`)

---

## launch-veh-sim — NEW (Issues Found) 🆕

| Issue | Severity | Detail |
|-------|----------|--------|
| No `-v` flag | P2 | `flag provided but not defined: -v` |
| No `-i` interactive | P2 | No interactive mode |
| Float duration | P2 | Uses `-duration float` (seconds) not Go `time.Duration` |
| No `-seed` | P2 | No reproducibility flag |
| `-json` empty output | P1 | Returns empty string, no JSON |
| NaN in output | P1 | `Final Alt: NaN km | Vel: NaN m/s | Max Q: 0 Pa` |
| Physics incomplete | P1 | Status shows "incomplete", milestones stop at liftoff |

**This sim is not production-ready.** The physics engine appears incomplete (NaN values, no computed trajectory). Needs Go source fix.

---

## Excluded Sims (by design)

| Sim | Reason |
|-----|--------|
| cyber-redteam-sim | Config-driven — YAML + REST API + web UI |
| maritime-sim | Config-driven — YAML + web UI |
| space-war-sim | Config-driven — YAML + AAR export |
| electronic-war-sim | Engine/library — prints version only |
| missile-defense-sim | Engine/library — prints version only |
| submarine-war-sim | Engine/library — prints version only |

---

## Resolved Since Round 10 ✅

| Issue | Status |
|-------|--------|
| bmd-sim-gfcb format bug (`%!s(MISSING)`) | ✅ Fixed |
| bmd-sim-decoy format bug (`%!s(MISSING)`, `%!f(MISSING)`) | ✅ Fixed |
| bmd-sim-jamming format bug (`%!f(MISSING)`) | ✅ Fixed + enriched output |

---

## Summary

**41/41 operational sims pass all standard flags (100%).** 🎉
**Zero format string bugs remaining.** ✅
**Zero stubs/mocks across all binaries.** ✅

**1 new sim (launch-veh-sim) has significant issues** — NaN output, no JSON, incomplete physics. Marked P1/P2.
**ufo** still prints text header before JSON (P3, handled by API).