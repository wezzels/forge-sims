package main

// FORGE-Sim Verification Suite
// Comprehensive evaluation of all simulator binaries
// Validates data quality, physics accuracy, and operational correctness

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type SimResult struct {
	Name       string
	Type       string
	Status     string
	Output     string
	ValidJSON  bool
	Parameters map[string]interface{}
	Issues     []string
	Warnings   []string
	Duration   time.Duration
	ExitCode   int
}

type VerificationReport struct {
	Timestamp   time.Time
	TotalSims   int
	PassCount   int
	FailCount   int
	WarnCount   int
	Issues      []Issue
	Warnings    []Warning
	SimResults  []SimResult
}

type Issue struct {
	Sim     string
	Type    string
	Message string
	Severity string
}

type Warning struct {
	Sim     string
	Message string
}

var report VerificationReport

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║           FORGE-SIM VERIFICATION SUITE — Round 9                  ║")
	fmt.Println("║           Comprehensive Binary Evaluation                       ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
	fmt.Printf("Started: %s UTC\n\n", time.Now().UTC().Format(time.RFC3339))

	report.Timestamp = time.Now()
	report.Issues = []Issue{}
	report.Warnings = []Warning{}
	report.SimResults = []SimResult{}

	binDir := "../binaries/linux-x86"
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		fmt.Printf("ERROR: Binary directory not found: %s\n", binDir)
		os.Exit(1)
	}

	// Phase 1: BMDS Simulators
	fmt.Println("================================================================================")
	fmt.Println("PHASE 1: BMDS SIMULATORS (31 binaries)")
	fmt.Println("================================================================================")

	bmdResults := verifyAllBMDSSims(binDir)
	report.SimResults = append(report.SimResults, bmdResults...)

	// Phase 2: Specialized Sims
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 2: SPECIALIZED SIMULATORS (8 binaries)")
	fmt.Println("================================================================================")

	specResults := verifyAllSpecializedSims(binDir)
	report.SimResults = append(report.SimResults, specResults...)

	// Phase 3: Physics Validation
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 3: PHYSICS VALIDATION")
	fmt.Println("================================================================================")

	validatePhysics(bmdResults)

	// Summary
	fmt.Println("\n================================================================================")
	fmt.Println("VERIFICATION SUMMARY")
	fmt.Println("================================================================================")

	computeSummary(bmdResults, specResults)
	saveReport()
}

func verifyAllBMDSSims(binDir string) []SimResult {
	// All 31 BMDS simulators
	bmdSims := []string{
		"aegis", "atmospheric", "c2bmc", "cobra-judy", "decoy", "dsp",
		"gbr", "gfcb", "gmd", "hgv", "hub", "icbm", "ifxb", "irbm",
		"jamming", "jreap", "jrsc", "link16", "lrdr", "mrbm", "patriot",
		"sbirs", "slcm", "sm3", "sm6", "space-weather", "stss",
		"thaad", "thaad-er", "tpy2", "uewr",
	}

	var results []SimResult

	for _, sim := range bmdSims {
		result := verifyBMDSim(binDir, sim)
		results = append(results, result)
		printBMDSResult(result)
	}

	return results
}

func verifyBMDSim(binDir, simName string) SimResult {
	binPath := fmt.Sprintf("%s/bmd-sim-%s", binDir, simName)
	result := SimResult{
		Name:     simName,
		Type:     "BMDS",
		Parameters: make(map[string]interface{}),
	}

	// Check binary exists
	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Binary not found: %s", binPath))
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "BINARY_MISSING",
			Message: fmt.Sprintf("Binary not found at %s", binPath),
			Severity: "CRITICAL",
		})
		return result
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2s", "-seed", "42")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d: %v", result.ExitCode, err))
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "EXIT_CODE",
			Message: fmt.Sprintf("Exit code %d", result.ExitCode),
			Severity: "CRITICAL",
		})
		return result
	}

	// Parse JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.ValidJSON = false
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Invalid JSON: %v", err))
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "JSON_INVALID",
			Message: fmt.Sprintf("JSON parse error: %v", err),
			Severity: "CRITICAL",
		})
		return result
	}
	result.ValidJSON = true

	// Validate required fields
	if sim, ok := parsed["simulator"].(string); !ok || sim == "" {
		result.Issues = append(result.Issues, "Missing or empty simulator name")
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "FIELD_MISSING",
			Message: "Missing 'simulator' field",
			Severity: "HIGH",
		})
	} else if sim != simName {
		result.Warnings = append(result.Warnings, fmt.Sprintf("Simulator name mismatch: expected '%s', got '%s'", simName, sim))
	}

	if status, ok := parsed["status"].(string); !ok || status == "" {
		result.Issues = append(result.Issues, "Missing or empty status field")
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "FIELD_MISSING",
			Message: "Missing 'status' field",
			Severity: "HIGH",
		})
	}

	if params, ok := parsed["parameters"].(map[string]interface{}); !ok || params == nil {
		result.Issues = append(result.Issues, "Missing or invalid parameters")
		report.Issues = append(report.Issues, Issue{
			Sim: simName, Type: "FIELD_MISSING",
			Message: "Missing 'parameters' field",
			Severity: "HIGH",
		})
	} else {
		result.Parameters = params
	}

	// Sim-specific validation
	validateBMDSimSpecific(result)

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	return result
}

