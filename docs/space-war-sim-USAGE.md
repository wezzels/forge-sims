# Space War Sim — Usage Guide

## Overview

Space War Sim is a high-fidelity orbital warfare simulation engine. It models satellite orbits, ASAT engagements, debris generation, and space domain awareness.

## Physics Packages

| Package | Description |
|---------|-------------|
| `orbital` | Keplerian mechanics, Hohmann transfers, orbital elements |
| `sgp4` | SGP4/SDP4 orbit propagation from TLE data |
| `propagator` | RK4 numerical integrator with J2/J3/J4 perturbations and drag |
| `atm` | NRLMSISE-00 atmospheric density model (solar flux, Kp) |
| `asat` | Monte Carlo ASAT engagement (DA-ASAT, Co-Orbital, Laser, EMP, Cyber) |
| `debris` | NASA Standard Breakup Model, Kessler syndrome cascade |
| `ssa` | Space Situational Awareness tracking and sensor management |
| `core` | Simulation engine, event bus, scenario loading |

## Command Line

```bash
# Basic run
space-war-sim --config configs/scenario.yaml

# With AAR (after-action review) export
space-war-sim --config configs/scenario.yaml --aar results.json

# Validate a scenario
space-war-sim --validate configs/scenario.yaml

# List available scenarios
space-war-sim --list-scenarios

# Initialize sample configs
space-war-sim --init

# Health check
space-war-sim --doctor

# Version info
space-war-sim --version
```

## Scenario Configuration

Scenarios are YAML files that define the simulation:

```yaml
name: "LEO Conflict Alpha"
description: "DA-ASAT engagement against hostile ISR satellite"
duration: 36000    # 10 hours in seconds
delta_t: 10.0      # 10 second time steps
time_scale: 100.0  # 100x real-time

environment:
  solar_flux_10: 150    # F10.7 in SFU
  kp_index: 3.0         # Geomagnetic Kp
  solar_wind_speed: 400  # km/s
  j2_enabled: true
  drag_enabled: true

entities:
  - id: 1
    name: "ISS"
    category: "satellite"
    side: "friendly"
    semi_major_axis: 6778  # km (408 km altitude)
    eccentricity: 0.001
    inclination: 51.6
    altitude: 408
    health: 1.0
    mass: 420000
    rcs: 50.0

  - id: 2
    name: "HOSTILE-ISR-1"
    category: "satellite"
    side: "hostile"
    semi_major_axis: 6871  # km (500 km altitude)
    eccentricity: 0.002
    inclination: 97.4
    altitude: 500
    health: 1.0
    mass: 800
    rcs: 3.0

msel:
  - time: 3600
    type: "asat_engage"
    target: 2
    data:
      weapon: "da-asat"
  - time: 7200
    type: "debris_assessment"
    data:
      altitude: 500
```

## Go API Usage

### SGP4 Orbit Propagation

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/sgp4"

// Parse TLE
line1 := "1 25544U 98067A   24001.50000000  .00016717  00000-0  30200-3 0  9993"
line2 := "2 25544  51.6412 200.2245 0006703  40.6465 319.5154 15.49560892434028"
tle, err := sgp4.ParseTLE(line1, line2)

// Or use well-known satellites
tle := sgp4.ISS()  // ISS, Hubble, GPS_IIF, GEO, HostileISR

// Propagate
prop, err := sgp4.NewSGP4(tle)
result, err := prop.Propagate(60) // 60 minutes from epoch

fmt.Printf("Position: (%.1f, %.1f, %.1f) km\n", result.X, result.Y, result.Z)
fmt.Printf("Velocity: (%.3f, %.3f, %.3f) km/s\n", result.XD, result.YD, result.ZD)

// Geodetic conversion
lat, lon, alt := sgp4.ToGeodetic(result)
fmt.Printf("Lat: %.2f° Lon: %.2f° Alt: %.1f km\n", lat, lon, alt)
```

### RK4 Numerical Propagation

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/propagator"

cfg := propagator.DefaultConfig()
cfg.J2Enabled = true
cfg.DragEnabled = true
cfg.BallisticCoeff = 100.0  // kg/m²
prop := propagator.NewRK4Propagator(cfg)

// Start from SGP4 state
state := propagator.StateVector{X: 6778, Y: 0, Z: 0, XD: 0, YD: 7.67, ZD: 0}

// Propagate 1 orbit
for i := 0; i < 556; i++ {
    state = prop.Propagate(state, 10.0)  // 10s steps
}

// Extract orbital elements
a, e, i, raan, argPer, nu := propagator.OrbitalElements(state)
lat, lon, alt := propagator.ToGeodetic(state)
```

