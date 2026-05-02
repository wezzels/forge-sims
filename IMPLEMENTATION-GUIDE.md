# FORGE-Sims Round 26 — Implementation Guide

**Created:** 2026-05-02
**Purpose:** Step-by-step guide to implement all Round 26 remediation fixes
**Prerequisites:** SSH access to darth (10.0.0.117), GitLab access

---

## Overview

| Phase | Task | Binaries | Est. Time |
|-------|------|----------|-----------|
| **Phase 1** | space-data-network-sim flags | 1 | 30 min |
| **Phase 2** | Flag standardization | 20 | 2 hours |
| **Phase 3** | Duplicate consolidation | 15 pairs | 30 min |
| **Phase 4** | Text cleanup | 3 | 30 min |
| **Testing** | Full validation | All | 1 hour |
| **Total** | — | — | **~4.5 hours** |

---

## Phase 1: space-data-network-sim (P1)

### Step 1: Clone Repo

```bash
ssh wez@darth
cd ~/src
git clone git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git
cd space-data-network-sim
```

### Step 2: Edit main.go

Open `main.go` and make these changes:

**Add import:**
```go
import (
    "fmt"  // Add this line
    "encoding/json"
    "flag"
    // ... rest of imports
)
```

**Add flags (around line 15, after `-scenario` flag):**
```go
var (
    scenario  = flag.String("scenario", "peacetime", "Scenario: peacetime, golden-dome, nuclear, custom")
    jsonOut   = flag.Bool("json", false, "Output JSON to stdout")        // ADD
    interactive = flag.Bool("i", false, "Interactive mode")              // ADD
    verbose   = flag.Bool("v", false, "Verbose output")                  // ADD
    duration  = flag.Float64("duration", 0, "Override duration (minutes)")
    step      = flag.Float64("step", 0, "Override step size (minutes)")
    output    = flag.String("output", "", "Output JSON file")
)
```

**Add handlers in main() (before existing scenario run logic):**
```go
func main() {
    flag.Parse()

    // ADD: JSON stdout output
    if *jsonOut {
        result := runScenario(*scenario, *duration, *step)
        jsonBytes, _ := json.MarshalIndent(result, "", "  ")
        fmt.Println(string(jsonBytes))
        return
    }

    // ADD: Interactive mode
    if *interactive {
        runInteractive(*scenario, *duration, *step)
        return
    }

    // ADD: Verbose logging
    if *verbose {
        fmt.Printf("FORGE-SDN v1.0.0 — Scenario: %s\n", *scenario)
        fmt.Printf("Duration: %v min, Step: %v min\n", *duration, *step)
    }

    // ... existing scenario run logic
}
```

### Step 3: Build & Test

```bash
# Build all platforms
make build

# Test -json flag
./space-data-network-sim -json -duration 0.1 | jq .

# Test -i flag
./space-data-network-sim -i -duration 3s

# Test -v flag
./space-data-network-sim -v -scenario nuclear

# Test backward compatibility (-output still works)
./space-data-network-sim -scenario peacetime -output test.json
cat test.json | jq .summary
```

### Step 4: Copy to forge-sims

```bash
cp space-data-network-sim ~/forge-sims/binaries/linux-x86/
cp space-data-network-sim-darwin-amd64 ~/forge-sims/binaries/darwin-amd64/
cp space-data-network-sim-darwin-arm64 ~/forge-sims/binaries/darwin-arm64/
cp space-data-network-sim-windows-amd64.exe ~/forge-sims/binaries/windows-amd64/
```

### Step 5: Commit

```bash
cd ~/src/space-data-network-sim
git add main.go
git commit -m "Add -json, -i, -v flags for BMDS CLI consistency"
git push origin main
```

---

## Phase 2: Flag Standardization (P2)

### Step 1: Clone All Repos

```bash
cd ~/src
for repo in bmd-sim-c2bmc bmd-sim-link16 bmd-sim-jreap bmd-sim-jrsc bmd-sim-gfcb bmd-sim-decoy bmd-sim-ifxb bmd-sim-ir-signature bmd-sim-kill-assessment bmd-sim-rcs bmd-sim-satellite-weapon bmd-sim-wta bmd-sim-boost-intercept bmd-sim-debris-field electronic-war-sim bmd-sim-emp-effects bmd-sim-engagement-chain bmd-sim-hgv bmd-sim-kill-chain-rt bmd-sim-tactical-net; do
    git clone git@idm.wezzel.com:crab-meat-repos/$repo.git
done
```

### Step 2: Apply Fix to Each Repo

For each repo, edit `main.go`:

**Change from:**
```go
duration = flag.Duration("d", 0, "Duration")
```

**Change to:**
```go
duration      = flag.Duration("duration", 0, "Duration")
durationAlias = flag.Duration("d", 0, "Duration (alias)")
```

**Quick script to do this:**
```bash
for repo in bmd-sim-c2bmc bmd-sim-link16 bmd-sim-jreap bmd-sim-jrsc bmd-sim-gfcb bmd-sim-decoy bmd-sim-ifxb bmd-sim-ir-signature bmd-sim-kill-assessment bmd-sim-rcs bmd-sim-satellite-weapon bmd-sim-wta bmd-sim-boost-intercept bmd-sim-debris-field electronic-war-sim bmd-sim-emp-effects bmd-sim-engagement-chain bmd-sim-hgv bmd-sim-kill-chain-rt bmd-sim-tactical-net; do
    cd ~/src/$repo
    sed -i 's/flag.Duration("d",/flag.Duration("duration",/g' main.go
    # Add alias after the duration line
    sed -i '/duration.*flag.Duration/a\    durationAlias = flag.Duration("d", 0, "Duration (alias)")' main.go
    git add main.go
    git commit -m "Add -duration flag alias for CLI consistency"
    git push origin main
done
```

