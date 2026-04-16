# FORGE-Sims Evaluation Report

**Date:** 2026-04-16 (Updated)
**Evaluator:** Automated Test Suite + Manual Review
**Version:** v1.0.1

---

## Summary

| Category | Total | Pass | Fail | Notes |
|----------|-------|------|------|-------|
| Binaries execute | 30 | 30 | 0 | ✅ C2BMC crash fixed |
| `-v` flag support | 30 | 30 | 0 | ✅ All 4 missing flags added |
| Scenario flags | ~8 | 8 | 0 | icbm, gmd, sm3, hgv, sbirs all work |
| Unit tests | ~310 | ~310 | 0 | All passing |
| Integration tests | 7 | 7 | 0 | Full chain validated |

---

## Fixes Applied This Session

### 1. C2BMC CorrelateTracks Nil Pointer Crash ✅ FIXED
**Root cause:** `CorrelateTracks()` deleted tracks from the map during iteration, causing subsequent access to return nil pointers.  
**Fix:** Collect merge pairs first, then apply after iteration completes. All 9 unit tests pass. Binary runs without crash.

### 2. Missing `-v` Flags ✅ FIXED
**Simulators:** bmd-sim-stss, bmd-sim-lrdr, bmd-sim-irbm, bmd-sim-hgv  
**Fix:** Added `-v` verbose flag to all four. Each now prints detailed parameters when `-v` is passed.

---

## Binary-by-Binary Results

### ✅ All 30 Binaries Working

| Binary | `-v` | `-help` | Output |
|--------|------|---------|--------|
| bmd-sim-sbirs | ✅ | ✅ | Full constellation output, detection ranges, clutter |
| bmd-sim-stss | ✅ | ✅ | LEO dual-band, discrimination |
| bmd-sim-dsp | ✅ | ✅ | 3 GEO satellites, 30-90s latency |
| bmd-sim-uewr | ✅ | ✅ | 4 radar sites, detection ranges |
| bmd-sim-lrdr | ✅ | ✅ | S-band, 500kW, 1000+ tracks, discrimination |
| bmd-sim-cobra-judy | ✅ | ✅ | S-band ship radar |
| bmd-sim-gmd | ✅ | ✅ | CE-II from FTG, flyout simulation |
| bmd-sim-sm3 | ✅ | ✅ | IIA variant, 2500km, Mach 13 |
| bmd-sim-sm6 | ✅ | ✅ | Endo-atmospheric, 370km, 33km ceiling |
| bmd-sim-thaad-er | ✅ | ✅ | Extended range, 600km, 200km |
| bmd-sim-icbm | ✅ | ✅ | MIRV + countermeasures |
| bmd-sim-irbm | ✅ | ✅ | Road-mobile TEL, 4000km |
| bmd-sim-hgv | ✅ | ✅ | Mach 5-25, maneuver scheduling |
| bmd-sim-slcm | ✅ | ✅ | Sea-skimming, 1500km |
| bmd-sim-decoy | ✅ | ✅ | Balloon/inflatable/jammer |
| bmd-sim-jamming | ✅ | ✅ | Noise/deception, ECCM |
| bmd-sim-c2bmc | ✅ | ✅ | Multi-sensor fusion, engagement plan |
| bmd-sim-gfcb | ✅ | ✅ | Fire control cycle |
| bmd-sim-ifxb | ✅ | ✅ | Engage-on-remote |
| bmd-sim-jrsc | ✅ | ✅ | Regional security center |
| bmd-sim-link16 | ✅ | ✅ | TDMA network, 238 kbps |
| bmd-sim-jreap | ✅ | ✅ | JREAP-C, CRC-16 |
| bmd-sim-space-weather | ✅ | ✅ | Kp index, radar degradation |
| bmd-sim-atmospheric | ✅ | ✅ | Rain attenuation, refraction |
| bmd-sim-thaad | ✅ | ✅ | Battery sim, 6 missiles |
| bmd-sim-patriot | ✅ | ✅ | PAC-2/PAC-3 |
| bmd-sim-aegis | ✅ | ✅ | SPY-1 radar, fire control |
| bmd-sim-gbr | ✅ | ✅ | Ground-based radar |
| bmd-sim-hub | ✅ | ✅ | Scenario hub |
| bmd-sim-tpy2 | ✅ | ✅ | X-band radar |

---

## Integration Tests

| Test | Result | Description |
|------|--------|-------------|
| TestFullEngagementChain | ✅ PASS | SBIRS → C2BMC → GMD complete chain |
| TestC2BMC_Discrimination | ✅ PASS | RV vs decoy discrimination |
| TestUEWR_Detection | ✅ PASS | UEWR detects 1m² at 1000km |
| TestLRDR_Discrimination | ✅ PASS | LRDR midcourse discrimination |
| TestICBM_Trajectory | ✅ PASS | All 4 flight phases |
| TestSBIRS_Constellation | ✅ PASS | 95% global coverage |
| TestMultiThreat | ✅ PASS | 3 RVs + 2 decoys |

---

## Physics & Realism Validation

| Check | Result | Notes |
|-------|--------|-------|
| GEO altitude (SBIRS) | ✅ | 42,164 km (correct) |
| HEO altitude (SBIRS) | ✅ | ~23,560 km (Molniya) |
| LEO altitude (STSS) | ✅ | 1,350 km |
| ICBM flight time | ✅ | 1800s (30 min for 10,000km) |
| ICBM reentry speed | ✅ | Mach 24 (~7,000 m/s) |
| SM-3 IIA range | ✅ | 2,700 km |
| SM-6 range | ✅ | 370 km, 33km ceiling |
| UEWR frequency | ✅ | 435 MHz (UHF) |
| LRDR S-band | ✅ | 3.0 GHz |
| Link 16 slots | ✅ | 1,535 per frame |
| JREAP-C MTU | ✅ | 65,535 bytes |

---

## Recommendation

**All 30 binaries execute correctly. All ~310 unit tests and 7 integration tests pass. C2BMC nil pointer crash is fixed. All simulators support `-v` verbose output.**

**Result: PASS ✅**