# submarine-war-sim — Submarine Warfare Simulation

## Purpose
Models submarine warfare: ASW helicopters, sonar systems, torpedoes, minefields, and hydrodynamics.

## Physics & Fidelity

### Hydrodynamics
- Submarine models: Virginia, Seawolf, 688 (US), Kilo, Akula (Russian)
- Cavitation threshold, noise level vs. depth/speed
- Hydrodynamic resistance, propulsive efficiency, shaft power

### Sonar
- Passive sonar equation: SNR = SL - TL - NL + DI
- Active sonar with reverberation
- Sound speed profiles: deep ocean, Arctic
- Convergence zone, thermocline, surface duct modeling
- Sonar systems: BQQ-5, TB-23, SQS-53, DIFAR sonobuoy

### Torpedoes
- Mk-48, Mk-46, Mk-54 (US), VA-111 Shkval, Type-53 (Chinese)
- Homing probability, damage radius, spread patterns

### Mines
- CAPTOR, moored influence, bottom contact, rising mines
- Minefield effectiveness, transit risk, sweep effectiveness

### ASW
- SH-60B/R Seahawk, P-8 Poseidon
- Sonobuoy field deployment, dipping sonar
- ASW engagement probability

### Scenarios
`asw-patrol`, `torpedo-attack`, `minefield`, `sonar-detection`

## Accuracy: Medium-High
**Reliable**: Sonar equation, hydrodynamic models, torpedo performance, mine effectiveness
**Needs improvement**: 3D geometry, oceanographic data integration, real bathymetry, multi-submarine tactics
