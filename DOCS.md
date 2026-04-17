# FORGE-Sims Documentation

Complete algorithm and API reference for all 30 BMDS simulators.

---

## Table of Contents

1. [Sensor Simulators](#sensor-simulators)
   - [SBIRS](#sbirs-satellite-constellation)
   - [STSS](#stss-space-tracking--surveillance-system)
   - [DSP](#dsp-defense-support-program)
   - [UEWR](#uewr-upgraded-early-warning-radar)
   - [LRDR](#lrdr-long-range-discrimination-radar)
   - [Cobra Judy](#cobra-judy-anspq-11-radar)
   - [TPY-2](#tpy-2-antpy-2-radar)
   - [GBR](#gbr-ground-based-radar)
2. [Interceptor Simulators](#interceptor-simulators)
   - [GMD](#gmd-ground-based-midcourse-defense)
   - [SM-3](#sm-3-standard-missile-3)
   - [SM-6](#sm-6-standard-missile-6)
   - [THAAD](#thaad-terminal-high-altitude-area-defense)
   - [THAAD-ER](#thaad-er-extended-range)
   - [Patriot](#patriot-pac-2pac-3)
   - [Aegis BMD](#aegis-bmd)
3. [Threat Simulators](#threat-simulators)
   - [ICBM](#icbm-intercontinental-ballistic-missile)
   - [IRBM](#irbm-intermediate-range-ballistic-missile)
   - [HGV](#hgv-hypersonic-glide-vehicle)
   - [SLCM](#slcm-submarine-launched-cruise-missile)
4. [Countermeasure Simulators](#countermeasure-simulators)
   - [Decoy](#decoypenetration-aid)
   - [Jamming](#ewjamming)
5. [C2 / Communications Simulators](#c2--communications-simulators)
   - [C2BMC](#c2bmc-command--control-battle-management)
   - [Hub](#c2bmc-integration-hub)
   - [GFCB](#gfcb-global-fire-control--battle)
   - [IFXB](#ifxb-integrated-fire-control-exchange)
   - [JRSC](#jrsc-joint-regional-security-center)
   - [Link 16](#link-16-jtids)
   - [JREAP](#jreap-mil-std-3011)
6. [Environmental Simulators](#environmental-simulators)
   - [Space Weather](#space-weather)
   - [Atmospheric](#atmospheric-propagation)

---

## Sensor Simulators

### SBIRS Satellite Constellation

**Package:** `github.com/forge-c2/bmd-sim-sbirs`
**Binary:** `bmd-sim-sbirs`

#### Purpose
Simulates the Space-Based Infrared System with 4 GEO and 2 HEO satellites for boost-phase missile detection and tracking.

#### Key Algorithms

- **Visibility Calculation**: Computes line-of-sight between each satellite and a target position on Earth using 3D geometry. Uses `closestPointToOrigin()` to find the closest approach of the satellite-to-target vector to Earth's center — if distance > Earth radius, the satellite can see the target.

- **Satellite Handoff**: When a track moves from one satellite's FOV to another, `PerformHandoff()` transfers the track maintaining continuity. Calculates which satellite has the best viewing angle.

- **Coverage Percentage**: `CoveragePercent()` sums the spherical area coverage of each satellite's FOV cone, accounting for overlap using solid angle calculations.

- **Constellation Health**: `UpdateHealth()` monitors satellite health status and recalculates effective coverage when satellites go offline.

#### API Reference

```go
// Create constellation (4 GEO + 2 HEO by default)
c := constellation.NewConstellation()

// Get satellite subsets
geos := c.GEOSatellites()   // 4 GEO satellites
heos := c.HEOSatellites()   // 2 HEO satellites

// Calculate visibility for a target position
results := c.CalculateVisibility(targetPos models.Vector3)
// Returns []VisibilityResult with satellite ID, visible bool, elevation, slant range

// Hand off a track between satellites
handoff, err := c.PerformHandoff(trackNum, fromSatID, targetPos, time.Now())

// Coverage metrics
pct := c.CoveragePercent()       // 0-100% global coverage
healthy := c.UpdateHealth()       // Update health status

// Sensor detection
s := sensor.NewGEO(id, name, position)
// SBIRSSensor has: Altitude, IRWavelengthMin/Max, ScanRate, NEFD, FieldOfView
```

#### CLI Flags
```
-v              Verbose: constellation details, satellite list
-i              Interactive: live satellite coverage
--json          Full JSON output with parameters
-site N         Site ID (default 1)
-scenario FILE  Scenario file
--duration DUR  Run for duration
--tick DUR      Simulation tick rate (default 1s)
--seed N        Random seed
```

#### Integration Example
```go
// Feed SBIRS detections into C2BMC
c := constellation.NewConstellation()
vis := c.CalculateVisibility(targetPos)
for _, v := range vis {
    if v.Visible {
        c2bmc.ProcessReport(c2bmc.SensorReport{
            SensorType: c2bmc.SensorSBIRS,
            SensorID:   v.SatelliteID,
            TrackID:    trackNum,
            Latitude:   targetLat,
            Longitude:  targetLon,
            Altitude:   targetAlt,
            Speed:      targetSpeed,
            Confidence: 0.95,
        })
    }
}
```

---

### STSS Space Tracking & Surveillance System

**Package:** `github.com/forge-c2/bmd-sim-stss`
**Binary:** `bmd-sim-stss`

#### Purpose
Simulates the LEO dual-band (IR + visible) tracking constellation for midcourse discrimination and birth-to-death tracking.

#### Key Algorithms

- **Midcourse Discrimination**: `Discriminate()` uses multi-parameter scoring to classify objects:
  - IR signature consistency (0.3 weight): RVs maintain high IR, decoys cool
  - Speed consistency (0.3 weight): RVs maintain ~7 km/s, decoys decelerate
  - RCS consistency (0.1 weight): RV ~0.1 m², decoys vary
  - Multi-source bonus (0.1): More sensors = higher confidence
  - Score ≥ 0.7 → RV, < 0.3 → DECOY

- **Birth-to-Death Tracking**: `BirthToDeathTrack` monitors an object from first detection through IR decay, tracking `IRDecayRate()` — real RVs cool slowly, balloons cool rapidly.

- **Constellation Coverage**: Two LEO satellites at 1350 km provide periodic coverage with handoff capability.

#### API Reference

```go
// Create constellation (2 LEO satellites)
c := constellation.NewSTSSConstellation()

// Coverage and visibility
pct := c.CoveragePercent()
visible := c.VisibleSatellites(targetLat, targetLon, targetAlt)
nextSat := c.HandoffSatellite(targetLat, targetLon, currentSatID)

// Discrimination
d := discrimination.NewMidcourseDiscriminator()
result := d.Discriminate(objectID, irSignature, speed, rcs, altitude)
// Returns: Classification, Confidence, Method

// Birth-to-death tracking
t := track.NewBirthToDeathTrack("threat-1")
t.AddObservation(irReading, speedReading)
decayRate := t.IRDecayRate()  // High = decoy, Low = RV
```

---

### DSP Defense Support Program

**Package:** `github.com/forge-c2/bmd-sim-dsp`
**Binary:** `bmd-sim-dsp`

#### Purpose
Legacy 3-GEO spinning scanner constellation for early warning. Precursor to SBIRS.

#### Key Algorithms

- **Spinning Scanner Detection**: `Detect()` models the rotating IR sensor — each scan covers the full Earth disk, with detection probability based on IR signature vs NEFD (Noise Equivalent Flux Density).

- **Detection Latency**: `DetectionLatency()` accounts for spin rate — objects may not be detected until the scanner sweeps past their position (up to 30s for 10 rpm spin).

- **SBIRS Correlation**: `CorrelateWithSBIRS()` cross-references DSP detections with SBIRS tracks using geographic proximity threshold.

#### API Reference

```go
s := sensor.NewDSPSensor()
// Fields: Altitude, SpinRate, IRWavelengthMin/Max, NEFD, FalseAlarmRate

detected, snr := s.Detect(irSignature, rangeToTarget)
latency := s.DetectionLatency(irSignature)
far := s.FalseAlarmRate()
maxRange := s.MaxDetectionRange(irSignature)

c := constellation.NewDSPConstellation()
gaps := c.CoverageGaps()      // Angular gaps in coverage
best := c.BestSatellite(lon)  // Best satellite for a longitude
active := c.ActiveCount()     // Number of active satellites

// Cross-reference with SBIRS
correlated := CorrelateWithSBIRS(dspLat, dspLon, sbirsLat, sbirsLon, threshold)
```

---

### UEWR Upgraded Early Warning Radar

**Package:** `github.com/forge-c2/bmd-sim-uewr`
**Binary:** `bmd-sim-uewr`

#### Purpose
UHF phased array radar with 4-site network (Beale, Fylingdales, Clear, Cape Cod) for early warning and track initialization.

#### Key Algorithms

- **Radar Range Equation**: `RadarRangeEquation(rcs, range)` computes received signal power using:
  ```
  Pr = (Pt * G² * λ² * σ) / ((4π)³ * R⁴ * L)
  ```
  Where Pt=transmit power, G=antenna gain, λ=wavelength, σ=RCS, R=range, L=system losses

- **Atmospheric Loss**: `AtmosphericLoss(elevation, range)` models tropospheric attenuation — lower elevation angles have longer atmospheric path lengths and more loss.

- **Site Visibility**: `CanSee()` checks if a site has line-of-sight to a target based on Earth curvature and radar horizon.

- **Track Correlation**: `CorrelateTrack()` determines which sites can see a track for multi-site triangulation.

#### API Reference

```go
r := radar.NewUEWRRadar("BEALE")
// Fields: Frequency, AvgPower, PeakPower, AntennaGain, TrackCapacity, SystemLoss

det := r.Detect(rcs, rangeToTarget)  // DetectionResult {Detected, SNR, Range, Azimuth}
maxRange := r.MaxDetectionRangeForRCS(rcs)
loss := r.AtmosphericLoss(elevationDeg, rangeToTarget)

// Site management
sm := sites.NewSiteManager()
active := sm.ActiveSites()
canSee := sm.CanSee(site, lat, lon, alt)
best := sm.BestSite(lat, lon, alt)
sites := sm.CorrelateTrack(lat, lon, alt)  // Which sites can see this track
```

---

### LRDR Long Range Discrimination Radar

**Package:** `github.com/forge-c2/bmd-sim-lrdr`
**Binary:** `bmd-sim-lrdr`

#### Purpose
S-band phased array radar for midcourse tracking and discrimination of RVs vs decoys at 5000+ km range.

#### Key Algorithms

- **Discrimination Scoring**: `Discriminate(rcs, irSig, speed)` scores objects on 3 axes:
  - RCS: 0.1 m² consistent with RV, >0.5 m² suggests decoy
  - IR: High IR = RV (reentry heating), Low IR = balloon
  - Speed: ≥4 km/s = RV orbital velocity, <2 km/s = debris/decoy
  - Combined score determines classification and confidence

- **Kill Assessment**: `KillAssessment(preRCS, postRCS)` detects intercept success — significant RCS drop indicates target destroyed.

- **Track Staleness**: `IsStale(maxAge)` flags tracks not updated within threshold for quality control.

#### API Reference

```go
r := radar.NewLRDR()
// Fields: Frequency, AvgPower, AntennaGain, TrackCapacity, SystemLoss

det := r.Detect(rcs, rangeToTarget)
maxRange := r.MaxDetectionRangeForRCS(rcs)
snr := r.radarEquation(rcs, rangeToTarget)  // Raw SNR calculation

// Discrimination
d := discrimination.NewDiscriminator()
result := d.Discriminate(rcs, irSig, speed)
// DiscriminationResult {Classification, Confidence, RCS, Reason}

killed := d.KillAssessment(preRCS, postRCS)

// Track management
t := track.NewLRDRTrack(num, lat, lon, alt, speed, rcs)
t.Update(lat, lon, alt, speed, rcs)
stale := t.IsStale(30 * time.Second)
```

---

### Cobra Judy (AN/SPQ-11) Radar

**Package:** `github.com/forge-c2/bmd-sim-cobra-judy`
**Binary:** `bmd-sim-cobra-judy`

#### Purpose
Ship-based S-band radar for missile test collection and tracking.

#### Key Algorithms

- **Collection Mission Scheduling**: `MissionSchedule` prioritizes missions by type (verification, tracking, signature collection) and manages time allocation.

- **Track Precision**: `TrackPrecision()` returns the angular measurement precision based on radar parameters and target range.

#### API Reference

```go
r := radar.NewCobraJudy()
// Fields: Frequency, AvgPower, AntennaGain, SystemLoss

detected, snr := r.Detect(rcs, rangeToTarget)
maxRange := r.MaxRangeForRCS(rcs)
precision := r.TrackPrecision()
r.SetPosition(lat, lon)  // Move the ship

// Mission scheduling
ms := collection.NewMissionSchedule()
ms.AddMission(collection.NewCollectionMission(id, mType, target, duration, priority))
active := ms.ActiveMissions()
totalTime := ms.TotalDuration()
```

---

### TPY-2 (AN/TPY-2) Radar

**Package:** `github.com/forge-c2/bmd-sim-tpy2`
**Binary:** `bmd-sim-tpy2`

#### Purpose
X-band transportable phased array radar for terminal/forward-based missile defense tracking.

#### Key Algorithms

- **Sector Search Management**: `SectorManager` divides the radar FOV into elevation/azimuth sectors and manages search patterns. `CalculateCoverage()` sums sector coverage with overlap detection.

- **Cued Search**: `CuedSector(az, el, width, height)` directs the radar to a specific sector based on external cueing (e.g., from SBIRS or UEWR).

- **Track Initiation**: `TrackInitiator.ProcessDetection()` implements a sequential detection logic — multiple detections on the same track number confirm the track.

#### API Reference

```go
// Radar state
rs := radar.RadarState{
    Mode:      radar.ModeSearch,
    Position:  models.NewVector3(x, y, z),
    Azimuth:   0,
    Elevation: 0.2,
}

// Detection simulator
sim := radar.NewDetectionSimulator()
rangeRV := sim.SearchRadarRange(0.1)    // Range for 0.1 m² target
rangeLarge := sim.SearchRadarRange(1.0)  // Range for 1.0 m² target

// Track initiation
initiator := track.NewTrackInitiator()
trackNum, confirmed := initiator.ProcessDetection(pos, vel, rcs, snr)

// Search sectors
sm := search.NewSectorManager()
coverage := sm.CalculateCoverage()
sector := sm.CuedSector(az, el, azWidth, elWidth)
nextSector := sm.GetNextSector()
```

---

### GBR Ground-Based Radar

**Package:** `github.com/forge-c2/bmd-sim-gbr`
**Binary:** `bmd-sim-gbr`

#### Purpose
X-band ground-based radar for midcourse acquisition, tracking, and discrimination.

#### Key Algorithms

- **Acquisition Scan**: `AcquisitionManager` manages scan patterns for initial target detection. `StartScan()` creates a raster or cued scan pattern.

- **Discrimination**: `Discriminator.Classify()` uses RCS, velocity, and altitude to classify targets. Different from LRDR discrimination — GBR has higher resolution at shorter range.

- **Scan Patterns**: `AcquisitionScanPattern()` generates raster/cued/spiral patterns for the phased array.

#### API Reference

```go
rs := radar.NewGBRRadarState(models.NewVector3(x, y, z))
range_ := rs.CalculateDetectionRange(rcs)
pattern := rs.AcquisitionScanPattern(targetAz, targetEl)

// Acquisition
am := acquisition.NewAcquisitionManager()
scan := am.StartScan(1, nil)
confirmed := am.AddDetection(pos, vel, rcs, snr)

// Discrimination
d := discrimination.NewDiscriminator()
result := d.Classify(pos, vel, rcs)
// DiscriminationResult {Classification, Confidence, RCS, Type}
```

---

## Interceptor Simulators

### GMD Ground-Based Midcourse Defense

**Package:** `github.com/forge-c2/bmd-sim-gmd`
**Binary:** `bmd-sim-gmd`

#### Purpose
3-stage Ground-Based Interceptor with EKV (Exoatmospheric Kill Vehicle) for midcourse interception of ICBMs.

#### Key Algorithms

- **Flyout Simulation**: `Flyout(targetRange, targetAlt)` models the 3-stage boost phase:
  - Stage 1: High thrust acceleration (60s burn)
  - Stage 2: Sustainer (60s burn)
  - Stage 3: Fine adjustment (45s CE-II / variable)
  - Returns: velocity, altitude, range at burnout, time-to-target

- **EKV Homing**: `HomeToIntercept(closingV, range, irSignature)` simulates proportional navigation:
  - Acquires target via IR seeker if SNR sufficient
  - Computes lateral divert maneuver requirement
  - Returns: Kill probability, divert needed, intercept feasible

- **Engagement Zone**: `EngagementZone()` returns max/min altitude and range boundaries where engagement is feasible.

- **Commit Point**: `CommitPoint()` calculates the latest point at which a GBI can be launched to intercept a given threat.

#### API Reference

```go
// Launch management
lm := launch.NewLaunchManager()
// Sites: SiteFortGreely (AK), SiteVandenberg (CA)

// Create interceptor
gbi := interceptor.NewGBI(interceptor.VariantCEII, "FTG")
result := gbi.Flyout(targetRange, targetAlt)
// FlyoutResult {FinalVelocity, FinalAltitude, FinalRange, TimeOfFlight, StagesUsed, Success}

maxRange, maxAlt, minAlt := gbi.EngagementZone()
commitRange, commitTime, feasible := gbi.CommitPoint(lat, lon, alt, speed)

// EKV guidance
ekv := guidance.NewEKVCEII()
acquired := ekv.AcquireTarget(rangeToTarget, targetIRSignature)
homing := ekv.HomeToIntercept(closingVelocity, rangeToTarget, irSignature)
// HomingResult {KillProb, DivertRequired, InterceptFeasible, TimeToIntercept}

releaseTime, releaseRange := ekv.ReleasePoint(speed, targetSpeed, range)
```

---

### SM-3 Standard Missile-3

**Package:** `github.com/forge-c2/bmd-sim-sm3`
**Binary:** `bmd-sim-sm3`

#### Purpose
Ship-based exo-atmospheric interceptor with LEAP (Lightweight Exo-Atmospheric Projectile) kill vehicle.

#### Key Algorithms

- **Flyout Model**: 3-stage solid rocket motor with variant-specific parameters:
  - Block IA: 700 km range, Mach 10.5
  - Block IB: 900 km range, Mach 11.5
  - Block IIA: 2500 km range, Mach 13 (extended for Aegis Ashore)

- **LEAP Homing**: `HomeToIntercept()` models IR seeker acquisition and kill vehicle divert. Lower divert capability than EKV but lighter and faster.

- **Variant Comparison**: `CompareVariants()` and `AllVariants()` provide side-by-side specifications.

#### API Reference

```go
i := interceptor.NewSM3(interceptor.BlockIIA)
// Fields: Variant, Range, Speed, PK, Stages, Stage1/2/3Burn

i.SetLaunchPosition(lat, lon)
result := i.Flyout(targetRange, targetAlt)
maxRange, maxAlt, minAlt := i.EngagementZone()
canEngage := i.CanEngage(range_, alt)

// LEAP kill vehicle
kv := guidance.NewLEAPIIA()
acquired := kv.AcquireTarget(range, irSig)
homing := kv.HomeToIntercept(closingV, range, irSig)

// Variant comparison
specs := variants.AllVariants()
radius := variants.EngagementRadius("IIA")
```

---

### SM-6 Standard Missile-6

**Package:** `github.com/forge-c2/bmd-sim-sm6`
**Binary:** `bmd-sim-sm6`

#### Purpose
Endo-atmospheric Mach 3.5 dual-role interceptor (BMD + anti-air/anti-surface) launched from Mk 41 VLS.

#### Key Algorithms

- **Engagement Feasibility**: `CanEngage(alt, speed, range)` checks:
  - Target altitude ≤ 33 km ceiling
  - Target within 370 km range
  - Closing geometry feasible

- **Kill Probability**: `KillProbability(speed, rcs, range)` models:
  - Blast-fragmentation warhead effectiveness
  - Proximity fuze engagement envelope
  - Range-dependent Pk degradation

- **VLS Launcher**: `Salvo(count)` fires multiple missiles, `ReloadCount()` tracks remaining inventory.

#### API Reference

```go
s := interceptor.NewSM6()
// Fields: MaxSpeed, MinAltitude, MaxAltitude, Range, RCS, WarheadWeight

canEngage := s.CanEngage(targetAlt, targetSpeed, rangeToTarget)
time := s.InterceptTime(rangeToTarget)
pk := s.KillProbability(targetSpeed, targetRCS, rangeToTarget)
endo := s.IsEndoAtmospheric()  // Always true for SM-6

// Launcher
l := launch.NewSM6Launcher()
// Fields: CellCount, LoadedSM6, MaxSalvoRate
fired := l.Fire()
hits := l.Salvo(count)
reloads := l.ReloadCount()
empty := l.IsEmpty()
```

---

### THAAD Terminal High Altitude Area Defense

**Package:** `github.com/forge-c2/bmd-sim-thaad`
**Binary:** `bmd-sim-thaad`

#### Purpose
Terminal phase interceptor with hit-to-kill technology for endo- to exo-atmospheric intercepts.

#### Key Algorithms

- **Radar Detection**: `CalculateDetectionRange(rcs)` uses the X-band AN/TPY-2 radar equation for X-band frequency (~10 GHz).

- **Battery Management**: `AssignLauncherToTrack()` allocates a launcher to a threat track, managing resources across multiple simultaneous engagements.

- **Hit-to-Kill**: THAAD uses kinetic energy intercept — no warhead. PK depends on divert capability and closing velocity.

#### API Reference

```go
// Radar
rs := radar.NewTHAADRadarState(position)
range_ := rs.CalculateDetectionRange(rcs)

// Battery
batt := battery.NewBattery(id, position)
batt.AddLauncher(launcherID, position)
launcher := batt.AssignLauncherToTrack(trackNum)

// Launcher
l := launcher.NewLauncher(id, position)
launched := l.LaunchInterceptor()
l.Reload()
// Fields: MissilesReady, MissileType

// Missile
m := launcher.NewTHAADMissile()
// Fields: MaxVelocity, MaxRange, PK
```

---

### THAAD-ER Extended Range

**Package:** `github.com/forge-c2/bmd-sim-thaad-er`
**Binary:** `bmd-sim-thaad-er`

#### Purpose
Extended range THAAD variant with 600 km range, 200 km ceiling, and enhanced KKV divert capability.

#### Key Algorithms

Same as THAAD but with:
- **Extended Engagement Envelope**: 600 km range (vs ~200 km), 200 km ceiling (vs ~150 km)
- **Enhanced KKV**: Higher thrust kill vehicle with 600 m/s divert capability
- **Exo-atmospheric**: `IsExoAtmospheric()` returns true — can engage targets above the atmosphere

#### API Reference

```go
t := interceptor.NewTHAADER()
// Fields: MaxSpeed, MinAltitude, MaxAltitude, Range, BoosterThrust, KKVThrust, RCS

canEngage := t.CanEngage(alt, speed, range)
time := t.InterceptTime(range)
pk := t.KillProbability(speed, rcs, range)
exo := t.IsExoAtmospheric()
divert := t.DivertCapability()  // 600 m/s
```

---

### Patriot PAC-2/PAC-3

**Package:** `github.com/forge-c2/bmd-sim-patriot`
**Binary:** `bmd-sim-patriot`

#### Purpose
Terminal defense with PAC-2 (blast-fragmentation) and PAC-3 (hit-to-kill) missile variants.

#### Key Algorithms

- **PAC-2 vs PAC-3**: Two distinct engagement models:
  - PAC-2: Proximity fuze + blast warhead, lower PK (0.4-0.6)
  - PAC-3: Hit-to-kill Ka-band seeker, higher PK (0.8-0.9)

- **Engagement Sector**: Defines azimuth/width/height for fire control.

#### API Reference

```go
// Radar
rs := radar.NewPatriotRadarState(position)
range_ := rs.CalculateDetectionRange(rcs)

// Battery
batt := battery.NewBattery(id, position)
batt.AddLauncher(launcherID, position)
total := batt.TotalMissiles()

// Missiles
pac3 := missile.NewPAC3()
pac2 := missile.NewPAC2()
// Fields: Type, PK, MaxRange, MaxSpeed, SeekerType

// Engagement sector
sector := missile.EngagementSector{CenterAzimuth: 0, Width: 1.0, Height: 0.5}
```

---

### Aegis BMD

**Package:** `github.com/forge-c2/bmd-sim-aegis`
**Binary:** `bmd-sim-aegis`

#### Purpose
Ship-based BMD system with SPY-1 phased array radar and SM-2/3/6 missile suite.

#### Key Algorithms

- **SPY-1 Radar**: `CalculateDetectionRange(rcs)` models the S-band phased array with multiple array faces for 360° coverage.

- **Naval Track Manager**: Manages track database with classification (Hostile, Friendly, Unknown) and engagement status.

- **Fire Control Solution**: `ComputeSolution()` calculates intercept geometry:
  - Launch azimuth and elevation
  - Predicted intercept point
  - Flight time to intercept
  - Probability of kill

- **Task Scheduling**: SPY-1 interleaves search and track tasks within its timeline.

#### API Reference

```go
// Radar
radar := aegis.NewSPY1State(position, heading)
range_ := radar.CalculateDetectionRange(rcs)
// Fields: Mode, ArrayFaces, Frequency, Power

// Track manager
ntm := aegis.NewNavalTrackManager(shipID)
ntm.AddTrack(trackNum, pos, vel, classification)
// Tracks have: TrackNumber, Position, Velocity, Classification, EngagementStatus

// Fire control
fcc := firecontrol.NewFireControlComputer()
sol := fcc.ComputeSolution(target, targetVel, launcher, weaponType)
// Solution: FlightTime, PkProbability, InterceptPoint, LaunchAzimuth
valid := fcc.ValidateSolution(sol, launcher)

// Weapons
sm2 := weapon.NewSM2()
sm3 := weapon.NewSM3()
sm6 := weapon.NewSM6()
// Each has: PK, MaxRange, MaxSpeed, Type
```

---

## Threat Simulators

### ICBM Intercontinental Ballistic Missile

**Package:** `github.com/forge-c2/bmd-sim-icbm`
**Binary:** `bmd-sim-icbm`

#### Purpose
3-stage 10,000 km range ICBM threat with MIRV and countermeasures.

#### Key Algorithms

- **Trajectory Calculation**: `Trajectory(numPoints)` computes the full flight path:
  1. **Boost phase** (0-300s): 3-stage acceleration to burnout velocity
  2. **Midcourse phase** (~20 min): Ballistic coast in vacuum, MIRV deployment
  3. **Reentry** (final 30s): Atmospheric deceleration from Mach 21

- **MIRV Deployment**: `DeployRVs()` disperses multiple reentry vehicles with configurable footprint. `DispersionPattern()` calculates the elliptical impact pattern.

- **Countermeasures**: Suite includes balloon decoys, chaff clouds, and jammers — each with specific discrimination difficulty scores.

#### API Reference

```go
icbm := threat.NewICBM()
icbm.SetLaunchTarget(launchLat, launchLon, targetLat, targetLon)

points := icbm.Trajectory(100)  // 100-point trajectory
// TrajectoryPoint: Time, Latitude, Longitude, Altitude, Speed, Phase

lat, lon := icbm.ImpactPoint()
flightTime := icbm.FlightTime()    // ~1800s (30 min)
apogee := icbm.Apogee()           // ~1200 km
reentryV := icbm.ReentrySpeed()   // ~7000 m/s (Mach 21)
rcs := icbm.DetectableRCS()       // 0.1 m²

// MIRV
ms := mirv.NewMIRVSystem(count, deployLat, deployLon, deployAlt)
ms.DeployRVs(targets)
pattern := ms.DispersionPattern(lat, lon, footprintKm)
active := ms.ActiveRVCount()
total := ms.TotalRVs()

// Countermeasures
cm := countermeasures.NewCountermeasureSuite()
cm.AddBalloon(count, irSig, rcs, deployAlt)
cm.AddChaff(deployAlt, duration)
cm.AddJammer(deployAlt)
```

---

### IRBM Intermediate Range Ballistic Missile

**Package:** `github.com/forge-c2/bmd-sim-irbm`
**Binary:** `bmd-sim-irbm`

#### Purpose
2-stage 4,000 km range IRBM with road-mobile TEL launch capability.

#### Key Algorithms

- **Trajectory**: Similar to ICBM but shorter range, 2 stages, ~20 min flight time, ~600 km apogee, Mach 16 reentry.

- **Mobile Launch (TEL)**: `SimulateMobileLaunch()` models:
  - TEL deployment from hidden position
  - Erection and fueling time
  - Launch preparation sequence
  - `IsVulnerable()` identifies window where TEL is exposed

- **Range Categories**: `RangeCategory()` classifies by range: SRBM (<1000 km), MRBM (1000-3000), IRBM (3000-5500), ICBM (>5500).

#### API Reference

```go
ir := threat.NewIRBM()
ir.SetLaunchTarget(launchLat, launchLon, targetLat, targetLon)

points := ir.Trajectory(100)
flightTime := ir.FlightTime()   // ~1200s (20 min)
apogee := ir.Apogee()           // ~600 km
reentryV := ir.ReentrySpeed()   // ~5500 m/s (Mach 16)
boostDuration := ir.BoostDuration()
cat := threat.RangeCategory(ir.Range / 1000)  // "IRBM"

// TEL
tel := threat.NewTEL(lat, lon)
tel.Deploy()                  // Expose TEL
tel.Hide()                    // Return to hidden
vulnerable := tel.IsVulnerable()
launchTime := threat.SimulateMobileLaunch(tel, prepTime)
```

---

### HGV Hypersonic Glide Vehicle

**Package:** `github.com/forge-c2/bmd-sim-hgv`
**Binary:** `bmd-sim-hgv`

#### Purpose
Mach 5-25 boost-glide hypersonic threat with maneuver scheduling.

#### Key Algorithms

- **Glide Trajectory**: `Trajectory(numPoints)` computes the boost-glide path:
  1. **Boost** (0-120s): Rocket acceleration to Mach 25
  2. **Pull-up**: Transition to glide altitude (30-80 km)
  3. **Glide** (~60 min): Hypersonic cruise with gradual deceleration
  4. **Terminal dive**: Final approach to target

- **Detection Assessment**: `EngagementSummary()` evaluates detectability against each BMDS sensor:
  - SBIRS: Boost phase IR detection (detectable during boost)
  - UEWR: Below radar horizon during glide (not detectable)
  - TPY-2: Short range for low-altitude HGV (limited)

- **Maneuver Scheduling**: `ManeuverScheduler` generates random maneuver events:
  - PULL-UP, DIVE, LATERAL, TERMINAL types
  - `ApplyManeuvers()` modifies trajectory with delta-V perturbations
  - `TotalDeltaV()` tracks cumulative fuel usage

#### API Reference

```go
hgv := glide.NewHGV()
hgv.SetLaunchTarget(launchLat, launchLon, targetLat, targetLon)

points := hgv.Trajectory(100)
// GlidePoint: Time, Latitude, Longitude, Altitude, Speed
flightTime := hgv.FlightTime()

// Detection
detectRange := hgv.DetectionRange(radarNEI, radarGain)
boostDet, glideDet := hgv.IsDetectableBySBIRS(sensorNEFD, threshold)

// Detection assessment per sensor
summary := detection.EngagementSummary(alt, speed, rcs, irSig)
// map[string]DetectionChallenge: SBIRS, UEWR, TPY-2
// Each: {Detectable, MaxRange, EngagementWindow}

// Maneuvers
ms := maneuver.NewManeuverScheduler(5)
ms.ApplyManeuvers(altitudes, interval)
totalDV := ms.TotalDeltaV()
```

---

### SLCM Submarine-Launched Cruise Missile

**Package:** `github.com/forge-c2/bmd-sim-slcm`
**Binary:** `bmd-sim-slcm`

#### Purpose
Stealthy sea-skimming cruise missile launched from submarine.

#### Key Algorithms

- **Sea-Skimming Flight**: `CruiseSegments(numPoints)` models terrain-following flight at 30m altitude, making radar detection extremely difficult.

- **Detection Difficulty**: `DetectionDifficulty()` rates how hard the SLCM is to detect — factoring in low RCS (0.01 m²), low altitude (below radar horizon for most sensors), and slow speed (Mach 0.8).

- **Submarine Launch**: `SubmarineLaunch` models the launch sequence — sub surfaces from depth, launches missile, submerges.

#### API Reference

```go
s := threat.NewSLCM()
s.SetLaunchTarget(launchLat, launchLon, targetLat, targetLon)

flightTime := s.FlightTime()
segments := s.CruiseSegments(100)  // CruisePoint: Time, Lat, Lon, Alt, Speed
difficulty := s.DetectionDifficulty()  // 0.0-1.0
seaSkimming := s.IsSeaSkimming()       // true

// Submarine
sub := threat.NewSubmarineLaunch(lat, lon, depth)
launched := sub.Launch()
surfaced := sub.SurfaceForLaunch()
// Fields: Depth, Speed, Position
```

---

## Countermeasure Simulators

### Decoy/Penetration Aid

**Package:** `github.com/forge-c2/bmd-sim-decoy`
**Binary:** `bmd-sim-decoy`

#### Purpose
Simulates balloon decoys, inflatable RVs, RCS enhancers, and jammers as midcourse penetration aids.

#### Key Algorithms

- **Thermal Decay**: `SimulateCooling(dt)` models IR signature decay over time:
  - Balloons cool rapidly (high CoolingRate) — reveals them as decoys
  - Real RVs maintain IR signature through reentry heating
  - Discrimination algorithms exploit this difference

- **Discrimination Difficulty**: `DiscriminationDifficulty()` scores how hard each decoy is to identify:
  - Balloon with low RCS/IR: 0.3-0.5 (easier to discriminate)
  - Inflatable RV replica: 0.7-0.9 (harder)
  - RCS enhancer: varies by configuration

- **RV Profile Matching**: `MatchesRVProfile()` checks if a decoy's RCS and IR match expected RV values.

#### API Reference

```go
ds := decoy.NewDecoySuite()

// Add decoys
ds.AddDecoy(decoy.NewBalloonDecoy(rcs, irSig))
ds.AddDecoy(decoy.NewInflatableDecoy())
ds.AddDecoy(decoy.NewRCSEnhancer())
ds.AddDecoy(decoy.NewJammerDecoy())

// Decoy types: BALLOON, INFLATABLE_RV, RCS_ENHANCER, JAMMER, CHAFF
// Decoy fields: Type, Name, RCS, IRSignature, Weight, DeployAltitude, CoolingRate, Active

// Operations
ds.SimulateAllCooling(dt)          // Cool all decoys by dt seconds
count := ds.ActiveCount()         // Still-active decoys
maxDiff := ds.MaxDifficulty()     // Hardest decoy difficulty

// Per-decoy
d.SimulateCooling(dt)
difficulty := d.DiscriminationDifficulty()
matches := d.MatchesRVProfile(rvRCS, rvIR)
active := d.IsActive()

// String representation
fmt.Println(d.Type)  // "BALLOON", "INFLATABLE_RV", etc.
```

---

### EW/Jamming

**Package:** `github.com/forge-c2/bmd-sim-jamming`
**Binary:** `bmd-sim-jamming`

#### Purpose
Simulates noise and deception jamming against BMDS radars, plus ECCM effectiveness.

#### Key Algorithms

- **JNR Calculation**: `JNR(radarGain, radarBandwidth)` computes Jammer-to-Noise Ratio:
  ```
  JNR = (Pj * Gj * Gr * λ²) / ((4π * Rj)² * Bj * k * T * Br)
  ```
  Where Pj=jammer power, Gj=jammer antenna gain, Rj=jammer range, Bj=jammer bandwidth, Br=radar bandwidth

- **Burnthrough Range**: `BurnthroughRange()` calculates the range at which radar signal power overcomes jamming — the minimum effective range.

- **Jamming Effect**: `JammingEffect(jnrDB)` classifies the impact:
  - JNR > 30 dB: DENIED (radar unusable)
  - JNR 10-30 dB: DEGRADED
  - JNR < 10 dB: MINIMAL

- **ECCM Effectiveness**: `ECCMEffectiveness()` rates countermeasures:
  - Frequency agility: 0.7 (reduces jamming by 70%)
  - Side-lobe cancellation: 0.5
  - Home-on-jam: 0.9 (uses jammer as beacon)

#### API Reference

```go
// Noise jammer (barrage)
j := jamming.NewNoiseJammer(power, freq, bandwidth, distance)

// Deception jammer (DRFM)
j := jamming.NewDeceptionJammer(power, freq, distance)

jnr := j.JNR(radarGain, radarBandwidth)
burnthrough := j.BurnthroughRange(radarPower, radarGain, targetRCS, marginDB)
effect := jamming.JammingEffect(jnr)

// ECCM
eff := jamming.ECCMEffectiveness(jamming.FreqAgility)
// Techniques: FreqAgility, SideLobeCancel, HomeOnJam, PulseCompression, CFAR
```

---

## C2 / Communications Simulators

### C2BMC Command & Control Battle Management

**Package:** `github.com/forge-c2/bmd-sim-c2bmc`
**Binary:** `bmd-sim-c2bmc`

#### Purpose
Multi-sensor fusion, track correlation, discrimination, and engagement planning for the BMDS.

#### Key Algorithms

- **Track Fusion**: `ProcessReport()` merges sensor reports into fused tracks:
  - New track: Creates entry with 0.5× confidence (single-source uncertainty)
  - Existing track: Updates position (weighted average), boosts confidence (+0.1)
  - Multi-source tracks accumulate higher confidence

- **Track Correlation**: `CorrelateTracks(thresholdDist)` merges nearby tracks from different sensors:
  - Collects merge pairs first (avoids nil pointer during iteration)
  - Distance metric: `sqrt(Δlat² + Δlon²)`
  - Merged tracks get +0.05 confidence bonus
  - Lower track ID survives

- **Discrimination**: `Discriminate()` classifies all tracks:
  - Thermal (RCS): 0.35 weight
  - Speed consistency: 0.35 weight
  - Multi-source: 0.1 bonus
  - Score ≥ 0.7 → RV, < 0.3 → DECOY

- **Engagement Planning**: `EngagementPlan()` prioritizes threats by speed (faster = higher priority) and confidence (≥0.4 threshold).

#### API Reference

```go
c := c2bmc.NewC2BMC()

// Process sensor reports
c.ProcessReport(c2bmc.SensorReport{
    SensorType: c2bmc.SensorSBIRS,
    SensorID:   "GEO-1",
    TrackID:    1,
    Latitude:   35, Longitude: 130,
    Altitude:   200000, Speed: 6000,
    Confidence: 0.95,
})

// Correlate and discriminate
merged := c.CorrelateTracks(1.0)  // Threshold distance in degrees
c.Discriminate()

// Query
count := c.TrackCount()
active := c.ActiveTrackCount(0.5)
threats := c.EngagementPlan()
// []FusedTrack sorted by speed, confidence ≥ 0.4
```

---

### C2BMC Integration Hub

**Package:** `github.com/forge-c2/bmd-sim-hub`
**Binary:** `bmd-sim-hub`

#### Purpose
Central hub for multi-sensor fusion, engagement coordination, data routing, and scenario orchestration.

#### Key Algorithms

- **Fusion Engine**: `FusionEngine.AddSensorTrack()` merges tracks with source-weighted averaging. `canAssociate()` determines if a new detection should merge with an existing track. `updateFusedTrack()` recomputes position from all sources.

- **Data Routing**: `Router` manages message routing between systems with priority-based queuing (High, Critical).

- **Engagement Coordination**: `Coordinator.CreateEngagement()` allocates a weapon to a threat and computes time-to-intercept and probability of kill.

- **Scenario Orchestration**: `Orchestrator.CreateDemoScenario()` generates scripted event sequences for testing.

#### API Reference

```go
// Fusion
fe := fusion.NewFusionEngine()
ft := fe.AddSensorTrack(siteID, sensorType, trackNum, pos, vel)
fe.Maint()           // Clean stale tracks
tracks := fe.GetAllTracks()
track := fe.GetTrack(id)

// Routing
router := routing.NewRouter()
router.AddRoute(from, to, msgType, name, priority)
count := router.GetRouteCount()

// Engagement
coord := engagement.NewCoordinator()
eng := coord.CreateEngagement(threatID, weaponID, weaponType, threatPos, shooterPos)
// Engagement: ID, Phase, TimeToIntercept, ProbOfKill

// Orchestration
orch := orchestration.NewOrchestrator()
scenario := orch.CreateDemoScenario()
// Scenario: Name, Events[]
```

---

### GFCB Global Fire Control & Battle

**Package:** `github.com/forge-c2/bmd-sim-gfcb`
**Binary:** `bmd-sim-gfcb`

#### Purpose
Fire control workflow: request → authorize → fire → assess engagement cycle.

#### Key Algorithms

- **Engagement Workflow**: 4-step process:
  1. `RequestEngagement()` — validates request, assigns ID
  2. `AuthorizeEngagement()` — human-in-the-loop approval
  3. `FireEngagement()` — executes launch
  4. `AssessEngagement(killed)` — post-intercept assessment

- **Prioritization**: `Prioritize()` sorts pending engagements by threat priority and weapon availability.

#### API Reference

```go
g := gfcb.NewGFCB("GFCB-1", "PACIFIC")
// Fields: NodeID, Region

req := gfcb.EngagementRequest{ThreatID: "ICBM-1", WeaponID: "GBI-1", Priority: 1}
accepted := g.RequestEngagement(req)
authorized := g.AuthorizeEngagement(reqID)
fired := g.FireEngagement(reqID)
status := g.AssessEngagement(reqID, killed)

pending := g.PendingCount()
total := g.TotalEngagements()
g.Prioritize()
```

---

### IFXB Integrated Fire Control Exchange

**Package:** `github.com/forge-c2/bmd-sim-ifxb`
**Binary:** `bmd-sim-ifxb`

#### Purpose
Engage-on-remote capability — enables a sensor to direct a shooter without direct communication.

#### Key Algorithms

- **Engage-on-Remote**: `EngageOnRemote(sensorID, shooterID, trackID)` computes whether a sensor-shooter pair can coordinate:
  - Checks both are active assets
  - Validates data route exists between them
  - Calculates latency and engagement feasibility

- **Data Routing**: Routes track data from sensors to shooters through the IFXB mesh.

#### API Reference

```go
i := ifxb.NewIFXB("IFXB-1")
// Fields: NodeID, Active, Assets, DataRoutes

i.AddAsset(ifxb.Asset{ID: "TPY2-1", Type: "SENSOR", Active: true})
i.AddRoute(ifxb.DataRoute{From: "TPY2-1", To: "THAAD-1", Priority: 1})

assets := i.ActiveAssets()
routes := i.ActiveRoutes()
latency := i.AverageLatency()
route := i.FindRoute(from, to)

feasible, latency := i.EngageOnRemote(sensorID, shooterID, trackID)
```

---

### JRSC Joint Regional Security Center

**Package:** `github.com/forge-c2/bmd-sim-jrsc`
**Binary:** `bmd-sim-jrsc`

#### Purpose
Regional command center aggregating sensor feeds, managing alert levels, and fusing data.

#### API Reference

```go
j := jrsc.NewJRSC("INDO-PACIFIC")
// Fields: RegionID, Active, Sensors, AlertLevel

j.AddSensorFeed(jrsc.SensorFeed{
    SourceID: "SBIRS-1", FeedType: "IR", Active: true, Latency: 5,
})

feeds := j.ActiveFeeds()
j.SetAlertLevel(3)              // 1-5 (DEFCON-like)
latency := j.AverageLatency()
trackCount, confidence := j.FuseData()
```

---

### Link 16 (JTIDS)

**Package:** `github.com/forge-c2/bmd-sim-link16`
**Binary:** `bmd-sim-link16`

#### Purpose
TDMA tactical data link with 1535 time slots per 12-second frame, 238 kbps throughput.

#### Key Algorithms

- **TDMA Scheduling**: Time slots divided among participants. `Throughput()` calculates effective data rate based on active participants.

- **Range Calculation**: `RangeBetweenParticipants()` uses Haversine formula for great-circle distance.

- **Message Delivery**: `SendMessage()` simulates slot allocation and contention.

#### API Reference

```go
n := link16.NewNetwork("BMDS-NET-1")
// Fields: NetID, TimeSlots, Participants, Active

n.AddParticipant(link16.Participant{
    ID: "AEGIS-1", Type: "SHOOTER", JTIDSUnit: 1, Active: true,
})

active := n.ActiveParticipants()
throughput := n.Throughput()
success, slotUsed := n.SendMessage(link16.Message{Type: "TRACK", Data: data})
latency := n.Latency()
range_ := link16.RangeBetweenParticipants(lat1, lon1, lat2, lon2)
```

---

### JREAP (MIL-STD-3011)

**Package:** `github.com/forge-c2/bmd-sim-jreap`
**Binary:** `bmd-sim-jreap`

#### Purpose
JREAP-A (serial), JREAP-B (dialup), JREAP-C (packet) transport protocol simulation with CRC-16 error detection.

#### Key Algorithms

- **JREAP-C Header**: 8-octet header with protocol version, message type, message length, and sequence number.

- **CRC-16**: Ethernet polynomial CRC for message integrity validation. `ValidateCRC()` checks received data against embedded CRC.

- **Transport Metrics**: `Latency()`, `Bandwidth()`, `MessagesPerSecond()` model protocol performance.

#### API Reference

```go
t := jreap.NewTransport(jreap.JREAPC, port)
// Fields: Protocol, Port, RemoteURL, Active

maxSize := t.MaxMessageSize()
latency := t.Latency()
bw := t.Bandwidth()
mps := t.MessagesPerSecond(avgMsgSize)

// Low-level encoding
header := jreap.EncodeHeader(jreap.Header{
    Version: 1, MessageType: 0x28, Length: 100, SeqNum: 1,
})
crc := jreap.CRC16(data)
valid := jreap.ValidateCRC(dataWithCRC)
```

---

## Environmental Simulators

### Space Weather

**Package:** `github.com/forge-c2/bmd-sim-space-weather`
**Binary:** `bmd-sim-space-weather`

#### Purpose
Models space weather effects (Kp index, solar flux, scintillation) on radar and IR sensor performance.

#### Key Algorithms

- **Radar Degradation**: `RadarDegradation()` models ionospheric scintillation effects:
  - Kp < 5: Minimal impact (0-2 dB)
  - Kp 5-7: Moderate (2-6 dB)
  - Kp > 7: Severe (6-15 dB)

- **IR Clutter**: `IRClutter()` models increased background IR from auroral activity during storms.

- **Tracking Error**: `TrackingError()` estimates angular error from ionospheric refraction.

#### API Reference

```go
c := spaceweather.NewQuietConditions()    // Kp=1, low flux
c := spaceweather.NewStormConditions()    // Kp=7, high flux
// Fields: KpIndex, SolarFlux, Density, Scintillation, Active

radarLoss := c.RadarDegradation()    // dB
irNoise := c.IRClutter()            // W/sr
trackErr := c.TrackingError()       // radians
rangeErr := c.RadarRangeError()     // meters
isStorm := c.IsStorm()             // bool
effect := c.EffectOnRadar()         // string description
```

---

### Atmospheric Propagation

**Package:** `github.com/forge-c2/bmd-sim-atmospheric`
**Binary:** `bmd-sim-atmospheric`

#### Purpose
Models rain attenuation, refraction, and turbulence effects on radar propagation.

#### Key Algorithms

- **Rain Attenuation**: `RainAttenuation(freqGHz)` uses the ITU-R P.838 model:
  - Attenuation increases with frequency² and rain rate
  - S-band (3 GHz): ~0.1 dB/km at 25 mm/hr
  - X-band (10 GHz): ~0.5 dB/km at 25 mm/hr

- **Radar Refraction**: `RadarRefraction(elevationDeg)` models atmospheric bending:
  - Low elevation (0-5°): Significant refraction (10-50 mrad)
  - High elevation (>30°): Negligible

- **Total Degradation**: `TotalRadarDegradation(freq, pathKm)` combines rain, refraction, and turbulence.

#### API Reference

```go
a := atmospheric.NewStandardAtmosphere()  // Standard conditions
a := atmospheric.NewSevereAtmosphere()    // Rain=25mm/hr, Wind=25m/s
// Fields: Refractivity, Humidity, WindSpeed, Turbulence, CloudCover, RainRate

rainLoss := a.RainAttenuation(freqGHz)    // dB/km
refraction := a.RadarRefraction(elevDeg)  // mrad
irLoss := a.IRAttenuation()              // dB
turbEffect := a.TurbulenceEffect()        // Cn² factor
wind := a.WindEffect()                    // string
totalLoss := a.TotalRadarDegradation(freqGHz, pathKm)  // dB
```

---

## Cross-Simulator Integration

**Package:** `github.com/forge-c2/bmd-sim-integration`

### Integration Tests

7 cross-simulator tests validating end-to-end BMDS chains:

| Test | Chain | Validates |
|------|-------|-----------|
| TestFullEngagementChain | SBIRS → C2BMC → GMD | Complete kill chain |
| TestC2BMC_Discrimination | Multiple sensors → C2BMC | RV/decoy classification |
| TestUEWR_Detection | UEWR + range equation | Radar detection physics |
| TestLRDR_Discrimination | LRDR discrimination | Midcourse discrimination |
| TestICBM_Trajectory | ICBM 4-phase flight | Trajectory physics |
| TestSBIRS_Constellation | SBIRS coverage | 95% global coverage |
| TestMultiThreat | 3 RVs + 2 decoys | Multi-target handling |

### Usage
```go
import "github.com/forge-c2/bmd-sim-integration/pkg/integration"

// Run all integration tests
go test -v ./pkg/integration/
```

---

## CLI Usage (All Sims)

All 30 simulators share a common CLI interface:

```bash
# Default: print summary and exit
bmd-sim-sbirs

# Verbose: full simulation details
bmd-sim-sbirs -v

# Interactive: live updating metrics
bmd-sim-sbirs -i -tick 1s -duration 60s

# JSON: machine-readable output with full parameters
bmd-sim-sbirs --json

# Timed run with custom tick
bmd-sim-icbm -v -mirvs 5 -cms -duration 30s -tick 2s

# Help: all available flags
bmd-sim-gmd --help
```

### Standard Flags (all sims)

| Flag | Description |
|------|-------------|
| `-v` | Verbose output with simulation details |
| `-i`, `--interactive` | Live updating metrics dashboard |
| `--json` | Machine-readable JSON with parameters |
| `--duration DUR` | Run for specified duration |
| `--tick DUR` | Simulation tick rate (default 1s) |
| `--seed N` | Deterministic random seed |
| `--help`, `-h` | Show all flags |

### Simulator-Specific Flags

| Simulator | Flags | Description |
|-----------|-------|-------------|
| `bmd-sim-gmd` | `-site`, `-variant` | Launch site (FTG/VAFB), variant (CE-I/CE-II) |
| `bmd-sim-sm3` | `-variant` | SM-3 variant (IA/IB/IIA) |
| `bmd-sim-icbm` | `-mirvs`, `-cms` | Number of MIRVs, include countermeasures |
| `bmd-sim-hgv` | `-maneuvers` | Number of maneuvers (default 5) |
| `bmd-sim-tpy2` | `-site`, `-scenario` | Site ID, scenario file |
| `bmd-sim-uewr` | `-scenario` | Scenario file |
| `bmd-sim-sbirs` | `-site` | Site ID |

### Go Module Import

```go
import (
    sbirs "github.com/forge-c2/bmd-sim-sbirs/pkg/constellation"
    c2bmc "github.com/forge-c2/bmd-sim-c2bmc/pkg/c2bmc"
    icbm "github.com/forge-c2/bmd-sim-icbm/pkg/threat"
    // ... etc
)
```