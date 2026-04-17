#!/bin/bash
set -e

echo "╔════════════════════════════════════════════╗"
echo "║   SPACE WAR SIM — Uninstaller               ║"
echo "╚════════════════════════════════════════════╝"
echo ""

if [ "$(id -u)" -ne 0 ]; then
    echo "Please run as root (sudo)"
    exit 1
fi

# Stop service if running
if command -v systemctl &>/dev/null; then
    systemctl stop space-war-sim 2>/dev/null || true
    systemctl disable space-war-sim 2>/dev/null || true
fi

# Remove binary
rm -f /usr/local/bin/space-war-sim
echo "✓ Binary removed"

# Remove systemd unit
rm -f /lib/systemd/system/space-war-sim.service
systemctl daemon-reload 2>/dev/null || true
echo "✓ Systemd unit removed"

# Remove completions
rm -f /etc/bash_completion.d/space-war-sim
echo "✓ Completions removed"

# Remove man page
rm -f /usr/share/man/man1/space-war-sim.1.gz
echo "✓ Man page removed"

# Ask about configs
echo ""
read -p "Remove config files at /etc/space-war-sim/? [y/N] " confirm
if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
    rm -rf /etc/space-war-sim
    echo "✓ Configs removed"
else
    echo "  Configs preserved at /etc/space-war-sim/"
fi

# Ask about data
read -p "Remove data directory at /var/lib/space-war-sim/? [y/N] " confirm
if [ "$confirm" = "y" ] || [ "$confirm" = "Y" ]; then
    rm -rf /var/lib/space-war-sim
    echo "✓ Data directory removed"
else
    echo "  Data preserved at /var/lib/space-war-sim/"
fi

echo ""
echo "✅ Space War Sim uninstalled."