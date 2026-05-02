#!/bin/bash
# FORGE-Sims JSON Validation - Round 26
cd /home/wez/.openclaw/workspace/forge-sims/binaries/linux-x86

echo "=== FORGE-SIMS JSON VALIDATION (Round 26) ==="
echo ""

PASS=0
FAIL=0

for bin in $(ls -1); do
    # Skip config-driven sims (they don't use -json flag)
    if echo "$bin" | grep -qE "^(cyber|launch|maritime|space-war)"; then
        echo "⊘ $bin (config-driven, skipped)"
        continue
    fi
    
    result=$($bin -json -duration 1s 2>&1)
    
    if echo "$result" | grep -q '"simulator"'; then
        echo "✓ $bin - JSON valid"
        ((PASS++))
    else
        echo "✗ $bin - ${result:0:60}"
        ((FAIL++))
    fi
done

echo ""
echo "=== RESULTS ==="
echo "JSON-valid: $PASS"
echo "Invalid: $FAIL"
echo "Total tested: $((PASS + FAIL))"
