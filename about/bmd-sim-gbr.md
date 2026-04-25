# bmd-sim-gbr — Ground-Based Radar (XBR)

## Purpose
Models the X-band Radar for midcourse discrimination and GMD fire control support.

## Physics & Fidelity

### Radar Parameters
- **Frequency**: 10 GHz (X-band)
- **Peak power**: High-power phased array
- **System losses**: 10 dB
- **ISLR**: -30 dB for discrimination

### Discrimination Model
- Object-by-object discrimination with confidence scoring
- RCS-based classification (RV: 0.01-0.1 m², decoy: 0.001-10 m²)
- Motion-based discrimination (tumbling, precession)
- Integrated result: RV, decoy, uncertain, unknown

### Scenarios
`midcourse-discrimination`, `rv-decoy-separation`, `kill-assessment`, `multi-target-track`

## Accuracy: High
**Reliable**: X-band radar equation, discrimination logic with confidence levels, acquisition/tracking/discrimination pipeline
**Needs improvement**: ISAR imaging, Doppler discrimination (RV spin), multi-function scheduling, debris tracking for kill assessment
