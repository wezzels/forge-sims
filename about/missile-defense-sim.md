# missile-defense-sim — Layered Missile Defense Simulation

## Purpose
Models layered ballistic missile defense with GMD, SM-3, THAAD, and Patriot interceptors against ICBM/IRBM threats.

## Physics & Fidelity
- ICBM threat modeling (Minuteman-III, DF-41, R-36M2, etc.)
- Interceptor models: GMD, SM-3 Block IIA, THAAD, PAC-3 MSE
- Atmospheric model: speed of sound, gravity, drag
- Seeker modeling: acquisition range, FOV, tracking accuracy
- Track management and engagement scheduling

### Scenarios
`layered-defense` (4-layer defense), `icbm-attack` (single ICBM), `saturation-raid` (8 threats), `hgv-defense`

## Accuracy: Medium
**Reliable**: Interceptor parameters, threat modeling, layered defense concept
**Needs improvement**: Simplified engagement geometry, no detailed guidance law, basic track management
