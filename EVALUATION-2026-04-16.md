# FORGE-Sims Full Test & Evaluation

**Date:** 2026-04-16
**Tester:** Wez (AI)
**Repo:** github.com/wezzels/forge-sims

---

## Summary

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute | 30 | 29 | 1 | C2BMC crash |
| `-v` flag support | 30 | 26 | 4 | See below |
| Scenario flags | ~8 | 8 | 0 | icbm, gmd, sm3, hgv, sbirs all work |

---

## Binary-by-Binary Results

### ✅ Working (29/30)

| Binary | `-v` | Output | Notes |
|--------|------|--------|-------|
| bmd-sim-sbirs | ✅ | Full constellation output | 4 GEO + 2 HEO, detection ranges, clutter analysis |
| bmd-sim-stss | ❌ | `-v` not defined | Works without flag. LEO dual-band, discrimination |
| bmd-sim-dsp | ✅ | 3 GEO satellites | Legacy satellite sim, 30-90s latency |
| bmd-sim-uewr | ✅ | 4 radar sites | UHF phased array, detection ranges |
| bmd-sim-lrdr | ❌ | `-v` not defined | S-band, 500kW, 1000+ tracks, discrimination |
| bmd-sim-cobra-judy | ✅ | S-band ship radar | Compact output |
| bmd-sim-gmd | ✅ | CE-II from FTG | 3-stage EKV, flyout simulation |
| bmd-sim-sm3 | ✅ | IIA variant | Range 2500km, Mach 13, PK 0.90 |
| bmd-sim-sm6 | ✅ | Endo-atmospheric | Range 370km, ceiling 33km |
| bmd-sim-thaad | ✅ | Battery sim | 2 launchers, 6 missiles, PK 90% |
| bmd-sim-thaad-er | ✅ | Extended range | 600km range, 200km ceiling |
| bmd-sim-patriot | ✅ | PAC-2/PAC-3 | 8 missiles, engagement sector |
| bmd-sim-aegis | ✅ | SPY-1 radar | Naval tracks, fire control solution |
| bmd-sim-icbm | ✅ | MIRV + countermeasures | 5 RVs, discrimination difficulty 0.8 |
| bmd-sim-irbm | ❌ | `-v` not defined | Works without flag |
| bmd-sim-hgv | ❌ | `-v` not defined | Mach 5-25, maneuver scheduling |
| bmd-sim-slcm | ✅ | Sea-skimming | Mach 0.8, 1500km range |
| bmd-sim-decoy | ✅ | Balloon/pen aids | RCS, IR, difficulty metrics |
| bmd-sim-jamming | ✅ | EW sim | Noise/deception jamming |
| bmd-sim-link16 | ✅ | TDMA network | BMDS-NET-1, latency 6ms |
| bmd-sim-jreap | ✅ | Protocol sim | JREAP-C, bandwidth 32kbps |
| bmd-sim-jrsc | ✅ | Regional security | INDO-PACIFIC feeds |
| bmd-sim-gfcb | ✅ | Fire control | GFCB-1, PACIFIC region |
| bmd-sim-ifxb | ✅ | Integrated fire | IFXB-1 active |
| bmd-sim-space-weather | ✅ | Solar effects | Kp=7.0, SEVERE_DEGRADATION |
| bmd-sim-atmospheric | ✅ | Propagation | Rain/wind degradation |
| bmd-sim-gbr | ✅ | Ground-based radar | Type 0, confidence 90% |
| bmd-sim-hub | ✅ | Scenario hub | TBM Intercept Demo, 4 events |
| bmd-sim-tpy2 | ✅ | X-band radar | Track confirmation, 300% sector coverage |

### ❌ Failed (1/30)

| Binary | Error | Details |
|--------|-------|---------|
| **bmd-sim-c2bmc** | **Panic: nil pointer dereference** | Crashes at `c2bmc.go:117` in `CorrelateTracks()` during track correlation. Segfault after processing 3 sensor reports. |

---

## C2BMC Crash Details

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x48d3b1]

goroutine 1 [running]:
github.com/forge-c2/bmd-sim-c2bmc/pkg/c2bmc.(*C2BMC).CorrelateTracks(0xc000119de0, 0x3ff0000000000000)
    /home/wez/bmd-sim-c2bmc/pkg/c2bmc/c2bmc.go:117 +0x1d1
main.main()
    /home/wez/bmd-sim-c2bmc/cmd/c2bmc/main.go:44 +0x61a
```

**Root cause:** Nil pointer when correlating 3 tracks from different sensors. Likely a missing nil check in `CorrelateTracks()` at line 117 of `c2bmc.go`. The function receives a correlation threshold (0x3ff... = 1.0 float64) and crashes accessing a track property.

---

## `-v` Flag Issues

4 binaries don't support `-v`:
- `bmd-sim-stss` — runs fine without it
- `bmd-sim-lrdr` — runs fine without it  
- `bmd-sim-irbm` — runs fine without it
- `bmd-sim-hgv` — runs fine without it

These exit with code 0 even with `-v`, just print usage and continue. Minor inconsistency, not a blocker.

---

## Physics & Realism Spot-Check

| Check | Result | Notes |
|-------|--------|-------|
| GEO altitude (SBIRS) | ✅ 42,164 km | Correct for geostationary orbit |
| HEO altitude (SBIRS) | ✅ ~23,560 km | Consistent with Molniya orbit |
| ICBM flight time | ✅ 1800s (30 min) | Realistic for 10,000 km |
| ICBM reentry speed | ✅ Mach 21 | Consistent (~7000 m/s) |
| SM-3 IIA range | ✅ 2500 km | Matches public data |
| UEWR frequency | ✅ 435 MHz | Correct UHF band |
| LRDR S-band | ✅ 3.0 GHz | Correct |
| TPY-2 X-band | ⚠️ Not explicitly stated | Should be ~10 GHz (X-band) |
| Patriot detection range | ⚠️ 2505 km for 0.1m² | Seems high for Patriot, but mode 0 may be search |

---

## Recommendations

1. **CRITICAL: Fix C2BMC nil pointer crash** — `CorrelateTracks()` needs nil check at `c2bmc.go:117`
2. **Low: Add `-v` flag to stss, lrdr, irbm, hgv** — for consistency
3. **Low: Verify TPY-2 frequency band** — should be X-band (~10 GHz)
4. **Low: Review Patriot detection range** — 2505km seems optimistic for 0.1m² RCS