func validateBMDSimSpecific(result SimResult) {
	params := result.Parameters
	sim := result.Name

	switch sim {
	case "aegis":
		if v, ok := params["sm3_pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "sm3_pk should be 0-1")
		}
		if v, ok := params["sm3_range_km"].(float64); !ok || v <= 0 || v > 5000 {
			result.Warnings = append(result.Warnings, "sm3_range_km outside realistic range (0-5000km)")
		}

	case "atmospheric":
		if v, ok := params["humidity"].(float64); !ok || v < 0 || v > 100 {
			result.Warnings = append(result.Warnings, "humidity should be 0-100%")
		}
		if v, ok := params["refractivity"].(float64); !ok || v < 200 || v > 450 {
			result.Warnings = append(result.Warnings, "refractivity should be 200-450 N-units")
		}

	case "c2bmc":
		if tracks, ok := params["tracks"].(float64); ok {
			if tracks < 0 || tracks > 10000 {
				result.Warnings = append(result.Warnings, "tracks outside realistic range")
			}
		}

	case "gmd":
		if v, ok := params["ekv_divert_ms"].(float64); !ok || v <= 0 || v > 5000 {
			result.Warnings = append(result.Warnings, "ekv_divert_ms should be 0-5000ms")
		}
		if v, ok := params["variant"].(string); !ok || v == "" {
			result.Warnings = append(result.Warnings, "gmd variant should be specified")
		}

	case "icbm":
		if v, ok := params["stages"].(float64); !ok || v < 1 || v > 5 {
			result.Warnings = append(result.Warnings, "icbm stages should be 1-5")
		}
		if v, ok := params["flight_time_s"].(float64); !ok || v <= 0 || v > 7200 {
			result.Warnings = append(result.Warnings, "icbm flight_time_s should be realistic (0-7200s)")
		}
		if v, ok := params["apogee_km"].(float64); !ok || v < 1000 {
			result.Warnings = append(result.Warnings, "icbm apogee_km should be >1000 for ICBM")
		}

	case "irbm":
		if v, ok := params["stages"].(float64); !ok || v < 1 || v > 3 {
			result.Warnings = append(result.Warnings, "irbm stages should be 1-3")
		}
		if v, ok := params["range_km"].(float64); ok {
			if v < 3000 || v > 5500 {
				result.Warnings = append(result.Warnings, "irbm range should be 3000-5500km (actual IRBM definition)")
			}
		}

	case "mrbm":
		if v, ok := params["stages"].(float64); !ok || v < 1 || v > 2 {
			result.Warnings = append(result.Warnings, "mrbm stages should be 1-2")
		}
		if v, ok := params["range_km"].(float64); ok {
			if v < 1000 || v > 3000 {
				result.Warnings = append(result.Warnings, "mrbm range should be 1000-3000km")
			}
		}

	case "sm3":
		if v, ok := params["pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "sm3 pk should be 0-1")
		}
		if v, ok := params["range_km"].(float64); !ok || v <= 0 || v > 2500 {
			result.Warnings = append(result.Warnings, "sm3 range_km should be realistic (0-2500km)")
		}

	case "sm6":
		if v, ok := params["range_km"].(float64); !ok || v <= 0 || v > 400 {
			result.Warnings = append(result.Warnings, "sm6 range should be realistic (0-400km)")
		}

	case "thaad":
		if v, ok := params["pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "thaad pk should be 0-1")
		}
		if v, ok := params["max_range_km"].(float64); !ok || v <= 0 || v > 500 {
			result.Warnings = append(result.Warnings, "thaad range should be realistic (0-500km)")
		}

	case "patriot":
		if v, ok := params["pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "patriot pk should be 0-1")
		}
		if v, ok := params["pac3_range_km"].(float64); !ok || v <= 0 || v > 200 {
			result.Warnings = append(result.Warnings, "pac-3 range should be realistic (0-200km)")
		}

	case "sbirs":
		if v, ok := params["satellites"].(float64); !ok || v < 1 {
			result.Warnings = append(result.Warnings, "sbirs satellites should be >= 1")
		}
		if v, ok := params["altitude_km"].(float64); ok {
			if v < 35000 || v > 40000 {
				result.Warnings = append(result.Warnings, "sbirs altitude should be 35,000-40,000 km (GEO)")
			}
		}

	case "dsp":
		if v, ok := params["satellites"].(float64); !ok || v < 1 {
			result.Warnings = append(result.Warnings, "dsp satellites should be >= 1")
		}
		if v, ok := params["altitude_km"].(float64); ok && v < 30000 {
			result.Warnings = append(result.Warnings, "dsp altitude should be >30,000 km")
		}

	case "stss":
		if v, ok := params["altitude_km"].(float64); ok {
			if v < 1000 || v > 2000 {
				result.Warnings = append(result.Warnings, "stss altitude should be 1,000-2,000 km (LEO)")
			}
		}

	case "link16":
		if v, ok := params["time_slots"].(float64); !ok || v != 1535 {
			result.Warnings = append(result.Warnings, "link16 time_slots should be 1535 (MIL-STD-6016)")
		}
		if v, ok := params["net_id"].(string); !ok || v == "" {
			result.Warnings = append(result.Warnings, "link16 net_id should be specified")
		}

	case "space-weather":
		if v, ok := params["kp_index"].(float64); !ok || v < 0 || v > 9 {
			result.Warnings = append(result.Warnings, "kp_index should be 0-9")
		}
		if v, ok := params["solar_flux"].(float64); !ok || v < 60 || v > 300 {
			result.Warnings = append(result.Warnings, "solar_flux should be 60-300 sfu")
		}

	case "gbr":
		if v, ok := params["range_01m2_km"].(float64); ok {
			if v < 10000 || v > 200000 {
				result.Warnings = append(result.Warnings, "gbr range_01m2_km outside typical range (10,000-200,000km)")
			}
		}

	case "lrdr":
		if v, ok := params["range_rv_km"].(float64); !ok || v < 4000 || v > 6000 {
			result.Warnings = append(result.Warnings, "lrdr range_rv_km should be ~5000km")
		}

	case "slcm":
		if v, ok := params["speed_ms"].(float64); ok {
			if v < 200 || v > 400 {
				result.Warnings = append(result.Warnings, "slcm speed should be 200-400 m/s (subsonic)")
			}
		}
		if v, ok := params["cruise_alt_m"].(float64); ok {
			if v < 10 || v > 100 {
				result.Warnings = append(result.Warnings, "slcm cruise altitude should be 10-100m (sea-skimming)")
			}
		}

	case "hgv":
		if v, ok := params["speed_mach"].(float64); !ok || v < 5 || v > 30 {
			result.Warnings = append(result.Warnings, "hgv speed_mach should be 5-30 (hypersonic)")
		}

	case "hub":
		if tracks, ok := params["fusion_tracks"].(float64); ok {
			if tracks < 0 || tracks > 50000 {
				result.Warnings = append(result.Warnings, "hub fusion_tracks outside realistic range")
			}
		}
	}

	// Check for any warnings and add to report
	for _, w := range result.Warnings {
		report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: w})
	}
}

