# FORGE-Sims Round 26 Remediation Plan

**Created:** 2026-05-02
**Priority:** P1/P2 fixes for 35 binaries
**Estimated Effort:** 2-3 days

---

## Executive Summary

Round 26 identified **35 binaries requiring fixes**:
- 1 space-data-network-sim (missing standard flags)
- 20 flag convention issues (`-d` vs `-duration`)
- 15 duplicate pairs (consolidation needed)

---

## P1 — Critical Fixes (1 sim)

### space-data-network-sim

**Current State:**
- Uses `-output <file>` instead of `-json` for stdout
- No interactive mode (`-i`)
- No verbose mode (`-v`)

**Required Changes:**

```go
// Add -json flag support
var jsonOutput = flag.Bool("json", false, "Output JSON to stdout")

// Add -i flag support
var interactive = flag.Bool("i", false, "Interactive mode (live dashboard)")

// Add -v flag support
var verbose = flag.Bool("v", false, "Verbose output")

// In main():
if *jsonOutput {
    // Output JSON to stdout instead of file
    fmt.Println(jsonResult)
    return
}
```

**Files to Modify:**
- Source: `git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git`
- File: `main.go` (flag parsing + output logic)

**Test Plan:**
```bash
# Test stdout JSON
./space-data-network-sim -json -duration 1s | jq .

# Test interactive mode
./space-data-network-sim -i -duration 5s

# Test verbose mode
./space-data-network-sim -v -scenario nuclear
```

**Acceptance Criteria:**
- [ ] `-json` outputs valid JSON to stdout
- [ ] `-i` shows live dashboard
- [ ] `-v` shows verbose logging
- [ ] Backward compatible with `-output` flag

---

## P2 — Flag Standardization (20 binaries)

### Issue: `-d` vs `-duration` Inconsistency

**Affected Binaries:**

| # | Binary | Current | Should Be |
|---|--------|---------|-----------|
| 1 | bmd-sim-c2bmc | `-d duration` | `-duration duration` |
| 2 | bmd-sim-link16 | `-d duration` | `-duration duration` |
| 3 | bmd-sim-jreap | `-d duration` | `-duration duration` |
| 4 | bmd-sim-jrsc | `-d duration` | `-duration duration` |
| 5 | bmd-sim-gfcb | `-d duration` | `-duration duration` |
| 6 | bmd-sim-decoy | `-d duration` | `-duration duration` |
| 7 | bmd-sim-ifxb | `-d duration` | `-duration duration` |
| 8 | bmd-sim-ir-signature | `-d duration` | `-duration duration` |
| 9 | bmd-sim-kill-assessment | `-d duration` | `-duration duration` |
| 10 | bmd-sim-rcs | `-d duration` | `-duration duration` |
| 11 | bmd-sim-satellite-weapon | `-d duration` | `-duration duration` |
| 12 | bmd-sim-wta | `-d duration` | `-duration duration` |
| 13 | boost-intercept | `-d duration` | `-duration duration` |
| 14 | debris-field | `-d duration` | `-duration duration` |
| 15 | electronic-war-sim | `-d duration` | `-duration duration` |
| 16 | emp-effects | `-d duration` | `-duration duration` |
| 17 | engagement-chain | `-d duration` | `-duration duration` |
| 18 | hgv | `-d duration` | `-duration duration` |
| 19 | kill-chain-rt | `-d duration` | `-duration duration` |
| 20 | tactical-net | `-d duration` | `-duration duration` |

**Required Changes:**

```go
// Change from:
var duration = flag.Duration("d", 0, "Duration")

// Change to:
var duration = flag.Duration("duration", 0, "Duration")
// Add alias for backward compatibility:
var durationAlias = flag.Duration("d", 0, "Duration (alias)")
```

**Files to Modify:**
Each binary's source repo on `idm.wezzel.com:crab-meat-repos/`

**Test Plan:**
```bash
# Test both flags work
./bmd-sim-c2bmc -json -duration 1s | jq .
./bmd-sim-c2bmc -json -d 1s | jq .

# Verify identical output
diff <(./bin -json -duration 1s) <(./bin -json -d 1s)
```

**Acceptance Criteria:**
- [ ] Both `-duration` and `-d` work (backward compatible)
- [ ] JSON output identical for both flags
- [ ] Help text shows both flags

---

## P2 — Duplicate Consolidation (15 pairs)

### Strategy

**Keep:** `bmd-sim-*` prefixed versions
**Remove:** Non-prefixed versions

| Keep | Remove | Notes |
|------|--------|-------|
| bmd-sim-air-traffic | air-traffic | Same sim |
| bmd-sim-space-debris | space-debris | Same sim |
| bmd-sim-satellite-tracker | satellite-tracker | Size differs |
| bmd-sim-boost-intercept | boost-intercept | Same sim |
| bmd-sim-debris-field | debris-field | Same sim |
| bmd-sim-electronic-attack | electronic-war-sim | Similar |
| bmd-sim-emp-effects | emp-effects | Same sim |
| bmd-sim-engagement-chain | engagement-chain | Same sim |
| bmd-sim-hgv | hgv | Same sim |
| bmd-sim-ir-signature | ir-signature | Same sim |
| bmd-sim-kill-assessment | kill-assessment | Same sim |
| bmd-sim-kill-chain-rt | kill-chain-rt | Same sim |
| bmd-sim-rcs | rcs | Same sim |
| bmd-sim-tactical-net | tactical-net | Same sim |
| bmd-sim-wta | wta | Same sim |

