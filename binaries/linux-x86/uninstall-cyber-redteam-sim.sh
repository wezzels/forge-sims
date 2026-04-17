#!/bin/bash
# Cyber Red Team Simulator — Linux uninstaller
set -e

BINARY="cyber-redteam-sim"

echo "Uninstalling $BINARY..."

# Check root
if [ "$(id -u)" -ne 0 ]; then
    echo "Error: Run as root (sudo ./uninstall.sh)"
    exit 1
fi

# Stop service if running
if systemctl is-active --quiet "$BINARY" 2>/dev/null; then
    echo "Stopping service..."
    systemctl stop "$BINARY"
fi
systemctl disable "$BINARY" 2>/dev/null || true

# Remove files
rm -f "/usr/local/bin/$BINARY"
rm -f "/lib/systemd/system/${BINARY}.service"
rm -f "/etc/bash_completion.d/$BINARY"
rm -f "/usr/share/man/man1/${BINARY}.1.gz"
systemctl daemon-reload

echo ""
echo "Config and data directories preserved:"
echo "  /etc/$BINARY/configs/"
echo "  /var/lib/$BINARY/"
echo ""
echo "To remove all data:"
echo "  sudo rm -rf /etc/$BINARY /var/lib/$BINARY"
echo ""
echo "✅ Uninstalled."