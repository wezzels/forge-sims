# bmd-sim-sbirs — SBIRS Constellation Simulator

## Purpose
Models the Space-Based Infrared System (SBIRS) for ballistic missile early warning from GEO and HEO orbits.

## Physics & Fidelity

### Constellation Model
- **4 GEO + 2 HEO satellites** (real: 6 GEO + 2 HEO + 5 DSP legacy)
- GEO at 35,786 km with fixed longitude slots
- HEO in Molniya orbits (63.4° inclination, 12-hour period)
- Each satellite has position, sensor footprint, and detection parameters

### IR Detection
- Single composite IR value (2.7μm MWIR + 4.3μm SWIR combined)
- Detection range ~40,000 km using inverse-square law
- Boost-phase plume detection via Planck radiation model
- Fixed scan rate (real SBIRS: 6-11 Hz staring, 1 Hz scanning)
- No atmospheric absorption modeling

### Scenarios
`regional-conflict`, `global-surge`, `constellation-health`, `arctic-launch`

## Accuracy: Medium

### Reliable
- Constellation geometry and coverage calculations
- GEO/HEO visibility windows and coverage gaps
- Boost-phase detection timing (10-30s from launch)

### Needs Improvement
1. No MWIR/SWIR spectral band separation
2. No scan/stare mode switching
3. No atmospheric path attenuation (MODTRAN-like)
4. No satellite-to-satellite track handoff
5. Missing 5 DSP legacy satellites
6. No C2BMC reporting latency (10-30s delay)
7. No cloud obscuration effects