# bmd-sim-patriot — PATRIOT Air Defense System

## Purpose
Models the PATRIOT air defense system (PAC-2/PAC-3 MSE) for terminal ballistic missile defense.

## Physics & Fidelity

### Battery Organization
- Fire unit: AN/MPQ-65 radar, ECS, 4-8 launchers (4 missiles each)
- PAC-3 MSE: Hit-to-kill, lethal radius ~15m, range 60 km, altitude 40 km
- PAC-2 GEM: Blast-frag, lethal radius ~25m, range 70 km

### Radar (AN/MPQ-65)
- **Frequency**: C-band (~5 GHz)
- **Peak power**: ~100 kW
- **System losses**: 11 dB
- **Range**: ~170 km (detection), ~100 km (tracking)

### Interceptors
- Proportional navigation guidance
- PK: PAC-3 MSE ~0.85, PAC-2 GEM ~0.65
- Launch sequencing and fire control logic

### Scenarios
`terminal-defense`, `multi-target-engagement`, `pac3-vs-pac2`, `saturation-raid`

## Accuracy: High
**Reliable**: Battery organization, PAC-3/PAC-2 dual model, C-band radar, proportional navigation
**Needs improvement**: Terrain masking, TBM/Cruise mode switching, PAC-3 MSE Ka-band seeker, datalink latency
