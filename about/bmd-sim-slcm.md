# bmd-sim-slcm — Submarine-Launched Cruise Missile

## Purpose
Models SLCM threats from submerged submarine launches.

## Physics & Fidelity
- Underwater launch phase (buoyancy, cavitation)
- Low-altitude cruise profile (terrain-following)
- Terminal dive attack
- Radar cross-section modeling (low observable)
- Range 1,500-2,500 km depending on variant

### Scenarios
`kalibr-launch`, `tomahawk-attack`, `low-observable-scrm`, `saturation-scrm`

## Accuracy: Medium
**Reliable**: Cruise profile, low-altitude flight, RCS modeling
**Needs improvement**: Underwater launch physics simplified, no terrain-following model, no waypoints
