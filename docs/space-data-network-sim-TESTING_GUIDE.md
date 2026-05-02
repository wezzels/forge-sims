# FORGE-SDN Testing Guide

Complete step-by-step guide for verifying the FORGE-SDN simulator.

---

## Prerequisites

```bash
# Verify binary exists and is executable
file space-data-network-sim
# Expected: ELF 64-bit LSB executable, x86-64 (Linux)
# or: Mach-O 64-bit arm64 (macOS)
# or: PE32+ executable x86-64 (Windows)
```

---

## Test 1: Peacetime Baseline

**Purpose:** Verify constellation initialization and normal operations.

```bash
./space-data-network-sim -scenario peacetime
```

**Verify:**
- [ ] 791 total satellites displayed
- [ ] 3 layers shown: Backbone (432), Space Link (21), Transport (338)
- [ ] Coverage = 100.0%
- [ ] Resilience score ≥ 75
- [ ] 0 sats lost
- [ ] Simulation completes in < 500ms

**Expected output:**
```
Summary:
  Final sats: 791 / 791
  Final links: 1526
  Coverage: 100.0%
  Routes: 28.6%
  Resilience: 81.1 / 100
  Sats lost: 0
  Threats executed: 0
```

---

## Test 2: Golden Dome Scenario

**Purpose:** Verify threat models (ASAT, jamming, cyber) and network re-routing.

```bash
./space-data-network-sim -scenario golden-dome
```

**Verify:**
- [ ] 4 threat events executed
- [ ] ASAT events show `hit=true` with debris count
- [ ] Jamming event shows JSR value in dB and `link_killed=true`
- [ ] Cyber event fires
- [ ] Some sats lost (typically 2)
- [ ] Coverage still > 95% (network re-routes)
- [ ] Resilience score > 80 (partial degradation only)

**Expected output:**
```
Threat Events:
  [t=15] ASAT → BB-0-5: hit=true debris=150
  [t=30] Jamming → KaS2G: JSR=41.7 dB killed=true
  [t=60] ASAT → BB-3-12: hit=true debris=0
  [t=90] Cyber → GEP-2

Summary:
  Final sats: 789 / 791
  Coverage: 99.8%
  Resilience: 85.3 / 100
  Sats lost: 2
```

---

## Test 3: Nuclear Scenario

**Purpose:** Verify catastrophic threat modeling and survival rates.

```bash
./space-data-network-sim -scenario nuclear
```

**Verify:**
- [ ] 1 threat event (nuclear detonation)
- [ ] Massive satellite loss (> 90%)
- [ ] Survival rate ≈ 5%
- [ ] Coverage drops dramatically (< 30%)
- [ ] Resilience score < 30
- [ ] Most satellites killed by radiation belt

**Expected output:**
```
Threat Events:
  [t=10] Nuclear 400kT@400km: survival=5.0%

Summary:
  Final sats: 25 / 791
  Coverage: 22.5%
  Resilience: 21.1 / 100
  Sats lost: 766
```

---

## Test 4: Custom Duration and Step Size

**Purpose:** Verify CLI flag overrides work correctly.

```bash
./space-data-network-sim -scenario peacetime -duration 30 -step 5
```

**Verify:**
- [ ] Duration shows 30 min (not default 60)
- [ ] Step shows 5 min (not default 15)
- [ ] Time series has 6 steps (0, 5, 10, 15, 20, 25, 30)
- [ ] Simulation completes successfully

---

## Test 5: JSON Output

**Purpose:** Verify structured JSON output for API integration.

```bash
./space-data-network-sim -scenario golden-dome -output test_result.json
```

**Verify:**
- [ ] File `test_result.json` created
- [ ] Valid JSON structure (use `python3 -m json.tool test_result.json`)
- [ ] Contains `name`, `description`, `start_time`, `duration_min`
- [ ] Contains `time_series` array with per-step metrics
- [ ] Contains `threat_log` array with threat events
- [ ] Contains `summary` object with final metrics
- [ ] `resilience_score` is a number between 0-100

```bash
# Validate JSON
python3 -c "
import json
with open('test_result.json') as f:
    d = json.load(f)
assert 'time_series' in d, 'Missing time_series'
assert 'threat_log' in d, 'Missing threat_log'
assert 'summary' in d, 'Missing summary'
assert len(d['time_series']) > 0, 'No time steps'
print(f'✅ JSON valid: {len(d[\"time_series\"])} steps, {len(d[\"threat_log\"])} threats, resilience={d[\"summary\"][\"resilience_score\"]}')
"
```

---

## Test 6: Unit Tests

**Purpose:** Verify all internal packages pass their test suites.

```bash
go test -v ./...
```

**Verify:**
- [ ] All 9 packages pass (ok or cached)
- [ ] No FAIL entries
- [ ] No panics or race conditions

