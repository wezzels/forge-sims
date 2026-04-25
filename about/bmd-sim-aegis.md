# bmd-sim-aegis — Aegis Combat System

## Purpose
Models the Aegis combat system (AN/SPY-1D + SM-3/SM-6) for sea-based BMD.

## Physics & Fidelity

### Radar (AN/SPY-1D)
- **Frequency**: S-band (~3.3 GHz)
- **Peak power**: 4-6 MW per array face
- **System losses**: 14 dB (highest in BMDS — ship environment)
- **4 array faces** for 360° coverage
- **Range**: ~320 km detection, ~170 km tracking

### Ship Model
- DDG-51 class destroyer
- Heading, roll, pitch affect radar horizon
- Sea state degradation model included

### Interceptors
- SM-3 Block IIA: Midcourse, 2,500 km range, 500 km altitude, PK 0.65
- SM-6: Terminal, 240 km range, 33 km altitude

### Scenarios
`sea-based-defense`, `sm3-engagement`, `multi-threat-tracking`, `sea-state-degradation`

## Accuracy: High
**Reliable**: SPY-1D 4-face model, sea state effects, SM-3/SM-6 performance, ship orientation
**Needs improvement**: CEC model, multi-ship task force, Aegis Ashore variant, detailed radar scheduling
