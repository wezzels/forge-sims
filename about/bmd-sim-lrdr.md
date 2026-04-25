# bmd-sim-lrdr — Long-Range Discrimination Radar

## Purpose
Models the LRDR at Clear, Alaska for midcourse discrimination and tracking of ballistic threats.

## Physics & Fidelity

### Radar Parameters
- **Frequency**: S-band (~2-4 GHz)
- **Range**: ~3,000+ km (tracking), ~5,000+ km (detection)
- **System losses**: 10 dB
- **ISLR**: -30 dB

### Discrimination
- Object-by-object with confidence scoring
- RCS-based (tumbling, precession signatures)
- Categories: RV, decoy, uncertain, unknown
- Post-intercept kill assessment

### Scenarios
`midcourse-discrimination`, `icbm-tracking`, `decoy-rejection`, `kill-assessment`

## Accuracy: High
**Reliable**: S-band radar equation, discrimination logic, long-range geometry
**Needs improvement**: Auroral clutter, multi-target resource allocation, ISAR imaging, debris tracking
