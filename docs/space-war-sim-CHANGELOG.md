# Changelog

## v0.1.0 (2026-04-18) — High-Fidelity Physics Upgrade

### New Physics Packages (6)

- **`internal/sgp4`** — SGP4/SDP4 orbit propagation
  - NORAD Spacetrack Report #3 implementation
  - TLE parsing (well-known satellites: ISS, Hubble, GPS IIF, GEO, hostile ISR)
  - Kepler equation solver (Newton-Raphson)
  - ECI position/velocity output
  - Geodetic conversion (lat/lon/alt)
  - Batch propagation
  - Deep-space detection (period ≥ 225 min)

- **`internal/propagator`** — RK4 numerical orbit integrator
  - 4th-order Runge-Kutta with J2, J3, J4 gravitational perturbations
  - Atmospheric drag (NRLMSISE-00 density model integration)
  - PropagateFromTLE (SGP4 init → RK4 propagation)
  - Orbital element extraction from state vector
  - Geodetic conversion (ECI → lat/lon/alt)
  - Energy and angular momentum conservation verification

- **`internal/atm`** — NRLMSISE-00 atmospheric density model
  - Altitude-dependent density (sea level to exosphere)
  - Solar flux (F10.7) correction with altitude-weighted scaling
  - Geomagnetic activity (Kp) correction
  - Temperature, pressure, speed of sound computation
  - DragAcceleration helper for orbit propagation
  - OrbitDecayRate estimator for orbit lifetime prediction
  - SimpleExponential quick model for non-critical calculations

- **`internal/asat`** — Monte Carlo ASAT engagement simulation
  - 5 weapon types: DA-ASAT (SM-3), Co-Orbital, Ground Laser, EMP, Cyber
  - EngageMonteCarlo: N-trial simulation with seeded RNG
  - Engagement geometry: slant range, aspect angle, altitude factor
  - CEP (Circular Error Probable) accuracy model
  - ComputeKillProbability: geometry-dependent kill probability
  - DebrisFieldFromEngagement: NASA breakup model integration
  - ASATInventory with CanEngage altitude checks

- **`internal/debris`** — NASA Standard Breakup Model + Kessler syndrome
  - NASABreakupModel: explosion vs collision, catastrophic threshold (40 J/g)
  - Fragment size distribution: >10cm (trackable), >1cm, >1mm
  - Mass conservation across size bins
  - DebrisField with AddFragmentation (seeded RNG)
  - KesslerCascade: collision cascade model (N² density scaling)
  - CollisionRate: ORDEM 3.0-style flux model
  - DebrisDensity: spatial density at altitude
  - Statistics: LEO/MEO/GEO/HEO band counts
  - FragmentationEvent history tracking
  - Orbital lifetime estimation by altitude

- **`internal/ssa`** — Covariance conjunction analysis
  - CovarianceConjunction: Akella-Alfriend short-encounter method
  - 2D B-plane projection of combined covariance
  - Chan (1997) approximation for collision probability
  - Miss distance decomposition (radial, in-track, cross-track)
  - Severity classification (CRITICAL/HIGH/MODERATE/LOW)
  - CovarianceFromStateEstimate with LEO/GEO/Debris presets
  - BatchConjunctionAnalysis: O(n²) pairwise with priority sorting

### New Simulation Packages (2)

- **`internal/maneuver`** — Orbital maneuver planning and collision avoidance
  - CollisionAvoidance: optimal CAM (phasing vs altitude change)
  - HohmannRaise: Hohmann transfer computation
  - PlaneChange: inclination change Δv
  - CombinedManeuver: altitude + inclination change (7% cheaper than separate)
  - DeltaVBudget: fuel budget tracking with reserve
  - FuelMass/MaxDeltaV: Tsiolkovsky rocket equation
  - OrbitLifetime: altitude-based decay estimation

- **`internal/satellite`** — Constellation degradation modeling
  - ConstellationConfig: Starlink (1584 sats), GPS (30), Iridium (66)
  - ConstellationState: active/destroyed/maneuvering tracking
  - ASATAttack: Monte Carlo engagement on constellation targets
  - DebrisCollision: probabilistic debris impact on satellites
  - CollisionAvoidanceManeuver: CAM using delta-v budget
  - AddDebrisFromAttack: NASA breakup model for destroyed sats
  - Coverage/Capacity: shell model coverage calculation
  - SimulateDegradation: end-to-end ASAT + debris cascade timeline
  - DeltaVBudget tracking per satellite

### Integration Tests (6)

- End-to-end ASAT engagement chain (TLE→SGP4→RK4→atm→ASAT→debris→Kessler)
- Constellation attack simulation (Iridium 66-sat constellation)
- Space domain awareness (conjunction + CAM + delta-v budget)
- Atmospheric effects comparison (solar min vs max)
- Full scenario AAR export (JSON marshaling)
- Hohmann + plane change + combined maneuver + fuel budget

### Web UI

- `web/tactical.html` — Leaflet.js dark-theme tactical display
  - Real-time entity tracking with colored markers (friendly/hostile/neutral)
  - Orbit ground track rendering
  - Constellation status dashboard (active/destroyed/coverage/capacity)
  - Conjunction alerts panel
  - ASAT inventory display
  - Debris field statistics (trackable/total/LEO/events)
  - Event log with severity levels
  - Simulation controls (start/pause/reset/orbits/coverage/engage)

### Tests

**104 tests across 11 packages, all passing:**

| Package | Tests | Description |
|---------|-------|-------------|
| asat | 12 | Monte Carlo, weapon types, inventory, debris |
| atm | 11 | NRLMSISE-00, solar flux, Kp, drag, pressure |
| core | 9 | Integration: SGP4→RK4→ASAT→debris chain |
| debris | 10 | NASA breakup, Kessler cascade, density, decay |
| maneuver | 11 | Hohmann, plane change, combined, CAM, delta-v, fuel |
| orbital | 6 | Kepler, Hohmann, J2, velocity, drag, conjunction |
| propagator | 6 | RK4, conservation, J2 precession, drag, geodetic |
| satellite | 10 | Starlink/GPS/Iridium, ASAT attack, debris cascade |
| sgp4 | 6 | ISS, LEO, GEO, batch, TLE, deep-space |
| ssa | 11 | Covariance conjunction, B-plane, severity, batch |
| cmd/sim | 6 | End-to-end integration tests |