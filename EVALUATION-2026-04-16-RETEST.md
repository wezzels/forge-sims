# FORGE-Sims Re-Test & Evaluation

**Date:** 2026-04-16 (Round 2)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Summary

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute | 30 | 30 | 0 | C2BMC crash **FIXED** ✅ |
| `-v` flag | 24 new + 2 old | 26 | 4 | aegis, gbr, hub, patriot, thaad, tpy2 still old |
| `-i` / `--interactive` | 24 new | 24 | 0 | Works but **minimal output** |
| `-json` | 24 new | 24 | 0 | Works but **minimal output** |
| `-seed`, `-duration`, `-tick` | 24 new | 24 | 0 | All work |
| Full TUI dashboard | 30 | **0** | 30 | Only prints `[T+Ns] Running...` |

---

## C2BMC Crash — FIXED ✅

Previously crashed with nil pointer dereference in `CorrelateTracks()`. Now runs clean, exit 0.

```
C2BMC (Command & Control Battle Management) Simulator
=====================================================
  Multi-sensor fusion, engagement planning

Simulator ready.
EXIT: 0
```

---

## Flag Coverage

### Updated Sims (24/30) — New Flags

All have: `-v`, `-i`/`-interactive`, `-json`, `-duration`, `-tick`, `-seed`, `-h`/`-help`

| # | Binary | `-v` | `-i` | `-json` | Special Flags | Notes |
|---|--------|------|------|---------|---------------|-------|
| 1 | bmd-sim-atmospheric | ✅ | ✅ | ✅ | — | |
| 2 | bmd-sim-c2bmc | ✅ | ✅ | ✅ | — | Crash fixed |
| 3 | bmd-sim-cobra-judy | ✅ | ✅ | ✅ | — | |
| 4 | bmd-sim-decoy | ✅ | ✅ | ✅ | — | ⚠️ Format bug |
| 5 | bmd-sim-dsp | ✅ | ✅ | ✅ | — | |
| 6 | bmd-sim-gfcb | ✅ | ✅ | ✅ | — | |
| 7 | bmd-sim-gmd | ✅ | ✅ | ✅ | `-site`, `-variant` | |
| 8 | bmd-sim-hgv | ✅ | ✅ | ✅ | `-maneuvers` | |
| 9 | bmd-sim-icbm | ✅ | ✅ | ✅ | `-mirvs`, `-cms` | |
| 10 | bmd-sim-ifxb | ✅ | ✅ | ✅ | — | |
| 11 | bmd-sim-irbm | ✅ | ✅ | ✅ | — | |
| 12 | bmd-sim-jamming | ✅ | ✅ | ✅ | — | |
| 13 | bmd-sim-jreap | ✅ | ✅ | ✅ | — | |
| 14 | bmd-sim-jrsc | ✅ | ✅ | ✅ | — | |
| 15 | bmd-sim-link16 | ✅ | ✅ | ✅ | — | |
| 16 | bmd-sim-lrdr | ✅ | ✅ | ✅ | — | |
| 17 | bmd-sim-sbirs | ✅ | ✅ | ✅ | — | |
| 18 | bmd-sim-slcm | ✅ | ✅ | ✅ | — | |
| 19 | bmd-sim-sm3 | ✅ | ✅ | ✅ | `-variant` | |
| 20 | bmd-sim-sm6 | ✅ | ✅ | ✅ | — | |
| 21 | bmd-sim-space-weather | ✅ | ✅ | ✅ | — | |
| 22 | bmd-sim-stss | ✅ | ✅ | ✅ | — | |
| 23 | bmd-sim-thaad-er | ✅ | ✅ | ✅ | — | |
| 24 | bmd-sim-uewr | ✅ | ✅ | ✅ | — | |

### Not Yet Updated (6/30) — Old Binaries

| # | Binary | Has `-i` | Has `-v` | Has `-json` | Status |
|---|--------|-----------|----------|-------------|--------|
| 1 | bmd-sim-aegis | ❌ | ❌ | ❌ | Old binary, still works |
| 2 | bmd-sim-gbr | ❌ | ❌ | ❌ | Old binary, still works |
| 3 | bmd-sim-hub | ❌ | ❌ | ❌ | Old binary, still works |
| 4 | bmd-sim-patriot | ❌ | ❌ | ❌ | Old binary, still works |
| 5 | bmd-sim-thaad | ❌ | ❌ | ❌ | Old binary, still works |
| 6 | bmd-sim-tpy2 | ❌ | ✅ | ❌ | Old binary, `-v` only |

---

## Interactive Mode — Functional But Minimal

All 24 updated sims support `-i` but only display:

```
SBIRS Satellite Constellation Simulator
=======================================
  Mode: INTERACTIVE (ctrl+C to quit)
  [T+1s] Running...
  [T+2s] Running...
  [T+3s] Running...
  Duration reached.
```

### Issues Found

1. **No live metrics** — just `[T+Ns] Running...` instead of updating dashboard
2. **Double printing** — each tick prints twice (likely ticker + render loop both logging)
3. **No ANSI/TUI rendering** — no panels, colors, or layout
4. **No keyboard controls** — only ctrl+C to quit, no pause/step

### Expected vs Actual

| Feature | Expected | Actual |
|---------|----------|--------|
| Live metrics | Track counts, detection rates, Pk | ❌ Just "Running..." |
| Scenario clock | T+0:00:00 progressing | ⚠️ T+Ns only |
| ANSI panels | Color-coded dashboard | ❌ Plain text |
| Keyboard | space/step/quit | ❌ Only ctrl+C |
| Per-sim data | Sensor readings, trajectory, etc. | ❌ Generic |

---

## JSON Output — Functional But Minimal

```json
{
  "Simulator": "sbirs",
  "Status": "ready"
}
```

Only outputs simulator name and status. Should include all simulation data (parameters, tracks, detections, results).

---

## Bugs Found

### ⚠️ bmd-sim-decoy Format String Bug

```
#1: %!s(decoy.DecoyType=0) RCS=0.10 IR=5000 diff=0.40
#2: %!s(decoy.DecoyType=1) RCS=0.10 IR=1000 diff=0.80
```

`DecoyType` enum doesn't implement `String()` — Go format verb `%s` fails, showing raw type value instead of name (should be "BALLOON", "INFLATABLE_RV", etc.)

---

## Corrections Needed

| Priority | Issue | Binary | Fix |
|----------|-------|--------|-----|
| **P0** | Interactive mode only shows "Running..." | All 24 new | Implement actual TUI dashboard with live metrics |
| **P0** | Double tick printing | All 24 new | Fix render loop to only print once per tick |
| **P1** | JSON output is minimal | All 24 new | Include full sim data in JSON output |
| **P1** | 6 binaries not rebuilt | aegis, gbr, hub, patriot, thaad, tpy2 | Rebuild with new flags |
| **P2** | DecoyType format string | decoy | Add `String()` method to DecoyType enum |
| **P2** | No keyboard controls in -i | All 24 new | Add space/step/quit handling |
| **P3** | SBIRS `-v` less detailed | sbirs | Old verbose had detection ranges/clutter, now minimal |