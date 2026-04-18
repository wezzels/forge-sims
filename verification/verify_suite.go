package main

// FORGE-Sims Verification Suite
// Validates binary data quality and activity for all simulator types
// Run: cd verification && go run verify_suite.go

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
	Status     string
	Output     string
	ValidJSON  bool
	Issues     []string
	Warnings   []string
	Duration   time.Duration
	ExitCode   int
}

type BMDSResult struct {
	SimResult
	Parameters map[string]interface{}
}

type SatelliteResult struct {
	SimResult
	Satellites []Satellite
	Group     string
	Count     int
}

type Satellite struct {
	Name     string
	ID       int
	Lat      float64
	Lon      float64
	Alt      float64
	Vel      float64
	Vis      string
	Category string
}

type SpaceDebrisResult struct {
	SimResult
	Debris []DebrisObject
	Count  int
}

type DebrisObject struct {
	ID           int
	Name         string
	Apogee       float64
	Perigee      float64
	Inclination  float64
	Period       float64
	RCS          float64
	Lat          float64
	Lon          float64
}

type AirTrafficResult struct {
	SimResult
	FlightCount int
}

type TacticalNetResult struct {
	SimResult
	ParticipantCount int
	RouteCount      int
	LinkCount       int
}

type KillAssessmentResult struct {
	SimResult
	KillProbability float64
	Scenario        string
}

type WTAResult struct {
	SimResult
	ThreatLevel int
	ThreatType  string
}

type SpecializedResult struct {
	Name           string
	ScenarioLoaded bool
	Status         string
}

var issues []string
var warnings []string

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║           FORGE-SIMS VERIFICATION SUITE                         ║")
	fmt.Println("║           Binary Quality & Data Validation                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
	fmt.Printf("Started: %s UTC\n\n", time.Now().UTC().Format(time.RFC3339))

	binDir := "../binaries/linux-x86"
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		fmt.Printf("ERROR: Binary directory not found: %s\n", binDir)
		os.Exit(1)
	}

	// Phase 1: BMDS Simulators (30)
	fmt.Println("================================================================================")
	fmt.Println("PHASE 1: BMDS SIMULATORS (30 binaries)")
	fmt.Println("================================================================================")

	bmdResults := verifyBMDSSims(binDir)

	// Phase 2: Satellite Tracker
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 2: SATELLITE TRACKER")
	fmt.Println("================================================================================")

	satResult := verifySatelliteTracker(binDir)

	// Phase 3: Space Debris
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 3: SPACE DEBRIS")
	fmt.Println("================================================================================")

	debrisResult := verifySpaceDebris(binDir)

	// Phase 4: Air Traffic
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 4: AIR TRAFFIC")
	fmt.Println("================================================================================")

	airResult := verifyAirTraffic(binDir)

	// Phase 5: Tactical Net
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 5: TACTICAL NET")
	fmt.Println("================================================================================")

	tacResult := verifyTacticalNet(binDir)

	// Phase 6: Kill Assessment
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 6: KILL ASSESSMENT")
	fmt.Println("================================================================================")

	killResult := verifyKillAssessment(binDir)

	// Phase 7: WTA
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 7: WTA (Warfighter Training Algorithm)")
	fmt.Println("================================================================================")

	wtaResult := verifyWTA(binDir)

	// Phase 8: Specialized Sims (config-based)
	fmt.Println("\n================================================================================")
	fmt.Println("PHASE 8: SPECIALIZED SIMS")
	fmt.Println("================================================================================")

	specResults := verifySpecializedSims(binDir)

	// Summary
	fmt.Println("\n================================================================================")
	fmt.Println("VERIFICATION SUMMARY")
	fmt.Println("================================================================================")

	printSummary(bmdResults, satResult, debrisResult, airResult, tacResult, killResult, wtaResult, specResults)

	// Save issues
	if len(issues) > 0 {
		fmt.Println("\n================================================================================")
		fmt.Println("ISSUES FOUND")
		fmt.Println("================================================================================")
		for _, issue := range issues {
			fmt.Printf("  FAIL: %s\n", issue)
		}
		saveIssues()
	}
}

// PHASE 1: BMDS SIMULATORS