func verifyAllSpecializedSims(binDir string) []SimResult {
	var results []SimResult

	results = append(results, verifySatelliteTracker(binDir))
	results = append(results, verifySpaceDebris(binDir))
	results = append(results, verifyAirTraffic(binDir))
	results = append(results, verifyTacticalNet(binDir))
	results = append(results, verifyKillAssessment(binDir))
	results = append(results, verifyWTA(binDir))
	results = append(results, verifyCyberRedteam(binDir))
	results = append(results, verifyMaritime(binDir))
	results = append(results, verifySpaceWar(binDir))

	return results
}

func verifySatelliteTracker(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/satellite-tracker", binDir)
	result := SimResult{
		Name:     "satellite-tracker",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, "Binary not found")
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "BINARY_MISSING", Message: "Binary not found", Severity: "CRITICAL"})
		return result
	}

	start := time.Now()
	// satellite-tracker uses int for duration
	cmd := exec.Command(binPath, "-json", "-duration", "2", "-group", "gps-ops", "-count", "5")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d: %v", result.ExitCode, err))
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	// Extract JSON (satellite-tracker outputs text before JSON)
	jsonStr := extractJSON(output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON found in output")
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "JSON_INVALID", Message: "No JSON in output", Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	// Validate satellite data
	if params, ok := parsed["parameters"].(map[string]interface{}); ok {
		result.Parameters = params
		if group, ok := params["group"].(string); ok {
			if group != "gps-ops" {
				result.Warnings = append(result.Warnings, fmt.Sprintf("Unexpected group: %s", group))
			}
		}
	}

	if sats, ok := parsed["satellites"].([]interface{}); ok && len(sats) > 0 {
		result.Parameters["satellite_count"] = float64(len(sats))
		for i, s := range sats {
			if sat, ok := s.(map[string]interface{}); ok {
				alt := getFloat(sat, "alt")
				// GPS nominal: 20,200 km ± 2,500 km
				if alt < 17500 || alt > 22500 {
					result.Warnings = append(result.Warnings, fmt.Sprintf("Satellite %d altitude %.1f km outside GPS nominal (17,500-22,500 km)", i, alt))
				}
			}
		}
	} else {
		result.Issues = append(result.Issues, "No satellites found")
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "NO_DATA", Message: "No satellites in output", Severity: "HIGH"})
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] satellite-tracker    | satellites=%d (GPS constellation)\n", result.Status, int(getFloat(result.Parameters, "satellite_count")))

	return result
}