### ASAT Engagement (Monte Carlo)

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/asat"

weapon := asat.DirectAscentASAT()

// Monte Carlo engagement with 1000 trials
result := asat.EngageMonteCarlo(weapon, 500, 800, 0, asat.MonteCarloConfig{
    Trials:        1000,
    RandomSeed:    42,     // 0 for random
    MinConfidence: 0.5,
})

fmt.Printf("Hit: %v Kill Prob: %.2f Damage: %.0f%% Debris: %d\n",
    result.Hit, result.KillProb, result.DamageFraction*100, result.DebrisGenerated)
```

### Atmospheric Density

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/atm"

// NRLMSISE-00 at ISS altitude during solar minimum
result := atm.NRLMSISE00(408, 70, 1)   // 408 km, F10.7=70, Kp=1
fmt.Printf("Density: %.2e kg/m³ Temp: %.0f K\n", result.Density, result.Temperature)

// During solar maximum + storm
result = atm.NRLMSISE00(408, 250, 8)   // F10.7=250, Kp=8

// Drag acceleration
accel := atm.DragAcceleration(408, 7.67, 100, 150, 3)  // km/s²

// Orbit decay rate
decay := atm.OrbitDecayRate(408, 100, 150, 3)  // km/day
```

### Debris Modeling (NASA Breakup)

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/debris"

// NASA Standard Breakup Model
result := debris.NASABreakupModel(1000, 50, 10, false)  // 1000kg target, 50kg impactor, 10km/s
fmt.Printf("Trackable: %d >1cm: %d >1mm: %d Catastrophic: %v\n",
    result.FragmentsGT10cm, result.FragmentsGT1cm, result.FragmentsGT1mm, result.IsCatastrophic)

// Add to debris field
df := debris.NewDebrisField()
df.AddFragmentation(500, 51.6, 1000, 50, 10, false, "asat", 0, 42)

// Statistics
stats := df.GetStatistics()
fmt.Printf("Objects: %d Tracked: %d Mass: %.1f kg\n",
    stats.TotalObjects, stats.TrackedObjects, stats.TotalMass)

// Kessler cascade projection
newDebris, _ := df.KesslerCascade(500, 10)  // 500 km, 10 years
fmt.Printf("Kessler cascade: %.0f new debris over 10 years\n", newDebris)
```

## ASAT Weapon Types

| Type | Altitude Range | Kill Prob | Flight Time | Notes |
|------|---------------|-----------|-------------|-------|
| DA-ASAT (SM-3) | 200–1000 km | 0.90 | 10 min | Kinetic kill vehicle |
| Co-Orbital | 400–36000 km | 0.75 | 2 hours | Phasing interceptor |
| Ground Laser | 200–1500 km | 0.65 | Instant | Speed of light |
| EMP | 100–2000 km | 0.55 | 5 min | Area effect, 50km CEP |
| Cyber | 200–50000 km | 0.40 | Instant | Low reliability |

## Key Constants

| Constant | Value | Unit |
|----------|-------|------|
| Earth radius (Re) | 6371.137 | km |
| Gravitational parameter (μ) | 398600.4418 | km³/s² |
| Earth rotation rate | 7.2921159e-5 | rad/s |
| ISS altitude | ~408 | km |
| ISS velocity | ~7.67 | km/s |
| ISS period | ~5560 | s |
| GEO altitude | 35786 | km |

## Test Coverage

63 tests across 7 packages:

- `asat`: 12 tests (Monte Carlo, weapon types, inventory, debris)
- `atm`: 11 tests (NRLMSISE-00, solar flux, Kp, drag, pressure)
- `core`: 9 integration tests (full engagement chain, cross-package consistency)
- `debris`: 10 tests (NASA breakup, Kessler cascade, density, decay)
- `orbital`: 6 tests (Kepler, Hohmann, J2, velocity, drag, conjunction)
- `propagator`: 6 tests (RK4, conservation, J2 precession, drag, geodetic)
- `sgp4`: 6 tests (ISS, LEO, GEO, batch, TLE, deep-space)
## Constellation Degradation

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/satellite"

// Create a Starlink-like constellation
cfg := satellite.StarlinkConfig() // 1584 sats, 550 km, 53°
cs := satellite.NewConstellationState(cfg, 1)

// ASAT attack on 5 satellites
weapon := asat.DirectAscentASAT()
targets := []uint64{cs.Satellites[0].ID, cs.Satellites[1].ID, cs.Satellites[2].ID, cs.Satellites[3].ID, cs.Satellites[4].ID}
events := cs.ASATAttack(weapon, targets, 42)

// Add debris from destroyed satellites
destroyed := []uint64{}
for _, sat := range cs.Satellites {
    if !sat.IsActive && sat.DestroyReason == "asat_hit" {
        destroyed = append(destroyed, sat.ID)
    }
}
df := debris.NewDebrisField()
cs.AddDebrisFromAttack(df, destroyed, 42)

// Full degradation simulation (30 days)
timeline := satellite.SimulateDegradation(cfg, weapon, 10, 86400*30, 42)
for _, state := range timeline.States {
    fmt.Printf("T+%.0fs: %d active, coverage=%.1f%%, debris=%d\n",
        state.Time, state.ActiveSats, state.Coverage*100, state.Debris)
}
```

