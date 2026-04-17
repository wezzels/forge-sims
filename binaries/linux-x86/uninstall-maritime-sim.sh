#!/bin/bash
set -e

echo "╔════════════════════════════════════════════╗"
echo "║   MARITIME SIM — Uninstaller               ║"
echo "╚════════════════════════════════════════════╝"
echo ""

if [ "$(id -u)" -ne 0 ]; then
    echo "Please run as root (sudo)"
    exit 1
fi

# Stop service if running
if command -v systemctl &>/dev/null; then
    systemctl stop maritime-sim 2>/dev/null || true
    systemctl disable maritime-sim 2>/dev/null || true
fi

# Remove binary
rm -f /usr/local/bin/maritime-sim
echo "✓ Binary removed"

# Remove systemd unit
rm -f /lib/systemd/system/maritime-sim.service
systemctl daemon-reload 2>/dev/null || true
echo "✓ Systemd unit removed"

# Remove completions
rm -f /etc/bash_completion.d/maritime-sim
echo "✓ Completions removed"

# Remove man page
rm -f /usr/share/man/man1/maritime-sim.1.gz
echo "✓ Man page removed"

# Ask about configs
echo ""
read -p "Remove config files at /etc/maritime-sim/? [y/N] " confirm
if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
    rm -rf /etc/maritime-sim
    echo "✓ Configs removed"
else
    echo "  Configs preserved at /etc/maritime-sim/"
fi

# Ask about data
read -p "Remove data directory at /var/lib/maritime-sim/? [y/N] " confirm
if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
    rm -rf /var/lib/maritime-sim
    echo "✓ Data directory removed"
else
    echo "  Data preserved at /var/lib/maritime-sim/"
fi

echo ""
echo "✅ Maritime Sim uninstalled."