# FORGE-Sims High-Fidelity Physics Roadmap

**Date:** 2026-04-18
**Status:** Planning — priority-ordered upgrade path

---

## Tier Classification

| Tier | Description | Sims |
|------|-------------|------|
| **A** | Real physics already (energy-on-target, Tsiolkovsky, atmospheric model) | uewr, lrdr, tpy2, mrbm, tactical-net, wta, kill-assessment |
| **B** | Partial physics — has structure but simplified/shortcut core models | icbm, irbm, hgv, gmd, sm3, sm6, sbirs, stss, dsp, cobra-judy, slcm |
| **C** | Radar/interceptor wrapper — depends on bmd-sim-core (fixed) or needs own physics | aegis, gbr, patriot, thaad, thaad-er, hub |
| **D** | C2/protocol/state-machine — not physics sims by design | c2bmc, gfcb, ifxb, jrsc, link16, jreap, decoy, jamming |
| **E** | Environment models — need deeper atmospheric/solar physics | space-weather, atmospheric |

---

## Priority 0 — Threat Trajectory Physics (3 sims)

These are the most visible: they model the incoming threat. Stubs = fake war.

### bmd-sim-icbm
**Current:** Sine-wave altitude, linear speed, "simplified ballistic trajectory"
**Target:** Same physics as mrbm (Tsiolkovsky staging, atmospheric density, reentry heating, great-circle with Earth curvature)
**Work:**
1. Add `ICBMVariant` struct with 3-stage configs (Minuteman-III, DF-41, RS-28 Sarmat)
2. Tsiolkovsky burnout velocity with gravity/drag losses per stage
3. Trajectory: boost quadratic → sin-shaped ballistic arc with Earth curvature
4. Atmospheric density model (US Std 1976) for dynamic pressure
5. Reentry heating (Sutton-Graves), deceleration from drag
6. Speed of sound vs altitude for Mach computation
7. Flight path angle (steep → 0° apogee → steep descent)
8. MIRV separation model (deployment altitude, dispersion pattern)
9. Add `ReentryHeatingRate()`, `DecelerationAt()`, `DetectableIR(alt)`

**Estimated:** 400 LOC, 15 new tests

### bmd-sim-irbm
**Current:** Same simplified model as ICBM
**Target:** Same physics framework, IRBM-specific variants
**Work:**
1. Add `IRBMVariant` (Asha, Qiam, DF-26, Pershing-II)
2. 2-stage Tsiolkovsky with IRBM-appropriate ISP/thrust
3. Same trajectory/atmospheric/reentry model as ICBM/MRBM
4. TEL mobile launch with erection time
5. Terminal guidance model for Pershing-II (active radar homing in terminal phase)
6. Add `ReentryHeatingRate()`, `DetectableIR(alt)`

**Estimated:** 350 LOC, 12 new tests

### bmd-sim-hgv
**Current:** "Simplified" detection range (500km × √RCS), basic glide model
**Target:** Full hypersonic glide physics
**Work:**
1. Lift-to-drag ratio model (L/D varies with altitude/speed for waverider)
2. Altitude-dependent drag coefficient (transition from high-altitude low-drag to terminal high-drag)
3. Trajectory: boost → pull-up → glide (altitude-dependent L/D) → terminal dive
4. Real atmospheric density for each trajectory point
5. Mach computation with local speed of sound
6. Dynamic pressure, reentry heating at each point
7. Replace simplified detection range with proper radar equation (energy-on-target from bmd-sim-core)
8. Maneuver scheduling: lateral weave, pull-up/dive, terminal S-turn
9. Plasma blackout zone (below ~40km at Mach 15+, radio attenuation)
10. Add `GlideLiftDrag(alt, speed)`, `PlasmaBlackout(alt, mach)`, `ReentryHeatingRate()`

**Estimated:** 500 LOC, 18 new tests

---

## Priority 1 — Interceptor Physics (4 sims)

These model the defense. Simplified PK/engagement = unreliable C2 decisions.

