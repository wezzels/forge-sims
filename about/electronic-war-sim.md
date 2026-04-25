# electronic-war-sim — Electronic Warfare Simulation

## Purpose
Models electronic warfare between jammers and BMDS radars: noise jamming, DRFM deception, spectrum warfare.

## Physics & Fidelity
- Noise jamming: J/S ratio, burnthrough range, effective range degradation
- DRFM: Range Gate Pull-Off (RGPO), Velocity Gate Pull-Off (VGPO)
- ESM: Passive detection, direction finding, threat assessment
- Radar systems: SPY-1D, AN/TPY-2, Big Bird, APG-81
- Jammer systems: ALQ-99, ALQ-249, Khibiny, SORB/TSP

### Scenarios
`noise-jamming` (ALQ-99 vs SPY-1D), `drfm-deception` (DRFM vs SPY-1D), `spectrum-warfare` (multi-system), `esm-intercept` (SLQ-32 passive DF)

## Accuracy: Medium-High
**Reliable**: J/S ratio calculations, DRFM deception modeling, radar equation, ESM passive detection
**Needs improvement**: Detailed waveform modeling, cognitive EW, multi-radar cross-jamming, spectrum occupancy
