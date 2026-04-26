#!/bin/bash
# Uninstall intel-collection-sim binary

set -e

BINARY="intel-collection-sim"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

echo "=== Uninstalling $BINARY ==="

if [ -f "$INSTALL_DIR/$BINARY" ]; then
    sudo rm "$INSTALL_DIR/$BINARY"
    echo "Removed: $INSTALL_DIR/$BINARY"
else
    echo "Not found: $INSTALL_DIR/$BINARY (already removed?)"
fi

echo "✓ $BINARY uninstalled"