### Step 3: Rebuild All

```bash
# Use the forge-sims build script
cd ~/forge-sims
./build/all-platforms.sh
```

### Step 4: Test

```bash
# Test both flags produce identical output
for bin in bmd-sim-c2bmc bmd-sim-link16 bmd-sim-jreap; do
    cd ~/forge-sims/binaries/linux-x86
    diff <(./$bin -json -duration 1s) <(./$bin -json -d 1s) && echo "✓ $bin" || echo "✗ $bin"
done
```

---

## Phase 3: Duplicate Consolidation (P2)

### Step 1: Run Consolidation Script

```bash
cd ~/forge-sims
chmod +x consolidate-duplicates.sh
./consolidate-duplicates.sh
```

### Step 2: Update Makefile

Edit `Makefile` to remove non-prefixed build targets:

```makefile
# Remove these lines (or comment out):
# air-traffic
# space-debris
# satellite-tracker
# boost-intercept
# debris-field
# electronic-war-sim
# emp-effects
# engagement-chain
# hgv
# ir-signature
# kill-assessment
# kill-chain-rt
# rcs
# tactical-net
# wta
```

### Step 3: Update Documentation

**README.md:**
```markdown
## Binary Count: 66 (was 81)

| Category | Count | Description |
|----------|-------|-------------|
| **BMDS Sensors** | 10 | Radars and space-based IR sensors |
| **BMDS Threats** | 9 | ICBM, IRBM, MRBM, HGV, SLCM, decoys, EW, nuclear effects, UAP |
...
```

**API_REFERENCE.md:**
Remove duplicate entries (air-traffic, space-debris, etc.)

### Step 4: Commit

```bash
cd ~/forge-sims
git add -A
git commit -m "Remove 15 duplicate binary pairs, consolidate to bmd-sim-* naming"
git push origin main
```

---

## Phase 4: Text Cleanup (P3)

### Affected Binaries

- air-traffic
- bmd-sim-air-traffic
- satellite-tracker

### Step 1: Fix Each Repo

For each repo, edit main.go to suppress text output when `-json` is set:

```go
// BEFORE (in air-traffic main.go):
fmt.Println("Generating flights...")
result := generateFlights()
jsonBytes, _ := json.Marshal(result)
fmt.Println(string(jsonBytes))

// AFTER:
if *jsonOut {
    result := generateFlights()
    jsonBytes, _ := json.Marshal(result)
    fmt.Println(string(jsonBytes))
    return
}
fmt.Println("Generating flights...")
result := generateFlights()
jsonBytes, _ := json.Marshal(result)
fmt.Println(string(jsonBytes))
```

### Step 2: Rebuild & Test

```bash
# Rebuild
cd ~/air-traffic && make build
cd ~/bmd-sim-air-traffic && make build
cd ~/bmd-sim-satellite-tracker && make build

# Test clean JSON
./air-traffic -json | jq .
./bmd-sim-air-traffic -json | jq .
./satellite-tracker -json | jq .

# Test text output (without -json)
./air-traffic
./bmd-sim-air-traffic
./satellite-tracker
```

---

## Testing Checklist

After all phases complete, run full validation:

```bash
cd ~/forge-sims/binaries/linux-x86

# 1. Test all standard JSON sims (46 binaries)
PASS=0
for bin in $(ls bmd-sim-* | head -46); do
    ./$bin -json -duration 1s | jq . >/dev/null 2>&1 && ((PASS++)) || echo "✗ $bin"
done
echo "Standard JSON: $PASS/46"

# 2. Test config-driven sims (15 binaries)
./space-data-network-sim -json -duration 0.1 | jq .
./cyber-redteam-sim -config configs/enterprise-ad.yaml
./space-war-sim -config configs/scenario.yaml

# 3. Test flag aliases (20 binaries)
for bin in bmd-sim-c2bmc bmd-sim-link16 bmd-sim-jreap; do
    diff <($bin -json -duration 1s) <($bin -json -d 1s) && echo "✓ $bin" || echo "✗ $bin"
done

# 4. Verify no duplicates remain
ls | grep -v "^bmd-sim\|^cyber\|^space-\|^launch\|^maritime\|^sensor\|^install\|^uninstall"
# Should return empty or only scripts

# 5. Check binary count
ls -1 | wc -l
# Should be ~66 (was 81)
```

---

## Rollback Plan

If anything breaks:

```bash
# Git rollback
cd ~/forge-sims
git log --oneline -5
git revert <commit-hash>

# Binary rollback (from backup)
cp /tmp/forge-sims-backup/* binaries/linux-x86/

# Restore individual repos
cd ~/src/<repo>
git reset --hard HEAD~1
```

---

## Success Metrics

| Metric | Before | After | Status |
|--------|--------|-------|--------|
| Total binaries | 81 | ~66 | ⏳ |
| JSON-capable | 61/81 (75%) | ~64/66 (97%) | ⏳ |
| Flag consistency | 20 issues | 0 | ⏳ |
| Duplicates | 15 pairs | 0 | ⏳ |
| Documentation | Partial | Complete | ✅ |

---

## Helper Scripts

All scripts are in `~/forge-sims/`:

- `fix-space-data-network-sim.sh` — Interactive guide for SDN fixes
- `fix-flag-aliases.sh` — List of repos needing flag fixes
- `consolidate-duplicates.sh` — Remove duplicate binaries
- `IMPLEMENTATION-GUIDE.md` — This document

---

**Owner:** Wez
**Status:** Ready to execute
**Next Step:** SSH to darth and begin Phase 1
