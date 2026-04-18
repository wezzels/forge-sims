# FORGE-Sims Verification Suite

Verification and validation test suite for all FORGE simulator binaries.

## Overview

This suite validates that all FORGE simulators produce:
- **Valid JSON output** — parseable by downstream systems
- **Realistic parameters** — physics-validated data
- **Correct binary execution** — proper exit codes, no crashes

## Running the Suite

```bash
cd verification
go run verify_suite.go
```

### Output

The suite prints results to stdout and saves issues to `../EVALUATION.md`.

```
╔════════════════════════════════════════════════════════════════════╗
║           FORGE-SIMS VERIFICATION SUITE                         ║
║           Binary Quality & Data Validation                      ║
╚════════════════════════════════════════════════════════════════════╝
Started: 2026-04-18T15:04:54Z UTC

================================================================================
PHASE 1: BMDS SIMULATORS (30 binaries)
================================================================================
  [PASS] bmd-sim-aegis           params=7
  [PASS] bmd-sim-atmospheric     params=4
  ...
```

## Adding New Simulators

### Step 1: Identify the Sim Type

FORGE simulators fall into two categories:

**BMDS Simulators (30)** — Use standard flags:
```bash
./bmd-sim-<name> -json -duration 2s -seed 42
```

**Specialized Simulators** — Have custom flags (see `verify_suite.go` for examples).

### Step 2: Add to the Verification Map

Edit `verify_suite.go` and add the sim to the appropriate function:

#### For BMDS Simulators (all use same interface)

Add to the `bmdSims` slice in `verifyBMDSSims()`:

```go
bmdSims := []string{
    "aegis", "atmospheric", "c2bmc", ...
    "your-new-sim",  // <-- ADD HERE
}
```

The BMDS verification automatically tests:
- JSON parseability
- Required fields: `simulator`, `status`, `parameters`
- Exit code = 0

#### For Specialized Simulators

Add a new verification function following this pattern:

```go
func verifyYourNewSim(binDir string) YourResult {
    binPath := fmt.Sprintf("%s/your-new-sim", binDir)
    result := YourResult{SimResult: SimResult{Name: "your-new-sim"}}

    start := time.Now()
    // Use correct flags for this sim
    cmd := exec.Command(binPath, "-json", "-duration", "2")
    output, err := cmd.CombinedOutput()
    result.Duration = time.Since(start)
    result.ExitCode = cmd.ProcessState.ExitCode()
    result.Output = string(output)

    if err != nil && result.ExitCode != 0 {
        result.Status = "ERROR"
        result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
        issues = append(issues, fmt.Sprintf("your-new-sim: exit code %d", result.ExitCode))
        return result
    }

    // Handle text-prefixed output (satellite-tracker pattern)
    jsonStr := findJSONStart(result.Output)
    if jsonStr == "" {
        result.Status = "FAIL"
        result.Issues = append(result.Issues, "No JSON found in output")
        issues = append(issues, "your-new-sim: no JSON in output")
        return result
    }

    var parsed map[string]interface{}
    if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
        result.Status = "FAIL"
        result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
        issues = append(issues, "your-new-sim: JSON parse error")
        return result
    }
    result.ValidJSON = true

    // Add domain-specific validation...
    // e.g., validate orbital mechanics, check parameter ranges

    if len(result.Issues) == 0 {
        result.Status = "PASS"
    }

    fmt.Printf("  [PASS] your-new-sim        | detail=%s", yourField)
    fmt.Println()

    return result
}
```

### Step 3: Call from main()

In `main()`, add your verification function to the appropriate phase:

```go
// Phase N: Your New Simulator
fmt.Println("\n================================================================================")
fmt.Println("PHASE N: YOUR NEW SIM")
fmt.Println("================================================================================")

yourResult := verifyYourNewSim(binDir)
```

### Step 4: Add to Summary

Update `printSummary()` to include your result:

```go
fmt.Printf("    Your New Sim:       %s (detail: %s)\n", yourResult.Status, yourResult.Field)
```

## Validation Rules

### JSON Structure

All sims must output valid JSON with:

```json
{
  "simulator": "sim-name",
  "status": "complete",
  "parameters": { ... }
}
```

### Physics Validation

| Sim Type | Validation |
|----------|-------------|
| Satellite Tracker | GPS altitude 18,000-25,000 km |
| Space Debris | Apogee > Perigee |
| Air Traffic | Altitude 3,000-45,000 ft |

### Parameter Ranges

| Parameter | Valid Range |
|-----------|-------------|
| pk (kill probability) | 0-1 |
| kp_index | 0-9 |
| time_slots (Link 16) | 1535 |
| humidity | 0-100% |

## Test Parameters Reference

| Simulator | Command |
|-----------|---------|
| BMDS (standard) | `-json -duration 2s -seed 42` |
| satellite-tracker | `-json -duration 2 -group gps-ops -count 5` |
| space-debris | `-json -duration 2` |
| air-traffic | `-json -duration 2` |
| tactical-net | `-json -duration 2` |
| kill-assessment | `-json -scenario single-icbm-ekv -seed 42` |
| wta | `-json -duration 2` |
| cyber-redteam-sim | `-doctor` |
| maritime-sim | `-list-scenarios` |
| space-war-sim | `-list-scenarios` |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid flag/argument |

## Continuous Integration

Add to CI pipeline:

```bash
#!/bin/bash
cd verification
go test -v ./... || exit 1
go run verify_suite.go
```

## Troubleshooting

### "No JSON found in output"
The sim outputs text before JSON. Use `findJSONStart()` to extract JSON from output:

```go
jsonStr := findJSONStart(output)
// then parse jsonStr
```

### "Invalid flag: -duration"
Some sims use different flag names. Check `--help` and update the test command.

### "Exit code 2"
Flag validation failed. Check sim's help output for correct flag names.

---

*Last Updated: 2026-04-18*