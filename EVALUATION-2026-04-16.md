# FORGE-Sims Full Test & Evaluation

**Date:** 2026-04-16 (Updated)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Summary

| Category | Total | Pass | Fail |
|----------|-------|------|------|
| Binaries execute | 30 | 30 | 0 |
| `-v` verbose | 30 | 30 | 0 |
| `-i` / `--interactive` | 30 | 30 | 0 |
| `--json` output | 30 | 30 | 0 |
| `--duration` timed run | 30 | 30 | 0 |
| `--tick` sim speed | 30 | 30 | 0 |
| `--seed` deterministic | 30 | 30 | 0 |
| Scenario flags work | 8 tested | 8 | 0 |

---

## Fixes Applied This Session

### 1. C2BMC CorrelateTracks Nil Pointer Crash ✅ FIXED
**Root cause:** `CorrelateTracks()` deleted tracks from map during iteration.  
**Fix:** Collect-then-merge pattern. All 9 unit tests pass. Binary exits 0.

### 2. Missing `-v` Flag ✅ FIXED (was 7/30, now 30/30)
**Fix:** Added sim-cli shared library with `-v` support to all 24 new sims.

### 3. Missing `-i`/`--interactive` Mode ✅ ADDED (was 0/30, now 30/30)
**Fix:** All 30 binaries now support interactive mode with:
- Live `[T+N]` timestamp progression
- `--duration` for timed runs
- `--tick` for simulation speed
- `ctrl+C` to quit

### 4. Missing `--json` Output ✅ ADDED (was 0/30, now 30/30)
**Fix:** Machine-readable JSON output for all sims.

### 5. Missing `--duration`, `--tick`, `--seed` ✅ ADDED
**Fix:** All sims support timed runs, custom tick rates, and deterministic seeds.

---

## Binary-by-Binary Results

### ✅ All 30 Binaries Working

| # | Binary | Exit | `-v` | `-i` | `--json` | Custom Flags |
|---|--------|------|------|------|----------|-------------|
| 1 | bmd-sim-sbirs | 0 | ✅ | ✅ | ✅ | `-site`, `-scenario` |
| 2 | bmd-sim-stss | 0 | ✅ | ✅ | ✅ | — |
| 3 | bmd-sim-dsp | 0 | ✅ | ✅ | ✅ | — |
| 4 | bmd-sim-uewr | 0 | ✅ | ✅ | ✅ | `-scenario` |
| 5 | bmd-sim-lrdr | 0 | ✅ | ✅ | ✅ | — |
| 6 | bmd-sim-cobra-judy | 0 | ✅ | ✅ | ✅ | — |
| 7 | bmd-sim-gmd | 0 | ✅ | ✅ | ✅ | `-site`, `-variant` |
| 8 | bmd-sim-sm3 | 0 | ✅ | ✅ | ✅ | `-variant` |
| 9 | bmd-sim-sm6 | 0 | ✅ | ✅ | ✅ | — |
| 10 | bmd-sim-thaad-er | 0 | ✅ | ✅ | ✅ | — |
| 11 | bmd-sim-thaad | 0 | ✅ | ✅ | ✅ | — |
| 12 | bmd-sim-patriot | 0 | ✅ | ✅ | ✅ | — |
| 13 | bmd-sim-aegis | 0 | ✅ | ✅ | ✅ | — |
| 14 | bmd-sim-icbm | 0 | ✅ | ✅ | ✅ | `-mirvs`, `-cms` |
| 15 | bmd-sim-irbm | 0 | ✅ | ✅ | ✅ | — |
| 16 | bmd-sim-hgv | 0 | ✅ | ✅ | ✅ | `-maneuvers` |
| 17 | bmd-sim-slcm | 0 | ✅ | ✅ | ✅ | — |
| 18 | bmd-sim-decoy | 0 | ✅ | ✅ | ✅ | — |
| 19 | bmd-sim-jamming | 0 | ✅ | ✅ | ✅ | — |
| 20 | bmd-sim-c2bmc | 0 | ✅ | ✅ | ✅ | — |
| 21 | bmd-sim-gfcb | 0 | ✅ | ✅ | ✅ | — |
| 22 | bmd-sim-ifxb | 0 | ✅ | ✅ | ✅ | — |
| 23 | bmd-sim-jrsc | 0 | ✅ | ✅ | ✅ | — |
| 24 | bmd-sim-link16 | 0 | ✅ | ✅ | ✅ | — |
| 25 | bmd-sim-jreap | 0 | ✅ | ✅ | ✅ | — |
| 26 | bmd-sim-space-weather | 0 | ✅ | ✅ | ✅ | — |
| 27 | bmd-sim-atmospheric | 0 | ✅ | ✅ | ✅ | — |
| 28 | bmd-sim-gbr | 0 | ✅ | ✅ | ✅ | — |
| 29 | bmd-sim-hub | 0 | ✅ | ✅ | ✅ | — |
| 30 | bmd-sim-tpy2 | 0 | ✅ | ✅ | ✅ | `-site`, `-scenario` |

---

## Physics & Realism Spot-Check

| Check | Result | Notes |
|-------|--------|-------|
| GEO altitude (SBIRS) | ✅ 42,164 km | Correct |
| HEO orbit (SBIRS) | ✅ ~23,560 km | Molniya orbit |
| ICBM flight time | ✅ 1800s (30 min) | Realistic |
| ICBM reentry speed | ✅ Mach 21 | Consistent |
| SM-3 IIA range/speed | ✅ 2500km / Mach 13 | Matches public data |
| UEWR frequency | ✅ 435 MHz UHF | Correct |
| LRDR S-band | ✅ 3.0 GHz | Correct |
| HGV detection by SBIRS | ✅ DETECTABLE | Boost phase |
| HGV detection by UEWR | ✅ NOT DETECTABLE | Below radar horizon |

---

## Shared Library: sim-cli

**Repo:** `git@idm.wezzel.com:crab-meat-repos/sim-cli.git`  
**Path:** `/home/wez/sim-cli/` on darth

Provides standardized CLI flags for all simulators:
- `cli.Parse()` — one-call setup for simple sims
- `cli.RegisterFlags()` — for sims with custom flags
- `cli.WriteJSON()` — machine-readable output
- `cli.Flags` struct — `Verbose`, `Interactive`, `JSON`, `Duration`, `Tick`, `Seed`

---

## Recommendations

### Completed ✅
1. ~~Fix C2BMC nil pointer crash~~
2. ~~Add `-i` / `--interactive` mode~~
3. ~~Standardize `-v` flag~~
4. ~~Add `--json` output flag~~

### Remaining
5. **Interactive TUI dashboard** — Current `-i` is a simple ticker. Full bubbletea/tcell dashboard with panels would be next step.
6. **Verify Aegis/Patriot detection ranges** — May be unreasonably high
7. **Add `--output` flag** for writing results to file

---

**Result: ALL PASS ✅**