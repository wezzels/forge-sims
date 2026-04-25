# emp-effects — EMP Effects on BMDS

## Purpose
Models nuclear EMP effects on BMDS sensors and infrastructure.

## Physics & Fidelity
- HEMP E1 (50 kV/m), E2, E3 components
- Sensor degradation: worst-case factor (not compounded product)
- Recovery times by system type
- 4 scenarios with varying burst altitudes and locations

### Scenarios
`high-altitude-emp`, `exo-atmospheric`, `surface-burst`, `multiple-burst`

## Accuracy: Medium-High
**Reliable**: HEMP E1/E2/E3 field strengths, sensor degradation model, recovery timing
**Needs improvement**: Electromagnetic tree modeling, specific system vulnerability data, cascading failure modes