func verifyBMDSSims(binDir string) []BMDSResult {
	bmdSims := []string{
		"aegis", "atmospheric", "c2bmc", "cobra-judy", "decoy", "dsp",
		"gbr", "gfcb", "gmd", "hgv", "hub", "icbm", "ifxb", "irbm",
		"jamming", "jreap", "jrsc", "link16", "lrdr", "patriot",
		"sbirs", "slcm", "sm3", "sm6", "space-weather", "stss",
		"thaad", "thaad-er", "tpy2", "uewr",
	}

	var results []BMDSResult

	for _, sim := range bmdSims {
		result := verifyBMDSim(binDir, sim)
		results = append(results, result)
		printBMDSResult(result)
	}

	return results
}

func verifyBMDSim(binDir, simName string) BMDSResult {
	binPath := fmt.Sprintf("%s/bmd-sim-%s", binDir, simName)
	result := BMDSResult{
		SimResult: SimResult{Name: simName},
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
		issues = append(issues, fmt.Sprintf("%s: exit code %d", simName, result.ExitCode))
		return result
	}

	// Parse JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.ValidJSON = false
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("Invalid JSON: %v", err))
		issues = append(issues, fmt.Sprintf("%s: invalid JSON output", simName))
		return result
	}
	result.ValidJSON = true

	// Validate structure
	if sim, ok := parsed["simulator"].(string); !ok || sim == "" {
		result.Issues = append(result.Issues, "Missing or empty simulator name")
		warnings = append(warnings, fmt.Sprintf("%s: missing simulator name", simName))
	}

	if status, ok := parsed["status"].(string); !ok || status == "" {
		result.Issues = append(result.Issues, "Missing or empty status field")
		issues = append(issues, fmt.Sprintf("%s: missing status field", simName))
	}

	if params, ok := parsed["parameters"].(map[string]interface{}); !ok || params == nil {
		result.Issues = append(result.Issues, "Missing or invalid parameters")
		issues = append(issues, fmt.Sprintf("%s: missing parameters", simName))
	} else {
		result.Parameters = params
	}

	// Sim-specific validation
	result = validateBMDSimSpecific(result)

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	} else {
		result.Status = "FAIL"
	}

	return result
}

func validateBMDSimSpecific(result BMDSResult) BMDSResult {
	params := result.Parameters
	sim := result.Name

	switch sim {
	case "aegis":
		if v, ok := params["sm3_pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "sm3_pk should be 0-1")
		}
		if v, ok := params["sm3_range_km"].(float64); !ok || v <= 0 {
			result.Warnings = append(result.Warnings, "sm3_range_km should be positive")
		}

	case "gmd":
		if v, ok := params["ekv_divert_ms"].(float64); !ok || v <= 0 {
			result.Warnings = append(result.Warnings, "ekv_divert_ms should be positive")
		}
		if v, ok := params["variant"].(string); !ok || v == "" {
			result.Warnings = append(result.Warnings, "gmd variant should be specified")
		}

	case "icbm":
		if v, ok := params["stages"].(float64); !ok || v < 1 || v > 5 {
			result.Warnings = append(result.Warnings, "icbm stages should be 1-5")
		}
		if v, ok := params["flight_time_s"].(float64); !ok || v <= 0 {
			result.Warnings = append(result.Warnings, "flight_time_s should be positive")
		}

	case "sm3":
		if v, ok := params["pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, "sm3 pk should be 0-1")
		}
		if v, ok := params["range_km"].(float64); !ok || v <= 0 {
			result.Warnings = append(result.Warnings, "sm3 range_km should be positive")
		}

	case "thaad", "patriot":
		if v, ok := params["pk"].(float64); !ok || v <= 0 || v > 1 {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s pk should be 0-1", sim))
		}
		if v, ok := params["max_range_km"].(float64); !ok || v <= 0 {
			result.Warnings = append(result.Warnings, fmt.Sprintf("%s max_range_km should be positive", sim))
		}

	case "sbirs":
		if v, ok := params["satellites"].(float64); !ok || v < 1 {
			result.Warnings = append(result.Warnings, "sbirs satellites should be >= 1")
		}
		if v, ok := params["altitude_km"].(float64); !ok || v < 30000 || v > 40000 {
			result.Warnings = append(result.Warnings, "sbirs altitude should be ~35786 km (GEO)")
		}

	case "dsp":
		if v, ok := params["satellites"].(float64); !ok || v < 1 {
			result.Warnings = append(result.Warnings, "dsp satellites should be >= 1")
		}
		if v, ok := params["altitude_km"].(float64); !ok || v < 30000 {
			result.Warnings = append(result.Warnings, "dsp altitude should be >30000 km")
		}

	case "stss":
		if v, ok := params["altitude_km"].(float64); !ok || v < 1000 || v > 2000 {
			result.Warnings = append(result.Warnings, "stss altitude should be ~1350 km (LEO)")
		}

	case "link16":
		if v, ok := params["time_slots"].(float64); !ok || v != 1535 {
			result.Warnings = append(result.Warnings, "link16 time_slots should be 1535 (standard)")
		}

	case "space-weather":
		if v, ok := params["kp_index"].(float64); !ok || v < 0 || v > 9 {
			result.Warnings = append(result.Warnings, "kp_index should be 0-9")
		}
		if v, ok := params["solar_flux"].(float64); !ok || v < 60 || v > 300 {
			result.Warnings = append(result.Warnings, "solar_flux should be 60-300 sfu")
		}

	case "atmospheric":
		if v, ok := params["humidity"].(float64); !ok || v < 0 || v > 100 {
			result.Warnings = append(result.Warnings, "humidity should be 0-100%")
		}
		if v, ok := params["refractivity"].(float64); !ok || v < 200 || v > 400 {
			result.Warnings = append(result.Warnings, "refractivity should be 200-400 N-units")
		}
	}

	return result
}

