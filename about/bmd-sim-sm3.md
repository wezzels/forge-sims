# bmd-sim-sm3 — Standard Missile 3

## Purpose
Models the SM-3 Block IIA interceptor for Aegis BMD midcourse intercept.

## Physics & Fidelity
- **SM-3 Block IIA**: 2-stage booster + KW kill vehicle
- **Range**: ~2,500 km, altitude ~500 km
- **Speed**: ~5.5 km/s burnout
- **PK**: 0.65 (Block IIA)
- **SM-6**: Terminal mode, 240 km range, 33 km altitude

### Scenarios
`sm3-vs-icbm`, `sm3-vs-irbm`, `sm3-vs-hgv`, `aegis-engagement`

## Accuracy: High
**Reliable**: SM-3 performance parameters, proportional navigation, PK modeling
**Needs improvement**: Detailed KW seeker, multi-stage burn profile, ship-based launch geometry
