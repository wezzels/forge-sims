# bmd-sim-sm6 — Standard Missile 6

## Purpose
Models the SM-6 interceptor for terminal air and missile defense.

## Physics & Fidelity
- **SM-6**: Active radar homing, extended-range SAM
- **Range**: ~240 km, altitude ~33 km
- **PK**: 0.75 (air targets), lower for BM targets
- **Dual role**: air defense + terminal BMD

### Scenarios
`sm6-air-defense`, `sm6-terminal-bmd`, `sm6-vs-cruise`, `sm6-saturation`

## Accuracy: Medium
**Reliable**: Range, altitude, PK for air targets
**Needs improvement**: Active seeker model, multi-path effects, over-the-horizon engagement
