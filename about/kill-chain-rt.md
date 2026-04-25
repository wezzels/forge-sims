# kill-chain-rt — Real-Time Kill Chain

## Purpose
Models the real-time kill chain with sensor allocation, engagement scheduling, and doctrine-based shoot-look-shoot.

## Physics & Fidelity
- Full kill chain: detect → track → discriminate → engage → assess
- Sensor allocation with track limits and fire control queuing
- 4 doctrine modes: shoot-shoot, shoot-look-shoot, adaptive, autonomous
- 5 ROE levels (0=peacetime → 5=autonomous)
- EW scenarios: barrage, DRFM, heavy EW (via -ew flag)
- Environment scenarios: solar storm, heavy rain, EMP, Arctic (via -env flag)
- Multi-threat scenarios: 8-ICBM saturation, mixed raid, regional, Pacific (via -multi flag)

### EW Results
- Barrage: −7% PK, DRFM: −20% PK, Heavy EW: −90% PK

### Environment Results
- Solar storm: −8%, Heavy rain: −35%, EMP high-alt: −55%, EMP exo: −60%, Arctic: −4%

### Multi-Threat Results
- 8-ICBM saturation: 76% PK, Mixed raid: 38% PK, Regional 4-MRBM: 70% PK, Pacific 6-IRBM: 88% PK

## Accuracy: High
**Reliable**: Kill chain timing, doctrine modeling, EW effects, environment degradation, multi-threat resource contention
**Needs improvement**: Detailed sensor scheduling optimization, real C2 latency, coalition coordination
