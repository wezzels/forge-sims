# FORGE-Sims Interactive Mode Evaluation

**Date:** 2026-04-16
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Interactive Mode (`-i` / `--interactive`) — MISSING ACROSS ALL SIMS

**Result: 0/30 binaries support interactive mode.**

Every simulator is a "run once, dump static output, exit" binary. No live dashboards, no updating metrics, no real-time scenario progression exist.

---

## Per-Binary Flag Audit

| Binary | `-i` | `--interactive` | `-v` | Other Flags | Current Behavior |
|--------|------|-----------------|------|-------------|------------------|
| bmd-sim-sbirs | ❌ | ❌ | ✅ | `-site`, `-scenario` | Static dump, exit |
| bmd-sim-stss | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-dsp | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-uewr | ❌ | ❌ | ✅ | `-scenario` | Static dump, exit |
| bmd-sim-lrdr | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-cobra-judy | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-gmd | ❌ | ❌ | ❌ | `-site`, `-variant` | Static dump, exit |
| bmd-sim-sm3 | ❌ | ❌ | ❌ | `-variant` | Static dump, exit |
| bmd-sim-sm6 | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-thaad | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-thaad-er | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-patriot | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-aegis | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-icbm | ❌ | ❌ | ❌ | `-mirvs`, `-cms` | Static dump, exit |
| bmd-sim-irbm | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-hgv | ❌ | ❌ | ❌ | `-maneuvers` | Static dump, exit |
| bmd-sim-slcm | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-decoy | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-jamming | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-link16 | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-jreap | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-jrsc | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-gfcb | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-ifxb | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-space-weather | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-atmospheric | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-gbr | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-hub | ❌ | ❌ | ❌ | none | Static dump, exit |
| bmd-sim-tpy2 | ❌ | ❌ | ✅ | `-site`, `-scenario` | Static dump, exit |
| bmd-sim-c2bmc | ❌ | ❌ | ❌ | none | **CRASHES** on default run |

---

## What Interactive Mode Should Provide

### Core Features (all sims)
- **Live metric dashboard** — track counts, detection rates, Pk, coverage %
- **Scenario clock** — T+0s, T+30s... progressing in real-time
- **ANSI terminal UI** — scrollable, color-coded panels (like `htop`/`lazygit`)
- **Keyboard controls** — `space` pause/resume, `s` step, `q` quit

### Sensor Sims (sbirs, stss, dsp, uewr, lrdr, tpy2, gbr, cobra-judy)
```
┌─ SBIRS Constellation ──────────────────────────────────┐
│ T+0:45:00  Coverage: 95%  Active Sats: 6/6             │
│                                                         │
│ GEO-1  ✅ DETECTED  Track-1  Pd=0.99  Alt=200km        │
│ GEO-2  ✅ TRACKING Track-1  Pd=0.99  Alt=199km        │
│ HEO-1  ⬜ SCANNING  —       Pd=n/a                     │
│                                                         │
│ Tracks: 3  Detections: 7  False Alarms: 0.02/s         │
│ Clutter: LOW  FAR: 0.01/s                              │
│                                                         │
│ [SPACE] Pause  [S] Step  [Q] Quit                      │
└─────────────────────────────────────────────────────────┘
```

### Interceptor Sims (gmd, sm3, sm6, thaad, thaad-er, patriot, aegis)
```
┌─ GMD CE-II @ FTG ──────────────────────────────────────┐
│ T+0:03:22  Phase: MIDCOURSE  Stages: 3/3              │
│                                                         │
│ Range to target: 2,450 km  Closing: -4,200 m/s        │
│ Altitude: 890 km  Speed: 6,800 m/s                    │
│ EKV Status: COASTING  Seeker: STANDBY                  │
│                                                         │
│ TTI: 287s  Pk: 0.90  Divert ΔV: 600 m/s              │
│                                                         │
│ [SPACE] Pause  [S] Step  [Q] Quit                      │
└─────────────────────────────────────────────────────────┘
```

### Threat Sims (icbm, irbm, hgv, slcm)
```
┌─ ICBM MIRV-5 ──────────────────────────────────────────┐
│ T+0:12:00  Phase: MIDCOURSE  RVs deployed: 5           │
│                                                         │
│ Alt: 1,200 km  Speed: 6,124 m/s  Lat: 37.0° Lon: 25° │
│ RV-1 → (55.0°, -28.4°)  ETA: 720s                     │
│ RV-2 → (55.9°, -29.5°)  ETA: 735s                     │
│ CM: 5 balloons, 1 chaff, 1 jammer                      │
│ Discrimination difficulty: 0.8                          │
│                                                         │
│ [SPACE] Pause  [S] Step  [Q] Quit                      │
└─────────────────────────────────────────────────────────┘
```

### C2 Sims (c2bmc, hub, gfcb, ifxb, jrsc, link16, jreap)
```
┌─ C2BMC Battle Management ──────────────────────────────┐
│ T+0:02:15  Tracks: 3  Correlated: 2  Engaged: 1       │
│                                                         │
│ Track-1: ICBM  CONFIRMED  Threat: HIGH  Engaged: GMD  │
│ Track-2: IRBM  PENDING    Threat: MED                  │
│ Track-3: UNKNOWN TENTATIVE Threat: LOW                  │
│                                                         │
│ Sensors: SBIRS✅ UEWR✅ TPY-2✅ LRDR⬜                  │
│ Weapons: GMD🔥 SM3 RDY  THAAD RDY                      │
│                                                         │
│ [SPACE] Pause  [S] Step  [Q] Quit                      │
└─────────────────────────────────────────────────────────┘
```

---

## Recommended Implementation Priority

| Priority | Sims | Rationale |
|----------|------|-----------|
| P0 | c2bmc, hub | C2 needs live track picture to be useful |
| P1 | sbirs, uewr, lrdr, tpy2 | Primary sensors — operators need live feed |
| P2 | gmd, sm3, thaad, aegis | Interceptors — flyout progression is key |
| P3 | icbm, hgv | Threats — trajectory/maneuver progression |
| P4 | All remaining | Environmental, comms, countermeasures |

### Suggested Tech Stack
- **Go library:** `charmbracelet/bubbletea` + `lipgloss` (TUI framework)
- **Alternative:** `gdamore/tcell` (lower-level, more control)
- **Shared package:** `pkg/tui/` with common dashboard components

### Suggested CLI
```
bmd-sim-sbirs -i              # Interactive mode (default refresh: 1s)
bmd-sim-sbirs -i -rate 5s     # Interactive, 5-second tick rate
bmd-sim-sbirs -i -scenario x.yaml  # Interactive with scenario file
```

---

## Additional Missing Flags

| Flag | Purpose | Priority |
|------|---------|----------|
| `--json` | Machine-readable output | High — enables piping/automation |
| `--duration 30s` | Run for N seconds then exit | Medium — useful for batch testing |
| `--tick 100ms` | Simulation speed | Medium — faster-than-realtime |
| `--output report.md` | Write results to file | Low |
| `--seed 42` | Deterministic random seed | Low — reproducibility |