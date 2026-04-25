# bmd-sim-thaad — Terminal High Altitude Area Defense

## Purpose
Models THAAD for upper-tier terminal missile defense.

## Physics & Fidelity
- **THAAD**: Hit-to-kill interceptor
- **Range**: ~200 km, altitude up to 150 km
- **PK**: 0.85 (within envelope)
- **Fire control**: AN/TPY-2 radar integration
- **HGV engagement**: PK = 0.15 (higher than GMD/SM-3 due to lower altitude)

### Scenarios
`thaad-terminal-defense`, `thaad-vs-mrbm`, `thaad-vs-hgv`, `thaad-saturation`

## Accuracy: High
**Reliable**: THAAD parameters, hit-to-kill guidance, TPY-2 integration
**Needs improvement**: Detailed seeker model, divert capability, multi-target engagement scheduling
