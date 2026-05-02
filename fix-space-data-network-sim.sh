#!/bin/bash
# Fix script for space-data-network-sim
# Run this on darth (build server) in the space-data-network-sim repo directory

set -e

echo "=== Fixing space-data-network-sim ==="
echo "Adding -json, -i, -v flags..."

# Check if we're in the right directory
if [ ! -f "main.go" ]; then
    echo "ERROR: main.go not found. Run this script in the space-data-network-sim repo root."
    exit 1
fi

# Backup original
cp main.go main.go.backup

# Create the fix patch
cat > main.go.patch << 'PATCH'
--- main.go.orig
+++ main.go
@@ -1,5 +1,6 @@
 package main
 
+import "fmt"
 import (
 	"encoding/json"
 	"flag"
@@ -10,6 +11,9 @@
 
 var (
 	scenario = flag.String("scenario", "peacetime", "Scenario: peacetime, golden-dome, nuclear, custom")
+	jsonOut  = flag.Bool("json", false, "Output JSON to stdout")
+	interactive = flag.Bool("i", false, "Interactive mode (live dashboard)")
+	verbose  = flag.Bool("v", false, "Verbose output")
 	duration = flag.Float64("duration", 0, "Override duration (minutes, 0=scenario default)")
 	step     = flag.Float64("step", 0, "Override step size (minutes, 0=scenario default)")
 	output   = flag.String("output", "", "Output JSON file")
@@ -20,6 +24,18 @@
 func main() {
 	flag.Parse()
 
+	// Handle -json flag (stdout output)
+	if *jsonOut {
+		result := runScenario(*scenario, *duration, *step)
+		jsonBytes, _ := json.MarshalIndent(result, "", "  ")
+		fmt.Println(string(jsonBytes))
+		return
+	}
+
+	// Handle -i flag (interactive mode)
+	if *interactive {
+		runInteractive(*scenario, *duration, *step)
+		return
+	}
+
+	// Handle -v flag (verbose)
+	if *verbose {
+		fmt.Printf("FORGE-SDN v1.0.0 — Scenario: %s\n", *scenario)
+		fmt.Printf("Duration: %v min, Step: %v min\n", *duration, *step)
+	}
+
 	// Run scenario
 	result := runScenario(*scenario, *duration, *step)
 
PATCH

echo "⚠️  Manual edit required - the patch above shows the changes needed"
echo ""
echo "Changes to make in main.go:"
echo "1. Add import \"fmt\" at the top"
echo "2. Add three new flags after line with -scenario flag:"
echo '   jsonOut     = flag.Bool("json", false, "Output JSON to stdout")'
echo '   interactive = flag.Bool("i", false, "Interactive mode (live dashboard)")'
echo '   verbose     = flag.Bool("v", false, "Verbose output")'
echo ""
echo "3. Add flag handling in main() before existing scenario run:"
echo '   if *jsonOut {'
echo '       result := runScenario(*scenario, *duration, *step)'
echo '       jsonBytes, _ := json.MarshalIndent(result, "", "  ")'
echo '       fmt.Println(string(jsonBytes))'
echo '       return'
echo '   }'
echo ""
echo "4. Add interactive mode handler (similar pattern)"
echo "5. Add verbose logging"
echo ""
echo "After editing, rebuild:"
echo "  make build"
echo "  # or: go build -o space-data-network-sim"
echo ""
echo "Test the fix:"
echo "  ./space-data-network-sim -json -duration 0.1 | jq ."
echo "  ./space-data-network-sim -i -duration 0.5"
echo "  ./space-data-network-sim -v -scenario nuclear"