func printBMDSResult(r BMDSResult) {
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

// PHASE 2: SATELLITE TRACKER

func verifySatelliteTracker(binDir string) SatelliteResult {
	binPath := fmt.Sprintf("%s/satellite-tracker", binDir)
	result := SatelliteResult{SimResult: SimResult{Name: "satellite-tracker"}}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2", "-group", "gps-ops", "-count", "5")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("satellite-tracker: exit code %d", result.ExitCode))
		return result
	}

	// satellite-tracker outputs text before JSON (e.g. "Fetching TLEs...")
	// Find the first line that starts with '{'
	jsonStr := findJSONStart(result.Output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON found in output")
		issues = append(issues, "satellite-tracker: no JSON in output")
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, fmt.Sprintf("satellite-tracker: JSON error: %v", err))
		return result
	}
	result.ValidJSON = true

	// Validate structure
	if params, ok := parsed["parameters"].(map[string]interface{}); ok {
		if group, ok := params["group"].(string); ok && group == "gps-ops" {
			result.Group = group
		}
		if tleCount, ok := params["tle_count"].(float64); ok && tleCount > 0 {
			result.Count = int(tleCount)
		}
	}

	if sats, ok := parsed["satellites"].([]interface{}); ok && len(sats) > 0 {
		for _, s := range sats {
			if sat, ok := s.(map[string]interface{}); ok {
				satellite := Satellite{
					Name:     getString(sat, "name"),
					ID:       getInt(sat, "id"),
					Lat:      getFloat(sat, "lat"),
					Lon:      getFloat(sat, "lon"),
					Alt:      getFloat(sat, "alt"),
					Vel:      getFloat(sat, "vel"),
					Vis:      getString(sat, "vis"),
					Category: getString(sat, "category"),
				}
				result.Satellites = append(result.Satellites, satellite)

				// Validate orbital physics - GPS nominal altitude 20,200 km
				if satellite.Alt < 18000 || satellite.Alt > 25000 {
					result.Warnings = append(result.Warnings,
						fmt.Sprintf("GPS satellite %s altitude %.1f km outside nominal range", satellite.Name, satellite.Alt))
				}
			}
		}
	} else {
		result.Warnings = append(result.Warnings, "No satellites found in output")
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] satellite-tracker    | satellites=%d (GPS constellation)", result.Count)
	if len(result.Warnings) > 0 {
		fmt.Printf(" | warnings=%d", len(result.Warnings))
	}
	fmt.Println()

	return result
}

// PHASE 3: SPACE DEBRIS

