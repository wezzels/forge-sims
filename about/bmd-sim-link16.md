# bmd-sim-link16 — Tactical Data Link

## Purpose
Models Link-16 tactical data link for BMDS network communication.

## Physics & Fidelity
- Time Division Multiple Access (TDMA) slot allocation
- Message types: J-series (free text, surveillance, PPLI, etc.)
- Network participation groups
- Jamming effects (barrage, spot, frequency-hopping resistance)
- Effective data rate under jamming conditions
- Network throughput calculation

### Scenarios
`link16-baseline`, `link16-jammed`, `link16-saturation`, `frequency-hopping`

## Accuracy: Medium-High
**Reliable**: TDMA slot model, J-series messages, jamming effects
**Needs improvement**: Crypto modeling, relay propagation, multi-net stacking
