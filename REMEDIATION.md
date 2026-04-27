# FORGE-Sims Remediation Plan — Round 18 Update

**Date:** 2026-04-26
**Goal:** Bring remaining sims to full pass with verified high-fidelity physics

---

## Completed ✅

| Sim | What was done | Round |
|-----|-------------|-------|
| launch-veh-sim | Max Q, -i/-v/-seed, JSON, stage sep | R17 |
| air-combat-sim | Full BVR kill chain, Pk>0, events | R18 |
| bmd-sim-jamming | -power/-freq/-range flags, data section | R18 |
| bmd-sim-jreap | -protocol/-messages flags, data section | R18 |
| bmd-sim-jrsc | -sensors/-alert flags, data section | R18 |
| bmd-sim-link16 | -participants/-messages flags, data section | R18 |

---

## Remaining Work

### Quick Wins (~30 min total)

| # | Sim | Fix | Lines |
|---|-----|-----|-------|
| 1 | bmd-sim-c2bmc | Move computed data from `parameters` to `data` section | ~30 |
| 2 | bmd-sim-decoy | Move computed data from `parameters` to `data` section | ~30 |
| 3 | bmd-sim-gfcb | Move computed data from `parameters` to `data` section | ~30 |
| 4 | bmd-sim-ifxb | Move computed data from `parameters` to `data` section | ~30 |
| 5 | hgv | Strip text preamble before JSON | ~5 |
| 6 | ir-signature | Strip text preamble before JSON | ~5 |
| 7 | kill-chain-rt | Strip text preamble before JSON | ~5 |

### Medium Effort (~1-2 hours each)

| # | Sim | Fix | Lines |
|---|-----|-----|-------|
| 8 | population-impact-sim | Add `-json` flag with data section | ~50-100 |
| 9 | emp-effects | Add `-scenario` default for `-json` mode | ~20 |
| 10 | rcs | Add `-threat` default for `-json` mode | ~20 |
| 11 | satellite-weapon | Add `-scenario` default for `-json` mode | ~20 |

### Low Priority (by design)

| # | Sim | Issue |
|---|-----|-------|
| 12 | cyber-redteam-sim | Config-driven, needs `-config` flag |
| 13 | maritime-sim | Config-driven, needs `-aar` flag |
| 14 | space-war-sim | Config-driven, needs `-aar` flag |