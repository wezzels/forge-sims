#!/bin/bash
# Install intel-collection-sim binary
# Part of FORGE-Sims: High-fidelity intelligence collection simulation engine

set -e

BINARY="intel-collection-sim"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
SOURCE_DIR="$(cd "$(dirname "$0")" && pwd)"

echo "=== Installing $BINARY ==="

if [ ! -f "$SOURCE_DIR/$BINARY" ]; then
    echo "ERROR: Binary not found at $SOURCE_DIR/$BINARY"
    exit 1
fi

# Check architecture
ARCH=$(uname -m)
if [ "$ARCH" != "x86_64" ]; then
    echo "WARNING: Binary built for x86_64, running on $ARCH"
fi

# Install
sudo cp "$SOURCE_DIR/$BINARY" "$INSTALL_DIR/$BINARY"
sudo chmod +x "$INSTALL_DIR/$BINARY"

echo "Installed: $INSTALL_DIR/$BINARY"
echo ""

# Verify
echo "Verification:"
$INSTALL_DIR/$BINARY version
echo ""

# Quick smoke test
echo "Running validation..."
$INSTALL_DIR/$BINARY validate 2>&1 | tail -3

echo ""
echo "✓ $BINARY installed successfully"
echo ""
echo "Commands:"
echo "  $BINARY run        Run a simulation scenario"
echo "  $BINARY platforms  List collector platforms"
echo "  $BINARY threats    List adversary targets"
echo "  $BINARY collect    Single discipline collection"
echo "  $BINARY scenario   Run all scenarios with AAR"
echo "  $BINARY validate   Run physics validation checks"
echo "  $BINARY version    Show version info"