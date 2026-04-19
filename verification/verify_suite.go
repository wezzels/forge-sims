package main

// FORGE-Sim Verification Suite — Round 10
// Comprehensive evaluation of all simulator binaries

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

type Issue struct {
	Sim      string
	Type     string
	Message  string
	Severity string
}

type Warning struct {
	Sim     string
	Message string
}

var report struct {
	Timestamp  time.Time
	TotalSims  int
	PassCount  int
	FailCount  int
	WarnCount  int
	Issues     []Issue
	Warnings   []Warning
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║           FORGE-SIM VERIFICATION SUITE — Round 10                 ║")
	fmt.Println("║           Comprehensive Binary Evaluation                          ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
	fmt.Printf("Started: %s UTC\n\n", time.Now().UTC().Format(time.RFC3339))

	report.Timestamp = time.Now()
	report.Issues = []Issue{}
	report.Warnings = []Warning{}

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

	// Phase 2: Specialized Sims
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 2: SPECIALIZED SIMULATORS (12 binaries)")
	fmt.Println("================================================================================")

	specResults := verifyAllSpecializedSims(binDir)

	// Summary
	fmt.Println("\n================================================================================")
	fmt.Println("VERIFICATION SUMMARY")
	fmt.Println("================================================================================")

	computeSummary(bmdResults, specResults)
	saveReport()
}

func verifyAllBMDSSims(binDir string) []SimResult {
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
	result := SimResult{Name: simName, Type: "BMDS", Parameters: make(map[string]interface{})}

	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Binary not found: %s", binPath))
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "BINARY_MISSING", Message: "Binary not found", Severity: "CRITICAL"})
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
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.ValidJSON = false
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Invalid JSON: %v", err))
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "JSON_INVALID", Message: err.Error(), Severity: "CRITICAL"})
		return result
	}
	result.ValidJSON = true

	if _, ok := parsed["simulator"].(string); !ok {
		result.Issues = append(result.Issues, "Missing simulator field")
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "FIELD_MISSING", Message: "Missing simulator field", Severity: "HIGH"})
	}

	if _, ok := parsed["status"].(string); !ok {
		result.Issues = append(result.Issues, "Missing status field")
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "FIELD_MISSING", Message: "Missing status field", Severity: "HIGH"})
	}

	if params, ok := parsed["parameters"].(map[string]interface{}); ok {
		result.Parameters = params
		validateBMDSimParams(result, params)
	} else {
		result.Issues = append(result.Issues, "Missing parameters")
		report.Issues = append(report.Issues, Issue{Sim: simName, Type: "FIELD_MISSING", Message: "Missing parameters", Severity: "HIGH"})
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}
	return result
}

func validateBMDSimParams(result SimResult, params map[string]interface{}) {
	sim := result.Name

	// Add physics validation warnings
	switch sim {
	case "aegis":
		if v, ok := params["sm3_pk"].(float64); ok && (v <= 0 || v > 1) {
			result.Warnings = append(result.Warnings, "sm3_pk should be 0-1")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "sm3_pk should be 0-1"})
		}
	case "gmd":
		if v, ok := params["ekv_divert_ms"].(float64); ok && (v <= 0 || v > 5000) {
			result.Warnings = append(result.Warnings, "ekv_divert_ms should be 0-5000ms")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "ekv_divert_ms should be 0-5000ms"})
		}
	case "icbm":
		if v, ok := params["stages"].(float64); ok && (v < 1 || v > 5) {
			result.Warnings = append(result.Warnings, "icbm stages should be 1-5")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "icbm stages should be 1-5"})
		}
	case "sm3", "sm6":
		if v, ok := params["pk"].(float64); ok && (v <= 0 || v > 1) {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s pk should be 0-1", sim))
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: fmt.Sprintf("%s pk should be 0-1", sim)})
		}
	case "thaad", "patriot":
		if v, ok := params["pk"].(float64); ok && (v <= 0 || v > 1) {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s pk should be 0-1", sim))
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: fmt.Sprintf("%s pk should be 0-1", sim)})
		}
	case "sbirs":
		if v, ok := params["altitude_km"].(float64); ok && (v < 35000 || v > 40000) {
			result.Warnings = append(result.Warnings, "sbirs altitude should be 35,000-40,000 km (GEO)")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "sbirs altitude should be 35,000-40,000 km (GEO)"})
		}
	case "stss":
		if v, ok := params["altitude_km"].(float64); ok && (v < 1000 || v > 2000) {
			result.Warnings = append(result.Warnings, "stss altitude should be 1,000-2,000 km (LEO)")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "stss altitude should be 1,000-2,000 km (LEO)"})
		}
	case "link16":
		if v, ok := params["time_slots"].(float64); ok && v != 1535 {
			result.Warnings = append(result.Warnings, "link16 time_slots should be 1535 (MIL-STD-6016)")
			report.Warnings = append(report.Warnings, Warning{Sim: sim, Message: "link16 time_slots should be 1535 (MIL-STD-6016)"})
		}
	}
}

