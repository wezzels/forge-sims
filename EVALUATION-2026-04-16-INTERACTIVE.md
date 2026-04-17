# FORGE-Sims Interactive Mode Evaluation

**Date:** 2026-04-16 (Updated)
**Tester:** Wez (AI)

---

## Status: ✅ ALL 30 BINARIES SUPPORT INTERACTIVE MODE

All simulators now support `-i` / `--interactive` via the shared `sim-cli` library.

---

## Per-Binary Flag Audit (Updated)

| Binary | `-i` | `-v` | `--json` | `--duration` | `--tick` | `--seed` | Custom Flags |
|--------|------|------|----------|-------------|----------|----------|-------------|
| bmd-sim-sbirs | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-site`, `-scenario` |
| bmd-sim-stss | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-dsp | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-uewr | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-scenario` |
| bmd-sim-lrdr | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-cobra-judy | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-gmd | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-site`, `-variant` |
| bmd-sim-sm3 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-variant` |
| bmd-sim-sm6 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-thaad-er | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-icbm | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-mirvs`, `-cms` |
| bmd-sim-irbm | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-hgv | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-maneuvers` |
| bmd-sim-slcm | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-decoy | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-jamming | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-c2bmc | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-gfcb | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-ifxb | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-jrsc | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-link16 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-jreap | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-space-weather | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-atmospheric | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-hub | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | — |
| bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | `-site`, `-scenario` |

---

## Interactive Mode Current Behavior

```
$ bmd-sim-c2bmc -i -duration 5s -tick 1s
C2BMC (Command & Control Battle Management) Simulator
=====================================================
  Mode: INTERACTIVE (ctrl+C to quit)
  [T+1s] Running...
  [T+2s] Running...
  [T+3s] Running...
  [T+4s] Running...
  [T+5s] Running...
  Duration reached.
```

---

## Next Step: Full TUI Dashboard

Current `-i` mode is a simple ticker. For operational use, implement:

- **charmbracelet/bubbletea** TUI framework
- Scrollable panels per sim type (sensor, interceptor, threat, C2)
- Real-time metrics (tracks, detections, Pk, coverage)
- Keyboard controls (pause/resume/step/quit)
- Color-coded status indicators

### Implementation Priority
| Priority | Sims | Rationale |
|----------|------|-----------|
| P0 | c2bmc, hub | C2 needs live track picture |
| P1 | sbirs, uewr, lrdr, tpy2 | Primary sensors |
| P2 | gmd, sm3, thaad, aegis | Interceptors |
| P3 | icbm, hgv | Threats |
| P4 | All remaining | Environmental/comms |