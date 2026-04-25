# bmd-sim-stss — Space Tracking & Surveillance System

## Purpose
Models the STSS constellation for midcourse tracking and discrimination from Low Earth Orbit.

## Physics & Fidelity

### Constellation
- 24 LEO satellites (~1,350 km altitude)
- Visible + LWIR dual-band composite detection
- Binary RV/decoy classification (no confidence scoring)

### Detection
- LEO altitude geometry with satellite visibility windows
- Midcourse tracking with 1-10 second update intervals
- Satellite handoff as targets move between coverage areas

### Scenarios
`midcourse-tracking`, `decoy-discrimination`, `constellation-gap`, `handoff-test`

## Accuracy: Low-Medium
**Reliable**: LEO constellation geometry, basic dual-band detection, visibility windows
**Needs improvement**: Confidence scoring, atmospheric clutter, tracking accuracy model, data latency