func verifyAllSpecializedSims(binDir string) []SimResult {
	return []SimResult{
		verifySatelliteTracker(binDir),
		verifySpaceDebris(binDir),
		verifyAirTraffic(binDir),
		verifyTacticalNet(binDir),
		verifyKillAssessment(binDir),
		verifyWTA(binDir),
		verifyElectronicWar(binDir),
		verifySubmarineWar(binDir),
		verifyCyberRedteam(binDir),
		verifyMaritime(binDir),
		verifySpaceWar(binDir),
	}
}

func verifySatelliteTracker(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/satellite-tracker", binDir)
	result := SimResult{Name: "satellite-tracker", Type: "Specialized", Parameters: make(map[string]interface{})}

	if _, err := os.Stat(binPath); os.IsNotExist(err) {
		result.Status = "ERROR"
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "BINARY_MISSING", Message: "Binary not found", Severity: "CRITICAL"})
		return result
	}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2", "-group", "gps-ops", "-count", "5")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "EXIT_CODE", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "CRITICAL"})
		return result
	}

	jsonStr := extractJSON(output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON in output")
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

	if params, ok := parsed["parameters"].(map[string]interface{}); ok {
		result.Parameters = params
	}

	if sats, ok := parsed["satellites"].([]interface{}); ok && len(sats) > 0 {
		result.Parameters["satellite_count"] = float64(len(sats))
	} else {
		result.Issues = append(result.Issues, "No satellites found")
		report.Issues = append(report.Issues, Issue{Sim: "satellite-tracker", Type: "NO_DATA", Message: "No satellites in output", Severity: "HIGH"})
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] satellite-tracker    | satellites=%d (GPS constellation)\n", result.Status, int(getFloat(result.Parameters, "satellite_count")))
	return result
}

func verifySpaceDebris(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/space-debris", binDir)
	result := SimResult{Name: "space-debris", Type: "Specialized", Parameters: make(map[string]interface{})}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

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
	} else {
		result.Issues = append(result.Issues, "No debris objects found")
		report.Issues = append(report.Issues, Issue{Sim: "space-debris", Type: "NO_DATA", Message: "No debris in output", Severity: "HIGH"})
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] space-debris        | debris_objects=%d\n", result.Status, int(getFloat(result.Parameters, "debris_count")))
	return result
}

func verifyAirTraffic(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/air-traffic", binDir)
	result := SimResult{Name: "air-traffic", Type: "Specialized", Parameters: make(map[string]interface{})}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

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
	} else {
		result.Issues = append(result.Issues, "No flights found")
		report.Issues = append(report.Issues, Issue{Sim: "air-traffic", Type: "NO_DATA", Message: "No flights in output", Severity: "HIGH"})
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] air-traffic          | flights=%d\n", result.Status, int(getFloat(result.Parameters, "flight_count")))
	return result
}

