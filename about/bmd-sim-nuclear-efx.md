# bmd-sim-nuclear-efx — Nuclear Detonation Effects

## Purpose
Models nuclear detonation effects on BMDS: EMP, thermal radiation, blast, ionospheric blackout, and satellite kill.

## Physics & Fidelity
- EMP: HEMP E1 (50 kV/m), E2, E3 components with range-dependent field strengths
- Thermal radiation: inverse-square law with atmospheric absorption
- Blast overpressure: modified Brode equation
- Ionospheric blackout: HF/VHF propagation disruption (minutes to hours)
- Satellite kill: direct radiation, Van Allen belt contamination
- Radar degradation: worst-case sensor factor (not compounded product)

### Scenarios
`high-altitude-emp`, `exo-atmospheric-burst`, `surface-burst`, `multiple-burst`

## Accuracy: High
**Reliable**: HEMP E1/E2/E3 modeling, blast overpressure, ionospheric blackout, satellite kill
**Needs improvement**: Detailed prompt gamma modeling, delayed fallout, electromagnetic tree modeling
