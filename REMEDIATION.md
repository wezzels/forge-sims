# FORGE-Sims Remediation Plan — 15 Sims to Full Pass

**Date:** 2026-04-23
**Author:** Wez (AI) — Round 14 evaluation
**Goal:** Bring all 15 non-fully-passing sims to 100% pass with verified high-fidelity physics

---

## Categories

| Category | Sims | Root Cause |
|----------|------|------------|
| C2/Infrastructure (minimal JSON) | 8 | Need scenario context flags + richer JSON output |
| Partial CLI (missing flags/JSON) | 2 | launch-veh-sim (regression), air-combat-sim (new) |
| Engine-only (no CLI runner) | 3 | Need `cmd/sim/main.go` runner |
| Config-driven (no standard CLI) | 2 | Need standard CLI flags alongside config mode |

---

## 1. C2/Infrastructure Sims (8 sims)

These compute real physics when given scenario data (proven in `-v` mode), but their JSON output only shows static config parameters. They need **sim-specific scenario flags** and **computed results in JSON**.

### 1a. bmd-sim-c2bmc — Command & Control Battle Management

**Current JSON output:**
```json
{"simulator":"c2bmc","status":"complete","parameters":{"correlated":0,"threats":0,"tracks":0}}
```

**What `-v` shows (real physics exists):**
- Processes sensor reports from SBIRS, UEWR, TPY-2
- Correlates tracks (merged 3 tracks into 2)
- Classifies targets: "Track 1: RV (conf=0.60, sources=[GEO-1 BEALE GEO-2])"
- Creates engagement plans: "Track 1: speed=5975 m/s, conf=0.60"

**Needed changes:**
1. Add `-threats int` flag (default: 0) — number of synthetic threats to generate
2. When `-threats > 0`, generate synthetic threat tracks and run full C2 pipeline
3. Add `data` section to JSON with computed results:
   ```json
   "data": {
     "tracks": [{"id":"T-1","type":"RV","confidence":0.60,"sources":["GEO-1","BEALE"],"speed_ms":5975,"alt_km":200}],
     "correlated": 1,
     "engagement_plan": [{"track_id":"T-1","interceptor":"GMD-GBI","priority":1}]
   }
   ```
4. Default `-threats 3` when `-json` flag is used (so API gets real data)

**Physics to add:**
- Track correlation algorithm (already exists in `-v` mode, expose to JSON)
- Threat classification confidence scoring
- Engagement priority computation

---

### 1b. bmd-sim-gfcb — Global Fire Control & Battle

**Current JSON:**
```json
{"simulator":"gfcb","status":"complete","parameters":{"node_id":"GFCB-1","region":"PACIFIC"}}
```

**Needed changes:**
1. Add `-threats int` flag (default: 0)
2. Add `-weapons string` flag (e.g., "GBI:2,SM3:4")
3. Generate fire control solution when threats provided
4. Add `data` section:
   ```json
   "data": {
     "fire_control_solutions": [{"threat_id":"T-1","weapon":"GBI","launch_time_s":340,"intercept_time_s":411,"pk":0.54}],
     "weapons_status": {"GBI":{"available":2,"committed":1},"SM3":{"available":4,"committed":0}},
     "engagement_decisions": [{"threat_id":"T-1","action":"ENGAGE","weapon":"GBI"}]
   }
   ```

**Physics to add:**
- Weapon-target pairing (use existing GMD/AEGIS PK data)
- Fire control timeline computation
- Shoot-look-shoot doctrine implementation

---

### 1c. bmd-sim-ifxb — Integrated Fire Exchange

**Current JSON:**
```json
{"simulator":"ifxb","status":"complete","parameters":{"assets":0,"node_id":"IFXB-1","routes":0}}
```

**Needed changes:**
1. Add `-assets int` flag (default: 0) — number of connected weapons/assets
2. Add `-routes int` flag (default: 0) — communication routes
3. Default to reasonable values in `-json` mode
4. Add `data` section:
   ```json
   "data": {
     "connected_assets": [{"id":"PATRIOT-1","type":"PAC-3","status":"ready"},{"id":"THAAD-1","type":"THAAD","status":"ready"}],
     "message_routes": [{"from":"C2BMC","to":"PATRIOT-1","latency_ms":45,"protocol":"JREAP-C"}],
     "network_status": {"bandwidth_kbps":256,"latency_ms":45,"reliability":0.99}
   }
   ```

