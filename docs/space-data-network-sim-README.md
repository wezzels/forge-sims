# FORGE-SDN — USSF Space Data Network Simulator

**Version:** 1.0.0  
**Binary:** `space-data-network-sim`  
**Language:** Go 1.22+  
**License:** MIT  
**Source:** `git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git`  
**Binaries:** [wezzels/forge-sims](https://github.com/wezzels/forge-sims) (Linux, macOS, Windows)

---

## Overview

FORGE-SDN simulates the United States Space Force's Space Data Network — the next-generation military satellite communications architecture that replaces the legacy SATCOM infrastructure. The simulator models all three SDN layers (Backbone, Space Link, Transport) with real orbital mechanics, mesh networking, protocol models, threat effects, and dynamic link scheduling.

**Key capabilities:**
- **791-satellite constellation** across 3 layers (MILNET Backbone, Golden Dome SBI, PWSA Transport)
- **SGP4 orbital propagation** with J2 perturbation and atmospheric drag
- **OSPF mesh routing** with Dijkstra shortest-path and QoS traffic classes
- **OCT/LDPC protocol stack** (Optical Crosslink Terminal, Low-Density Parity-Check codes)
- **Dynamic link scheduling** — satellites hand off ISLs as they orbit (ACQUIRE → ACTIVE → HANDOFF → DEGRADED → TEARDOWN → LOST)
- **Threat modeling** — ASAT kinetic kill, ground jamming (JSR calculation), nuclear EMP/DELTA belt, cyber attack
- **Walker constellation generator** (delta-pattern T/P/F notation)
- **Kessler syndrome** debris cascade modeling
- **REST API** for programmatic access and integration with FORGE-C2

---

## Quick Start

### Download Binary

| Platform | Binary | Size |
|----------|--------|------|
| Linux x86_64 | `space-data-network-sim` | 4.5 MB |
| macOS ARM64 | `space-data-network-sim` | 4.4 MB |
| Windows x86_64 | `space-data-network-sim.exe` | 4.6 MB |

Download from [wezzels/forge-sims/binaries](https://github.com/wezzels/forge-sims/tree/main/binaries).

### Build from Source

```bash
git clone git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git
cd space-data-network-sim
go build -o space-data-network-sim ./cmd/sim/
```

### Run Your First Simulation

```bash
# Peacetime baseline — no threats
./space-data-network-sim -scenario peacetime

# Golden Dome scenario — ASAT, jamming, cyber threats
./space-data-network-sim -scenario golden-dome

# Nuclear effects — 400 kT detonation at 400 km
./space-data-network-sim -scenario nuclear

# Custom duration (30 min with 5-min steps)
./space-data-network-sim -scenario peacetime -duration 30 -step 5

# Save JSON output to file
./space-data-network-sim -scenario golden-dome -output result.json
```

### Expected Output

```
FORGE-SDN v1.0.0 — USSF Space Data Network Simulator
====================================================

Scenario: Peacetime Baseline
Description: Normal SDN operations without threats
Duration: 60 min, Step: 15 min
Threats: 0

--- SDN Architecture ---
  Backbone (MILNET):   432 sats, 12 planes, 1000km
  Space Link (GD SBI): 21 sats, 3 planes, 800km
  Transport (PWSA):    T0=20 + T1=126 + T2=192 = 338 sats
  Total:               791 satellites

=== Scenario: Peacetime Baseline ===
Summary:
  Final sats: 791 / 791
  Final links: 1526
  Coverage: 100.0%
  Routes: 28.6%
  Resilience: 81.1 / 100
  Sats lost: 0
  Threats executed: 0

Simulation completed in 94ms
```

---

## Command-Line Reference

```
./space-data-network-sim [flags]

Flags:
  -scenario string   Scenario to run (default "peacetime")
                     Options: peacetime, golden-dome, nuclear, custom
  -duration float     Override scenario duration in minutes (0 = scenario default)
  -step float         Override simulation step size in minutes (0 = scenario default)
  -output string      Save JSON result to file
```

### Scenarios

| Scenario | Flag | Duration | Description |
|----------|------|----------|-------------|
| Peacetime | `-scenario peacetime` | 60 min | Normal operations, no threats. Baseline coverage and routing. |
| Golden Dome | `-scenario golden-dome` or `-scenario gd` | 120 min | SBI kill chain: ASAT strikes, Ka-band jamming, cyber attack on GEP. |
| Nuclear | `-scenario nuclear` or `-scenario nuke` | 60 min | 400 kT exo-atmospheric detonation at 400 km. DELTA radiation belt kills 97% of LEO sats. |
| Custom | `-scenario custom` | 60 min | User-defined threat scenario (modify `customScenario()` in source). |

---

## SDN Architecture Model

The simulator models three distinct layers of the USSF Space Data Network:

### Layer 1: Backbone (MILNET)
- **432 satellites** across **12 orbital planes** at **1,000 km altitude**
- Optical inter-satellite links (OCT) at 10 Gbps per link
- OSPF mesh routing with 7 ground entry points (GEPs)
- Provides the resilient backbone for all SDN traffic

### Layer 2: Space Link (Golden Dome SBI)
- **21 satellites** across **3 orbital planes** at **800 km altitude**
- Supports Space-Based Interceptor communications
- Ka-band space-to-ground links (S2G) at 1.5 Gbps
- Lower altitude for lower-latency interceptor command and control

### Layer 3: Transport (PWSA)
- **338 satellites** across three tranches:
  - **T0:** 20 sats, 2 planes (initial capability)
  - **T1:** 126 sats, 6 planes (expanded capability)
  - **T2:** 192 sats, multi-plane (full constellation)
- Multi-vendor transport layer for tactical communications
- Varies by altitude and inclination per tranche

### Ground Infrastructure
- **7 Ground Entry Points (GEPs)** — CONUS-based gateway stations
- OSPF routing from GEPs through backbone to any satellite
- GEP naming: `GEP-0` through `GEP-6`

### Total: 791 Satellites
- 1,590+ active inter-satellite links (peacetime)
- 3.18 Tbps aggregate bandwidth
- 100% global coverage (peacetime)

---

## Simulation Engine

### Orbital Mechanics (SGP4)
- Full SGP4 propagator with J2 perturbation and atmospheric drag
- Walker constellation generator supports delta-pattern T/P/F notation
- Satellite naming convention: `<Layer>-<Plane>-<Index>` (e.g., `BB-0-5`, `SL-2-3`, `PW-T1-3-12`)
- Orbital period: ~105 min at 1,000 km, ~101 min at 800 km

### Mesh Networking
- Bidirectional inter-satellite links (ISLs) with configurable bandwidth
- Dijkstra shortest-path-first (SPF) routing
- QoS traffic classes: CRITICAL, HIGH, NORMAL, LOW
- Link cost model: distance-weighted with bandwidth factor

### Dynamic Link Scheduling
As satellites orbit, inter-satellite links change. The `LinkScheduler` manages the full link lifecycle:

| State | Description |
|-------|-------------|
| ACQUIRE | Satellite entering range, link establishment in progress |
| ACTIVE | Link established, carrying traffic |
| HANDOFF | Satellite approaching edge of range, preparing transition |
| DEGRADED | Link quality declining (increasing distance/angle) |
| TEARDOWN | Controlled link shutdown in progress |
| LOST | Link failed (satellite out of range or destroyed) |

**Scheduling constraints:**
- Maximum 4 concurrent links per satellite
- Minimum link elevation angle: 10°
- Handoff timeout: 60 seconds
- Link budget threshold: SNR ≥ 10 dB for OCT, C/N₀ ≥ 14 dB-Hz for Ka-band

### Protocol Stack
- **OCT (Optical Crosslink Terminal):** 10 Gbps per link, 1,550 nm wavelength
- **LDPC encoding:** Rate 1/2 and 3/4, block sizes 648–64800 bits
- **OCT framing:** HDLC-like frame structure with CRC-32
- **Link 182 spec:** Space-to-space optical link model with pointing loss

---

## Threat Modeling

### ASAT (Direct-Ascent Kinetic Kill)
- Removes target satellite from constellation
- Generates debris field (Kessler cascade modeling)
- Debris count depends on impact energy and altitude
- Example: ASAT on `BB-0-5` → 150 trackable debris fragments

### Ground Jamming (Ka-band)
- Models JSR (Jamming-to-Signal Ratio) in dB
- Link killed when JSR exceeds threshold (typically >15 dB)
- Parameters: jammer power (W), distance to satellite, antenna gain
- Example: 10 kW ground jammer → JSR = 41.7 dB → Ka-band S2G link killed

### Nuclear EMP / DELTA Belt
- Exo-atmospheric detonation creates persistent radiation belt
- Kills satellites by altitude band (lower = more exposure)
- 400 kT at 400 km: 5% survival rate (25/791 sats survive)
- Models prompt EMP, trapped radiation, and DELTA belt formation
- Long-term degradation: surviving sats experience increased single-event upsets

### Cyber Attack
- Targets ground entry points (GEPs)
- Can disable routing through affected GEP
- Network re-routes traffic through remaining GEPs
- Duration and severity configurable in scenario

---

## JSON Output Format

When using `-output filename.json`, the simulator writes a structured JSON result:

```json
{
  "name": "Golden Dome - SBI Kill Chain",
  "description": "Simulates Space-Based Interceptor communications...",
  "start_time": "2026-05-02T19:13:58Z",
  "duration_min": 120,
  "time_series": [
    {
      "time_min": 5,
      "total_sats": 791,
      "active_links": 1590,
      "total_bw_gbps": 3182.25,
      "coverage_pct": 100,
      "routes_valid": 4,
      "routes_total": 7
    }
  ],
  "threat_log": [
    {
      "time_min": 15,
      "type": "asat",
      "target": "BB-0-5",
      "hit": true,
      "debris_count": 150
    },
    {
      "time_min": 30,
      "type": "jamming",
      "target": "KaS2G",
      "jsr_db": 41.7,
      "link_killed": true
    }
  ],
  "summary": {
    "final_sats": 789,
    "final_links": 1570,
    "final_coverage_pct": 99.8,
    "final_routes_pct": 42.9,
    "resilience_score": 85.3,
    "threats_executed": 4
  }
}
```

### Field Reference

**TimeStepResult (per step):**
| Field | Type | Description |
|-------|------|-------------|
| `time_min` | float | Elapsed time in minutes |
| `total_sats` | int | Active satellite count |
| `active_links` | int | Active inter-satellite links |
| `total_bw_gbps` | float | Aggregate bandwidth (Gbps) |
| `coverage_pct` | float | Global coverage percentage |
| `routes_valid` | int | Valid OSPF routes |
| `routes_total` | int | Total possible routes |

**ThreatResult (per threat event):**
| Field | Type | Description |
|-------|------|-------------|
| `time_min` | float | Time of event |
| `type` | string | `asat`, `jamming`, `nuclear`, `cyber` |
| `target` | string | Target satellite or link ID |
| `hit` | bool | Whether ASAT struck target |
| `debris_count` | int | Trackable debris fragments |
| `jsr_db` | float | Jamming-to-Signal Ratio (jamming only) |
| `link_killed` | bool | Whether link was disabled |
| `survival_pct` | float | Satellite survival rate (nuclear only) |

**ScenarioSummary:**
| Field | Type | Description |
|-------|------|-------------|
| `final_sats` | int | Remaining active satellites |
| `final_links` | int | Remaining active links |
| `final_coverage_pct` | float | Final global coverage |
| `final_routes_pct` | float | Valid route percentage |
| `resilience_score` | float | 0–100 resilience score |
| `threats_executed` | int | Number of threats that fired |

---

## REST API

FORGE-SDN includes an HTTP REST API server for programmatic access and integration with FORGE-C2.

### Starting the API Server

```go
// In code:
api := scenario.NewAPIServer(8080)
api.Start()
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/status` | Server status, version, sat count |
| GET | `/api/architecture` | SDN layer details (backbone, space link, transport) |
| POST | `/api/step` | Advance simulation one step |
| POST | `/api/run` | Run full simulation with `duration_min` and `step_min` |
| GET | `/api/route?from=GEP-0&to=GEP-2` | Find shortest path between nodes |
| GET | `/api/scenarios` | List available scenarios |
| POST | `/api/scenario/run` | Run scenario by ID: `{"scenario": "golden-dome"}` |

### Example: Run Golden Dome via API

```bash
curl -X POST http://localhost:8080/api/scenario/run \
  -H "Content-Type: application/json" \
  -d '{"scenario": "golden-dome"}'
```

### Example: Find Route Between GEPs

```bash
curl "http://localhost:8080/api/route?from=GEP-0&to=GEP-3"
# Returns: {"from":"GEP-0","to":"GEP-3","path":["GEP-0","BB-0-0","BB-1-5","BB-3-0","GEP-3"],
#           "cost":3.2,"bandwidth":10000,"hops":5,"valid":true}
```

---

## Testing

### Run All Tests

```bash
go test -v ./...
```

**Test packages and coverage:**

| Package | Tests | Description |
|---------|-------|-------------|
| `internal/constellation` | 8+ | Walker generator, orbital mechanics, coverage |
| `internal/mesh` | 4+ | Dijkstra routing, link add/remove, bidirectional |
| `internal/oct` | 3+ | OCT terminal specs, link budget calculation |
| `internal/protocol` | 6+ | LDPC encoding, OCT framing, link model |
| `internal/routing` | 4+ | OSPF SPF, QoS class prioritization |
| `internal/scenario` | 3+ | Scenario config, runner, format output |
| `internal/sgp4` | 4+ | SGP4 propagation, J2 perturbation, drag |
| `internal/simulation` | 8+ | SDN network, link scheduler, visibility |
| `internal/threat` | 4+ | Kessler debris, ASAT, nuclear, jamming |

### Run Specific Package

```bash
go test -v ./internal/simulation/...
go test -v ./internal/threat/...
```

### Benchmark

```bash
go test -bench=. ./internal/simulation/
```

---

## Project Structure

```
space-data-network-sim/
├── cmd/sim/                      # CLI entry point
│   └── main.go                   # Flag parsing, scenario selection, output
├── internal/
│   ├── constellation/            # Walker constellation generator
│   │   ├── constellation.go      # Orbital mechanics, coverage calculation
│   │   ├── walker.go             # Walker delta-pattern T/P/F generator
│   │   └── *_test.go
│   ├── mesh/                     # Mesh network graph
│   │   ├── mesh.go               # Dijkstra SPF, link management
│   │   └── *_test.go
│   ├── oct/                      # Optical Crosslink Terminal
│   │   ├── oct.go                # OCT specs, link budget
│   │   └── *_test.go
│   ├── protocol/                 # Communication protocols
│   │   ├── ldpc.go               # LDPC error correction codes
│   │   ├── linkmodel.go          # Link budget model (FSPL, pointing loss)
│   │   ├── octframe.go          # OCT frame structure (HDLC-like)
│   │   └── *_test.go
│   ├── routing/                   # Network routing
│   │   ├── ospf.go               # OSPF shortest-path-first
│   │   ├── qos.go                # QoS traffic classes
│   │   └── *_test.go
│   ├── scenario/                  # Scenario management
│   │   ├── scenario.go           # Config, runner, result types
│   │   ├── api.go                # REST API server
│   │   └── *_test.go
│   ├── sgp4/                     # SGP4 orbital propagator
│   │   ├── sgp4.go               # J2 perturbation, drag, mean motion
│   │   └── *_test.go
│   ├── simulation/               # Core simulation engine
│   │   ├── sdn.go                # SDN network model (3 layers)
│   │   ├── simulation.go         # Simulator, time stepping, state
│   │   ├── linkscheduler.go      # Dynamic link lifecycle (627 LOC)
│   │   ├── visibility.go         # Satellite visibility, link geometry
│   │   └── *_test.go
│   └── threat/                    # Threat models
│       ├── threat.go             # ASAT, jamming, cyber
│       ├── kessler.go            # Kessler syndrome debris
│       └── *_test.go
├── PAPER.md                      # Research paper / technical deep-dive
├── RESEARCH.md                   # SDN/SDA/MILNET intel sources
├── ROADMAP.md                    # Build roadmap and phases
├── go.mod
└── go.sum
```

---

## Cross-Platform Build

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o space-data-network-sim ./cmd/sim/

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -o space-data-network-sim ./cmd/sim/

# Windows
GOOS=windows GOARCH=amd64 go build -o space-data-network-sim.exe ./cmd/sim/
```

---

## FORGE-C2 Integration

FORGE-SDN is registered in the FORGE-C2 BMDS bridge as `space-data-network-sim`. The bridge runs it periodically (default 60s) and feeds track data into the C2 system.

### NORAD Integration

The binary is deployed to `/host-bin/space-data-network-sim` on the NORAD Forge server. Access via:

```bash
curl "https://norad.stsgym.com/api/public/sdn?scenario=peacetime"
```

### FORGE-C2 Bridge Configuration

```go
// In internal/bmdsbridge/bridge.go
{"space-data-network-sim", "space-data-network-sim",
 []string{"-scenario", "peacetime", "-output", "/dev/stdout"},
 60 * time.Second}
```

---

## Resilience Score

The simulator calculates a 0–100 resilience score based on:

| Factor | Weight | Description |
|--------|--------|-------------|
| Satellite survival | 30% | Percentage of satellites still active |
| Link availability | 25% | Active links vs initial links |
| Coverage | 25% | Global coverage percentage |
| Route validity | 20% | Valid OSPF routes vs total routes |

**Score interpretation:**
- 80–100: Network fully operational, minor degradation
- 60–79: Significant degradation, mission-capable with limitations
- 40–59: Severe degradation, critical functions at risk
- 0–39: Network failure, most capabilities lost

**Example scores:**
- Peacetime: 81.1 — baseline with no stress
- Golden Dome (2 ASAT + jamming + cyber): 85.3 — network reroutes around damage
- Nuclear (400 kT): 21.1 — catastrophic, only 25/791 sats survive

---

## Technical References

1. **USSF Space Data Network (SDN)** — DoD architecture for next-gen military SATCOM
2. **Golden Dome / SBI** — Space-Based Interceptor communications layer
3. **PWSA (Protected Tactical Satcom Architecture)** — Multi-vendor transport layer
4. **Walker Constellations** — Walker, J.G. "Satellite Constellations" (1984)
5. **SGP4** — Hoots & Roehrich, "Models for Propagation of NORAD Element Sets" (1988)
6. **OSPF** — RFC 2328, Open Shortest Path First routing protocol
7. **LDPC** — Gallager, "Low-Density Parity-Check Codes" (1962)
8. **Kessler Syndrome** — Kessler & Cour-Palais, "Collision Frequency of Artificial Satellites" (1978)

See `RESEARCH.md` for full source bibliography and `PAPER.md` for the technical deep-dive.

---

## Contributing

1. Fork the repo on IDM
2. Create a feature branch
3. Ensure all tests pass: `go test ./...`
4. Submit merge request

---

## Changelog

### v1.0.0 (2026-05-02)
- Initial release
- 791-satellite constellation (3-layer SDN architecture)
- SGP4 orbital propagation
- OSPF mesh routing with QoS
- OCT/LDPC protocol stack
- Dynamic link scheduling (6-state lifecycle)
- 4 threat models (ASAT, jamming, nuclear, cyber)
- Walker constellation generator
- Kessler syndrome debris modeling
- REST API server
- 3 built-in scenarios + custom support
- Cross-platform binaries (Linux, macOS, Windows)
- FORGE-C2 BMDS bridge integration