### bmd-sim-gmd
**Current:** "Simplified" PN guidance (HomingGain × closingVelocity × 0.001), approximate miss distance
**Target:** Real proportional navigation with proper kinematics
**Work:**
1. 3-stage flyout with actual thrust curves (not constant acceleration)
2. PN guidance: N=3-5, proper line-of-sight rate computation
3. Miss distance from PN with noise (acceleration saturation, seeker errors)
4. EKV endgame: seeker acquisition range, divert capability budget
5. Hit-to-kill assessment: miss < lethal radius
6. Atmospheric effects on booster trajectory (gravity turn, max-Q)
7. Add `FlyoutTrajectory()`, `EKVEndgame()`, `MissDistance(noise)`, `AccelerationBudget(alt)`

**Estimated:** 400 LOC, 15 tests

### bmd-sim-sm3
**Current:** "Approximate" 50 m/s² acceleration, basic PK formula
**Target:** Full SM-3 flyout with staging
**Work:**
1. 3-stage flyout with thrust curves per variant (IA/IB/IIA)
2. LEAP kill vehicle: seeker acquisition, divert budget
3. Exo-atmospheric intercept geometry (closing velocity, aspect angle)
4. Engagement zone computation with altitude constraints
5. PK model based on miss distance vs lethal radius
6. Add `FlyoutTrajectory()`, `LEAPAcquisition()`, `EngagementGeometry()`, `MissDistance()`

**Estimated:** 350 LOC, 12 tests

### bmd-sim-sm6
**Current:** Basic CanEngage/KillProbability, no flyout
**Target:** Dual-mode (SAM/SSM) flyout with active radar homing
**Work:**
1. Rocket motor thrust curve, atmospheric flight profile
2. Active radar seeker model (detection range vs RCS)
3. Terminal guidance: proportional navigation with active homing
4. Two modes: anti-air (high PK vs aircraft/cruise) vs ballistic (lower PK)
5. Add `FlyoutTrajectory()`, `ActiveSeekerRange(rcs)`, `TerminalGuidance()`

**Estimated:** 250 LOC, 10 tests

### bmd-sim-thaad-er
**Current:** Static parameters, basic CanEngage/KillProbability
**Target:** Extended-range flyout with larger booster
**Work:**
1. 2-stage flyout (larger booster than standard THAAD)
2. Exo-atmospheric intercept capability (200km altitude)
3. KKVI kill vehicle with divert budget
4. Engagement zone with extended range
5. Add `FlyoutTrajectory()`, `KKVIEndgame()`, `EngagementZone()`

**Estimated:** 250 LOC, 10 tests

---

## Priority 2 — Sensor Fidelity (3 sims)

### bmd-sim-sbirs
**Current:** "Simplified" visibility (within coverage cone), no scan model
**Target:** Proper IR detection with scan mechanics
**Work:**
1. Replace simplified visibility with slant-range + atmospheric IR transmission
2. IR detection model: NEFD-based (noise equivalent flux density)
3. Scan pattern: scanning vs staring mode, revisit rate
4. Atmospheric IR absorption (CO2, H2O bands) vs altitude
5. Constellation handoff timing
6. Add `SlantRange(lat,lon,alt)`, `IRTransmission(zenith_angle)`, `DetectionSNR(irSig, range)`

**Estimated:** 300 LOC, 10 tests

### bmd-sim-stss
**Current:** "Approximate" coverage percentage, basic discrimination
**Target:** LEO tracking with proper orbit mechanics
**Work:**
1. Keplerian orbit propagation for each satellite
2. Ground track / footprint computation
3. Track handoff between satellites (timing, gap analysis)
4. Multi-spectral discrimination (MWIR + LWIR + visible)
5. Add `KeplerPropagate(t)`, `GroundFootprint()`, `HandoffTiming()`, `SpectralDiscrimination()`

**Estimated:** 300 LOC, 10 tests

