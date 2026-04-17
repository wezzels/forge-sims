# FORGE-Sims Evaluation — Final

**Date:** 2026-04-16 (Final Round)
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Overall Status

| Category | Total | Pass | Fail |
|----------|-------|------|------|
| Binaries execute (exit 0) | 30 | 30 | 0 |
| `-v` verbose flag | 30 | 30 | 0 |
| `-i` / `--interactive` flag | 30 | 30 | 0 |
| `-json` output flag | 30 | 30 | 0 |
| `-duration` / `-tick` / `-seed` | 30 | 30 | 0 |
| Interactive shows live metrics | 30 | **0** | 30 |
| JSON includes full sim data | 30 | **0** | 30 |
| Full TUI dashboard rendering | 30 | **0** | 30 |

---

## All 30 Sims — Confirmed Working

Every binary executes, accepts flags, and exits cleanly (exit 0).

| # | Binary | Default Run | `-v` | `-i` | `-json` | Special Flags |
|---|--------|-------------|------|------|---------|---------------|
| 1 | bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ | — |
| 2 | bmd-sim-stss | ✅ | ✅ | ✅ | ✅ | — |
| 3 | bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ | — |
| 4 | bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ | — |
| 5 | bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ | — |
| 6 | bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ | — |
| 7 | bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ | `-site`, `-scenario` |
| 8 | bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ | — |
| 9 | bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ | — |
| 10 | bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ | — |
| 11 | bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ | — |
| 12 | bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ | — |
| 13 | bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ | `-site`, `-variant` |
| 14 | bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ | `-variant` |
| 15 | bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ | — |
| 16 | bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ | `-mirvs`, `-cms` |
| 17 | bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ | — |
| 18 | bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ | `-maneuvers` |
| 19 | bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ | — |
| 20 | bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ | — ⚠️ format bug |
| 21 | bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | — |
| 22 | bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ | — |
| 23 | bmd-sim-hub | ✅ | ✅ | ✅ | ✅ | — |
| 24 | bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ | — |
| 25 | bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ | — |
| 26 | bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ | — |
| 27 | bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ | — |
| 28 | bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | — |
| 29 | bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ | — |
| 30 | bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ | — |

---

## C2BMC Crash — FIXED ✅

Was crashing with nil pointer dereference in `CorrelateTracks()`. Now runs clean.

---

## Corrections Needed

### P0 — Interactive Mode Is a Stub

All 30 sims accept `-i` but only output:

```
Mode: INTERACTIVE (ctrl+C to quit)
[T+1s] Running...
[T+2s] Running...
[T+3s] Running...
```

**Missing:**
- No live metrics (tracks, detections, Pk, coverage)
- No ANSI/TUI dashboard rendering
- No keyboard controls (pause/step/quit)
- Double-prints each tick

### P0 — JSON Output Is Minimal

All 30 sims return only:

```json
{"Simulator": "sbirs", "Status": "ready"}
```

Should include full sim data: parameters, tracks, detections, discrimination results, trajectory points, etc.

### P1 — Verbose Output Regressed

Several sims now show less detail with `-v` than before. Example:

**Old SBIRS `-v`:** Full detection ranges, clutter analysis, visibility per satellite
**New SBIRS `-v`:** `Seed: 0, Tick: 1s` — just framework info, no sim data

Same pattern in aegis, gbr, hub, patriot, thaad — `-v` now shows seed/tick instead of simulation details.

### P2 — Decoy Format Bug

```
%!s(decoy.DecoyType=0)  ← should be "BALLOON"
%!s(decoy.DecoyType=1)  ← should be "INFLATABLE_RV"
```

`DecoyType` enum missing `String()` method.

### P2 — Double Tick Printing

Interactive mode prints each tick twice:

```
[T+1s] Running...
[T+1s] Running...
[T+2s] Running...
[T+2s] Running...
```

Likely ticker channel + render loop both printing.

---

## What Needs To Happen

| Priority | Task | Affects |
|----------|------|---------|
| **P0** | Implement real TUI dashboard in `-i` mode | All 30 |
| **P0** | Populate `-json` with full sim data | All 30 |
| **P1** | Fix `-v` to show sim-specific details again | All 30 (regression) |
| **P2** | Fix DecoyType `String()` method | decoy |
| **P2** | Fix double tick printing in interactive | All 30 |
| **P3** | Add keyboard controls (pause/step/quit) | All 30 |