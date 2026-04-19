# FORGE-Sims Evaluation — Round 7 (Post-Fix Verification)

**Date:** 2026-04-19 (Round 7 — P0+P1 Bugs Fixed)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 37 | 37 | 0 | All clean |
| `-v` verbose flag | 37 | 34 | 3 | Missing: cyber-redteam, maritime, space-war (by design) |
| `-i` interactive mode | 37 | 34 | 3 | Missing: cyber-redteam, maritime, space-war (by design) |
| `-json` JSON output | 37 | 34 | 3 | Missing: cyber-redteam, maritime, space-war (by design) |
| `-duration` | 37 | 34 | 3 | Missing: cyber-redteam, maritime, space-war (by design) |
| `-seed` reproducibility | 37 | 34 | 3 | Missing: cyber-redteam, maritime, space-war (by design) |

**37 total simulators: 30 BMDS + 3 new space + 4 specialized**

---

## Bugs Fixed (Round 7)

### P0 — Missing Core Flags ✅ FIXED

| # | Binary | Issue | Fix |
|---|--------|-------|-----|
| 1 | **bmd-sim-gmd** | Missing `-v`, `-i`, `-json`, `-duration`, `-seed` | Rebuilt with sim-cli integration. All flags now work. |
| 2 | **bmd-sim-lrdr** | Missing `-v`, `-i`, `-json`, `-duration`, `-seed` | Rebuilt with sim-cli integration. All flags now work. |
| 3 | **bmd-sim-tpy2** | Missing `-i`, `-json`, `-duration`, `-seed` | Rewrote main.go to use sim-cli. All flags now work. |

### P1 — Broken JSON Output ✅ FIXED

| # | Binary | Issue | Fix |
|---|--------|-------|-----|
| 4 | **bmd-sim-aegis** | `-json` printed text+JSON mixed | Rewrote main.go with sim-cli. Clean JSON output. |
| 5 | **bmd-sim-gbr** | `-json` printed text only | Rewrote main.go with sim-cli. Clean JSON output. |
| 6 | **bmd-sim-patriot** | `-json` printed text+JSON mixed | Rewrote main.go with sim-cli. Clean JSON output. |
| 7 | **bmd-sim-thaad** | `-json` printed text+JSON mixed | Rewrote main.go with sim-cli. Clean JSON output. |
| 8 | **bmd-sim-hub** | `-json` printed text only | Rewrote main.go with sim-cli. Clean JSON output. |

---

## Flag Compatibility Matrix (Post-Fix)

| Sim | `-v` | `-i` | `-json` | `-duration` | `-seed` | Notes |
|-----|------|------|---------|-------------|---------|-------|
| air-traffic | ✅ | ✅ | ✅ | ✅ (int) | ✅ | |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ (int) | ✅ | |
| space-debris | ✅ | ✅ | ✅ | ✅ (int) | ✅ | |
| tactical-net | ✅ | ✅ | ✅ | ✅ (int) | ✅ | |
| bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-stss | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-hub | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | **FIXED** |
| bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | |
| cyber-redteam-sim | ❌ | ❌ | ❌ | ❌ | ❌ | Config-driven model (by design) |
| maritime-sim | ❌ | ❌ | ❌ | ❌ | ❌ | Config-driven model (by design) |
| space-war-sim | ❌ | ❌ | ❌ | ❌ | ❌ | Config-driven model (by design) |

---

## Remaining Issues

### P2 — Duration Format Inconsistency (Low Priority)

BMDS sims (via sim-cli) use Go `time.Duration` format: `-duration 5s`, `-duration 1m`
New sims use integer seconds: `-duration 5`

Both work correctly, but the UX is different. Should be unified in a future release.

### P3 — Specialized Sims (By Design, Document)

| Sim | CLI Model | Has Instead |
|-----|-----------|-------------|
| cyber-redteam-sim | Config-driven | `-config`, `-doctor`, `-init`, `-no-server`, `-rest-port`, `-version` |
| maritime-sim | Config-driven | `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version`, `-web` |
| space-war-sim | Config-driven | `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version` |

These are scenario-based simulators with web UIs and REST APIs. They don't fit the "run once, get output" pattern. **This is by design.**

---

## Previously Resolved Issues ✅

| Issue | Status |
|-------|--------|
| C2BMC nil pointer crash | ✅ Fixed |
| Interactive mode was stub | ✅ Fixed |
| Double-tick printing | ✅ Fixed |
| JSON was minimal | ✅ Fixed |
| DecoyType format string | ✅ Fixed |
| `-v` regression on 3 new sims | ✅ Fixed |
| 6 binaries missing new flags | ✅ Fixed |

---

## Source Repos Updated

All 8 fixed sims had their source repos updated on IDM (`idm.wezzel.com:crab-meat-repos/`):
- bmd-sim-gmd, bmd-sim-lrdr, bmd-sim-tpy2 — rewrote main.go with sim-cli
- bmd-sim-aegis, bmd-sim-gbr, bmd-sim-patriot, bmd-sim-thaad, bmd-sim-hub — rewrote main.go with sim-cli

Build environment: `/home/wez/{sim-cli,bmd-sim-core,bmd-sim-*}` on miner

---

## Sample Outputs (Post-Fix)

### bmd-sim-gmd -json (FIXED — was missing, now clean)
```json
{
  "simulator": "gmd",
  "status": "complete",
  "parameters": {
    "available_silos": 44,
    "burnout_ms": 5810.7,
    "ekv_divert_ms": 800,
    "ekv_lethal_m": 0.5,
    "ekv_seeker_km": 700,
    "max_alt_km": 1700,
    "pk_hgv": 0.03,
    "pk_icbm": 0.54,
    "pk_irbm": 0.54,
    "pk_slcm": 0,
    "site": "FTG",
    "variant": "CE-II GBI @ FTG"
  }
}
```

### bmd-sim-aegis -json (FIXED — was text+JSON, now clean)
```json
{
  "simulator": "aegis",
  "status": "complete",
  "parameters": {
    "array_faces": 4,
    "range_01m2_km": 266.24,
    "range_1m2_km": 473.46,
    "sm2_pk": 0.8,
    "sm2_range_km": 200,
    "sm3_pk": 0.85
  }
}
```

### bmd-sim-gmd -i (FIXED — was missing, now works)
```
GMD (Ground-Based Midcourse Defense) Simulator
=======================================================
  Mode: INTERACTIVE | CE-II GBI @ FTG | FTG

[T+1s] CE-II GBI @ FTG  Burnout: 5811 m/s  EKV Divert: 800 m/s
```