# rcs — Radar Cross Section Simulator

## Purpose
Models aspect-dependent radar cross-section for BMDS threat characterization.

## Physics & Fidelity
- 13 threat profiles (ICBM RV, MRBM RV, HGV, decoy, etc.)
- 8 frequency bands (VHF through Ka-band)
- Aspect-dependent RCS (nose, broadside, tail, 360°)
- Swerling fluctuation models (I-IV)

### Scenarios
Run with `-threat <type> -band <band> -json`

## Accuracy: Medium-High
**Reliable**: Aspect-dependent RCS, frequency variation, Swerling models
**Needs improvement**: Complex body RCS (cavity scattering), material absorption, active RCS reduction
