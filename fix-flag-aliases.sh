#!/bin/bash
# Fix script for flag standardization (-d → -duration alias)
# Run this on darth in each affected repo

set -e

# List of repos to fix
REPOS=(
    "bmd-sim-c2bmc"
    "bmd-sim-link16"
    "bmd-sim-jreap"
    "bmd-sim-jrsc"
    "bmd-sim-gfcb"
    "bmd-sim-decoy"
    "bmd-sim-ifxb"
    "bmd-sim-ir-signature"
    "bmd-sim-kill-assessment"
    "bmd-sim-rcs"
    "bmd-sim-satellite-weapon"
    "bmd-sim-wta"
    "bmd-sim-boost-intercept"
    "bmd-sim-debris-field"
    "electronic-war-sim"
    "bmd-sim-emp-effects"
    "bmd-sim-engagement-chain"
    "bmd-sim-hgv"
    "bmd-sim-kill-chain-rt"
    "bmd-sim-tactical-net"
)

echo "=== Flag Standardization Fix Script ==="
echo "This script adds -duration as an alias for -d in 20 repos"
echo ""
echo "For each repo, you need to:"
echo "1. Clone: git clone git@idm.wezzel.com:crab-meat-repos/<repo>.git"
echo "2. Edit main.go"
echo "3. Rebuild"
echo ""
echo "The change in main.go:"
echo ""
echo "BEFORE:"
echo '  duration = flag.Duration("d", 0, "Duration")'
echo ""
echo "AFTER:"
echo '  duration = flag.Duration("duration", 0, "Duration")'
echo '  durationAlias = flag.Duration("d", 0, "Duration (alias)")'
echo ""
echo "This maintains backward compatibility while standardizing on -duration"
echo ""
echo "Repos to fix (${#REPOS[@]} total):"
for repo in "${REPOS[@]}"; do
    echo "  - $repo"
done
echo ""
echo "Test after fix:"
echo '  ./binary -json -duration 1s | jq .'
echo '  ./binary -json -d 1s | jq .'
echo "  # Both should produce identical output"
