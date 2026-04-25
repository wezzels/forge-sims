# bmd-sim-gmd — Ground-Based Midcourse Defense

## Purpose
Models the GMD system (GBI interceptor) for midcourse intercept of ICBM-class threats.

## Physics & Fidelity

### Interceptor Parameters
- **GBI**: 3-stage booster, EKV kill vehicle
- **Range**: ~5,000 km, altitude up to 2,000 km
- **Speed**: ~7.8 km/s burnout velocity
- **PK**: CE-I 0.55, CE-II 0.60
- **Divert capability**: ~800 km crossrange

### Guidance
- Proportional navigation with divert maneuvers
- Seeker acquisition at ~200-400 km range
- Hit-to-kill (kinetic energy interceptor)

### HGV Engagement
- GMD vs HGV: PK = 0.05 (extremely low due to maneuvering)

### Scenarios
`gmd-vs-icbm`, `gmd-vs-hgv`, `shoot-shoot-look`, `multi-intercept`

## Accuracy: High
**Reliable**: EKV guidance, proportional navigation, realistic PK values, divert capability
**Needs improvement**: Detailed seeker model (MWIR IR), kill vehicle homing, debris avoidance
