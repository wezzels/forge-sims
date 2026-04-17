# FORGE-Sims Evaluation — Final

**Date:** 2026-04-17 (Round 4)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail |
|----------|-------|------|------|
| Binaries execute (exit 0) | 33 | 33 | 0 |
| `-v` verbose flag | 33 | 30 | 3 |
| `-i` interactive with live metrics | 33 | 33 | 0 |
| `-json` output | 33 | 33 | 0 |
| `-duration` / `-tick` / `-seed` | 33 | 33 | 0 |

**33 total simulators: 30 BMDS + 3 new (air-traffic, satellite-tracker, space-debris)**

---

## All 33 Sims — Verified

### New Simulators (3)

| # | Binary | Default | `-v` | `-i` | `-json` | Special Flags | Issues |
|---|--------|---------|------|------|---------|---------------|--------|
| 31 | air-traffic | ✅ | ❌ | ✅ | ✅ | `-count`, `-seed` | Missing `-v` |
| 32 | satellite-tracker | ✅ | ❌ | ✅ | ✅ | `-group`, `-count` | Missing `-v`; `active`/`starlink` groups return 0 sats |
| 33 | space-debris | ✅ | ❌ | ✅ | ✅ | `-count`, `-collisions`, `-threshold`, `-seed` | Missing `-v`; `-collisions` shows no close approaches |

#### air-traffic
Generates flights across 40 major hubs. Shows callsigns, aircraft types (B763, A388, B52, C172), altitude, speed, heading, phase (CRUISE/CLIMB/DESCENT/APPROACH). JSON includes per-flight data with origin, dest, lat, lon, progress.

#### satellite-tracker
Fetches live TLE data from Celestrak. Supports groups: `stations` (ISS, CSS - 30 sats), `gps-ops` (32 sats), `weather`, `starlink`, `active`. Shows NORAD ID, inclination, altitude, period, lat/lon. JSON includes per-satellite orbital data. ⚠️ `active` and `starlink` groups return 0 satellites (likely timeout or size limit on Celestrak fetch).

#### space-debris
Generates up to 25,000 debris objects with full orbital elements (inclination, RAAN, eccentricity, arg_perigee, mean_anomaly, mean_motion). Categories: LEO, GEO, MEO, HIGH. Types: DEBRIS, ROCKET_BODY, PAYLOAD. JSON includes per-object orbital data with lat/lon/RCS. ⚠️ `-collisions` flag accepted but no close approaches reported in output.

### BMDS Simulators (30)

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

## Previously Reported Issues — All Resolved ✅

| Issue | Status |
|-------|--------|
| C2BMC nil pointer crash | ✅ Fixed |
| Interactive mode was stub (`Running...` only) | ✅ Fixed — shows live metrics |
| Double-tick printing | ✅ Fixed |
| JSON was minimal (`{simulator, status}`) | ✅ Fixed — includes parameters/data |
| DecoyType format string (`%!s(...)`) | ✅ Fixed — shows BALLOON/INFLATABLE_RV |
| `-v` regression | ✅ Fixed — shows sim-specific details |
| 6 binaries missing new flags | ✅ Fixed — all rebuilt |

---

## Corrections Needed

| Priority | Issue | Binary | Fix |
|----------|-------|--------|-----|
| **P1** | Missing `-v` flag | air-traffic, satellite-tracker, space-debris | Add verbose mode |
| P2 | `active`/`starlink` groups return 0 sats | satellite-tracker | Likely Celestrak timeout on large groups; add timeout/retry or paginate |
| P2 | `-collisions` shows no close approaches | space-debris | Flag accepted but no collision data in output; may need more objects or propagation time |
| P3 | Missing `--interactive` long flag | air-traffic, satellite-tracker, space-debris | BMDS sims have both `-i` and `--interactive` |
| P3 | `-duration` format inconsistent | all | BMDS sims use Go duration (`3s`), new sims use int seconds (`3`) |

---

## Remaining Enhancements (Not Bugs)

| Priority | Enhancement | Notes |
|----------|-------------|-------|
| P1 | Full TUI/ANSI dashboard in `-i` mode | Currently line-by-line metrics, not panel layout |
| P2 | Richer JSON output for BMDS sims | Add tracks, detections, trajectory data |
| P2 | Keyboard controls in `-i` | pause/step/quit keys beyond ctrl+C |
| P3 | Scenario-driven dynamic metrics | Interactive metrics currently static per tick |

---

## Sample Outputs

### air-traffic -i
```
[T+0s] 100 flights
  COMMERCIAL      39    CARGO           17
  PRIVATE         14    MILITARY        30
  CRUISE     70   CLIMB      20
  DESCENT    8    APPROACH   2
```

### satellite-tracker -i (stations)
```
[T+0s] Tracking 30 satellites
  Visible: 30 | Eclipsed: 0
  MANNED               6
  OTHER                24
```

### space-debris -i
```
[T+0s] 5000 debris objects
  LEO        3030  GEO         510
  HIGH        953  MEO         507
```

### satellite-tracker -json
```json
{
  "parameters": {"group": "gps-ops", "time": "2026-04-17T18:06:33Z", "tle_count": 5},
  "satellites": [{"name": "GPS BIIR-2 (PRN 13)", "id": 24876, "lat": 0.015, "lon": 101.0, "alt": 20040.3, "vis": "visible"}],
  "simulator": "satellite-tracker", "status": "complete"
}
```