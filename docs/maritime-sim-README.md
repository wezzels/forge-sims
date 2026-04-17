# Maritime-Sim

High-fidelity maritime warfare simulation engine in Go.

## Quick Start

```bash
go build -o maritime-sim ./cmd/sim/

# Run default ASW scenario
./maritime-sim --config configs/scenario.yaml

# Run with web tactical display
./maritime-sim --config configs/engagement.yaml --web :8080

# Run with AAR export
./maritime-sim --config configs/engagement.yaml --aar output.json
```

Open `http://localhost:8080` for the tactical map.

## Architecture

```
cmd/sim/                    — Full simulation CLI with AI, sensors, weapons, web UI
internal/
  ai/                       — Threat assessment, CPA, tactical controller
  comms/c2/                 — ROE engine, formations, shoot-look-shoot
  comms/rf/                 — RF propagation (free-space, 2-ray, ITU-R P.1546)
  core/                     — Engine, config, WGS84 kinematics, event bus
  env/ocean/                — Bathymetry, sound channels, thermocline
  ew/ecm/                   — Chaff, jamming, decoy models
  ew/esm/                   — ESM bearing estimation, emitter ID
  fed/dis/                  — DIS IEEE 1278 PDU bridge for HLA federation
  hydro/                    — Pierson-Moskowitz spectrum, sea state
  platform/submarine/       — SSN/SSK: depth, cavitation, battery, sprint-and-drift
  platform/surface/         — DDG/FFG/CVN
  scenario/aar/             — After-action review, replay, scoring, JSON export
  scenario/msel/            — MSEL timeline (loaded from YAML config)
  sensor/acoustic/          — Sonar equation, Mackenzie, CZ, duct, bottom bounce, SOFAR
  sensor/radar/             — Radar equation, horizon, Swerling, detection probability
  sensor/sensormgr/         — Sensor scheduling, towed array, sonobuoy fields
  sensor/trackmgr/          — Track management, classification, TMA
  weapon/ashm/              — AShM flyout (Harpoon, YJ-83, Exocet) with CIWS/chaff
  weapon/torpedo/            — Mk48 ADCAP runout/search/homing with proportional nav
  weapon/damage/             — Compartment flooding, fire, magazine explosion, capsize
  weapon/engagement/        — Engagement manager, ROE, launch, hit/miss
  webui/                    — Leaflet.js tactical display with SSE real-time push
pkg/
  geom/                     — Haversine, bearing, WGS84 position propagation
  types/                    — Core types (EntityState with Name/Side/Category)
  units/                    — Unit conversions (knots, NM, dB)
tests/                      — Validation tests against published reference values
```

## Scenarios

| File | Description |
|------|-------------|
| `configs/scenario.yaml` | South China Sea ASW patrol (3 entities, 2hr) |
| `configs/engagement.yaml` | Close-range ASW engagement (sub at 11km, full kill chain) |

## Models

### Physics (Validated Against References)

| Model | Reference | Error |
|-------|-----------|-------|
| WGS84 position propagation | 1° lat = 111320m | <1m |
| Pierson-Moskowitz peak frequency | ω₀ = g/U₁₀ | matches |
| Mackenzie sound speed | 5 reference conditions | within tolerance |
| Sonar equation at TL=100dB | SE ≈ 0 dB | ±0.1 dB |
| Radar horizon (4/3 Earth) | 20m/10m → 31.4km | 31.5km |
| Great circle distance | London-Paris ≈ 340km | <10km |
| Unit conversions | 30 kts = 15.433 m/s | exact |

### Acoustic Propagation
- **Direct path**: Spherical spreading + absorption
- **Convergence zone**: 55-65km spacing, ~57dB gain over direct at CZ
- **Surface duct**: Cylindrical spreading in mixed layer (100x MLD range)
- **Bottom bounce**: Shallow water reflection path
- **Deep sound channel (SOFAR)**: Very long range in deep water
- **Sea state noise**: Wenz ambient noise model integrated into sonar equation

### Submarines
- **Depth control**: Smooth approach to ordered depth, max depth rate
- **Cavitation**: Speed threshold increases with depth (√(1 + D/10.3))
- **Radiated noise**: Below/above cavitation, 110 dB quiet, +20 dB penalty
- **Sprint-and-drift**: Tactical pattern controller
- **Battery management**: Power ∝ speed³, snorkel recharge at periscope depth
- **Presets**: Virginia SSN, Kilo SSK

### Weapons
- **AShM**: Boost → cruise → terminal, CIWS and chaff countermeasures
- **Torpedo**: Runout → search (snake/circle/spiral) → homing (proportional nav)
- **Damage**: Compartment flooding, fire spread, magazine explosion, capsize check

### C2
- **ROE**: Weapons Free / Tight / Hold with self-defense override
- **Formations**: ASW screen, column, line abreast, wedge
- **Shoot-look-shoot**: Optimal engagement for Pk ≥ 90%

### AI
- **Threat assessment**: Range, CPA, closing rate, platform type, speed, classification
- **Decision tree**: attack/pursue/track/evade/patrol based on threat score + ROE
- **Tactical controller**: Per-platform maneuvering (DDG 3°/s, SSK 2°/s)
- **Auto-classification**: Tracks classified by entity side (HO/FR/NE)

### Sensors
- **Active sonar**: Scheduled pinging, SQS-53C
- **Passive sonar/towed array**: Bearing-only tracking
- **Radar**: Surface target detection with horizon calculation
- **ESM**: Bearing to radar emitters
- **Sonobuoy fields**: DIFAR/LOFAR/DICASS with 4-hour lifetime
- **Search patterns**: Creeping, expanding box

## Test Status

20 packages, 100+ tests — all green.

```bash
go test ./...
```

## Configuration

```yaml
name: "South China Sea ASW Scenario"
duration: 7200
delta_t: 0.1
time_scale: 1.0

environment:
  water_temp: 28
  salinity: 34.5
  sea_state: 3
  wind_speed: 8

entities:
  - id: 1
    name: "USS Halsey (DDG-97)"
    category: "surface"
    lat: 18.5
    lon: 114.0
    heading: 180
    speed: 6.17
    threat_level: "friendly"
```

## Federation

DIS (IEEE 1278) bridge for HLA federation interoperability:
- Entity State PDUs with ECEF positioning
- Fire/Detonation PDUs
- Configurable exercise/site/application IDs

## License

Internal use — crab-meat-repos