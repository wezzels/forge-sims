# hgv — Hypersonic Glide Vehicle

## Purpose
Models hypersonic glide vehicles with 5-phase trajectory and maneuvering.

## Physics & Fidelity
- 5-phase trajectory: boost, pull-up, glide, maneuver, terminal dive
- Glide at Mach 5-25 at 30-60 km altitude
- Lateral maneuvering (S-turns, jinking)
- Very low RCS (0.001-0.01 m²)
- Detection PK: GMD 0.05, SM-3 0.05, THAAD 0.15, SM-6 0.10

### Scenarios
`avangard`, `df-zf`, `boost-glide-maneuver`, `low-observable`

## Accuracy: High
**Reliable**: 5-phase trajectory, glide aerodynamics, maneuver modeling
**Needs improvement**: Plasma sheath effects, multi-bounce skip trajectories, detailed aero-thermodynamics
