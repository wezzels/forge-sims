# FORGE-Sims Physics Remediation — High-Fidelity Fixes Needed

**Date:** 2026-04-23
**Total Binaries:** 55
**Passing (all flags + high-fidelity):** 44
**Need Work:** 11 sims

---

## 1. launch-veh-sim — Physics Calibration Needed

**Current text output is good, but several physics values need correction:**

### Max Q is 2.7x too high
- **Current:** 94,004 Pa (94 kPa)
- **Real F9 Block 5:** ~35,000 Pa (35 kPa)
- **Cause:** Likely using sea-level density for entire ascent instead of exponential atmosphere model
- **Fix:** Apply barometric formula ρ(h) = ρ₀ × exp(-h/H) where H=8.5km scale height. Max Q occurs where dynamic pressure ½ρv² peaks, which depends on the interplay of increasing velocity and decreasing density

### Orbital velocity slightly low
- **Current:** 7,519 m/s
- **Real LEO circular at 400km:** ~7,672 m/s
- **Fix:** Verify gravitational model uses μ = 398,600.4418 km³/s² and correct orbital velocity formula v = √(μ/r) where r = R_E + 400km = 6778 km

### Missing standard flags
- **Need:** `-i` (interactive), `-v` (verbose), `-seed` (reproducibility)
- **Need:** Change `-duration float` to Go `time.Duration` format (e.g., `5s`, `2m`)
- **These were present in R12 but removed** — restore them

### Inconsistent JSON keys
- Mixed case: `Apogee`, `Perigee`, `Incl` alongside `circ_alt_km`, `circ_dv_ms`
- Standardize to snake_case matching other sims

---

## 2. air-combat-sim — Weapons Employment Logic Missing

**The sim runs a BVR scenario but no missiles are ever fired (Pk=0.00):**

### Weapons never employed
- **Current:** 300s BVR engagement, both sides close to 9.4 km then separate, zero missiles fired
- **Real behavior:** F-22 would employ AIM-120Ds at ~80+ km, Su-57 would employ R-77Ms at ~60+ km
- **Fix:** Implement weapons employment decision logic:
  1. Radar detection triggers track
  2. Track classification triggers weapons commit
  3. Within missile NEZ → fire
  4. Missile flyout model (AIM-120D NEZ=261km, R-77M NEZ=~80km)
  5. PK calculation and kill determination

### Events list always empty
- **Current:** `"events": []` in every run
- **Fix:** Populate with radar detections, missile launches, impacts, kills:
  ```json
  {"time":80,"type":"radar_detection","entity":1,"target":2,"range_km":65}
  {"time":95,"type":"missile_launch","entity":1,"missile":"AIM-120D","target":2}
  {"time":120,"type":"missile_impact","missile":"AIM-120D","target":2,"result":"KILL"}
  ```

### F-22 fuel slightly off
- **Current:** fuel_kg = 6100
- **Real:** Internal fuel ~8200 kg (F-22 carries 18,000 lb internal = 8,165 kg)
- **Fix:** Update fuel constant

---

## 3. population-impact-sim — Needs JSON Output

**Currently text-only output. No `-json` flag, no standard sim-cli flags.**

### Add `-json` flag
- **Fix:** Add JSON output mode producing structured output per scenario:
  ```json
  {
    "simulator": "population-impact-sim",
    "status": "complete",
    "scenario": "hiroshima-recompute",
    "parameters": {
      "yield_kt": 15, "hob_km": 0.58, "population": 700000,
      "killed": 105000, "killed_pct": 15.0,
      "injured": 140000, "injured_pct": 20.0,
      "displaced": 280000, "displaced_pct": 40.0,
      "debris_m3": 150000, "debris_clearance_days": 3,
      "power_recovery_days": 72, "water_recovery_days": 168,
      "hospital_beds": 2650, "burn_bed_shortfall": 6959,
      "excess_cancer_per_100k": 12, "ptsd_rate_pct": 35,
      "gdp_loss_pct": 30, "societal_disruption": 7.5,
      "destroyed_radius_km": 1.6, "severe_radius_km": 2.2, "thermal_radius_km": 1.3
    }
  }
  ```

### Add `-scenario` flag
- Currently runs all 6 scenarios sequentially
- Need ability to run single scenario: `-scenario hiroshima`
- Default should be single scenario (not all 6)

### Add standard flags
- `-i`, `-v`, `-duration`, `-seed`

---

## 4. C2/Infrastructure Sims (8) — Need Richer JSON Data

All 8 C2/infra sims output only static parameters in JSON. Their `-v` mode shows real computation, but JSON doesn't expose it.

### bmd-sim-c2bmc
- **Fix:** Add `-threats int` flag. When threats>0, generate synthetic threat tracks, run correlation/classification, output results in JSON `data` section
- **JSON needs:** tracks, correlated tracks, engagement plans, classification confidence

### bmd-sim-gfcb
- **Fix:** Add `-threats int` flag. Generate fire control solutions with weapon-target pairing
- **JSON needs:** fire_control_solutions, weapons_status, engagement_decisions

### bmd-sim-ifxb
- **Fix:** Add `-assets int`, `-routes int` flags
- **JSON needs:** connected_assets, message_routes, network_status

### bmd-sim-jrsc
- **Fix:** Add `-sensors int` flag
- **JSON needs:** sensor_feeds, alert_decisions, regional_assessment

### bmd-sim-link16
- **Fix:** Add `-participants int` flag
- **JSON needs:** participant_list, network_stats, messages

