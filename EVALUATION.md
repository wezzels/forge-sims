# FORGE-Sims Evaluation — Round 7 (Full Flag Audit)

**Date:** 2026-04-18 (Round 7 — Full P0/P1 Fix Audit)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 36 | 36 | 0 | All clean |
| `-v` verbose flag | 30 | 30 | 0 | ✅ All 30 BMDS sims |
| `-i` interactive mode | 30 | 30 | 0 | ✅ All 30 BMDS sims |
| `-json` JSON output (clean) | 30 | 30 | 0 | ✅ All produce valid JSON only |
| `-duration`/`-tick` | 30 | 30 | 0 | ✅ All accept duration |
| `-seed` reproducibility | 30 | 30 | 0 | ✅ All accept seed |

**36 total entries** (30 BMDS + 3 space + 3 specialized)

---

## All 30 BMDS Sims — Full Flag Compliance ✅

Every BMDS simulator now produces **clean JSON only** (no text banner) when `-json` is specified, and supports all standard flags (`-v`, `-i`, `-json`, `-duration`, `-seed`) via `sim-cli`.

| Sim | `-v` | `-i` | `-json` | `-duration` | `-seed` | Fixed In |
|-----|------|------|---------|-------------|---------|----------|
| sbirs | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| stss | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| dsp | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| uewr | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| lrdr | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P0) |
| cobra-judy | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| tpy2 | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P0) |
| gbr | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P1) |
| aegis | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P1) |
| patriot | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P1) |
| thaad | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P1) |
| hub | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P1) |
| gmd | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 (was P0) |
| sm3 | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| sm6 | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| icbm | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| irbm | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| hgv | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| slcm | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| decoy | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| jamming | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| c2bmc | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| gfcb | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| ifxb | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| jrsc | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| link16 | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| jreap | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| space-weather | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| atmospheric | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |
| thaad-er | ✅ | ✅ | ✅ | ✅ | ✅ | Round 7 |

### New Sims (already compliant)

| Sim | `-v` | `-i` | `-json` | `-duration` | `-seed` |
|-----|------|------|---------|-------------|---------|
| air-traffic | ✅ | ✅ | ✅ | ✅ (int) | ✅ |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ (int) | ✅ |
| space-debris | ✅ | ✅ | ✅ | ✅ (int) | ✅ |
| tactical-net | ✅ | ✅ | ✅ | ✅ (int) | ✅ |
| wta | ✅ | ✅ | ✅ | ✅ (int) | ✅ |
| kill-assessment | ✅ | ✅ | ✅ | ✅ (int) | ✅ |

### Specialized Sims (config-based, by design)

| Sim | Notes |
|-----|-------|
| cyber-redteam | Uses `-config`, `-doctor`, `-init`, `-no-server`, `-rest-port`, `-version` |
| maritime | Uses `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version`, `-web` |
| space-war | Uses `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version` |

---

## Bug Fixes Applied in Round 7

| # | Sim | Fix |
|---|-----|-----|
| 1 | **gmd** | Added sim-cli integration (was missing all standard flags) |
| 2 | **lrdr** | Moved JSON check before text output (was printing banner then JSON) |
| 3 | **tpy2** | Replaced hand-rolled flags with sim-cli (was missing -i, -json, -duration, -seed) |
| 4 | **aegis** | Added sim-cli, clean JSON output |
| 5 | **gbr** | Added sim-cli, clean JSON output |
| 6 | **patriot** | Added sim-cli, clean JSON output |
| 7 | **thaad** | Added sim-cli, clean JSON output |
| 8 | **hub** | Added sim-cli, clean JSON output |
| 9 | **c2bmc** | Fixed CorrelateTracks nil pointer crash (collect-then-merge pattern) + sim-cli |
| 10-30 | **All 22 remaining** | Added sim-cli integration, clean JSON output, all standard flags |

---

## Previously Resolved Issues ✅

| Issue | Status |
|-------|--------|
| C2BMC nil pointer crash | ✅ Fixed (again — collect-then-merge) |
| Interactive mode was stub | ✅ Fixed — shows live metrics |
| Double-tick printing | ✅ Fixed |
| JSON was minimal | ✅ Fixed — includes parameters/data |
| DecoyType format string | ✅ Fixed |
| 6 binaries missing new flags | ✅ Fixed |
| All 8 P0/P1 issues from Round 6 | ✅ Fixed |

---

## Remaining Notes

| Priority | Item | Notes |
|----------|------|-------|
| P2 | Duration format inconsistency | BMDS sims use Go duration (`5s`), new sims use int seconds (`5`). Minor UX issue. |
| P3 | Specialized sims different CLI | cyber-redteam, maritime, space-war use config-driven CLI — by design |
| P3 | mrbm not built | No source repo exists yet |
| P3 | Interactive metrics mostly static | Most sims show constant values in -i mode |