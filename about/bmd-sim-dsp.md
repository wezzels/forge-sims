# bmd-sim-dsp — Defense Support Program

## Purpose
Models the legacy DSP satellite constellation for strategic missile warning from geostationary orbit.

## Physics & Fidelity

### Constellation
- 3 GEO DSP satellites (~120° each)
- Lower resolution than SBIRS (single IR band)
- Coverage gap analysis between satellite footprints
- Correlation with SBIRS for dual-confirmation tracking

### Detection
- IR plume detection at GEO altitude
- Single composite IR band
- Detection range ~40,000 km
- SBIRS cross-referencing for confirmation

### Scenarios
`strategic-warning`, `coverage-gap`, `sbirs-correlation`, `degraded-mode`

## Accuracy: Low-Medium
**Reliable**: GEO constellation geometry, coverage gaps, SBIRS correlation
**Needs improvement**: Spectral modeling, atmospheric effects, scan rate differentiation
