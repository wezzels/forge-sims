# FORGE-Sims Full Test & Evaluation

**Date:** 2026-04-16
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Summary

| Category | Total | Pass | Fail |
|----------|-------|------|------|
| Binaries execute | 30 | 29 | 1 |
| Scenario flags work | 8 tested | 8 | 0 |
| `-i` / `--interactive` support | 30 | **0** | 30 |

---

## Binary-by-Binary Results

### ❌ CRITICAL: C2BMC Crash

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x48d3b1]

github.com/forge-c2/bmd-sim-c2bmc/pkg/c2bmc.(*C2BMC).CorrelateTracks()
    /home/wez/bmd-sim-c2bmc/pkg/c2bmc/c2bmc.go:117 +0x1d1
```

Crashes every time during track correlation. Processes 3 sensor reports then segfaults. **Exit code 2.**

### ✅ Working (29/30) — Full Confirmation

| # | Binary | Exit | Flags Tested | Output Confirmed |
|---|--------|------|-------------|------------------|
| 1 | bmd-sim-sbirs | 0 | `-site 3 -v` | ✅ Constellation, detection ranges, clutter, visibility |
| 2 | bmd-sim-stss | 0 | (no flags) | ✅ LEO dual-band, discrimination, 2-sat constellation |
| 3 | bmd-sim-dsp | 0 | (no flags) | ✅ 3 GEO sats, spin rate, single-band IR |
| 4 | bmd-sim-uewr | 0 | `-v` | ✅ 4 radar sites, UHF, detection ranges, track correlation |
| 5 | bmd-sim-lrdr | 0 | (no flags) | ✅ S-band 500kW, 1000+ tracks, discrimination |
| 6 | bmd-sim-cobra-judy | 0 | (no flags) | ✅ S-band ship radar |
| 7 | bmd-sim-gmd | 0 | `-site VAFB -variant CE-II` | ✅ Flyout sim, EKV guidance, homing test, 44 silos |
| 8 | bmd-sim-sm3 | 0 | `-variant IB` and `-variant IIA` | ✅ IA/IB/IIA comparison, LEAP kill vehicle |
| 9 | bmd-sim-sm6 | 0 | (no flags) | ✅ 370km range, 33km ceiling |
| 10 | bmd-sim-thaad | 0 | (no flags) | ✅ 2 launchers, 6 missiles, PK 90% |
| 11 | bmd-sim-thaad-er | 0 | (no flags) | ✅ 600km range, 200km ceiling, 600 m/s divert |
| 12 | bmd-sim-patriot | 0 | (no flags) | ✅ PAC-2/PAC-3, 8 missiles, engagement sector |
| 13 | bmd-sim-aegis | 0 | (no flags) | ✅ SPY-1, fire control solution, SM-2/3/6 models |
| 14 | bmd-sim-icbm | 0 | `-mirvs 5 -cms` | ✅ 5 MIRVs, balloons+chaff+jammers, discrimination 0.8 |
| 15 | bmd-sim-irbm | 0 | (no flags) | ✅ Runs clean |
| 16 | bmd-sim-hgv | 0 | `-maneuvers 8` | ✅ 8 maneuvers, PULL-UP/DIVE/LATERAL/TERMINAL, detection assessment |
| 17 | bmd-sim-slcm | 0 | (no flags) | ✅ Mach 0.8, 1500km, sea-skimming |
| 18 | bmd-sim-decoy | 0 | (no flags) | ✅ Balloon, RCS/IR/difficulty metrics |
| 19 | bmd-sim-jamming | 0 | (no flags) | ✅ Noise jammer, JNR, DENIED effect |
| 20 | bmd-sim-link16 | 0 | (no flags) | ✅ TDMA, BMDS-NET-1, latency 6ms |
| 21 | bmd-sim-jreap | 0 | (no flags) | ✅ JREAP-C, latency/bandwidth |
| 22 | bmd-sim-jrsc | 0 | (no flags) | ✅ INDO-PACIFIC regional |
| 23 | bmd-sim-gfcb | 0 | (no flags) | ✅ GFCB-1, PACIFIC |
| 24 | bmd-sim-ifxb | 0 | (no flags) | ✅ IFXB-1 active |
| 25 | bmd-sim-space-weather | 0 | (no flags) | ✅ Kp=7.0, SEVERE_DEGRADATION |
| 26 | bmd-sim-atmospheric | 0 | (no flags) | ✅ Rain/wind/radar degradation |
| 27 | bmd-sim-gbr | 0 | (no flags) | ✅ Acquisition scan, discrimination, confidence 90% |
| 28 | bmd-sim-hub | 0 | (no flags) | ✅ Track fusion, engagement sim, 4-event scenario |
| 29 | bmd-sim-tpy2 | 0 | `-site 2 -v` | ✅ Track confirmation, sector coverage 300% |

---

## Interactive Mode (`-i` / `--interactive`)

**None of the 30 binaries support `-i` or `--interactive`.** All binaries are "run once, dump output, exit" — no live/dynamic mode exists.

This is a gap for operational use. Interactive mode should provide:
- Live updating metrics (track counts, detection rates, Pk updates)
- Scrollable stats dashboard (similar to `htop`)
- Real-time scenario progression with timestamps
- Keyboard controls (pause/resume, step, quit)

---

## `-v` Flag Inconsistency

| Has `-v` | Missing `-v` |
|----------|-------------|
| sbirs, tpy2, uewr, dsp, gbr, patriot, aegis, hub | stss, lrdr, irbm, hgv, c2bmc, gmd, icbm, sm3, sm6, thaad, thaad-er, decoy, jamming, link16, jreap, jrsc, gfcb, ifxb, space-weather, atmospheric, cobra-judy, slcm |

Only 7/30 have `-v`. The rest print their output regardless or only show flags with `-help`.

---

## Physics & Realism Spot-Check

| Check | Result | Notes |
|-------|--------|-------|
| GEO altitude (SBIRS) | ✅ 42,164 km | Correct |
| HEO orbit (SBIRS) | ✅ ~23,560 km | Molniya orbit |
| ICBM flight time | ✅ 1800s (30 min) | Realistic |
| ICBM reentry speed | ✅ Mach 21 | Consistent |
| SM-3 IIA range/speed | ✅ 2500km / Mach 13 | Matches public data |
| SM-3 IB range/speed | ✅ 700km / Mach 12 | Good |
| UEWR frequency | ✅ 435 MHz UHF | Correct |
| LRDR S-band | ✅ 3.0 GHz | Correct |
| GMD EKV seeker range | ✅ 500 km | Reasonable |
| GMD divert ΔV | ✅ 600 m/s | Consistent |
| HGV detection by SBIRS | ✅ DETECTABLE | Boost phase |
| HGV detection by UEWR | ✅ NOT DETECTABLE | Below radar horizon |
| TPY-2 X-band | ⚠️ Not explicitly stated | Should be ~10 GHz |
| Patriot detection range | ⚠️ 2505km for 0.1m² | Optimistic |
| Aegis detection range | ⚠️ 31605km for 0.1m² | Very high, likely uncorrected |

---

## Recommendations

### Critical
1. **Fix C2BMC nil pointer crash** — `CorrelateTracks()` at `c2bmc.go:117`

### High
2. **Add `-i` / `--interactive` mode** to all simulators — live dashboard with:
   - Real-time metrics (tracks, detections, Pk, coverage)
   - Scenario time progression
   - Pause/resume/step controls
   - ANSI terminal UI (similar to `htop`/`lazygit`)

### Medium
3. **Standardize `-v` flag** across all 30 binaries
4. **Review Aegis/Patriot detection ranges** — values seem unreasonably high

### Low
5. **Verify TPY-2 frequency band** — should explicitly show X-band (~10 GHz)
6. **Add `--json` output flag** for machine-readable results
7. **Add `--duration` flag** for timed scenario runs