**Physics to add:**
- Network latency computation (distance + protocol overhead)
- Message routing simulation
- Asset status tracking

---

### 1d. bmd-sim-jrsc — Joint Regional Security Center

**Current JSON:**
```json
{"simulator":"jrsc","status":"complete","parameters":{"alert":3,"region":"JRSC-PAC","sensors":0}}
```

**Needed changes:**
1. Add `-sensors int` flag (default: 0) — connected sensor count
2. Default to 5+ sensors in `-json` mode
3. Add `data` section:
   ```json
   "data": {
     "sensor_feeds": [{"id":"SBIRS-GEO-1","type":"IR","status":"active","tracks":3},{"id":"UEWR-BEALE","type":"Radar","status":"active","tracks":2}],
     "alert_decisions": [{"level":3,"reason":"detected_launch","threats":2}],
     "regional_assessment": {"threat_level":"MODERATE","defense_posture":"READY"}
   }
   ```

---

### 1e. bmd-sim-link16 — Tactical Data Link

**Current JSON:**
```json
{"simulator":"link16","status":"complete","parameters":{"net_id":"N-1","participants":0,"time_slots":1535}}
```

**Needed changes:**
1. Add `-participants int` flag (default: 0)
2. Default to 8 participants in `-json` mode
3. Add `data` section:
   ```json
   "data": {
     "participant_list": [{"id":"F-22-1","type":"fighter","time_slots":64,"msgs_per_frame":12},{"id":"AEGIS-1","type":"ship","time_slots":128}],
     "network_stats": {"utilization_pct":45.2,"throughput_kbps":128,"update_rate_s":12},
     "messages": [{"type":"J3.2","from":"AEGIS-1","track_count":3,"timestamp":"T+45s"}]
   }
   ```

**Physics to add:**
- TDMA time slot allocation (1535 slots/frame, 12s/frame)
- Message throughput calculation
- Update rate vs participants tradeoff

---

### 1f. bmd-sim-jreap — JREAP-C Transport

**Current JSON:**
```json
{"simulator":"jreap","status":"complete","parameters":{"port":4000,"protocol":0}}
```

**Needed changes:**
1. Add `-connections int` flag (default: 0)
2. Add `-protocol int` flag (0=JREAP-C, 1=JREAP-A, 2=JREAP-T)
3. Default to 3 connections in `-json` mode
4. Add `data` section:
   ```json
   "data": {
     "connections": [{"id":"C2BMC-PAC","protocol":"JREAP-C","port":4000,"status":"active","latency_ms":50}],
     "messages_sent": 45,
     "bandwidth_used_kbps": 32,
     "protocol_overhead_pct": 12.5
   }
   ```

---

### 1g. bmd-sim-decoy — Penetration Aids

**Current JSON:**
```json
{"simulator":"decoy","status":"complete","parameters":{"active_count":2,"rcs":0.1,"type":"balloon"}}
```

**What `-v` shows:**
- "Decoy #1: BALLOON RCS=0.10 IR=5000 diff=0.40"
- "Decoy #2: INFLATABLE_RV RCS=0.10 IR=1000 diff=0.80"

**Needed changes:**
1. Add `-count int` flag (default: 2) — number of decoys to deploy
2. Add `data` section with individual decoy details:
   ```json
   "data": {
     "decoys": [
       {"id":1,"type":"BALLOON","rcs_m2":0.10,"ir_w":5000,"discrimination_difficulty":0.40,"deployed":true},
       {"id":2,"type":"INFLATABLE_RV","rcs_m2":0.10,"ir_w":1000,"discrimination_difficulty":0.80,"deployed":true}
     ],
     "discrimination_stats": {"avg_difficulty":0.60,"best_decoy":"INFLATABLE_RV","sensor_confusion_pct":35}
   }
   ```

**Physics to add:**
- IR signature modeling (already exists in `-v`, expose to JSON)
- Discrimination difficulty scoring
- Sensor confusion probability

---

### 1h. bmd-sim-jamming — EW/Jamming

**Current JSON:**
```json
{"simulator":"jamming","status":"complete","parameters":{"effect":"DENIED","freq_ghz":3,"jnr_db":109.01,"power_w":1000000}}
```