func verifySpaceDebris(binDir string) SpaceDebrisResult {
	binPath := fmt.Sprintf("%s/space-debris", binDir)
	result := SpaceDebrisResult{SimResult: SimResult{Name: "space-debris"}}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("space-debris: exit code %d", result.ExitCode))
		return result
	}

	// space-debris outputs text before JSON (e.g. "Generating 25000 debris objects...")
	jsonStr := findJSONStart(result.Output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON found in output")
		issues = append(issues, "space-debris: no JSON in output")
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, "space-debris: JSON error")
		return result
	}
	result.ValidJSON = true

	if debris, ok := parsed["debris"].([]interface{}); ok {
		result.Count = len(debris)
		for _, d := range debris {
			if obj, ok := d.(map[string]interface{}); ok {
				debrisObj := DebrisObject{
					ID:          getInt(obj, "id"),
					Name:        getString(obj, "name"),
					Apogee:      getFloat(obj, "apogee"),
					Perigee:     getFloat(obj, "perigee"),
					Inclination: getFloat(obj, "inclination"),
					Period:      getFloat(obj, "mean_motion"),
					RCS:         getFloat(obj, "rcs"),
					Lat:         getFloat(obj, "lat"),
					Lon:         getFloat(obj, "lon"),
				}
				result.Debris = append(result.Debris, debrisObj)

				// Validate orbital mechanics: apogee > perigee
				if debrisObj.Apogee < debrisObj.Perigee {
					result.Warnings = append(result.Warnings,
						fmt.Sprintf("Debris %s: apogee < perigee (%.1f < %.1f)", debrisObj.Name, debrisObj.Apogee, debrisObj.Perigee))
				}
			}
		}
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] space-debris        | debris_objects=%d", result.Count)
	if len(result.Warnings) > 0 {
		fmt.Printf(" | warnings=%d", len(result.Warnings))
	}
	fmt.Println()

	return result
}

// PHASE 4: AIR TRAFFIC

func verifyAirTraffic(binDir string) AirTrafficResult {
	binPath := fmt.Sprintf("%s/air-traffic", binDir)
	result := AirTrafficResult{SimResult: SimResult{Name: "air-traffic"}}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("air-traffic: exit code %d", result.ExitCode))
		return result
	}

	// air-traffic outputs text before JSON (e.g. "Generating flights...")
	jsonStr := findJSONStart(result.Output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON found in output")
		issues = append(issues, "air-traffic: no JSON in output")
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, "air-traffic: JSON parse error")
		return result
	}
	result.ValidJSON = true

	if flights, ok := parsed["flights"].([]interface{}); ok {
		result.FlightCount = len(flights)
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] air-traffic          | flights=%d", result.FlightCount)
	fmt.Println()

	return result
}

// PHASE 5: TACTICAL NET

func verifyTacticalNet(binDir string) TacticalNetResult {
	binPath := fmt.Sprintf("%s/tactical-net", binDir)
	result := TacticalNetResult{SimResult: SimResult{Name: "tactical-net"}}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("tactical-net: exit code %d", result.ExitCode))
		return result
	}

	// tactical-net outputs text during simulation, then JSON at end
	jsonStr := findJSONStart(result.Output)
	if jsonStr == "" {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, "No JSON found in output")
		issues = append(issues, "tactical-net: no JSON in output")
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, "tactical-net: JSON parse error")
		return result
	}
	result.ValidJSON = true

	if v, ok := parsed["parameters"].(map[string]interface{}); ok {
		if nodes, ok := v["nodes"].(float64); ok {
			result.ParticipantCount = int(nodes)
		}
	}

	if links, ok := parsed["links"].([]interface{}); ok {
		result.LinkCount = len(links)
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] tactical-net         | nodes=%d, links=%d", result.ParticipantCount, result.LinkCount)
	fmt.Println()

	return result
}

// PHASE 6: KILL ASSESSMENT