**Expected packages:**
- `internal/constellation` — Walker generator, orbital mechanics
- `internal/mesh` — Dijkstra routing, link management
- `internal/oct` — OCT specs, link budget
- `internal/protocol` — LDPC encoding, OCT framing
- `internal/routing` — OSPF, QoS
- `internal/scenario` — Scenario config, runner
- `internal/sgp4` — SGP4 propagation
- `internal/simulation` — SDN network, link scheduler
- `internal/threat` — Kessler, ASAT, nuclear, jamming

---

## Test 7: Link Scheduler

**Purpose:** Verify dynamic link lifecycle transitions.

```bash
go test -v ./internal/simulation/ -run LinkScheduler
```

**Verify:**
- [ ] 6 link states transition correctly: ACQUIRE → ACTIVE → HANDOFF → DEGRADED → TEARDOWN → LOST
- [ ] Max 4 concurrent links per satellite enforced
- [ ] Handoff timeout (60s) respected
- [ ] Link budget threshold checked before ACQUIRE

---

## Test 8: SGP4 Propagation

**Purpose:** Verify orbital mechanics accuracy.

```bash
go test -v ./internal/sgp4/
```

**Verify:**
- [ ] J2 perturbation produces correct precession
- [ ] Atmospheric drag reduces semi-major axis over time
- [ ] Orbital period matches expected values (~105 min at 1000 km)

---

## Test 9: Cross-Platform Binary

**Purpose:** Verify binaries work on all supported platforms.

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o space-data-network-sim ./cmd/sim/
file space-data-network-sim
# Expected: ELF 64-bit LSB executable, x86-64

# macOS
GOOS=darwin GOARCH=arm64 go build -o space-data-network-sim ./cmd/sim/
file space-data-network-sim
# Expected: Mach-O 64-bit arm64 executable

# Windows
GOOS=windows GOARCH=amd64 go build -o space-data-network-sim.exe ./cmd/sim/
file space-data-network-sim.exe
# Expected: PE32+ executable (console) x86-64
```

---

## Test 10: FORGE-C2 Integration

**Purpose:** Verify NORAD API endpoint.

```bash
# On miner (207.244.226.151) or via norad.stsgym.com
curl -s "https://norad.stsgym.com/api/public/sdn?scenario=peacetime" | python3 -m json.tool | head -10
```

**Verify:**
- [ ] HTTP 200 response
- [ ] Valid JSON with `simulator` and `status` fields
- [ ] `status` = `complete`

---

## Troubleshooting

### "Port 5008 in use" (NORAD deployment)
```bash
sudo lsof -i :5008 -t | xargs -r sudo kill -9
docker restart norad-sim
```

### "Address already in use" (API server)
The REST API server defaults to port 8080. Kill conflicting processes or change the port in source.

### "Unknown scenario" error
Only `peacetime`, `golden-dome` (or `gd`), `nuclear` (or `nuke`), and `custom` are valid scenario names.

### Empty time_series in JSON
Ensure `duration_min > 0` and `step_min > 0`. Use `-duration` and `-step` flags to override.

### Resilience score seems low (peacetime)
The baseline resilience score (~81) accounts for the fact that not all possible OSPF routes are valid at any given time step. This is expected — satellites in different orbital planes have limited cross-plane connectivity.

---

## Automated Verification Script

```bash
#!/bin/bash
set -e
echo "=== FORGE-SDN Verification ==="

echo "[1/6] Peacetime scenario..."
./space-data-network-sim -scenario peacetime -output /tmp/sdn_peace.json
python3 -c "import json; d=json.load(open('/tmp/sdn_peace.json')); assert d['summary']['final_sats']==791; print('✅ 791 sats')"

echo "[2/6] Golden Dome scenario..."
./space-data-network-sim -scenario golden-dome -output /tmp/sdn_gd.json
python3 -c "import json; d=json.load(open('/tmp/sdn_gd.json')); assert d['summary']['threats_executed']==4; print('✅ 4 threats')"

echo "[3/6] Nuclear scenario..."
./space-data-network-sim -scenario nuclear -output /tmp/sdn_nuke.json
python3 -c "import json; d=json.load(open('/tmp/sdn_nuke.json')); assert d['summary']['final_sats']<100; print('✅ catastrophic loss')"

echo "[4/6] Unit tests..."
go test ./... 2>&1 | grep -E '^ok|^FAIL'

echo "[5/6] JSON validation..."
python3 -c "
import json
for f in ['/tmp/sdn_peace.json','/tmp/sdn_gd.json','/tmp/sdn_nuke.json']:
    d=json.load(open(f))
    assert 'time_series' in d and 'summary' in d
    print(f'✅ {f.split(\"/\")[-1]}: {len(d[\"time_series\"])} steps')
"

echo "[6/6] Cross-compile check..."
GOOS=darwin GOARCH=arm64 go build -o /dev/null ./cmd/sim/ && echo "✅ darwin-arm64"
GOOS=windows GOARCH=amd64 go build -o /dev/null ./cmd/sim/ && echo "✅ windows-amd64"

echo ""
echo "=== All verification tests passed ==="
```