func verifyTacticalNet(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/tactical-net", binDir)
	result := SimResult{Name: "tactical-net", Type: "Specialized", Parameters: make(map[string]interface{})}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

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
			result.Parameters["node_count"] = nodes
		}
	}

	if links, ok := parsed["links"].([]interface{}); ok {
		result.Parameters["link_count"] = float64(len(links))
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] tactical-net         | nodes=%d, links=%d\n",
		result.Status, int(getFloat(result.Parameters, "node_count")), int(getFloat(result.Parameters, "link_count")))
	return result
}

func verifyKillAssessment(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/kill-assessment", binDir)
	result := SimResult{Name: "kill-assessment", Type: "Specialized", Parameters: make(map[string]interface{})}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-scenario", "single-icbm-ekv", "-seed", "42")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

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
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] kill-assessment      | kill_probability=%.2f\n", result.Status, getFloat(result.Parameters, "kill_probability"))
	return result
}

func verifyWTA(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/wta", binDir)
	result := SimResult{Name: "wta", Type: "Specialized", Parameters: make(map[string]interface{})}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()

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
		result.Parameters["threat_level"] = v
	}
	if v, ok := parsed["threat_type"].(string); ok {
		result.Parameters["threat_type"] = v
	}

	result.Status = getStatus(result.Issues)
	fmt.Printf("  [%s] wta                  | threat_level=%d, threat_type=%s\n",
		result.Status, int(getFloat(result.Parameters, "threat_level")), result.Parameters["threat_type"])
	return result
}

func verifyElectronicWar(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/electronic-war-sim", binDir)
	result := SimResult{Name: "electronic-war-sim", Type: "Specialized", Parameters: make(map[string]interface{})}

	// electronic-war-sim is a Go app that outputs package info, not JSON
	// Verify it executes without error
	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil || result.ExitCode != 0 {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "electronic-war-sim", Type: "EXECUTION", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] electronic-war-sim  | exit=%d\n", result.Status, result.ExitCode)
	return result
}

func verifySubmarineWar(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/submarine-war-sim", binDir)
	result := SimResult{Name: "submarine-war-sim", Type: "Specialized", Parameters: make(map[string]interface{})}

	// submarine-war-sim is a Go app that outputs package info, not JSON
	// Verify it executes without error
	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil || result.ExitCode != 0 {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		report.Issues = append(report.Issues, Issue{Sim: "submarine-war-sim", Type: "EXECUTION", Message: fmt.Sprintf("Exit code %d", result.ExitCode), Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] submarine-war-sim   | exit=%d\n", result.Status, result.ExitCode)
	return result
}

func verifyCyberRedteam(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/cyber-redteam-sim", binDir)
	result := SimResult{Name: "cyber-redteam-sim", Type: "Specialized", Parameters: make(map[string]interface{})}

	cmd := exec.Command(binPath, "-doctor")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "Doctor check failed")
		report.Issues = append(report.Issues, Issue{Sim: "cyber-redteam-sim", Type: "HEALTH_CHECK", Message: "Doctor check failed", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] cyber-redteam-sim    | doctor_exit=%d\n", result.Status, result.ExitCode)
	return result
}

func verifyMaritime(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/maritime-sim", binDir)
	result := SimResult{Name: "maritime-sim", Type: "Specialized", Parameters: make(map[string]interface{})}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "List scenarios failed")
		report.Issues = append(report.Issues, Issue{Sim: "maritime-sim", Type: "LIST_SCENARIOS", Message: "Failed to list scenarios", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] maritime-sim         | list_scenarios_exit=%d\n", result.Status, result.ExitCode)
	return result
}

func verifySpaceWar(binDir string) SimResult {
	binPath := fmt.Sprintf("%s/space-war-sim", binDir)
	result := SimResult{Name: "space-war-sim", Type: "Specialized", Parameters: make(map[string]interface{})}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "List scenarios failed")
		report.Issues = append(report.Issues, Issue{Sim: "space-war-sim", Type: "LIST_SCENARIOS", Message: "Failed to list scenarios", Severity: "HIGH"})
	} else {
		result.Status = "PASS"
	}

	fmt.Printf("  [%s] space-war-sim        | list_scenarios_exit=%d\n", result.Status, result.ExitCode)
	return result
}

