# FORGE-SDN API Reference

Complete reference for the FORGE-SDN REST API and programmatic interfaces.

---

## REST API Server

### Starting the Server

The API server is included in the source code but not exposed via the CLI by default. To run it:

```go
package main

import "gitlab.com/crab-meat-repos/forge-sdn/internal/scenario"

func main() {
    api := scenario.NewAPIServer(8080)
    api.Start()
}
```

Or start both CLI and API by adding to `cmd/sim/main.go`:
```go
go api.Start()  // non-blocking
```

---

## Endpoints

### GET /api/status

Returns server status and constellation overview.

**Response:**
```json
{
  "version": "1.0.0",
  "status": "running",
  "sats": 791,
  "timestamp": "2026-05-02T19:13:51Z"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `version` | string | Simulator version |
| `status` | string | `running` or `error` |
| `sats` | int | Total satellite count |
| `timestamp` | string | ISO 8601 timestamp |

---

### GET /api/architecture

Returns detailed SDN architecture breakdown.

**Response:**
```json
{
  "backbone": {
    "sats": 432,
    "planes": 12,
    "altitude_km": 1000
  },
  "space_link": {
    "sats": 21,
    "planes": 3,
    "altitude_km": 800
  },
  "transport": {
    "t0": 20,
    "t1": 126,
    "t2": 192,
    "total": 338
  },
  "total_sats": 791,
  "geps": 7
}
```

---

### POST /api/step

Advance the simulation by one time step.

**Request:** (none — empty POST)

**Response:** Current `SDNNetworkState` JSON:
```json
{
  "time_min": 15,
  "total_sats": 791,
  "active_links": 1590,
  "total_bw_gbps": 3182.25,
  "coverage_pct": 100,
  "routes_valid": 4,
  "routes_total": 7,
  "scheduled_links": {
    "active": 1520,
    "handoff": 15,
    "acquiring": 12,
    "degraded": 3,
    "lost": 0
  }
}
```

---

### POST /api/run

Run a complete simulation from start to finish.

**Request:**
```json
{
  "duration_min": 60,
  "step_min": 15
}
```

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `duration_min` | float | 60 | Total simulation duration |
| `step_min` | float | 15 | Time step size |

**Response:**
```json
{
  "duration_min": 60,
  "step_min": 15,
  "steps": 4,
  "final": { ... }
}
```

---

### GET /api/route

Find the shortest path between two nodes in the SDN mesh.

**Query Parameters:**
| Param | Default | Description |
|-------|---------|-------------|
| `from` | `GEP-0` | Source node ID |
| `to` | `GEP-2` | Destination node ID |

**Request:**
```
GET /api/route?from=GEP-0&to=GEP-3
```

**Response:**
```json
{
  "from": "GEP-0",
  "to": "GEP-3",
  "path": ["GEP-0", "BB-0-0", "BB-1-5", "BB-3-0", "GEP-3"],
  "cost": 3.2,
  "bandwidth": 10000,
  "hops": 5,
  "valid": true
}
```

| Field | Type | Description |
|-------|------|-------------|
| `path` | [string] | Ordered list of nodes along route |
| `cost` | float | OSPF path cost (lower = better) |
| `bandwidth` | float | Minimum bandwidth along path (Mbps) |
| `hops` | int | Number of hops (path length - 1) |
| `valid` | bool | `true` if path exists and cost is finite |

**Node ID formats:**
- GEPs: `GEP-0` through `GEP-6`
- Backbone: `BB-<plane>-<index>` (e.g., `BB-0-5`, `BB-11-35`)
- Space Link: `SL-<plane>-<index>` (e.g., `SL-0-6`)
- Transport: `PW-T<tranche>-<plane>-<index>` (e.g., `PW-T1-3-12`)

---

### GET /api/scenarios

List all available scenarios.

**Response:**
```json
{
  "scenarios": [
    {
      "id": "peacetime",
      "name": "Peacetime Baseline",
      "description": "Normal SDN operations without threats"
    },
    {
      "id": "golden-dome",
      "name": "Golden Dome - SBI Kill Chain",
      "description": "Simulates Space-Based Interceptor communications under threat"
    },
    {
      "id": "nuclear",
      "name": "Nuclear Effects on LEO Constellation",
      "description": "400 kT detonation at 400 km — DELTA belt formation"
    }
  ]
}
```

---

### POST /api/scenario/run

Run a named scenario.

**Request:**
```json
{
  "scenario": "golden-dome"
}
```

| Parameter | Values | Default |
|-----------|--------|---------|
| `scenario` | `peacetime`, `golden-dome`, `nuclear` | `peacetime` |

**Response:** Full `ScenarioResult` JSON (see README.md JSON Output Format section).

---

## FORGE-C2 / NORAD API

When deployed on the NORAD Forge server, FORGE-SDN is accessible via:

```
GET https://norad.stsgym.com/api/public/sdn?scenario=peacetime
```

**Response:** Same `ScenarioResult` format as `-output` JSON.

---

## Programmatic Go API

### Creating an SDN Network

```go
import "gitlab.com/crab-meat-repos/forge-sdn/internal/simulation"

