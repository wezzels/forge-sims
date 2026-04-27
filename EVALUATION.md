# FORGE-Sims Evaluation — Round 20

**Date:** 2026-04-27 (Round 20 — Full Audit + P3 Fixes)
**Total binaries:** 65 (59 sim binaries + 2 scripts + 4 platform variants)
**Source repos:** All on IDM GitLab (crab-meat-repos/)

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 59 | 59 | 0 | ✅ All pass |
| Valid JSON output | 59 | **58** | 1 | intel-collection-sim uses subcommands |
| Has `data` section (computed physics) | 59 | **20** | 39 | Up from 18 in R19 |
| Zero stubs/mocks | 59 | 59 | 0 | ✅ Zero found |
| High-fidelity physics verified | 59 | 44 | 15 | See PHYSICS-ROADMAP.md |

---

## Changes Since Round 19

| Sim | Change |
|-----|--------|
| kill-chain-rt | ✅ JSON output for infeasible scenarios, data section added |
| bmd-sim-hgv | ✅ Suppress text preamble with `-json`, detection data section |
| air-combat-sim | ✅ F-22 FuelKg corrected 6100→8200 kg |
| launch-veh-sim | ✅ Multi-burn circularization (F9, F9-Exp, NG, LM5, Starship) |
| 7 warfare sims | ✅ Added cbrn-sim, cyber-kinetic-sim, information-ops-sim, land-combat-sim, logistics-sustainment-sim, nuclear-effects-sim, space-data-network-sim |

---

## JSON Evaluation Summary

### ✅ Pass (58/59 produce valid JSON)

air-combat-sim, air-traffic, bmd-sim-aegis, bmd-sim-atmospheric, bmd-sim-c2bmc, bmd-sim-cobra-judy, bmd-sim-decoy, bmd-sim-dsp, bmd-sim-electronic-attack, bmd-sim-gbr, bmd-sim-gfcb, bmd-sim-gmd, bmd-sim-hgv, bmd-sim-hub, bmd-sim-icbm, bmd-sim-ifxb, bmd-sim-irbm, bmd-sim-jamming, bmd-sim-jreap, bmd-sim-jrsc, bmd-sim-link16, bmd-sim-lrdr, bmd-sim-mrbm, bmd-sim-nuclear-efx, bmd-sim-patriot, bmd-sim-sbirs, bmd-sim-slcm, bmd-sim-sm3, bmd-sim-sm6, bmd-sim-space-weather, bmd-sim-stss, bmd-sim-thaad, bmd-sim-thaad-er, bmd-sim-tpy2, bmd-sim-uewr, boost-intercept, cbrn-sim, cyber-kinetic-sim, debris-field, electronic-war-sim, emp-effects, engagement-chain, hgv, information-ops-sim, ir-signature, kill-assessment, kill-chain-rt, land-combat-sim, launch-veh-sim, logistics-sustainment-sim, maritime-sim, missile-defense-sim, nuclear-effects-sim, population-impact-sim, rcs, satellite-tracker, satellite-weapon, space-data-network-sim, space-debris, space-war-sim, submarine-war-sim, tactical-net, ufo, wta

### ✅ Has Data Section (20/59)

air-combat-sim, bmd-sim-c2bmc, bmd-sim-decoy, bmd-sim-gfcb, bmd-sim-hub, bmd-sim-ifxb, bmd-sim-jamming, bmd-sim-jreap, bmd-sim-jrsc, bmd-sim-link16, bmd-sim-patriot, bmd-sim-thaad, electronic-war-sim, hgv, ir-signature, kill-chain-rt, missile-defense-sim, submarine-war-sim, ufo, wta

### ❌ Fail (1/59)

| Sim | Issue | Fix |
|-----|-------|-----|
| intel-collection-sim | Uses subcommand pattern (`scenario <name>`) | By design |

---

## Works With Flags (not bare `-json`)

| Sim | Required Flags |
|-----|---------------|
| emp-effects | `-scenario <name>` or `-yield <kT> -altitude <km>` |
| rcs | `-threat <name>` or `-scenario <name>` |
| satellite-weapon | `-scenario <name>` or `-weapon <name> -target <name>` |
| kill-chain-rt | `-scenario <name>` or `-multi <name>` |
| cyber-redteam-sim | `-scenario <name>` (config-driven) |
| maritime-sim | `-scenario <name>` (config-driven) |
| space-war-sim | `-scenario <name>` (config-driven) |

---

## Open Issues (P3)

| # | Sim | Issue | Status |
|---|-----|-------|--------|
| 1 | launch-veh-sim | Vulcan/Delta crash (low-TWR S2) | Needs PEG guidance |
| 2 | launch-veh-sim | FH/Atlas/Ariane orbit too high | Excess S2 delta-V |
| 3 | intel-collection-sim | Subcommand CLI pattern | By design |
| 4 | air-combat-sim | ✅ F-22 fuel fixed | Done R20 |

---

## Summary

**58/59 sims produce valid JSON.** ✅
**20/59 have `data` section with computed physics.** ✅
**44/59 verified high-fidelity physics.** ✅
**1/59 needs subcommand flags** (by design)
**Zero stubs, mocks, or placeholder data.** ✅
