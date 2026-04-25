# bmd-sim-electronic-attack — Electronic Attack / ECM

## Purpose
Models electronic countermeasures against BMDS radars: jamming, DRFM, chaff, and their effects on detection.

## Physics & Fidelity
- Noise jamming: barrage, spot, swept (J/S ratio degradation)
- DRFM: Range Gate Pull-Off (RGPO), Velocity Gate Pull-Off (VGPO)
- Chaff: frequency-dependent scattering, deployment geometry
- ECCM techniques: frequency agility, sidelobe cancellation, STC
- Radar-specific degradation factors per radar type

### EW Scenarios (via -ew flag on kill-chain-rt)
`barrage-jamming` (−7% PK), `drfm-deception` (−20% PK), `heavy-ew` (−90% PK), `frequency-hopping`, `sidelobe-blanking`

## Accuracy: Medium-High
**Reliable**: J/S ratio calculations, DRFM deception modeling, per-radar degradation
**Needs improvement**: Detailed waveform modeling, cognitive jamming, multi-radar cross-jamming
