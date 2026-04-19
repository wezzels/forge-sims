# FORGE-Sims Evaluation — Round 6 (Full Audit)

**Date:** 2026-04-18 (Round 6 — Comprehensive Binary Audit)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute (exit 0) | 37 | 37 | 0 | All clean |
| `-v` verbose flag | 37 | 33 | 4 | Missing: gmd, lrdr, cyber-redteam, maritime, space-war |
| `-i` interactive mode | 37 | 30 | 7 | Missing: gmd, lrdr, tpy2, cyber-redteam, maritime, space-war; 2 print text not JSON (gbr, patriot) |
| `-json` JSON output | 37 | 28 | 9 | Missing: gmd, lrdr, tpy2, cyber-redteam, maritime, space-war; 3 print text+JSON (aegis, gbr, hub, patriot, thaad) |
| `-duration`/`-tick` | 37 | 33 | 4 | Missing: gmd, lrdr, tpy2; BMDS uses `5s` format (not int) |
| `-seed` reproducibility | 37 | 33 | 4 | Missing: gmd, lrdr, cyber-redteam, maritime, space-war |

**37 total entries** (30 BMDS + 3 new space + 4 specialized sims)

---

## Flag Compatibility Matrix

| Sim | `-v` | `-i` | `-json` | `-duration` | `-seed` | Notes |
|-----|------|------|---------|-------------|---------|-------|
| air-traffic | ✅ | ✅ | ✅ | ✅ (int) | ✅ | Full |
| satellite-tracker | ✅ | ✅ | ✅ | ✅ (int) | ✅ | Full |
| space-debris | ✅ | ✅ | ✅ | ✅ (int) | ✅ | Full |
| tactical-net | ✅ | ✅ | ✅ | ✅ (int) | ✅ | Full |
| bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-stss | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-lrdr | ❌ | ❌ | ❌ | ❌ | ❌ | **Only `-help`** |
| bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-tpy2 | ✅ | ❌ | ❌ | ❌ | ❌ | **Missing `-i`, `-json`, `-duration`, `-seed`** |
| bmd-sim-gbr | ✅ | ✅ | ⚠️ | ✅ (dur) | ✅ | `-json` prints text+JSON mixed |
| bmd-sim-aegis | ✅ | ✅ | ⚠️ | ✅ (dur) | ✅ | `-json` prints text+JSON mixed |
| bmd-sim-patriot | ✅ | ✅ | ⚠️ | ✅ (dur) | ✅ | `-json` prints text+JSON mixed |
| bmd-sim-thaad | ✅ | ✅ | ⚠️ | ✅ (dur) | ✅ | `-json` prints text+JSON mixed |
| bmd-sim-hub | ✅ | ✅ | ⚠️ | ✅ (dur) | ✅ | `-json` prints text+JSON mixed |
| bmd-sim-gmd | ❌ | ❌ | ❌ | ❌ | ❌ | **Only `-site`, `-variant`, `-help`** |
| bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ (dur) | ✅ | Full |
| cyber-redteam-sim | ❌ | ❌ | ❌ | ❌ | ❌ | **Different CLI model** (config-based) |
| maritime-sim | ❌ | ❌ | ❌ | ❌ | ❌ | **Different CLI model** (config-based) |
| space-war-sim | ❌ | ❌ | ❌ | ❌ | ❌ | **Different CLI model** (config-based) |

---

## New Issues Found (Round 6)

### P0 — Broken / Missing Core Flags

| # | Binary | Issue | Detail |
|---|--------|-------|--------|
| 1 | **bmd-sim-gmd** | Missing all standard flags | Only has `-site`, `-variant`, `-help`. No `-v`, `-i`, `-json`, `-duration`, `-seed`. Completely different CLI from all other BMDS sims. |
| 2 | **bmd-sim-lrdr** | Missing all standard flags | Only has `-help`. No `-v`, `-i`, `-json`, `-duration`, `-seed`. Outputs text only. |
| 3 | **bmd-sim-tpy2** | Missing `-i`, `-json`, `-duration`, `-seed` | Has `-v` and `-site`, `-scenario`, but lacks interactive/JSON/duration. |

### P1 — Mixed Text+JSON Output (not valid JSON)