### bmd-sim-jreap
- **Fix:** Add `-connections int` flag
- **JSON needs:** connections, messages_sent, bandwidth_used

### bmd-sim-decoy
- **Fix:** Add `-count int` flag. Expose individual decoy data already shown in `-v`
- **JSON needs:** decoys array (type, rcs, ir, discrimination_difficulty)

### bmd-sim-jamming
- **Fix:** Add `-radar string` flag. Compute burnthrough range, range denial
- **JSON needs:** jamming_effectiveness, eccm_response, radar_impact

---

## 5. Engine-Only Sims (3) — Need CLI Runners

### electronic-war-sim
- **Has:** radar, ecm, esm, spectrum, core packages
- **Needs:** `cmd/sim/main.go` with standard flags + `-scenario` flag
- **Default scenario:** radar-jamming (ECM vs AN/TPY-2)

### missile-defense-sim
- **Has:** ballistic, interceptor, seeker, track, countermeasures, layer, core packages
- **Needs:** `cmd/sim/main.go` with standard flags + `-scenario` flag
- **Default scenario:** layered-defense (SBIRS→UEWR→GBI→SM-3→Patriot)

### submarine-war-sim
- **Has:** hydro, sonar, torpedo, asw, mine, submarine, core packages
- **Needs:** `cmd/sim/main.go` with standard flags + `-scenario` flag
- **Default scenario:** asw-patrol (Virginia vs Kilo)

---

## 6. Config-Driven Sims (2) — Need Standard CLI Flags

### cyber-redteam-sim
- **Has:** `-config`, `-doctor`, `-init`, `-no-server`, `-rest-port`
- **Needs:** `-json` (stdout output), `-duration`, `-seed`, `-v`
- **When `-no-server -json`:** run simulation to completion, output JSON to stdout
- **Embed default scenario** so it works without `-config` file

### maritime-sim / space-war-sim
- **Current:** Crash with "Failed to load scenario: open configs/scenario.yaml"
- **Fix:** Handle missing config gracefully, use embedded defaults
- **Add:** `-json`, `-duration`, `-seed`, `-v` flags

---

## 7. New Sims Needing Standardization (3)

### rcs
- **Has:** `-aspect`, `-band`, `-compare`, `-json`, `-list`, `-threat`
- **Needs:** `-seed` for reproducibility (if random elements added)
- **Low priority** — lookup tool, not a simulation runner
- **Note:** Without `-threat`, produces no output (by design)

### satellite-weapon
- **Has:** `-json`, `-list`, `-scenario`, `-target`, `-weapon`, `-weapons`, `-targets`
- **Issue:** Without `-scenario`, outputs zero-filled JSON (all values = 0)
- **Fix:** Either output nothing or output an error message when no scenario selected
- **Needs:** `-seed` for reproducibility

### emp-effects
- **Has:** `-altitude`, `-json`, `-list`, `-scenario`, `-yield`
- **Needs:** Standard flags (`-i`, `-v`, `-duration`, `-seed`)
- **Without `-scenario`:** No JSON output (need default or error message)

---

## 8. tactical-net — JSON Format Issue

### Text before JSON
- **Current:** `-json` flag still prints interactive dashboard text before JSON output
- **Fix:** When `-json` is specified, only output JSON to stdout. Move interactive text to stderr or suppress it entirely.
- **Workaround:** `_run_sim_binary` in norad container already handles this by finding the last `{` and parsing from there

---

## Priority Order for Fixes

| Priority | Sim | Effort | Impact |
|----------|-----|--------|--------|
| **P1** | launch-veh-sim: Max Q calibration | Low | Physics accuracy |
| **P1** | air-combat-sim: weapons employment | Medium | Sim produces no results |
| **P2** | population-impact-sim: add `-json` | Low | Enables API integration |
| **P2** | launch-veh-sim: restore `-i`/`-v`/`-seed` | Medium | Restores lost functionality |
| **P2** | air-combat-sim: populate events | Medium | Engagement transparency |
| **P2** | 8 C2/infra sims: add scenario flags + data | Medium (each ~50-100 lines) | Rich API data |
| **P2** | satellite-weapon: handle no-scenario gracefully | Low | Prevents zero-filled JSON |
| **P2** | emp-effects: add default scenario + standard flags | Medium | CLI standardization |
| **P3** | 3 engine-only sims: add CLI runner | Medium (each ~200 lines) | Full simulation capability |
| **P3** | 2 config-driven sims: add standard CLI | Medium | CLI standardization |
| **P3** | tactical-net: suppress text in JSON mode | Low | Clean API output |
| **P3** | air-combat-sim: F-22 fuel value | Low | Accuracy |
| **P3** | launch-veh-sim: JSON key casing | Low | Consistency |
| **P3** | rcs/satellite-weapon: add `-seed` | Low | Reproducibility |

---

## Estimated Total Effort

| Category | Lines of Go | Priority |
|----------|------------|----------|
| launch-veh-sim physics + flags | ~100-200 | P1-P2 |
| air-combat-sim weapons logic | ~200-400 | P1-P2 |
| population-impact-sim JSON | ~100-150 | P2 |
| 8 C2/infra scenario flags | ~400-800 | P2 |
| satellite-weapon/emp-effects fixes | ~100-200 | P2 |
| 3 engine-only CLI runners | ~450-900 | P3 |
| 2 config-driven standard CLI | ~200-400 | P3 |
| Small fixes (tactical-net, fuel, keys) | ~50-100 | P3 |
| **Total** | **~1,600-3,150 lines** | |