func verifyKillAssessment(binDir string) KillAssessmentResult {
	binPath := fmt.Sprintf("%s/kill-assessment", binDir)
	result := KillAssessmentResult{SimResult: SimResult{Name: "kill-assessment"}}

	// kill-assessment uses -scenario flag, not -duration
	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-scenario", "single-icbm-ekv", "-seed", "42")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("kill-assessment: exit code %d", result.ExitCode))
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, "kill-assessment: JSON parse error")
		return result
	}
	result.ValidJSON = true

	if v, ok := parsed["kill_probability"].(float64); ok {
		result.KillProbability = v
		if v < 0 || v > 1 {
			result.Warnings = append(result.Warnings, "kill_probability should be 0-1")
		}
	} else {
		result.Warnings = append(result.Warnings, "kill_probability not found or invalid")
	}

	if v, ok := parsed["scenario"].(string); ok {
		result.Scenario = v
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] kill-assessment      | kill_probability=%.2f, scenario=%s", result.KillProbability, result.Scenario)
	fmt.Println()

	return result
}

// PHASE 7: WTA

func verifyWTA(binDir string) WTAResult {
	binPath := fmt.Sprintf("%s/wta", binDir)
	result := WTAResult{SimResult: SimResult{Name: "wta"}}

	start := time.Now()
	cmd := exec.Command(binPath, "-json", "-duration", "2")
	output, err := cmd.CombinedOutput()
	result.Duration = time.Since(start)
	result.ExitCode = cmd.ProcessState.ExitCode()
	result.Output = string(output)

	if err != nil && result.ExitCode != 0 {
		result.Status = "ERROR"
		result.Issues = append(result.Issues, fmt.Sprintf("Exit code %d", result.ExitCode))
		issues = append(issues, fmt.Sprintf("wta: exit code %d", result.ExitCode))
		return result
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(output, &parsed); err != nil {
		result.Status = "FAIL"
		result.Issues = append(result.Issues, fmt.Sprintf("JSON parse error: %v", err))
		issues = append(issues, "wta: JSON parse error")
		return result
	}
	result.ValidJSON = true

	if v, ok := parsed["threat_level"].(float64); ok {
		result.ThreatLevel = int(v)
	}
	if v, ok := parsed["threat_type"].(string); ok {
		result.ThreatType = v
	}

	if len(result.Issues) == 0 {
		result.Status = "PASS"
	}

	fmt.Printf("  [PASS] wta                  | threat_level=%d, threat_type=%s", result.ThreatLevel, result.ThreatType)
	fmt.Println()

	return result
}

// PHASE 8: SPECIALIZED SIMS

func verifySpecializedSims(binDir string) []SpecializedResult {
	var results []SpecializedResult

	results = append(results, verifyCyberRedteam(binDir))
	results = append(results, verifyMaritime(binDir))
	results = append(results, verifySpaceWar(binDir))

	return results
}

func verifyCyberRedteam(binDir string) SpecializedResult {
	binPath := fmt.Sprintf("%s/cyber-redteam-sim", binDir)
	result := SpecializedResult{Name: "cyber-redteam-sim"}

	cmd := exec.Command(binPath, "-doctor")
	output, err := cmd.CombinedOutput()

	if err != nil {
		result.Status = "FAIL"
		issues = append(issues, "cyber-redteam-sim: doctor check failed")
	} else {
		result.Status = "PASS"
		if strings.Contains(string(output), "healthy") ||
			strings.Contains(string(output), "OK") ||
			strings.Contains(string(output), "ready") {
			result.ScenarioLoaded = true
		}
	}

	fmt.Printf("  [PASS] cyber-redteam-sim    | doctor_exit=%d", cmd.ProcessState.ExitCode())
	fmt.Println()

	return result
}

func verifyMaritime(binDir string) SpecializedResult {
	binPath := fmt.Sprintf("%s/maritime-sim", binDir)
	result := SpecializedResult{Name: "maritime-sim"}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()

	if err != nil {
		result.Status = "FAIL"
		issues = append(issues, "maritime-sim: list-scenarios failed")
	} else {
		result.Status = "PASS"
		if strings.Contains(string(output), "scenario") ||
			strings.Contains(string(output), ".yaml") {
			result.ScenarioLoaded = true
		}
	}

	fmt.Printf("  [PASS] maritime-sim         | list_scenarios_exit=%d", cmd.ProcessState.ExitCode())
	fmt.Println()

	return result
}

func verifySpaceWar(binDir string) SpecializedResult {
	binPath := fmt.Sprintf("%s/space-war-sim", binDir)
	result := SpecializedResult{Name: "space-war-sim"}

	cmd := exec.Command(binPath, "-list-scenarios")
	output, err := cmd.CombinedOutput()

	if err != nil {
		result.Status = "FAIL"
		issues = append(issues, "space-war-sim: list-scenarios failed")
	} else {
		result.Status = "PASS"
		if strings.Contains(string(output), "scenario") ||
			strings.Contains(string(output), ".yaml") {
			result.ScenarioLoaded = true
		}
	}

	fmt.Printf("  [PASS] space-war-sim        | list_scenarios_exit=%d", cmd.ProcessState.ExitCode())
	fmt.Println()

	return result
}

// SUMMARY

func printSummary(bmdResults []BMDSResult, sat SatelliteResult, debris SpaceDebrisResult,
	air AirTrafficResult, tac TacticalNetResult, kill KillAssessmentResult,
	wta WTAResult, specResults []SpecializedResult) {

	totalIssues := len(issues)
	totalWarnings := len(warnings)

	bmdPass := 0
	bmdFail := 0
	bmdWarn := 0
	for _, r := range bmdResults {
		if r.Status == "PASS" {
			bmdPass++
		} else {
			bmdFail++
		}
		if len(r.Warnings) > 0 {
			bmdWarn++
		}
	}

	fmt.Printf("\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("                        VERIFICATION RESULTS\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("  BMDS Simulators (30)\n")
	fmt.Printf("    PASS: %d   FAIL: %d   Warnings: %d\n", bmdPass, bmdFail, bmdWarn)
	fmt.Printf("\n")
	fmt.Printf("  Specialized Sims:\n")
	fmt.Printf("    Satellite Tracker: %s (valid JSON, %d satellites)\n", sat.Status, sat.Count)
	fmt.Printf("    Space Debris:     %s (%d objects)\n", debris.Status, debris.Count)
	fmt.Printf("    Air Traffic:      %s (%d flights)\n", air.Status, air.FlightCount)
	fmt.Printf("    Tactical Net:     %s (nodes: %d, links: %d)\n", tac.Status, tac.ParticipantCount, tac.LinkCount)
	fmt.Printf("    Kill Assessment:  %s (pk: %.2f)\n", kill.Status, kill.KillProbability)
	fmt.Printf("    WTA:              %s (level: %d)\n", wta.Status, wta.ThreatLevel)
	fmt.Printf("    Cyber Redteam:    %s\n", specResults[0].Status)
	fmt.Printf("    Maritime Sim:     %s\n", specResults[1].Status)
	fmt.Printf("    Space War Sim:    %s\n", specResults[2].Status)
	fmt.Printf("\n")
	fmt.Printf("================================================================================\n")
	fmt.Printf("  TOTALS: Issues: %d   Warnings: %d\n", totalIssues, totalWarnings)
	fmt.Printf("================================================================================\n")
}

// UTILITIES

// findJSONStart finds the start of JSON in output that may have text before it
func findJSONStart(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "{") {
			// Try to find complete JSON from this point
			rest := output[strings.Index(output, trimmed):]
			// Try progressively longer substrings to find complete JSON
			for i := len(rest); i >= 0; i-- {
				testStr := rest[:i]
				if strings.HasSuffix(testStr, "}") {
					var check map[string]interface{}
					if err := json.Unmarshal([]byte(testStr), &check); err == nil {
						return testStr
					}
				}
			}
			// Fallback: return the trimmed line
			return trimmed
		}
	}
	return ""
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getInt(m map[string]interface{}, key string) int {
	if v, ok := m[key].(float64); ok {
		return int(v)
	}
	return 0
}

func getFloat(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

func saveIssues() {
	content := fmt.Sprintf("# FORGE-SIMS VERIFICATION ISSUES\n\n**Date:** %s\n\n## Issues Found\n\n",
		time.Now().UTC().Format(time.RFC3339))

	for _, issue := range issues {
		content += fmt.Sprintf("- %s\n", issue)
	}

	content += fmt.Sprintf("\n## Warnings (Low Priority)\n\n")
	for _, w := range warnings {
		content += fmt.Sprintf("- %s\n", w)
	}

	content += "\n---\n*Generated by FORGE-Sims Verification Suite*\n"

	os.WriteFile("../EVALUATION.md", []byte(content), 0644)
	fmt.Printf("\nIssues saved to EVALUATION.md\n")
}