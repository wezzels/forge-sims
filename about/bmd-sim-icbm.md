# bmd-sim-icbm — Intercontinental Ballistic Missile

## Purpose
Models ICBM threat trajectories (Minuteman-III, DF-41, R-36M2) with MIRV deployment, countermeasures, and RV aerodynamics.

## Physics & Fidelity

### Trajectory Model
- 3-stage boost with realistic thrust profiles
- Keplerian midcourse with J2 perturbation
- Re-entry aerodynamics (ballistic coefficient, drag)
- MIRV bus dispensing with dispense pattern

### Threat Parameters
- Minuteman-III: 13,000 km range, 3 stages, MIRV up to 3 RVs
- DF-41: 14,000 km range, 3 stages, up to 10 MIRV
- R-36M2 (Satan): 16,000 km range, 2 stages, up to 10 RVs
- CEP modeling, RV separation, decoy deployment

### Countermeasures
- Balloon decoys, replica decoys, chaff clouds, jammers
- Maneuvering RV (MaRV), stealth RV (low-observable)
- MIRV bus with post-boost vehicle

### Scenarios
`minuteman-iii`, `df-41`, `r-36m2`, `mirv-raid`, `maneuvering-rv`

## Accuracy: High
**Reliable**: Keplerian midcourse, MIRV deployment, countermeasure types, CEP modeling
**Needs improvement**: Atmospheric re-entry heating, ablation modeling, wind effects on CEP, detailed MIRV dispense patterns
