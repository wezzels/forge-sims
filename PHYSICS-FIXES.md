# FORGE-Sims Physics Fixes — Status Round 18

**Date:** 2026-04-26
**Total Binaries:** 58
**High-Fidelity Physics:** 44/58

---

## ✅ Fixed Since Round 17

### launch-veh-sim — FIXED (Round 17)
- Max Q: 94 kPa → 32.6 kPa (real F9 ~35 kPa)
- -i/-v/-seed: restored
- -duration: Go time.Duration format
- JSON: NaN/Inf fixed
- Stage separation: 60km/2525m/s

### air-combat-sim — FIXED (Round 18)
- Pk was 0.00 (no weapons employment)
- Full BVR kill chain implemented: detect→track→identify→decide→launch→guide→kill
- Radar detection with aspect-dependent RCS
- APG-77 detects Su-57 at ~90km; N036 detects F-22 at ~43km
- AIM-120D/AIM-9X/R-77M/R-74M2 loadouts
- Proportional navigation guidance
- Events populated: detection, launch, kill, miss

### C2/Infrastructure Sims — 4/8 UPDATED (Round 18)
- bmd-sim-jamming: Added -power/-freq/-range flags, JNR/burnthrough/ECCM data section
- bmd-sim-jreap: Added -protocol/-messages flags, transport/throughput/CRC data section
- bmd-sim-jrsc: Added -sensors/-alert flags, fusion/correlation data section
- bmd-sim-link16: Added -participants/-messages flags, TDMA/slot data section

---

## Remaining Physics Fixes (P2/P3)

### P2 — Moderate Effort

| # | Sim | Issue | Effort |
|---|-----|-------|--------|
| 1 | population-impact-sim | No `-json` flag at all | ~50 lines |
| 2 | bmd-sim-c2bmc | JSON data in `parameters` not `data` section | ~30 lines |
| 3 | bmd-sim-decoy | JSON data in `parameters` not `data` section | ~30 lines |
| 4 | bmd-sim-gfcb | JSON data in `parameters` not `data` section | ~30 lines |
| 5 | bmd-sim-ifxb | JSON data in `parameters` not `data` section | ~30 lines |
| 6 | emp-effects | Needs `-scenario` flag for JSON | ~20 lines |
| 7 | rcs | Needs `-threat` flag for JSON | ~20 lines |
| 8 | satellite-weapon | Needs `-scenario` flag for JSON | ~20 lines |
| 9 | hgv | Text preamble before JSON | ~5 lines |
| 10 | ir-signature | Text preamble before JSON | ~5 lines |
| 11 | kill-chain-rt | Text preamble before JSON | ~5 lines |

### P3 — Minor/Cosmetic

| # | Sim | Issue |
|---|-----|-------|
| 1 | air-combat-sim | F-22 fuel 6000 kg (real ~8200 kg internal) |
| 2 | launch-veh-sim | Eccentricity always 0 (perfectly circular) |
| 3 | launch-veh-sim | Orbital velocity slightly low (7187 vs ~7660 m/s) |
| 4 | 3 config-driven sims | cyber-redteam, maritime, space-war need standard CLI flags |

### Config-Driven Sims (by design)

These require config files or interactive sessions rather than standard CLI flags:
- cyber-redteam-sim (needs `-config`)
- maritime-sim (needs `-aar`)
- space-war-sim (needs `-aar`)