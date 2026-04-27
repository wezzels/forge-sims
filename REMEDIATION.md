# FORGE-Sims Remediation Plan — Round 20 Update

**Date:** 2026-04-27
**Goal:** Bring remaining sims to full pass with verified high-fidelity physics

---

## Completed ✅

| Sim | What was done | Round |
|-----|-------------|-------|
| launch-veh-sim | Max Q, -i/-v/-seed, JSON, stage sep, circularization | R17-R20 |
| air-combat-sim | Full BVR kill chain, Pk>0, events, F-22 fuel fix | R18-R20 |
| bmd-sim-jamming | -power/-freq/-range flags, data section | R18 |
| bmd-sim-jreap | -protocol/-messages flags, data section | R18 |
| bmd-sim-jrsc | -sensors/-alert flags, data section | R18 |
| bmd-sim-link16 | -participants/-messages flags, data section | R18 |
| bmd-sim-c2bmc | Moved computed data to `data` section | R19 |
| bmd-sim-decoy | Moved computed data to `data` section | R19 |
| bmd-sim-gfcb | Moved computed data to `data` section | R19 |
| bmd-sim-ifxb | Moved computed data to `data` section | R19 |
| hgv/bmd-sim-hgv | Fixed JSON preamble, added detection data section | R19-R20 |
| ir-signature | Fixed JSON preamble, added detection data section | R19 |
| kill-chain-rt | Fixed JSON for infeasible scenarios, added data section | R20 |

---

## Remaining Work (P3)

| # | Sim | Issue | Priority |
|---|-----|-------|----------|
| 1 | Vulcan Centaur | Low-TWR S2 crashes (needs guidance algorithm) | P3 |
| 2 | Delta IV | Low-TWR S2 crashes (needs guidance algorithm) | P3 |
| 3 | Falcon Heavy | Orbit at 793 km (excess delta-V for payload) | P3 |
| 4 | Atlas V | Orbit at 754 km (excess delta-V) | P3 |
| 5 | Ariane 6 | Orbit at 1569 km (excess delta-V) | P3 |
| 6 | intel-collection-sim | Subcommand pattern, not standard `-json` | P3 |
| 7 | cyber-redteam-sim | Config-driven by design | P3 |
| 8 | maritime-sim | Config-driven by design | P3 |
| 9 | space-war-sim | Config-driven by design | P3 |