// Default 791-satellite architecture
net := simulation.NewSDNNetwork()

// Custom architecture
sdn := simulation.SDNArchitecture{
    Backbone:   simulation.WalkerConfig{T: 432, P: 12, F: 3, AltitudeKm: 1000, Inclination: 55},
    SpaceLink:  simulation.WalkerConfig{T: 21, P: 3, F: 1, AltitudeKm: 800, Inclination: 60},
    TransportT0: simulation.WalkerConfig{T: 20, P: 2, F: 1, AltitudeKm: 1000, Inclination: 28.5},
    TransportT1: simulation.WalkerConfig{T: 126, P: 6, F: 2, AltitudeKm: 900, Inclination: 55},
    TransportT2: simulation.WalkerConfig{T: 192, P: 8, F: 4, AltitudeKm: 1100, Inclination: 70},
}
net := simulation.NewSDNNetworkWithConfig(sdn)
```

### Running a Simulation

```go
// Single step
state := net.Step()
fmt.Printf("Active links: %d, Coverage: %.1f%%\n",
    state.ActiveLinks, state.CoveragePct)

// Full run with callback
net.RunWithCallback(60, 15, func(state simulation.SDNNetworkState) {
    fmt.Printf("t=%dmin: %d sats, %d links\n",
        int(state.TimeMin), state.TotalSats, state.ActiveLinks)
})
```

### Scenario Runner

```go
import "gitlab.com/crab-meat-repos/forge-sdn/internal/scenario"

cfg := scenario.GoldenDomeScenario()
runner := scenario.NewScenarioRunner(cfg)
result := runner.Run()

fmt.Printf("Resilience: %.1f / 100\n", result.Summary.ResilienceScore)
fmt.Printf("Sats lost: %d\n", 791 - result.Summary.FinalSats)

// Save to file
scenario.SaveResult(result, "golden_dome_result.json")
```

### Querying Routes

```go
result := net.Router.RunSPF("GEP-0")
path, cost, bw := result.PathTo("GEP-3")
fmt.Printf("Route: %v, Cost: %.2f, BW: %.0f Mbps\n", path, cost, bw)
```

### Applying Threats

```go
import "gitlab.com/crab-meat-repos/forge-sdn/internal/threat"

// ASAT kill
asat := threat.ASATAttack{Target: "BB-0-5", WeaponType: "DA-ASAT"}
net.ApplyThreat(asat)

// Jamming
jam := threat.JammingAttack{Target: "KaS2G", JammerPowerW: 10000}
net.ApplyThreat(jam)

// Nuclear detonation
nuke := threat.NuclearDetonation{YieldKT: 400, AltitudeKm: 400}
net.ApplyThreat(nuke)
```

---

## Data Types Reference

### SDNNetworkState

| Field | Type | Description |
|-------|------|-------------|
| TimeMin | float64 | Elapsed time (minutes) |
| TotalSats | int | Active satellite count |
| ActiveLinks | int | Active ISL count |
| TotalBWGbps | float64 | Aggregate bandwidth |
| CoveragePct | float64 | Global coverage % |
| RoutesValid | int | Valid OSPF routes |
| RoutesTotal | int | Total possible routes |

### WalkerConfig

| Field | Type | Description |
|-------|------|-------------|
| T | int | Total satellites |
| P | int | Number of orbital planes |
| F | int | Phase factor (Walker delta) |
| AltitudeKm | float64 | Orbital altitude (km) |
| Inclination | float64 | Orbital inclination (degrees) |

### LinkInfo

| Field | Type | Description |
|-------|------|-------------|
| From | string | Source satellite ID |
| To | string | Destination satellite ID |
| BandwidthMbps | float64 | Link bandwidth |
| LatencyMs | float64 | Link latency |
| Type | string | `ISL` (inter-satellite) or `S2G` (space-to-ground) |
| State | string | Link scheduler state |

### ScenarioConfig

| Field | Type | Description |
|-------|------|-------------|
| Name | string | Scenario name |
| Description | string | Human-readable description |
| DurationMin | float64 | Duration (minutes) |
| StepMin | float64 | Step size (minutes) |
| Threats | []ThreatEvent | Ordered list of threat events |
| Metrics | []string | Metrics to track |
| OutputFile | string | JSON output file path |