**Needed changes:**
1. Add `-radar string` flag (target radar, e.g., "TPY-2")
2. Add `-ecm-type string` flag (NOISE, DRFM, SPOT, SWEEP)
3. Add `data` section:
   ```json
   "data": {
     "jamming_effectiveness": {"range_denied_km":350,"range_degraded_km":500,"burnthrough_km":1.15},
     "eccm_response": {"frequency_hopping":true,"side_lobe_cancellation":true,"recovery_pct":0},
     "radar_impact": {"original_range_km":6304,"effective_range_km":1.15,"tracking_lost":true}
   }
   ```

**Physics to add:**
- Burnthrough range calculation (JNR vs radar power)
- ECCM effectiveness modeling
- Range denial zone computation

---

## 2. Partial CLI Sims (2 sims)

### 2a. launch-veh-sim — REGRESSION

**Current state:**
- Text output: ✅ High-fidelity (orbit achieved, realistic values)
- `-json`: ❌ Returns completely empty (was working in R12)
- `-i` interactive: ❌ Removed (was working in R12)
- `-v` verbose: ❌ Removed (was working in R12)
- `-seed`: ❌ Removed (was working in R12)
- `-duration`: Changed from Go `time.Duration` to `float` seconds

**What needs to happen:**

1. **Fix `-json` (P1):** Restore JSON output. The physics engine works (text output proves it), but the JSON serialization is broken. Need to debug `main.go` — likely the JSON flag is parsed but the output writer is not being called. Expected output:
   ```json
   {
     "simulator": "launch-veh-sim",
     "status": "complete",
     "parameters": {
       "vehicle": "Falcon 9 Block 5",
       "payload_mass_kg": 22800,
       "target_altitude_km": 400,
       "flight_time_s": 277.3,
       "max_velocity_ms": 7191,
       "max_altitude_km": 400,
       "max_q_Pa": 30433,
       "max_q_time_s": 61,
       "status": "orbit_achieved"
     },
     "data": {
       "orbit": {"apogee_km":6049,"perigee_km":6049,"eccentricity":0.0,"period_min":78.0},
       "milestones": [...],
       "stage_events": [...]
     }
   }
   ```

2. **Restore `-i` flag (P2):** Add interactive dashboard back. Show real-time trajectory, altitude, velocity, Max Q during ascent.

3. **Restore `-v` flag (P2):** Add verbose text output back.

4. **Restore `-seed` flag (P2):** Add reproducibility. The trajectory computation uses numerical integration — seed the random components (atmospheric perturbations, etc.).

5. **Fix `-duration` format (P2):** Change back to Go `time.Duration` (e.g., `5s`, `2m`) for consistency with all other sims. Currently uses float seconds which is inconsistent.

**Physics verification needed:**
- ✅ Payload 22,800 kg (matches real F9 Block 5)
- ✅ Max Q 30,433 Pa (close to real ~35 kPa)
- ✅ Max Q time 61s (close to real ~70-80s)
- ✅ Perigee 6,049 km (circular orbit, realistic)
- ⚠️ Orbital velocity 7,191 m/s (real LEO ~7,660 m/s — verify if this accounts for Earth rotation)
- ⚠️ Apogee = Perigee = 6,049 km and Eccentricity = 0.0000 (perfectly circular — suspicious, real orbits have some eccentricity)

---

### 2b. air-combat-sim — Missing JSON stdout + weapons logic

**Current state:**
- Runs BVR scenario with real physics (F-22 vs Su-57)
- Exports to `air-combat-result.json` file (not stdout)
- No `-json` stdout flag
- No `-seed` flag
- No `-duration` flag (fixed 300s)
- Pk = 0.00 in default run (no weapons employment)

**What needs to happen:**

1. **Add `-json` stdout output (P1):** Instead of/in addition to file export, output JSON to stdout like all other sims. This is required for API integration. The `air-combat-result.json` already has good structure — just needs to go to stdout.

2. **Fix weapons employment logic (P2):** In the default BVR Shootout scenario, neither side fires missiles despite closing to within 9.4 km. The F-22 should employ AIM-120Ds at ~80+ km range. Possible issues:
   - Radar detection not triggering weapon release
   - Rules of engagement too restrictive
   - Missile employment decision logic not implemented
   - Add `-roe string` flag (e.g., "weapons-free", "weapons-tight", "weapons-hold")

3. **Add `-seed` flag (P2):** For reproducibility of engagement outcomes.

4. **Add `-duration` flag (P2):** Allow variable scenario length instead of fixed 300s.

