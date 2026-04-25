# bmd-sim-c2bmc — Command & Control Battle Management

## Purpose
Models the C2BMC (Command and Control, Battle Management, and Communications) network for BMDS command integration.

## Physics & Fidelity
- Sensor-to-shooter kill chain timing
- Network latency and data fusion
- Doctrine modes: shoot-shoot-look, shoot-look-shoot, adaptive
- ROE levels (0-5): peacetime through autonomous
- Engagement authorization and deconfliction

### Scenarios
`kill-chain-timing`, `doctrine-comparison`, `roe-escalation`, `sensor-fusion`

## Accuracy: Medium
**Reliable**: Kill chain timing, doctrine modes, ROE levels
**Needs improvement**: Real C2BMC network protocol details, SATCOM latency modeling, multi-domain command structure