**Action Items:**

1. **Update build scripts** — Stop building non-prefixed versions
2. **Update documentation** — Reference only `bmd-sim-*` names
3. **Update API_REFERENCE.md** — Remove duplicate entries
4. **Remove old binaries** — From `binaries/linux-x86/`, `binaries/darwin-*`, `binaries/windows-*`
5. **Update EVALUATION.md** — Reflect consolidated count

**Files to Modify:**
- `Makefile`
- `.gitlab-ci.yml`
- `API_REFERENCE.md`
- `README.md`
- `about/*.md` (for non-prefixed sims)

---

## P3 — Text-Before-JSON Issues (3 binaries)

### Issue

These binaries output text before JSON, breaking `jq` parsing:

| Binary | Current Output | Expected |
|--------|---------------|----------|
| air-traffic | `Generating flights...\n{...}` | `{...}` |
| bmd-sim-air-traffic | `Generating flights...\n{...}` | `{...}` |
| satellite-tracker | `Fetching TLE...\n{...}` | `{...}` |

**Required Changes:**

```go
// Suppress text output when -json flag is set
if *jsonOutput {
    // Skip all text output, only print JSON
    fmt.Println(jsonResult)
    return
}
```

**Test Plan:**
```bash
# Verify clean JSON output
./air-traffic -json | jq .

# Verify text output without -json
./air-traffic
```

---

## Implementation Checklist

### Phase 1: space-data-network-sim (P1)
- [ ] Clone repo: `git clone git@idm.wezzel.com:crab-meat-repos/space-data-network-sim.git`
- [ ] Add `-json` flag support
- [ ] Add `-i` interactive mode
- [ ] Add `-v` verbose mode
- [ ] Test all 3 scenarios with new flags
- [ ] Rebuild binaries (linux/darwin/windows)
- [ ] Update binaries in forge-sims repo
- [ ] Update API_REFERENCE.md

### Phase 2: Flag Standardization (P2)
- [ ] Create list of 20 repos to update
- [ ] Clone all repos to darth
- [ ] Apply `-d` → `-duration` fix with alias
- [ ] Rebuild all binaries
- [ ] Test both flag variants
- [ ] Update forge-sims binaries

### Phase 3: Duplicate Consolidation (P2)
- [ ] Update Makefile to skip non-prefixed builds
- [ ] Remove old binaries from all platforms
- [ ] Update API_REFERENCE.md
- [ ] Update README.md binary count
- [ ] Update EVALUATION.md
- [ ] Commit and push

### Phase 4: Text Cleanup (P3)
- [ ] Fix air-traffic
- [ ] Fix bmd-sim-air-traffic
- [ ] Fix satellite-tracker
- [ ] Test clean JSON output
- [ ] Rebuild and update

---

## Testing Matrix

After all fixes, verify:

```bash
# Standard JSON test (all 46 standard sims)
for bin in $(ls binaries/linux-x86/bmd-sim-* | head -46); do
    $bin -json -duration 1s | jq . >/dev/null && echo "✓ $bin" || echo "✗ $bin"
done

# Config-driven test (all 15 config sims)
for bin in space-data-network-sim cyber-redteam-sim space-war-sim; do
    # Test per-binary requirements
done

# Flag alias test (20 standardized binaries)
for bin in bmd-sim-c2bmc bmd-sim-link16 bmd-sim-jreap; do
    diff <($bin -json -duration 1s) <($bin -json -d 1s) && echo "✓ $bin" || echo "✗ $bin"
done

# Duplicate check (ensure no duplicates remain)
ls binaries/linux-x86/ | grep -v "^bmd-sim\|^cyber\|^space-\|^launch\|^maritime\|^sensor"
# Should return empty or only scripts
```

---

## Success Metrics

| Metric | Before | After | Target |
|--------|--------|-------|--------|
| Total binaries | 81 | ~66 | <70 |
| JSON-capable | 61/81 (75%) | ~64/66 (97%) | >95% |
| Flag consistency | 20 issues | 0 | 0 |
| Duplicates | 15 pairs | 0 | 0 |
| Documentation | Partial | Complete | 100% |

---

## Timeline

| Phase | Duration | Dependencies |
|-------|----------|--------------|
| P1: space-data-network-sim | 2-3 hours | Source access |
| P2: Flag standardization | 4-6 hours | 20 repos access |
| P3: Duplicate consolidation | 2-3 hours | Build system |
| P4: Text cleanup | 1-2 hours | 3 repos access |
| Testing | 2-3 hours | All phases complete |
| **Total** | **11-17 hours** | — |

---

## Rollback Plan

If issues arise:
1. Keep backup of old binaries in `/tmp/forge-sims-backup/`
2. Tag current state: `git tag pre-remediation-r26`
3. Revert individual changes as needed
4. Test after each rollback

---

**Owner:** Wez (AI)
**Status:** Plan created, awaiting implementation
**Next Step:** Begin Phase 1 (space-data-network-sim fixes)
