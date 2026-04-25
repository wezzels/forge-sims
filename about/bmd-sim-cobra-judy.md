# bmd-sim-cobra-judy — Cobra Judy Radar Ship

## Purpose
Models the AN/SPQ-11 S-band phased array radar on USNS Observation Island for strategic missile data collection.

## Physics & Fidelity

### Collection Model
- S-band (~2-4 GHz) ship-mounted phased array
- Mission scheduling with priority queue
- Multiple simultaneous collection missions
- Data quality scoring based on geometry and SNR

### Scenarios
`collection-mission`, `priority-tracking`, `signature-measurement`, `multi-mission`

## Accuracy: Medium
**Reliable**: Collection mission scheduling, priority-based resource allocation, ship geometry
**Needs improvement**: Detailed radar range equation (simplified), sea state degradation, signature measurement fidelity