func printBMDSResult(r SimResult) {
	statusIcon := "PASS"
	if r.Status != "PASS" {
		statusIcon = "FAIL"
	}

	fmt.Printf("  [%s] bmd-sim-%-15s params=%d", statusIcon, r.Name, len(r.Parameters))
	if len(r.Warnings) > 0 {
		fmt.Printf(" warnings=%d", len(r.Warnings))
	}
	fmt.Println()
}

func getStatus(issues []string) string {
	if len(issues) == 0 {
		return "PASS"
	}
	return "FAIL"
}

func computeSummary(bmdResults, specResults []SimResult) {
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
	}

	report.TotalSims = total
	report.PassCount = pass
	report.FailCount = fail
	report.WarnCount = warn

	bmdPass := 0
	bmdFail := 0
	for _, r := range bmdResults {
		if r.Status == "PASS" {
			bmdPass++
		} else {
			bmdFail++
		}
	}

	specPass := 0
	specFail := 0
	for _, r := range specResults {
		if r.Status == "PASS" {
			specPass++
		} else {
			specFail++
		}
	}

	fmt.Printf("\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("                        VERIFICATION RESULTS\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("  BMDS Simulators:     %d total | %d PASS | %d FAIL\n", len(bmdResults), bmdPass, bmdFail)
	fmt.Printf("  Specialized Sims:    %d total | %d PASS | %d FAIL\n", len(specResults), specPass, specFail)
	fmt.Printf("\n")
	fmt.Printf("  TOTAL: %d simulators | %d PASS | %d FAIL | %d WARNINGS\n", total, pass, fail, warn)
	fmt.Printf("\n")

	if fail == 0 {
		fmt.Printf("  ✅ ALL SIMULATORS PASSED VERIFICATION\n")
	} else {
		fmt.Printf("  ❌ %d failures detected - review EVALUATION.md\n", fail)
	}
	fmt.Printf("================================================================================\n")
}

func extractJSON(output []byte) string {
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "{") {
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
		fmt.Println("\n  No critical issues or warnings - EVALUATION.md not updated.")
		return
	}

	content := fmt.Sprintf("# FORGE-SIM VERIFICATION REPORT\n\n**Date:** %s UTC\n**Total Simulators:** %d\n**Passed:** %d\n**Failed:** %d\n**Warnings:** %d\n\n",
		report.Timestamp.UTC().Format(time.RFC3339), report.TotalSims, report.PassCount, report.FailCount, report.WarnCount)

	if len(report.Issues) > 0 {
		content += "## Issues\n\n"
		for _, issue := range report.Issues {
			icon := "🔴"
			if issue.Severity == "HIGH" {
				icon = "🟡"
			}
			content += fmt.Sprintf("- %s **%s**: %s\n", icon, issue.Sim, issue.Message)
		}
		content += "\n"
	}

	if len(report.Warnings) > 0 {
		content += "## Warnings\n\n"
		for _, w := range report.Warnings {
			content += fmt.Sprintf("- **%s**: %s\n", w.Sim, w.Message)
		}
		content += "\n"
	}

	content += "---\n*Generated by FORGE-Sim Verification Suite*\n"

	os.WriteFile("../EVALUATION.md", []byte(content), 0644)
	fmt.Printf("\n  Report saved to EVALUATION.md (%d issues, %d warnings)\n", len(report.Issues), len(report.Warnings))
}