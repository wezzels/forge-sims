# bmd-sim-irbm — Intermediate-Range Ballistic Missile

## Purpose
Models IRBM threats (DF-26, Agni-V, etc.) with 3,000-5,500 km range trajectories.

## Physics & Fidelity
- 2-stage boost with realistic thrust and Isp
- Midcourse Keplerian trajectory with Earth rotation
- Re-entry with ballistic coefficient
- MIRV capability (1-3 warheads)
- Maneuvering RV variants

### Scenarios
`df-26`, `agni-v`, `irbm-raid`, `maneuvering-irbm`

## Accuracy: Medium-High
**Reliable**: Trajectory physics, MIRV, maneuvering RV
**Needs improvement**: Less validated threat data than ICBM, limited countermeasure modeling
