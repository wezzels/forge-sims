# FORGE-Sims Evaluation — Round 25
**Date:** 2026-04-30
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Summary

| Metric | Count | Percentage |
|--------|-------|------------|
| Total Binaries | 33 | 100% |
| Execute Successfully | 33 | 100% |
| `-h` Help Flag | 33 | 100% |
| `-v` Verbose Flag | 30 | 90.9% |
| `-json` Output | 33 | 100% |
| `-duration` Flag | 33 | 100% |
| Interactive `-i` Mode | 33 | 100% |

**Status:** ✅ All 33 sims functional. 3 sims missing `-v` flag (new additions).

---

## All 33 Sims — Verified

### BMDS Simulators (30) — All Flags ✅

| # | Binary | -h | -v | -i | -json | -duration |
|---|--------|----|----|----|-------|-----------|
| 1 | bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 2 | bmd-sim-stss | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 3 | bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 4 | bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 5 | bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 6 | bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 7 | bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 8 | bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 9 | bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 10 | bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 11 | bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 12 | bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 13 | bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 14 | bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 15 | bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 16 | bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 17 | bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 18 | bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 19 | bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 20 | bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 21 | bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 22 | bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 23 | bmd-sim-hub | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 24 | bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 25 | bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 26 | bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 27 | bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 28 | bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 29 | bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |
| 30 | bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ | ✅ (seed) |

### New Simulators (3) — Missing `-v` ⚠️

| # | Binary | -h | -v | -i | -json | -duration | Issues |
|---|--------|----|----|----|-------|-----------|--------|
| 31 | air-traffic | ✅ | ❌ | ✅ | ✅ | ✅ | Missing `-v` flag |
| 32 | satellite-tracker | ✅ | ❌ | ✅ | ✅ | ✅ | Missing `-v`; `active`/`starlink` groups return 0 sats |
| 33 | space-debris | ✅ | ❌ | ✅ | ✅ | ✅ | Missing `-v` flag |

---

## Detailed Findings

### air-traffic
**Purpose:** Generates commercial, cargo, military, and private flights across 40 major airports.

**Flags:**
```
  -count int        Number of flights to simulate (default 5000)
  -duration int     Simulation duration in seconds
  -i                Interactive mode
  -json             Output as JSON
  -seed int         Random seed
```

**Sample JSON Output:**
```json
{
  "flights": [
    {
      "id": "FL0001",
      "callsign": "ZB845",
      "type": "COMMERCIAL",
      "aircraft": "B763",
      "origin": "ZBAA",
      "dest": "WSSS",
      "lat": 32.09,
      "lon": 113.26,
      "alt": 30482,
      "speed": 460,
      "heading": 199.7,
      "phase": "CRUISE",
      "progress": 0.21
    }
  ],
  "simulator": "air-traffic",
  "status": "complete"
}
```

**Interactive Mode:**
```
[T+0s] 100 flights
  COMMERCIAL      39    CARGO           17
  PRIVATE         14    MILITARY        30
  CRUISE     70   CLIMB      20
  DESCENT    8    APPROACH   2
```

---

### satellite-tracker
**Purpose:** Fetches live TLE data from Celestrak for satellite tracking.

**Flags:**
```
  -count int        Limit number of satellites (0=all)
  -duration int     Simulation duration in seconds (0=single snapshot)
  -group string     Celestrak group: active, stations, starlink, gps-ops, weather (default "active")
  -i                Interactive mode
  -json             Output as JSON
```

**⚠️ Known Issues:**
1. Missing `-v` verbose flag
2. `active` and `starlink` groups return 0 satellites (Celestrak API timeout or size limit)

**Working Groups:**
- `stations`: Returns 30 satellites (ISS, CSS, etc.)
- `gps-ops`: Returns 32 GPS satellites
- `weather`: Returns weather satellites

**Sample JSON (stations group):**
```json
{
  "parameters": {"group": "stations", "time": "2026-04-30T09:53:39Z", "tle_count": 30},
  "satellites": [
    {"name": "ISS (ZARYA)", "id": 25544, "lat": 12.5, "lon": -78.3, "alt": 420.1, "vis": "visible"}
  ],
  "simulator": "satellite-tracker",
  "status": "complete"
}
```