### bmd-sim-dsp
**Current:** Simple false alarm rate, basic detection range
**Target:** Legacy early warning with proper scan model
**Work:**
1. Spin-scan mechanics (6 RPM, 10s revisit)
2. IR detection with atmospheric transmission
3. False alarm model (Pfa vs threshold, Swerling targets)
4. Correlation with SBIRS (cross-check logic)
5. Add `ScanRevisitRate()`, `DetectionProbability(snr, Pfa)`, `SBIRSCorrelation()`

**Estimated:** 200 LOC, 8 tests

---

## Priority 3 — Platform Fidelity (3 sims)

### bmd-sim-aegis
**Current:** Has fire control but no flyout, basic PK
**Target:** Full Aegis BMD engagement model
**Work:**
1. SPY-1 track quality model (track accuracy vs range)
2. SM-3/SM-6 flyout integration (call sm3/sm6 packages or embed)
3. Illumination timing (SPY-1 uplink for midcourse)
4. Terminal engagement: active seeker for SM-6, LEAP for SM-3
5. Salvo management (fire-watch-fire-assess)
6. Add `TrackQuality(range)`, `IlluminationSchedule()`, `SalvoManager()`

**Estimated:** 350 LOC, 10 tests

### bmd-sim-thaad
**Current:** Battery state machine, no flyout, basic AssignLauncherToTrack
**Target:** Full THAAD engagement
**Work:**
1. THAAD interceptor flyout (hit-to-kill, 200km max alt)
2. TPY-2 radar track quality for engagement
3. Fire-control timeline (track → engage → launch → assess)
4. SLS (shoot-look-shoot) doctrine
5. Add `InterceptorFlyout()`, `EngagementTimeline()`, `SLSDecision()`

**Estimated:** 250 LOC, 8 tests

### bmd-sim-patriot
**Current:** Battery struct, basic missile, CalculateDetectionRange only
**Target:** Full PAC-3 engagement with blast-frag warhead
**Work:**
1. PAC-3 MSE blast-frag lethal radius model (5m radius)
2. PAC-2 fragment spray pattern
3. MPQ-65 radar track quality
4. Engagement timeline (detect → classify → engage)
5. Add `BlastFragLethalRcs(range)`, `FragmentPattern()`, `EngagementTimeline()`

**Estimated:** 250 LOC, 8 tests

---

## Priority 4 — Environment Models (2 sims)

### bmd-sim-space-weather
**Current:** "Simplified TECU", basic Kp/solar flux
**Target:** Full space weather model
**Work:**
1. International Reference Ionosphere (IRI) model for electron density
2. Kp→Ap magnetic index conversion
3. TEC from IRI (not simplified)
4. Scintillation: S4 index from ROT (rate of TEC change)
5. Solar wind parameters (density, velocity, Bz)
6. Add `IRIProfile(lat,lon,alt,hour)`, `ApFromKp()`, `ScintillationS4()`, `SolarWindImpact()`

**Estimated:** 300 LOC, 8 tests

### bmd-sim-atmospheric
**Current:** ITU-R P.838 rain attenuation, basic refractivity
**Target:** Full tropospheric propagation model
**Work:**
1. ITU-R P.676 gaseous absorption (O2 + H2O)
2. ITU-R P.453 radio refractivity (N-units from T/P/H)
3. Ducting model (surface, elevated, evaporation)
4. ITU-R P.530 multipath fading
5. Cloud/fog attenuation (ITU-R P.840)
6. Add `GaseousAbsorption(freq, alt, humidity)`, `DuctingProfile()`, `MultipathFading()`, `CloudAttenuation()`

**Estimated:** 350 LOC, 10 tests

---

## Priority 5 — Minor Upgrades (5 sims)

| Sim | Current Issue | Fix |
|-----|---------------|-----|
| **slcm** | "Simplified cruise trajectory" | Add terrain-following flight profile, sea-skimming alt variation, waypoint navigation |
| **cobra-judy** | Basic Detect/TrackPrecision | Add S-band radar equation, collection scheduling, multi-pass tracking |
| **decoy** | Static RCS/IR | Add time-varying IR (cooling model exists), radar echo modulation, wake modeling |
| **gbr** | Basic CalculateDetectionRange | Add X-band radar equation from bmd-sim-core, track quality model, discrimination mode |
| **hub** | State machines only | Add sensor fusion weighting, engagement arbitration, doctrine-based weapon assignment |

