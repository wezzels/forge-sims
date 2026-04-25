# tactical-net — BMDS RF Mesh Network

## Purpose
Models the BMDS tactical RF mesh network with propagation effects and EW degradation.

## Physics & Fidelity
- 38 BMDS nodes across 9 radio types
- Link-16, JREAP, SATCOM propagation models
- Signal propagation with terrain and atmospheric effects
- Jamming/EW effects on network throughput
- 4 scenarios with varying threat levels

### Scenarios
`baseline-network`, `jammed-network`, `saturate`, `multi-hop`

## Accuracy: Medium-High
**Reliable**: 38-node network model, 9 radio types, propagation physics, jamming effects
**Needs improvement**: Detailed crypto modeling, multi-net stacking, relay propagation
