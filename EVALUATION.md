# FORGE-Sims Evaluation — Round 16

**Date:** 2026-04-24 (Round 16)
**Repo:** github.com/wezzels/forge-sims
**Total binaries:** 54

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 54 | 52 | 2 | maritime, space-war (config-driven) |
| All standard flags | 54 | 46 | 8 | 3 config-driven, 3 engine-only, 2 partial |
| Clean JSON output | 54 | 50 | 4 | 3 config-driven, 1 engine-only |
| Zero stubs/mocks | 54 | 54 | 0 | ✅ Zero found |
| Zero format string bugs | 54 | 54 | 0 | ✅ All fixed |

---

## Changes Since Round 15

| Sim | Change |
|-----|--------|
| air-combat-sim | ✅ **FIXED**: Duration default=300s, Pk=0.25, 4 missiles fired, 8 events |
| air-combat-sim | ✅ **FIXED**: F-22 fuel 6100→8200 kg |
| launch-veh-sim | ✅ **FIXED**: JSON tags (snake_case: max_q_pa, perigee_km, etc.) |
| launch-veh-sim | ✅ **FIXED**: Added -v, -i, -seed flags |
| launch-veh-sim | ✅ **FIXED**: F9 max Q reduced from 94→89 kPa (F9-specific pitch profile) |
| launch-veh-sim | ✅ **FIXED**: DragArea per vehicle (Atlas=4.0, Vulcan=4.5, Ariane=3.5, etc.) |
| launch-veh-sim | ⚠️ Max Q still 89 kPa vs real F9 ~33-40 kPa |
| tactical-net | ✅ **FIXED**: `-json` now starts with `{` (clean JSON) |
| bmd-sim-rcs | 🆕 NEW — Aspect-dependent radar cross section simulator |

---

## bmd-sim-rcs — NEW 🆕

13 threat profiles, 8 frequency bands, 5 aspect angles per band, continuous interpolation, detection range factor, discrimination factor, IR signature, stealth flag. 12 tests passing.

---

## Open Issues Summary

| # | Sim | Issue | Severity | Status |
|---|-----|-------|----------|--------|
| 1 | launch-veh-sim | Max Q 89 kPa (real F9 ~33-40 kPa) | P3 | Open — pitch profile improved, atmospheric model needs tuning |
| 2 | launch-veh-sim | `-duration` is float not Go Duration | P3 | Low priority |
| 3 | kill-chain-rt | Requires `-scenario` for JSON output | P3 | By design |

### Resolved ✅ (since R15)

| Issue | Round |
|-------|------|
| air-combat-sim Pk=0.00, no weapons | R16 ✅ |
| air-combat-sim events empty | R16 ✅ |
| air-combat-sim F-22 fuel 6100 kg | R16 ✅ |
| launch-veh-sim no `-i`, `-v`, `-seed` | R16 ✅ (flags added) |
| launch-veh-sim inconsistent JSON keys | R16 ✅ (snake_case tags) |
| tactical-net JSON after text | R16 ✅ |
| launch-veh-sim DragArea hardcoded | R16 ✅ (per-vehicle DragArea) |

---

## Vehicle Results (launch-veh-sim)

| Vehicle | Status | Max Q | Perigee | Apogee | Ecc |
|---------|--------|-------|---------|--------|-----|
| Falcon 9 | ✅ orbit | 89 kPa | 442 km | 500 km | 0.004 |
| F9 Expendable | ✅ orbit | 94 kPa | 672 km | 731 km | 0.004 |
| Falcon Heavy | ✅ orbit | 101 kPa | 873 km | 940 km | 0.004 |
| Starship | ✅ orbit* | 100 kPa | 231 km | 1928 km | 0.114 |
| New Glenn | ✅ orbit | 75 kPa | 379 km | 445 km | 0.005 |
| Vulcan | ✅ orbit | 115 kPa | 1696 km | 1773 km | 0.005 |
| Atlas V | ✅ orbit | 111 kPa | 1850 km | 1932 km | 0.005 |
| Delta IV | ✅ orbit** | 91 kPa | 150 km | 999 km | 0.061 |
| Ariane 6 | ✅ orbit | 123 kPa | 1932 km | 2014 km | 0.005 |
| Long March 5 | ✅ orbit | 105 kPa | 1040 km | 1114 km | 0.005 |

\* Starship: fuel-limited circularization (ecc 0.114 realistic)
\** Delta IV: single-burn only (no circularization)

---

**Zero stubs, mocks, or placeholder data.** ✅
**Zero format string bugs.** ✅