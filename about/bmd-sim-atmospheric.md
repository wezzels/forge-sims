# bmd-sim-atmospheric — Atmospheric Effects on BMD

## Purpose
Models atmospheric effects (rain, fog, cloud, ionospheric scintillation) on BMDS sensor performance.

## Physics & Fidelity
- Rain attenuation: frequency-dependent (X-band worst, UHF minimal)
- Fog and cloud attenuation
- Ionospheric scintillation at UHF/VHF
- 6 environment presets: clear, moderate rain, heavy rain, fog, solar storm, Arctic

### Scenarios
`clear`, `moderate-rain`, `heavy-rain`, `fog`, `solar-storm`, `arctic`

## Accuracy: Medium-High
**Reliable**: ITU rain attenuation model, fog/cloud absorption, frequency-dependent effects
**Needs improvement**: Real-time weather data integration, seasonal variations, geographic-specific models
