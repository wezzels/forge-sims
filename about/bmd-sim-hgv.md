# bmd-sim-hgv — Hypersonic Glide Vehicle

## Purpose
Models hypersonic glide vehicles (AVangard, DF-ZF/HGV-DF) with boost-glide trajectory and maneuvering.

## Physics & Fidelity

### Trajectory Model
- 5-phase trajectory: boost, pull-up, glide, maneuver, terminal dive
- Boost phase on ICBM/IRBM booster (same as corresponding missile)
- Pull-up at 60-100 km altitude
- Glide at Mach 5-25 at 30-60 km altitude
- Lateral maneuvering (S-turns, jinking)
- Terminal dive at steep angles

### Detection Challenge
- Very low radar cross-section (~0.001-0.01 m²)
- Below traditional radar horizon until terminal phase
- Maneuvering defeats predictive intercept
- PK against HGV: GMD 0.05, SM-3 0.05, THAAD 0.15, SM-6 0.10

### Scenarios
`avangard`, `df-zf`, `boost-glide-maneuver`, `low-observable-hgv`

## Accuracy: High
**Reliable**: 5-phase trajectory, glide aerodynamics, lateral maneuver, detection difficulty modeling
**Needs improvement**: Plasma sheath effects, detailed aero-thermodynamics, multi-bounce skip trajectories
