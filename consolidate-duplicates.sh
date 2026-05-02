#!/bin/bash
# Duplicate Binary Consolidation Script
# Run this in forge-sims repo to remove duplicate binaries

set -e

cd /home/wez/.openclaw/workspace/forge-sims

echo "=== Duplicate Binary Consolidation ==="
echo ""
echo "This script removes non-prefixed duplicate binaries."
echo "We keep bmd-sim-* versions, remove plain names."
echo ""

# List of duplicates to remove (keep bmd-sim-* version)
DUPLICATES=(
    "air-traffic"
    "space-debris"
    "satellite-tracker"
    "boost-intercept"
    "debris-field"
    "electronic-war-sim"
    "emp-effects"
    "engagement-chain"
    "hgv"
    "ir-signature"
    "kill-assessment"
    "kill-chain-rt"
    "rcs"
    "tactical-net"
    "wta"
)

echo "Binaries to remove (${#DUPLICATES[@]} total):"
for dup in "${DUPLICATES[@]}"; do
    echo "  - $dup (keeping bmd-sim-$dup)"
done
echo ""

# Confirm before proceeding
read -p "Continue with removal? (y/N): " confirm
if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
    echo "Aborted."
    exit 0
fi

echo ""
echo "Removing duplicates from all platforms..."

# Remove from linux-x86
for dup in "${DUPLICATES[@]}"; do
    if [ -f "binaries/linux-x86/$dup" ]; then
        rm "binaries/linux-x86/$dup"
        echo "✓ Removed binaries/linux-x86/$dup"
    fi
done

# Remove from linux-arm64
for dup in "${DUPLICATES[@]}"; do
    if [ -f "binaries/linux-arm64/$dup" ]; then
        rm "binaries/linux-arm64/$dup"
        echo "✓ Removed binaries/linux-arm64/$dup"
    fi
done

# Remove from darwin-amd64
for dup in "${DUPLICATES[@]}"; do
    if [ -f "binaries/darwin-amd64/$dup" ]; then
        rm "binaries/darwin-amd64/$dup"
        echo "✓ Removed binaries/darwin-amd64/$dup"
    fi
done

# Remove from darwin-arm64
for dup in "${DUPLICATES[@]}"; do
    if [ -f "binaries/darwin-arm64/$dup" ]; then
        rm "binaries/darwin-arm64/$dup"
        echo "✓ Removed binaries/darwin-arm64/$dup"
    fi
done

# Remove from windows-amd64
for dup in "${DUPLICATES[@]}"; do
    if [ -f "binaries/windows-amd64/${dup}.exe" ]; then
        rm "binaries/windows-amd64/${dup}.exe"
        echo "✓ Removed binaries/windows-amd64/${dup}.exe"
    fi
done

echo ""
echo "✅ Duplicates removed!"
echo ""
echo "Next steps:"
echo "1. Update Makefile to stop building non-prefixed versions"
echo "2. Update API_REFERENCE.md to remove duplicate entries"
echo "3. Update README.md binary count"
echo "4. Commit changes: git add -A && git commit -m 'Remove duplicate binaries'"
