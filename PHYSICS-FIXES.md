# FORGE-Sims Physics Fixes — Status Round 20

**Date:** 2026-04-27
**Total Binaries:** 65 (58 original + 7 new warfare sims)
**High-Fidelity Physics:** 44/58 original + 7 new = 51/65

---

## ✅ Fixed Through Round 19

### All items from R17-R19 are complete:
- launch-veh-sim: Max Q fixed, -i/-v/-seed, JSON, stage sep, circularization
- air-combat-sim: Full BVR kill chain, Pk>0, events
- C2 sims: Data sections added (c2bmc, decoy, gfcb, ifxb, jamming, jreap, jrsc, link16)
- hgv/ir-signature: JSON preamble fixed, detection data sections
- population-impact-sim: -json flag working
- emp-effects: -scenario flag working
- rcs: -threat flag working
- satellite-weapon: -scenario flag working

## Round 20 Fixes

### air-combat-sim — FIXED (Round 20)
- F-22 FuelKg: 6100 → 8200 kg (matches real internal fuel capacity)
- F-22 thrust already correct (2×116 kN dry, 2×156 kN AB)

### launch-veh-sim — CIRCULARIZATION ADDED (Round 20)
- Multi-burn S2 with circularization for orbit vehicles
- F9: 449×391 km, ecc 0.004 ✅
- F9-Expendable: 457×403 km, ecc 0.004 ✅
- New Glenn: 480×417 km, ecc 0.005 ✅
- Long March 5: 449×382 km, ecc 0.005 ✅
- Starship: 512×454 km, ecc 0.004 ✅
- Falcon Heavy: 793×723 km, ecc 0.005 (circularized, excess delta-V for payload)
- Atlas V: 754×684 km, ecc 0.005 (circularized, excess delta-V)
- Ariane 6: 1569×1493 km, ecc 0.005 (circularized, excess delta-V)
- Vulcan Centaur: NEEDS GUIDANCE (low-TWR S2 crashes)
- Delta IV: NEEDS GUIDANCE (low-TWR S2 crashes)

### 7 New Warfare Sims (Round 20)
- cbrn-sim: 6 scenarios, zone modeling ✅
- cyber-kinetic-sim: 5 scenarios ✅
- information-ops-sim: 5 scenarios ✅
- land-combat-sim: 6 scenarios ✅
- logistics-sustainment-sim: 5 scenarios ✅
- nuclear-effects-sim: 6 scenarios ✅
- space-data-network-sim: 5 scenarios ✅

## Remaining Issues (P3)

| # | Sim | Issue | Priority |
|---|-----|-------|----------|
| 1 | Vulcan Centaur | Low-TWR S2 crashes — needs guidance algorithm | P3 |
| 2 | Delta IV | Low-TWR S2 crashes — needs guidance algorithm | P3 |
| 3 | Falcon Heavy | Orbit at 793 km (excess S2 delta-V for 26t payload) | P3 |
| 4 | Atlas V | Orbit at 754 km (excess S2 delta-V) | P3 |
| 5 | Ariane 6 | Orbit at 1569 km (excess S2 delta-V) | P3 |
| 6 | 3 config sims | cyber-redteam, maritime need configs deployed | P3 |
