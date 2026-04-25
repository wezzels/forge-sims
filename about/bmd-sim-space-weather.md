# bmd-sim-space-weather — Space Weather Effects

## Purpose
Models space weather (solar flares, geomagnetic storms, ionospheric disturbances) on BMDS operations.

## Physics & Fidelity
- Kp index modeling (0-9 geomagnetic storm scale)
- Ionospheric TEC (Total Electron Content) effects on radar
- HF/VHF propagation disruption during storms
- Satellite drag during geomagnetic storms
- GPS accuracy degradation

### Scenarios
`quiet-sun`, `moderate-storm`, `severe-storm`, `carrington-event`

## Accuracy: Medium
**Reliable**: Kp index effects, TEC modeling, basic propagation disruption
**Needs improvement**: Real-time SWPC data integration, auroral zone effects, satellite anomaly rates
