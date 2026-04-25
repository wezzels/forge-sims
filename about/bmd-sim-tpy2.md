# bmd-sim-tpy2 — AN/TPY-2 Radar

## Purpose
Models the AN/TPY-2 X-band transportable radar for BMDS forward-based detection and THAAD fire control.

## Physics & Fidelity

### Radar Parameters
- **Frequency**: 8.55-10 GHz (X-band)
- **Peak power**: ~100 kW per element
- **Antenna gain**: 50+ dBi (phased array)
- **System losses**: 10 dB
- **Range**: ~1,000-2,000 km (detection), ~600 km (tracking)

### Detection Model
- Full radar range equation with energy-on-target
- X-band atmospheric attenuation (rain, fog)
- Sector-based search management
- Multiple simultaneous tracking tasks with prioritization
- Task scheduling and resource allocation

### Scenarios
`forward-based-detection`, `thaad-fire-control`, `multi-sector-search`, `rain-degradation`

## Accuracy: High
**Reliable**: X-band detection model, task scheduling, atmospheric attenuation, sector search optimization
**Needs improvement**: Terrain masking, Doppler processing, NCTR, detailed ECM/ECCM
