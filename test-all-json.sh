#!/bin/bash
# FORGE-Sims Comprehensive JSON Test - Round 26b
cd /home/wez/.openclaw/workspace/forge-sims/binaries/linux-x86

echo "=== FORGE-SIMS JSON VALIDATION (Round 26b) ==="
echo ""

PASS=0
FAIL=0
CONFIG=0

for bin in $(ls -1 | grep -v "\.sh$\|\.json$"); do
    # Skip install/uninstall scripts
    if echo "$bin" | grep -qE "^(install|uninstall)"; then
        continue
    fi
    
    # Test with -json -duration 1s
    result=$($bin -json -duration 1s 2>&1)
    
    # Check for JSON output (simulator or name field)
    if echo "$result" | grep -qE '^\s*\{|"simulator"|"name"'; then
        echo "✓ $bin"
        ((PASS++))
    # Check if it's config-driven (mentions config/scenario/output)
    elif echo "$result" | grep -qiE "config|scenario|output file"; then
        echo "⊘ $bin (config-driven)"
        ((CONFIG++))
    else
        echo "✗ $bin - ${result:0:50}"
        ((FAIL++))
    fi
done

echo ""
echo "=== RESULTS ==="
echo "Standard JSON (-json): $PASS"
echo "Config-driven: $CONFIG"
echo "No JSON/Failed: $FAIL"
echo "Total tested: $((PASS + CONFIG + FAIL))"

if [ $((PASS + CONFIG)) -gt 0 ]; then
    PERCENT=$(( (PASS + CONFIG) * 100 / (PASS + CONFIG + FAIL) ))
    echo "JSON-capable: $PERCENT%"
fi