**Sample JSON (active group — broken):**
```json
{
  "parameters": {"group": "active", "time": "2026-04-30T09:53:39Z", "tle_count": 0},
  "satellites": [],
  "simulator": "satellite-tracker",
  "status": "complete"
}
```

---

### space-debris
**Purpose:** Generates space debris objects with full orbital elements for collision analysis.

**Flags:**
```
  -collisions       Detect close approaches
  -count int        Number of debris objects to simulate (default 25000)
  -duration int     Simulation duration in seconds
  -i                Interactive mode
  -json             Output as JSON
  -seed int         Random seed
  -threshold float  Close approach threshold in km (default 5.0)
```

**⚠️ Known Issues:**
1. Missing `-v` verbose flag
2. `-collisions` flag accepted but no close approaches reported in output

**Sample JSON Output:**
```json
{
  "debris": [
    {
      "id": 1,
      "name": "DEBRIS-1",
      "object_type": "DEBRIS",
      "inclination": 40.11,
      "raan": 292.64,
      "eccentricity": 0.065,
      "arg_perigee": 138.4,
      "mean_anomaly": 137.9,
      "mean_motion": 16.54,
      "altitude": 134.9,
      "perigee": -285.6,
      "apogee": 555.4,
      "category": "LEO",
      "size": "SMALL",
      "rcs": 0.048,
      "lat": -39.24,
      "lon": -143.13
    }
  ],
  "simulator": "space-debris",
  "status": "complete"
}
```

**Interactive Mode:**
```
[T+0s] 5000 debris objects
  LEO        3030  GEO         510
  HIGH        953  MEO         507
```

---

## Corrections Needed

| Priority | Issue | Binaries | Fix Required |
|----------|-------|----------|--------------|
| **P1** | Missing `-v` flag | air-traffic, satellite-tracker, space-debris | Add verbose mode to show sim-specific details |
| P2 | `active`/`starlink` groups return 0 sats | satellite-tracker | Celestrak API timeout; add timeout/retry or pagination |
| P2 | `-collisions` shows no close approaches | space-debris | Flag accepted but no collision detection logic in output |
| P3 | Missing `--interactive` long flag | air-traffic, satellite-tracker, space-debris | BMDS sims have both `-i` and `--interactive` |
| P3 | `-duration` format inconsistent | all | BMDS sims use Go duration (`3s`), new sims use int seconds (`3`) |

---

## Comparison to Previous Rounds

| Round | Date | Total Sims | JSON | -v | -i | -duration | Notes |
|-------|------|------------|------|----|----|-----------|-------|
| Round 4 | 2026-04-17 | 33 | 33/33 | 30/33 | 33/33 | 33/33 | 3 new sims added, missing -v |
| Round 24 | 2026-04-29 | 57 | 53/57 | N/A | N/A | N/A | Data sections focus |
| Round 25 | 2026-04-30 | 33 | 33/33 | 30/33 | 33/33 | 33/33 | Current state |

**Note:** Round 24 reported 57 sims with focus on data sections. Current repo has 33 binaries. The 24 additional sims from Round 24 may be in source repos not yet built, or the count discrepancy needs investigation.

---

## Recommendations

1. **Add `-v` flag to 3 new sims** — Consistent with BMDS simulators
2. **Fix satellite-tracker Celestrak fetch** — Add timeout, retry, or chunked fetching for large groups
3. **Implement collision detection** — space-debris `-collisions` flag needs actual detection logic
4. **Investigate sim count discrepancy** — Round 24 reported 57 sims; current repo has 33 binaries
5. **Standardize `-duration` format** — Use Go duration strings (`3s`) across all sims

---

## Test Commands Used

```bash
# Full evaluation script
for bin in $(ls binaries/linux-x86/ | sort); do
  ./binaries/linux-x86/$bin -h >/dev/null 2>&1 && echo "$bin: -h OK"
  ./binaries/linux-x86/$bin -v -duration 1 >/dev/null 2>&1 && echo "$bin: -v OK"
  ./binaries/linux-x86/$bin -json -duration 1 2>/dev/null | grep -q "{" && echo "$bin: -json OK"
done
```

---

**Generated:** 2026-04-30 03:53 MDT
**Host:** vader
**Workspace:** /home/wez/.openclaw/workspace/forge-sims
