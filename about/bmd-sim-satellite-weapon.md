# satellite-weapon — ASAT Weapon Simulator

## Purpose
Models anti-satellite weapons (DA-ASAT, co-orbital, directed energy) against orbital targets.

## Physics & Fidelity
- 7 ASAT weapon types: DA-ASAT (SM-3 variant), co-orbital interceptor, kinetic kill vehicle, directed energy, cyber, nuclear EMP, laser
- 8 target satellites: GPS, GEO comms, LEO recon, Starlink, SBIRS, DSP, ISS, Hubble
- Orbital mechanics for intercept geometry
- Kill assessment: kinetic, directed energy, EMP effects

### Scenarios
`da-asat-gps`, `co-orbital-geo`, `laser-leo`, `emp-constellation`, `full-asat`, `starlink-denial`, `sbirs-kill`, `mixed-asat`

## Accuracy: Medium
**Reliable**: Orbital intercept geometry, ASAT weapon types, kill assessment
**Needs improvement**: Debris cloud modeling, orbital perturbation after kill, Kessler cascade, real satellite catalog