| # | Binary | Issue | Detail |
|---|--------|-------|--------|
| 4 | **bmd-sim-aegis** | `-json` prints text banner + JSON | First 6 lines are text, then JSON. Not valid JSON parse. |
| 5 | **bmd-sim-gbr** | `-json` prints text output, no JSON | `-json` flag accepted but output is plain text. |
| 6 | **bmd-sim-patriot** | `-json` prints text banner + JSON | Same as aegis — text header then partial JSON. |
| 7 | **bmd-sim-thaad** | `-json` prints text banner + JSON | Same issue — text header then data. |
| 8 | **bmd-sim-hub** | `-json` prints text, no JSON | Output is entirely text with no JSON structure. |

### P2 — Duration Format Inconsistency

| # | Issue | Detail |
|---|-------|--------|
| 9 | BMDS sims use Go `time.Duration` format | `-duration 5s` works but `-duration 5` fails with parse error. New sims (air-traffic, satellite-tracker, space-debris, tactical-net) use int seconds: `-duration 5`. This is a UX inconsistency. |

### P3 — Specialized Sims (Different CLI Model)

| # | Binary | Issue | Detail |
|---|--------|-------|--------|
| 10 | **cyber-redteam-sim** | No standard flags | Uses `-config`, `-doctor`, `-init`, `-no-server`, `-rest-port`, `-version`. No `-v`, `-i`, `-json`, `-duration`, `-seed`. |
| 11 | **maritime-sim** | No standard flags | Uses `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version`, `-web`. No standard flags. |
| 12 | **space-war-sim** | No standard flags | Uses `-aar`, `-config`, `-doctor`, `-init`, `-list-scenarios`, `-validate`, `-version`. No standard flags. |

**Note:** The specialized sims (cyber-redteam, maritime, space-war) use a fundamentally different CLI model — they're config-driven, scenario-based simulators with web UIs and REST APIs. They don't fit the simple "run once, get output" pattern. This is **by design** but should be documented.

---

## Previously Resolved Issues ✅

| Issue | Status |
|-------|--------|
| C2BMC nil pointer crash | ✅ Fixed |
| Interactive mode was stub (`Running...` only) | ✅ Fixed — shows live metrics |
| Double-tick printing | ✅ Fixed |
| JSON was minimal (`{simulator, status}`) | ✅ Fixed — includes parameters/data |
| DecoyType format string (`%!s(...)`) | ✅ Fixed — shows BALLOON/INFLATABLE_RV |
| `-v` regression on 3 new sims | ✅ Fixed |
| 6 binaries missing new flags | ✅ Fixed |

---

## Enhancement Suggestions (Not Bugs)

| Priority | Enhancement | Notes |
|----------|-------------|-------|
| P0 | Add standard flags to gmd, lrdr, tpy2 | These 3 BMDS sims are missing most standard flags |
| P0 | Fix `-json` on aegis, gbr, patriot, thaad, hub | Output should be valid JSON only when `-json` is set |
| P1 | Standardize `-duration` format | BMDS uses Go duration (`5s`), new sims use int seconds (`5`). Pick one. |
| P2 | Add `-v`, `-i`, `-json` to specialized sims | Or document them as a separate category with different CLI |
| P2 | Richer JSON for aegis, gbr, patriot, thaad | Currently text output even with `-json` |
| P3 | Scenario-driven dynamic metrics in `-i` | Interactive metrics are mostly static per tick |
| P3 | Keyboard controls in `-i` | pause/step/quit beyond ctrl+C |
| P3 | TUI/ANSI dashboard in `-i` | Currently line-by-line, not panel layout |

---

## Sample Outputs

### air-traffic -i
```
[T+0s] 100 flights
  COMMERCIAL      39    CARGO           17
  PRIVATE         14    MILITARY        30
  CRUISE     70   CLIMB      20
  DESCENT    8    APPROACH   2
```

### satellite-tracker -i
```
[T+0s] Tracking 30 satellites
  Visible: 30 | Eclipsed: 0
  MANNED               6
  OTHER                24
```

### space-debris -i
```
[T+0s] 5000 debris objects
  LEO        3030  GEO         510
  HIGH        953  MEO         507
```

### bmd-sim-gmd (current — flags only)
```
Usage of bmd-sim-gmd:
  -help           Show help
  -site string    Launch site (FTG or VAFB) (default "FTG")
  -variant string GBI variant (CE-I or CE-II) (default "CE-I")
```

### bmd-sim-lrdr (current — flags only)
```
Usage of bmd-sim-lrdr:
  -help           Show help
```

### cyber-redteam-sim (config-driven model)
```
-config string    scenario config file (default "configs/enterprise-ad.yaml")
-doctor           run health check diagnostics
-init             initialize config directory with sample scenarios
-no-server        disable API servers (run simulation only)
-rest-port int    REST API port (default 8080)
-version          print version and exit
```