# space-data-network-sim

**FORGE-SDN v1.0.0** — USSF Space Data Network Simulator

---

## Overview

Simulates the Space Force's Space Data Network (SDN) — a constellation of 791 satellites providing communications, navigation, and data relay for missile defense operations.

**Binary Size:** 4.5 MB (linux-x86)
**Platforms:** linux-x86, linux-arm64, darwin-arm64, windows-amd64

---

## Architecture

| Constellation | Satellites | Planes | Altitude | Purpose |
|--------------|------------|--------|----------|---------|
| **Backbone (MILNET)** | 432 | 12 | 1000km | Primary comms backbone |
| **Space Link (GD SBI)** | 21 | 3 | 800km | SBI kill chain comms |
| **Transport (PWSA)** | 338 | — | Various | Proliferated Warfighter Space Architecture |
| **Total** | **791** | — | — | — |

### PWSA Breakdown
- T0 (Tracking): 20 satellites
- T1 (Transport Layer Alpha): 126 satellites
- T2 (Transport Layer Beta): 192 satellites

---

## Usage

```bash
# Run peacetime scenario (default)
./space-data-network-sim -scenario peacetime -output result.json

# Run golden dome (SBI kill chain) scenario
./space-data-network-sim -scenario golden-dome -output result.json

# Run nuclear attack scenario
./space-data-network-sim -scenario nuclear -output result.json

# Custom scenario with duration override
./space-data-network-sim -scenario custom -duration 60 -step 5 -output result.json
```

---

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-scenario` | string | `peacetime` | Scenario: peacetime, golden-dome, nuclear, custom |
| `-duration` | float | scenario default | Duration override (minutes, 0=use scenario default) |
| `-step` | float | scenario default | Time step size (minutes, 0=use scenario default) |
| `-output` | string | — | Output JSON file path |

---

## JSON Output Schema

```json
{
  "name": "Peacetime Baseline",
  "description": "Normal SDN operations without threats",
  "start_time": "2026-05-02T13:04:48-06:00",
  "duration_min": 0.1,
  "time_series": [
    {
      "time_min": 15,
      "total_sats": 791,
      "active_links": 1590,
      "total_bw_gbps": 3192,
      "coverage_pct": 100,
      "routes_valid": 5,
      "routes_total": 7
    }
  ],
  "threat_log": null,
  "summary": {
    "final_sats": 791,
    "final_links": 1590,
    "final_coverage_pct": 100,
    "final_routes_pct": 71.4,
    "resilience_score": 92.9,
    "threats_executed": 0,
    "sats_lost": 0,
    "peak_latency_ms": 0
  }
}
```

### Summary Fields

| Field | Type | Description |
|-------|------|-------------|
| `final_sats` | int | Satellites remaining at end of sim |
| `final_links` | int | Active communication links |
| `final_coverage_pct` | float | Global coverage percentage |
| `final_routes_pct` | float | Valid route percentage |
| `resilience_score` | float | Overall network resilience (0-100) |
| `threats_executed` | int | Number of threat events executed |
| `sats_lost` | int | Satellites destroyed/lost |
| `peak_latency_ms` | int | Maximum network latency (ms) |

---

## Scenarios

### Peacetime

**Description:** Normal SDN operations without threats

**Expected Results:**
- Satellites: 791/791 (100%)
- Links: 1590
- Coverage: 100%
- Routes: 71.4%
- Resilience: 92.9/100
- Losses: 0

### Golden Dome

**Description:** Simulates Space-Based Interceptor communications under various threat conditions

**Expected Results:**
- Satellites: 791/791 (100%)
- Links: 1590
- Coverage: 100%
- Routes: 57.1% (reduced due to SBI tasking)
- Resilience: 89.3/100
- Losses: 0

### Nuclear

**Description:** Single nuclear detonation at 400km altitude (0kT test — EMP effects only)

**Expected Results:**
- Satellites: 25/791 (3.2% survive)
- Links: 50
- Coverage: 22.5%
- Routes: 28.6%
- Resilience: 21.1/100
- Losses: 766 (96.8%)
- Threats: 1

---

## Physics Fidelity

| Aspect | Fidelity | Notes |
|--------|----------|-------|
| Orbital mechanics | Medium | Simplified two-body propagation |
| Link budget | Medium | Free-space path loss, basic antenna gain |
| Routing | Medium | Graph-based pathfinding |
| Threat effects | High | Nuclear EMP damage model based on altitude/yield |
| Constellation geometry | High | Realistic plane/count distributions |

---

## Known Issues

| Priority | Issue | Status |
|----------|-------|--------|
| P1 | Missing `-json` flag (stdout output) | Open — uses `-output` file instead |
| P1 | Missing `-i` interactive mode | Open — no live dashboard |
| P2 | Missing `-v` verbose flag | Open |
| P2 | Inconsistent with BMDS flag convention | Open — should support `-json -duration` |

---

## Source

**Repo:** `git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git`
**Branch:** main

---

## Related Sims

- `bmd-sim-stss` — Space Tracking and Surveillance System
- `bmd-sim-sbirs` — Space-Based Infrared System
- `bmd-sim-space-weather` — Space weather effects
- `satellite-tracker` — NORAD TLE tracking

---

## Changelog

### v1.0.0 (2026-05-02)
- Initial release
- 791-satellite constellation
- 3 scenarios: peacetime, golden-dome, nuclear
- JSON output via `-output` flag