5. **Add events to JSON (P2):** The `events` list in the result JSON is always empty. Need to populate with:
   ```json
   "events": [
     {"time":80,"type":"radar_detection","entity":1,"target":2,"range_km":65},
     {"time":95,"type":"missile_launch","entity":1,"missile":"AIM-120D","target":2},
     {"time":120,"type":"missile_impact","missile":"AIM-120D","target":2,"result":"KILL"}
   ]
   ```

6. **Fix F-22 mass (P3):** MassKg=24000 should be ~27200 (loaded). FuelKg=5900 should be ~8200 (internal).

---

## 3. Engine-Only Sims (3 sims)

These have library code with real physics packages but no `cmd/sim/main.go` CLI runner. They just print version and package list.

### 3a. electronic-war-sim

**Current output:**
```
electronic-war-sim v0.1.0
Electromagnetic warfare simulation engine
Packages: radar, ecm, esm, spectrum, core
```

**What needs to happen:**
1. Create `cmd/sim/main.go` with standard sim-cli flags (`-json`, `-i`, `-v`, `-duration`, `-seed`, `-help`)
2. Add `-scenario string` flag (default: "radar-jamming")
3. Run a scenario using the existing library packages
4. JSON output example:
   ```json
   {
     "simulator": "electronic-war-sim",
     "status": "complete",
     "parameters": {
       "scenario": "radar-jamming",
       "radar": {"type":"AN/TPY-2","freq_ghz":9.5,"power_kw":100,"range_km":6304},
       "jammer": {"type":"NOISE_BARRAGE","power_w":10000,"freq_ghz":9.5},
       "results": {"jsr_db":15.2,"range_denied_km":3500,"burnthrough_km":120}
     }
   }
   ```

**Physics packages already exist:** radar, ecm, esm, spectrum, core. Just need a runner.

---

### 3b. missile-defense-sim

**Current output:**
```
missile-defense-sim v0.1.0
Ballistic missile defense simulation engine
Packages: ballistic, interceptor, seeker, track, countermeasures, layer, core
```

**What needs to happen:**
1. Create `cmd/sim/main.go` with standard sim-cli flags
2. Add `-scenario string` flag (default: "layered-defense")
3. Run a layered BMD scenario: SBIRS detect → UEWR track → GBI engage → SM-3 backup → Patriot terminal
4. JSON output example:
   ```json
   {
     "simulator": "missile-defense-sim",
     "status": "complete",
     "parameters": {
       "scenario": "layered-defense",
       "threats": [{"id":"T-1","type":"ICBM","range_km":10000,"pk_vs":0.54}],
       "defense_layers": [{"name":"GMD","interceptors":2,"pk":0.54},{"name":"AEGIS-SM3","interceptors":4,"pk":0.74}],
       "results": {"threats_killed":1,"threats_leaked":0,"interceptors_used":1,"overall_pk":0.78}
     }
   }
   ```

**Physics packages already exist:** ballistic, interceptor, seeker, track, countermeasures, layer, core.

---

### 3c. submarine-war-sim

**Current output:**
```
submarine-war-sim v0.1.0
Underwater warfare simulation engine
Packages: hydro, sonar, torpedo, asw, mine, submarine, core
Use: go test ./... to run tests
     go run ./cmd/sim/ to run scenarios
```

**What needs to happen:**
1. Create or fix `cmd/sim/main.go` with standard sim-cli flags
2. Add `-scenario string` flag (default: "asw-patrol")
3. Run an ASW scenario: sonar search → torpedo attack
4. JSON output example:
   ```json
   {
     "simulator": "submarine-war-sim",
     "status": "complete",
     "parameters": {
       "scenario": "asw-patrol",
       "submarine": {"class":"Virginia","depth_m":200,"speed_kn":12},
       "sonar": {"type":"BQQ-10","freq_khz":3,"detection_range_km":45},
       "results": {"contacts":2,"classified":1,"torpedoes_fired":1,"pk":0.65}
     }
   }
   ```

**Physics packages already exist:** hydro, sonar, torpedo, asw, mine, submarine, core.

---

## 4. Config-Driven Sims (2 sims)

These use YAML config files and have their own architecture. They need standard CLI flags alongside their config mode.

### 4a. cyber-redteam-sim

**Current flags:**
```
-config string    scenario config file (default "configs/enterprise-ad.yaml")
-doctor           run health check diagnostics
-init             initialize config directory with sample scenarios
-no-server        disable API servers (run simulation only)
-rest-port int    REST API port
```