---

## No-Action (Correct as-is)

These sims are C2/protocol/state-machine by design — they don't model physics:

| Sim | What They Actually Do | Why No Physics Needed |
|-----|---------------------|----------------------|
| c2bmc | Track correlation, discrimination, engagement planning | Decision engine, not physics |
| gfcb | Fire control authorization chain | Policy engine |
| ifxb | Data routing between BMDS elements | Network model |
| jrsc | Regional sensor feed aggregation | State machine |
| link16 | TDMA slot allocation, J-series messages | Protocol model |
| jreap | JREAP-C encode/decode, UDP/TCP transport | Transport protocol |
| jamming | JNR, burnthrough range (already physics-based) | ✅ Has real math |
| tactical-net | RF propagation, Link 16, JREAP, jamming | ✅ Full physics |
| wta | Weapon-target assignment optimization | ✅ Math/optimization |
| kill-assessment | Assessment sensors, debris analysis | ✅ Full model |

---

## Shared Physics Library

Create `bmd-sim-core/pkg/physics/` modules to avoid duplication:

| Module | Provides | Used By |
|--------|----------|----------|
| `trajectory.go` | Boost→ballistic arc, atmospheric density, reentry heating, Mach, FPA | icbm, irbm, mrbm, hgv, gmd, sm3, sm6 |
| `atmosphere.go` | US Std 1976 density, speed of sound, gaseous absorption | all sims |
| `radar.go` | Energy-on-target SNR, detection probability, search range | uewr, lrdr, tpy2, gbr, aegis, patriot, cobra-judy |
| `reentry.go` | Sutton-Graves heating, stagnation pressure, drag deceleration | icbm, irbm, mrbm, hgv, gmd, sm3 |
| `guidance.go` | PN guidance, miss distance, acceleration budget | gmd, sm3, sm6, thaad-er |
| `ir.go` | NEFD, atmospheric IR transmission, scan mechanics | sbirs, stss, dsp |

---

## Implementation Order

| Phase | Sims | Est LOC | Est Tests | Dependencies |
|-------|------|---------|-----------|-------------|
| **1** | icbm, irbm | 750 | 27 | mrbm (already done as template) |
| **2** | hgv | 500 | 18 | trajectory lib |
| **3** | gmd, sm3 | 750 | 27 | guidance lib |
| **4** | sm6, thaad-er | 500 | 20 | sm3/gmd flyout model |
| **5** | sbirs, stss, dsp | 800 | 28 | ir lib |
| **6** | aegis, thaad, patriot | 850 | 26 | radar lib, interceptor flyout |
| **7** | space-weather, atmospheric | 650 | 18 | atmosphere lib |
| **8** | slcm, cobra-judy, decoy, gbr, hub | 600 | 20 | minor upgrades |
| **Total** | **20 sims** | **5,400** | **184** | |

**Already done (A-tier):** uewr, lrdr, tpy2, mrbm, tactical-net, wta, kill-assessment = 7 sims ✅
**No-action (by design):** 10 sims ✅
**Remaining:** 20 sims need physics upgrades

---

## Quality Gates

Each sim upgrade must pass before moving to next:

1. ✅ All existing tests still pass
2. ✅ New physics tests with known-answer verification (e.g., ICBM apogee ~1200km at 7km/s burnout)
3. ✅ No "simplified", "approximate", or "stub" in non-comment code
4. ✅ `-json` output includes physical quantities (Mach, dynamic pressure, heating rate, etc.)
5. ✅ Trajectory altitude never exceeds physically possible for the variant
6. ✅ Reentry heating follows Sutton-Graves (W/m²)
7. ✅ Dynamic pressure consistent with density × v²
8. ✅ Flight time within 10% of analytical estimate
9. ✅ Cross-sim consistency (ICBM at same range as NORAD scenario produces same trajectory)