### Predefined Constellations

| Constellation | Altitude | Planes | Sats/Plane | Total | Δv Budget |
|---------------|----------|--------|------------|-------|------------|
| Starlink      | 550 km   | 72     | 22         | 1584  | 0.35 km/s  |
| GPS            | 20200 km | 6      | 5          | 30    | 0.50 km/s  |
| Iridium        | 780 km   | 6      | 11         | 66    | 0.20 km/s  |

## Covariance Conjunction Analysis

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/ssa"

// ISS vs debris
iss := ssa.StateVector{X: 6778, Y: 0, Z: 0, XD: 0, YD: 7.67, ZD: 0}
debris := ssa.StateVector{X: 6778.1, Y: 0.5, Z: 0.2, XD: -0.01, YD: -7.67, ZD: 0.01}
covP := ssa.TypicalLEOCovariance()
covS := ssa.TypicalDebrisCovariance()

cfg := ssa.DefaultConjunctionConfig()
cfg.HardBodyRadius = 0.05 // 50m (ISS is big)

result := ssa.CovarianceConjunction(iss, debris, covP, covS, cfg)
fmt.Printf("Miss: %.3f km, Pc: %.2e, Severity: %s\n",
    result.MissDistance, result.Pc, result.Severity)
```

## Orbital Maneuvers & Collision Avoidance

```go
import "gitlab.com/crab-meat-repos/space-war-sim/internal/maneuver"

// Hohmann transfer
plan := maneuver.HohmannRaise(408, 800) // 408→800 km
fmt.Printf("Δv: %.3f km/s, Time: %.0f s\n", plan.DeltaV, plan.TransferTime)

// Inclination change
plan := maneuver.PlaneChange(408, 23.1) // 23.1° at 408 km
fmt.Printf("Δv: %.3f km/s\n", plan.DeltaV)

// Combined maneuver (cheaper than separate)
plan := maneuver.CombinedManeuver(408, 800, 23.1)

// Collision avoidance
cam := maneuver.CollisionAvoidance(conjunction, primaryState, 408, 0.1)
fmt.Printf("CAM: feasible=%v safe=%v Δv=%.4f km/s\n", cam.Feasible, cam.Safe, cam.Maneuver.DeltaV)

// Delta-v budget
budget := maneuver.NewDeltaVBudget(1.0, 0.1) // 1 km/s total, 0.1 reserve
fmt.Printf("Available: %.3f km/s\n", budget.Available())

// Fuel mass (Tsiolkovsky)
fuel := maneuver.FuelMass(0.1, 2.94, 8000) // 100 m/s, Isp 300s, 8000 kg dry
```

## Test Coverage

98 tests across 10 packages:

- `asat`: 12 tests (Monte Carlo, weapon types, inventory, debris)
- `atm`: 11 tests (NRLMSISE-00, solar flux, Kp, drag, pressure)
- `core`: 9 integration tests (full engagement chain, cross-package consistency)
- `debris`: 10 tests (NASA breakup, Kessler cascade, density, decay)
- `maneuver`: 11 tests (Hohmann, plane change, combined, CAM, delta-v, fuel mass)
- `orbital`: 6 tests (Kepler, Hohmann, J2, velocity, drag, conjunction)
- `propagator`: 6 tests (RK4, conservation, J2 precession, drag, geodetic)
- `satellite`: 10 tests (Starlink/GPS/Iridium, ASAT attack, debris cascade, coverage)
- `sgp4`: 6 tests (ISS, LEO, GEO, batch, TLE, deep-space)
- `ssa`: 11 tests (covariance conjunction, B-plane, severity, batch analysis)