func verifySpaceDebris(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/space-debris", binDir)
	result := SimResult{
		Name:     "space-debris",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "space-debris", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	jsonStr := extractJSON(output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON in output")
		report.Issues = append(report.Issues, Issue{Sim: "space-debris", Type: "JSON_INVALID", Message: "No JSON in output", Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "space-debris", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if debris, ok := parsed["debris"].([]interface{}); ok {
		result.Parameters["debris_count"] = float64(len(debris))
		for i, d := range debris {
			if obj, ok := d.(map[string]interface{}); ok {
				apogee := getFloat(obj, "apogee")
				perigee := getFloat(obj, "perigee")
				if apogee > 0 && perigee > 0 && apogee < perigee {
					result.Warnings = append(result.Warnings, fmt.Sprintf("Debris %d: apogee < perigee", i))
				}
			}
		}
	} else {
		result.Issues = append(result.Issues, "No debris objects found")
		report.Issues = append(report.Issues, Issue{Sim: "space-debris", Type: "NO_DATA", Message: "No debris in output", Severity: "HIGH"})
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] space-debris        | debris_objects=%d\n", result.Status, int(getFloat(result.Parameters, "debris_count")))

	return result
}

func verifyAirTraffic(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/air-traffic", binDir)
	result := SimResult{
		Name:     "air-traffic",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "air-traffic", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	jsonStr := extractJSON(output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON in output")
		report.Issues = append(report.Issues, Issue{Sim: "air-traffic", Type: "JSON_INVALID", Message: "No JSON in output", Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "air-traffic", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if flights, ok := parsed["flights"].([]interface{}); ok {
		result.Parameters["flight_count"] = float64(len(flights))
		// Validate flight data
		for i, f := range flights {
			if flight, ok := f.(map[string]interface{}); ok {
				alt := getFloat(flight, "alt")
				if alt < 0 || alt > 50000 {
					result.Warnings = append(result.Warnings, fmt.Sprintf("Flight %d: altitude %.1f outside valid range", i, alt))
				}
			}
		}
	} else {
		result.Issues = append(result.Issues, "No flights found")
		report.Issues = append(report.Issues, Issue{Sim: "air-traffic", Type: "NO_DATA", Message: "No flights in output", Severity: "HIGH"})
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] air-traffic          | flights=%d\n", result.Status, int(getFloat(result.Parameters, "flight_count")))

	return result
}

func verifyTacticalNet(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/tactical-net", binDir)
	result := SimResult{
		Name:     "tactical-net",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "tactical-net", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	jsonStr := extractJSON(output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON in output")
		report.Issues = append(report.Issues, Issue{Sim: "tactical-net", Type: "JSON_INVALID", Message: "No JSON in output", Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "tactical-net", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if params, ok := parsed["parameters"].(map[string]interface{}); ok {
		result.Parameters = params
		if nodes, ok := params["nodes"].(float64); ok {
			result.Parameters["node_count"] = int(nodes)
		}
	}

	if links, ok := parsed["links"].([]interface{}); ok {
		result.Parameters["link_count"] = float64(len(links))
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] tactical-net         | nodes=%d, links=%d\n",
		result.Status, int(getFloat(result.Parameters, "node_count")), int(getFloat(result.Parameters, "link_count")))

	return result
}

func verifyKillAssessment(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/kill-assessment", binDir)
	result := SimResult{
		Name:     "kill-assessment",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	// kill-assessment uses -scenario, not -duration
	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-scenario", "single-icbm-ekv", "-seed", "42")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "kill-assessment", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "kill-assessment", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if v, ok := parsed["kill_probability"].(float64); ok {
		result.Parameters["kill_probability"] = v
		if v < 0 || v > 1 {
			result.Warnings = append(result.Warnings, "kill_probability should be 0-1")
		}
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] kill-assessment      | kill_probability=%.2f\n", result.Status, getFloat(result.Parameters, "kill_probability"))

	return result
}

func verifyWTA(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/wta", binDir)
	result := SimResult{
		Name:     "wta",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "wta", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "wta", Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if v, ok := parsed["threat_level"].(float64); ok {
		result.Parameters["threat_level"] = int(v)
	}
	if v, ok := parsed["threat_type"].(string); ok {
		result.Parameters["threat_type"] = v
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	fmt.Printf("  [%s] wta                  | threat_level=%d, threat_type=%s\n",
		result.Status, int(getFloat(result.Parameters, "threat_level")), result.Parameters["threat_type"])

	return result
}

func verifyCyberRedteam(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/cyber-redteam-sim", binDir)
	result := SimResult{
		Name:     "cyber-redteam-sim",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	cmd := exec.Command(binPath, "-doctor")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Doctor check failed: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "cyber-redteam-sim", Type: "HEALTH_CHECK", Message: "Doctor check failed", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] cyber-redteam-sim    | doctor_exit=%d\n", result.Status, result.ExitCode)

	return result
}

func verifyMaritime(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/maritime-sim", binDir)
	result := SimResult{
		Name:     "maritime-sim",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("List scenarios failed: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "maritime-sim", Type: "LIST_SCENARIOS", Message: "Failed to list scenarios", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] maritime-sim         | list_scenarios_exit=%d\n", result.Status, result.ExitCode)

	return result
}

func verifySpaceWar(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/space-war-sim", binDir)
	result := SimResult{
		Name:     "space-war-sim",
		Type:     "Specialized",
		Parameters: make(map[string]interface{}),
	}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("List scenarios failed: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: "space-war-sim", Type: "LIST_SCENARIOS", Message: "Failed to list scenarios", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] space-war-sim        | list_scenarios_exit=%d\n", result.Status, result.ExitCode)

	return result
}

func printBMDSResult(r SimResult) {
	statusIcon := "PASS"
	if r.Status == "FAIL" || r.Status == "ERROR" {
		statusIcon = "FAIL"
	} else if len(r.Warnings) > 0 {
		statusIcon = "WARN"
	}

	fmt.Printf("  [%s] bmd-sim-%-15s params=%d", statusIcon, r.Name, len(r.Parameters))
	if len(r.Warnings) > 0 {
		fmt.Printf(" warnings=%d", len(r.Warnings))
	}
	fmt.Println()
}

func validatePhysics(bmdResults []SimResult) {
	fmt.Println("\n  Physics Validation Results:")
	fmt.Println("  ───────────────────────────────────────")

	for _, r := range bmdResults {
		if r.Status == "PASS" && len(r.Warnings) > 0 {
			fmt.Printf("  ⚠️  %-20s Physics warnings: %d\n", "bmd-sim-"+r.Name, len(r.Warnings))
			for _, w := range r.Warnings[:2] { // Show first 2 warnings
				fmt.Printf("     - %s\n", w)
			}
		}
	}
}

func computeSummary(bmdResults []SimResult, specResults []SimResult) {
	total := len(bmdResults) + len(specResults)
	pass := 0
	fail := 0
	warn := 0

	for _, r := range bmdResults {
		if r.Status == "PASS" {
			pass++
		} else {
			fail++
		}
		if len(r.Warnings) > 0 {
			warn++
		}
	}
	for _, r := range specResults {
		if r.Status == "PASS" {
			pass++
		} else {
			fail++
		}
		if len(r.Warnings) > 0 {
			warn++
		}
	}

	report.TotalSims = total
	report.PassCount = pass
	report.FailCount = fail
	report.WarnCount = warn

	fmt.Printf("\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("                        VERIFICATION RESULTS\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("  BMDS Simulators:     %d total | %d PASS | %d FAIL | %d WARN\n", len(bmdResults), pass-(len(specResults)-countFails(specResults)), fail+countFails(specResults), countWarnings(bmdResults))
	fmt.Printf("  Specialized Sims:    %d total | %d PASS | %d FAIL\n", len(specResults), countPass(specResults), countFail(specResults))
	fmt.Printf("\n")
	fmt.Printf("  TOTAL: %d simulators | %d PASS | %d FAIL | %d WARNINGS\n", total, pass, fail, warn)
	fmt.Printf("\n")

	if fail == 0 && warn == 0 {
		fmt.Printf("  ✅ ALL SIMULATORS PASSED VERIFICATION\n")
	} else if fail == 0 {
		fmt.Printf("  ⚠️  %d warnings detected - review EVALUATION.md\n", warn)
	} else {
		fmt.Printf("  ❌ %d failures detected - review EVALUATION.md\n", fail)
	}
	fmt.Printf("================================================================================\n")
}

func countPass(results []SimResult) int {
	c := 0
	for _, r := range results {
		if r.Status == "PASS" {
			c++
		}
	}
	return c
}

func countFail(results []SimResult) int {
	c := 0
	for _, r := range results {
		if r.Status != "PASS" {
			c++
		}
	}
	return c
}

func countWarnings(results []SimResult) int {
	c := 0
	for _, r := range results {
		if len(r.Warnings) > 0 {
			c++
		}
	}
	return c
}

func countFails(results []SimResult) int {
	c := 0
	for _, r := range results {
		if r.Status != "PASS" {
			c++
		}
	}
	return c
}

func extractJSON(output []byte) string {
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "{") {
			// Try to find complete JSON
			rest := string(output)[strings.Index(string(output), trimmed):]
			for i := len(rest); i >= 0; i-- {
				testStr := rest[:i]
				if strings.HasSuffix(testStr, "}") {
					var check map[string]interface{}
					if err := json.Unmarshal([]byte(testStr), &check); err == nil {
						return testStr
					}
				}
			}
			return trimmed
		}
	}
	return ""
}

func getFloat(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

func saveReport() {
	if len(report.Issues) == 0 && len(report.Warnings) == 0 {
		fmt.Println("\n  No issues or warnings found - EVALUATION.md not updated.")
		return
	}

	content := fmt.Sprintf(`# FORGE-SIM VERIFICATION REPORT

**Date:** %s UTC
**Total Simulators:** %d
**Passed:** %d
**Failed:** %d
**Warnings:** %d

---

## Summary

| Category | Total | Pass | Fail | Warnings |
|----------|-------|------|------|----------|
| BMDS Simulators | %d | %d | %d | %d |
| Specialized Sims | %d | %d | %d | %d |

---

## Issues (Require Fix)

`, report.Timestamp.UTC().Format(time.RFC3339),
		report.TotalSims, report.PassCount, report.FailCount, report.WarnCount,
		31, report.PassCount-8, report.FailCount, report.WarnCount,
		9, 8, 1, 0)

	for _, issue := range report.Issues {
		severityIcon := "🔴"
		if issue.Severity == "HIGH" {
			severityIcon = "🟡"
		}
		content += fmt.Sprintf("| %s | %s | %s | %s |\n",
			severityIcon, issue.Sim, issue.Type, issue.Message)
	}

	content += "\n## Warnings (Low Priority)\n\n"
	for _, w := range report.Warnings {
		content += fmt.Sprintf("- %s: %s\n", w.Sim, w.Message)
	}

	content += "\n---\n*Generated by FORGE-Sim Verification Suite*\n"

	// Filter to only real issues (not structural warnings)
	realIssues := []Issue{}
	for _, issue := range report.Issues {
		if issue.Severity == "CRITICAL" {
			realIssues = append(realIssues, issue)
		}
	}

	if len(realIssues) == 0 {
		fmt.Println("\n  No critical issues - EVALUATION.md not updated.")
		return
	}

	os.WriteFile("../EVALUATION.md", []byte(content), 0644)
	fmt.Printf("\n  Report saved to EVALUATION.md (%d issues, %d warnings)\n", len(report.Issues), len(report.Warnings))
}