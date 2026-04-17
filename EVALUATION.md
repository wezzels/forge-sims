# FORGE-Sims Evaluation — Final

**Date:** 2026-04-16 (Round 3 — Final)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail |
|----------|-------|------|------|
| Binaries execute (exit 0) | 30 | 30 | 0 |
| `-v` verbose with sim details | 30 | 30 | 0 |
| `-i` / `--interactive` with live metrics | 30 | 30 | 0 |
| `-json` with parameters | 30 | 30 | 0 |
| `-duration` / `-tick` / `-seed` | 30 | 30 | 0 |
| Double-tick bug | 30 | 30 ✅ | 0 |
| Decoy format string bug | 1 | 1 ✅ | 0 |

**All previous corrections resolved.**

---

## All 30 Sims — Confirmed Working

| # | Binary | Default | `-v` | `-i` Metrics | `-json` | Special Flags |
|---|--------|---------|------|-------------|---------|---------------|
| 1 | bmd-sim-sbirs | ✅ | ✅ | Sats: 6, GEO: 4, HEO: 2 | ✅ | — |
| 2 | bmd-sim-stss | ✅ | ✅ | ✅ | ✅ | — |
| 3 | bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ | — |
| 4 | bmd-sim-uewr | ✅ | ✅ | Sites: 4, Freq: 435 MHz | ✅ | — |
| 5 | bmd-sim-lrdr | ✅ | ✅ | Tracks: 1000, Freq: 3.0 GHz | ✅ | — |
| 6 | bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ | — |
| 7 | bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ | `-site`, `-scenario` |
| 8 | bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ | — |
| 9 | bmd-sim-aegis | ✅ | ✅ | Tracks: 2, Range(0.1m²): 31605 km | ✅ | — |
| 10 | bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ | — |
| 11 | bmd-sim-thaad | ✅ | ✅ | Launchers: 2, Missiles: 6, PK: 90% | ✅ | — |
| 12 | bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ | — |
| 13 | bmd-sim-gmd | ✅ | ✅ | Site: FTG, EKV Divert: 800 m/s | ✅ | `-site`, `-variant` |
| 14 | bmd-sim-sm3 | ✅ | ✅ | Variant: IIA, Range: 2500 km, Mach: 13 | ✅ | `-variant` |
| 15 | bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ | — |
| 16 | bmd-sim-icbm | ✅ | ✅ | Range: 10000 km, MIRVs: 3, Flight: 30 min | ✅ | `-mirvs`, `-cms` |
| 17 | bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ | — |
| 18 | bmd-sim-hgv | ✅ | ✅ | Mach: 5-25, Range: 8000 km, Maneuvers: 5 | ✅ | `-maneuvers` |
| 19 | bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ | — |
| 20 | bmd-sim-decoy | ✅ | ✅ | Decoys: 2, Active: 2 | ✅ | — |
| 21 | bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | — |
| 22 | bmd-sim-c2bmc | ✅ | ✅ | Tracks: 0, Active: 0, Threats: 0 | ✅ | — |
| 23 | bmd-sim-hub | ✅ | ✅ | ✅ | ✅ | — |
| 24 | bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ | — |
| 25 | bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ | — |
| 26 | bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ | — |
| 27 | bmd-sim-link16 | ✅ | ✅ | Net: BMDS-NET-1, Slots: 1535 | ✅ | — |
| 28 | bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | — |
| 29 | bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ | — |
| 30 | bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ | — |

---

## Previously Reported Issues — All Resolved

| Issue | Status | Notes |
|-------|--------|-------|
| C2BMC nil pointer crash | ✅ Fixed | Runs clean, exit 0 |
| Interactive mode was stub (`Running...` only) | ✅ Fixed | Now shows live per-tick metrics |
| Double-tick printing | ✅ Fixed | Single print per tick |
| JSON was minimal (`{simulator, status}`) | ✅ Fixed | Now includes `parameters` object with sim data |
| DecoyType format string (`%!s(...)`) | ✅ Fixed | Shows `BALLOON`, `INFLATABLE_RV` |
| `-v` regression (showed only seed/tick) | ✅ Fixed | Shows sim-specific details again |
| 6 binaries missing new flags | ✅ Fixed | aegis, gbr, hub, patriot, thaad, tpy2 rebuilt |

---

## Remaining Enhancements (Not Bugs)

| Priority | Enhancement | Notes |
|----------|-------------|-------|
| P1 | Full TUI/ANSI dashboard in `-i` mode | Currently line-by-line metrics, not panel layout |
| P2 | Richer JSON output | Add tracks, detections, trajectory data (not just parameters) |
| P2 | Keyboard controls in `-i` | pause/step/quit keys beyond ctrl+C |
| P3 | Scenario-driven dynamic metrics | Interactive metrics currently static per tick; should evolve with scenario time |

---

## Sample Interactive Output

```
SBIRS Satellite Constellation Simulator
=======================================
  Mode: INTERACTIVE (ctrl+C to quit)
[T+1s]   Sats: 6  GEO: 4  HEO: 2
[T+2s]   Sats: 6  GEO: 4  HEO: 2
[T+3s]   Sats: 6  GEO: 4  HEO: 2
  Duration reached.
```

```
THAAD Battery Simulator
========================
  Mode: INTERACTIVE (ctrl+C to quit)
[T+1s] Launchers: 2  Missiles: 6  PK: 90%
[T+2s] Launchers: 2  Missiles: 6  PK: 90%
[T+3s] Launchers: 2  Missiles: 6  PK: 90%
  Duration reached.
```

## Sample JSON Output

```json
{
  "simulator": "icbm",
  "status": "complete",
  "parameters": {
    "apogee_km": 1200,
    "flight_time_min": 30,
    "mirvs": 3,
    "range_km": 10000,
    "reentry_mach": 20.59,
    "stages": 3
  }
}
```