**What needs to happen:**
1. Add `-json` flag — output results as JSON to stdout (in addition to/instead of API)
2. Add `-duration duration` flag — limit simulation run time
3. Add `-seed int` flag — reproducibility
4. Add `-v` flag — verbose output
5. When `-no-server -json` is used, run simulation to completion and output JSON to stdout
6. Embed a default scenario so it works without `-config` file (similar to how other sims have default scenarios)

**Example JSON output:**
```json
{
  "simulator": "cyber-redteam-sim",
  "status": "complete",
  "parameters": {
    "scenario": "enterprise-ad",
    "network_nodes": 50,
    "attack_vectors": ["phishing","lateral-movement","privilege-escalation"],
    "results": {
      "initial_access": {"method":"phishing","success_rate":0.35,"time_to_compromise_s":7200},
      "lateral_movement": {"technique":"pass-the-hash","hosts_compromised":12},
      "detection": {"ids_alerts":3,"siem_correlated":1,"mean_time_detect_s":14400}
    }
  }
}
```

---

### 4b. maritime-sim

**Current state:** Exits with error "Failed to load scenario: open configs/scenario.yaml: no such file or directory"

**What needs to happen:**
1. Add `-json` flag — output results as JSON to stdout
2. Add `-duration duration` flag
3. Add `-seed int` flag
4. Add `-v` flag
5. Embed a default scenario so it works without config file (ships, submarines, aircraft in a naval theater)
6. Handle missing config gracefully — don't crash, use defaults
7. Standard CLI flags alongside config mode

**Example JSON output:**
```json
{
  "simulator": "maritime-sim",
  "status": "complete",
  "parameters": {
    "scenario": "default-pacific",
    "vessels": [{"id":"DDG-1","type":"destroyer","position":{"lat":35.0,"lon":140.0}}],
    "results": {"contacts_tracked":5,"weapons_engaged":2,"pk":0.72}
  }
}
```

**Note:** space-war-sim also exits with error for missing config — same treatment as maritime-sim.

---

## Priority Order

| Priority | Sim | Effort | Impact |
|----------|-----|--------|--------|
| P1 | launch-veh-sim: fix `-json` | Low | Restores working JSON (regression) |
| P1 | air-combat-sim: add `-json` stdout | Low | Enables API integration |
| P2 | C2BMC: add `-threats` flag + data | Medium | Richest C2 sim, already computes in `-v` |
| P2 | air-combat-sim: fix weapons logic | Medium | Makes sim actually produce kills |
| P2 | launch-veh-sim: restore `-i`, `-v`, `-seed` | Medium | Restores lost functionality |
| P2 | All 8 C2/infra: add scenario flags + data | Medium | Each needs ~50 lines of Go |
| P2 | 3 engine sims: add CLI runner | Medium | Each needs `cmd/sim/main.go` |
| P3 | Config-driven: add standard CLI flags | Medium | Architecture change |
| P3 | air-combat-sim: fix mass values | Low | Simple constant fix |

---

## Effort Estimate

| Category | Per Sim | Total |
|----------|---------|-------|
| C2/Infrastructure (8) | ~50-100 lines Go each | ~400-800 lines |
| launch-veh-sim fixes | ~20-50 lines (debug JSON) | ~50 lines |
| air-combat-sim fixes | ~100-200 lines | ~200 lines |
| Engine-only CLI runners (3) | ~150-300 lines each | ~450-900 lines |
| Config-driven (2) | ~100-200 lines each | ~200-400 lines |
| **Total** | | **~1300-2350 lines of Go** |

---

## Verification Checklist (per sim after changes)

For each sim, verify:
- [ ] `./sim -json -duration 5s` produces valid JSON with computed physics values
- [ ] `./sim -i -duration 5s` shows interactive dashboard
- [ ] `./sim -v -duration 5s` shows verbose physics output
- [ ] `./sim -seed 42 -duration 5s` produces reproducible output
- [ ] `./sim -duration 5s` uses Go time.Duration format
- [ ] JSON has `simulator`, `status`, `parameters`, and `data` sections
- [ ] `data` section contains computed physics (not just config)
- [ ] Numeric values are realistic (not zero, not suspiciously round)
- [ ] No stubs, mocks, or placeholder data in output
- [ ] All values match published real-world data where available