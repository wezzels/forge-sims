# bmd-sim-decoy — Decoy & Penetration Aids

## Purpose
Models ballistic missile penetration aids: balloon decoys, replica decoys, chaff, active jammers.

## Physics & Fidelity
- Balloon decoys: lightweight, large RCS, no thermal mass (discriminated by midcourse cooling)
- Replica decoys: matched RCS and thermal signature, harder to discriminate
- Chaff clouds: RF scattering, rapid deployment, frequency-dependent RCS
- Active jammers: noise jamming, DRFM, frequency-agile
- Penetration aid deployment timing and spatial distribution

### Scenarios
`balloon-decoys`, `replica-decoys`, `chaff-cloud`, `active-jammer`, `mixed-penaids`

## Accuracy: Medium
**Reliable**: Decoy types, deployment patterns, discrimination difficulty scaling
**Needs improvement**: Detailed thermal modeling (midcourse cooling), chaff cloud dynamics, DRFM waveform modeling
