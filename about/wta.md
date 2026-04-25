# wta — Weapon-Target Assignment Optimizer

## Purpose
Models weapon-target assignment with doctrine and ROE constraints.

## Physics & Fidelity
- 9 interceptor types with PK by threat type (GMD vs ICBM=0.55, vs HGV=0.05)
- 8 WTA scenarios with varying threat sizes
- 4 solvers: greedy, Hungarian algorithm, SLS (stochastic local search), cost-optimal
- 5 ROE levels (0=peacetime → 5=autonomous)
- 3 doctrine modes: shoot-shoot, shoot-look-shoot, adaptive

### Scenarios
`single-icbm`, `5-icbm-raid`, `mixed-threat`, `pacific-surge`, `saturation`, `regional-4-mrbm`, `hgv-raid`, `full-spectrum`

## Accuracy: High
**Reliable**: PK matrices, Hungarian algorithm, ROE modeling, doctrine constraints
**Needs improvement**: Real-time assignment updates, multi-objective optimization, resource replenishment
