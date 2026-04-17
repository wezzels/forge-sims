#!/bin/bash
# Cyber Red Team Simulator — Linux installer
# Usage: sudo ./install.sh
set -e

BINARY="cyber-redteam-sim"
BINDIR="/usr/local/bin"
ETCDIR="/etc/$BINARY"
VARDIR="/var/lib/$BINARY"
UNITFILE="/lib/systemd/system/${BINARY}.service"

echo "╔════════════════════════════════════════════╗"
echo "║   CYBER RED TEAM SIM — Installer          ║"
echo "╚════════════════════════════════════════════╝"
echo ""

# Check root
if [ "$(id -u)" -ne 0 ]; then
    echo "Error: Run as root (sudo ./install.sh)"
    exit 1
fi

# Detect script directory
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

# Find binary
if [ -f "$SCRIPT_DIR/usr/local/bin/$BINARY" ]; then
    BINARY_SRC="$SCRIPT_DIR/usr/local/bin/$BINARY"
elif [ -f "$SCRIPT_DIR/$BINARY" ]; then
    BINARY_SRC="$SCRIPT_DIR/$BINARY"
elif [ -f "$SCRIPT_DIR/${BINARY}-0.2.0-linux-amd64" ]; then
    BINARY_SRC="$SCRIPT_DIR/${BINARY}-0.2.0-linux-amd64"
else
    echo "Error: Cannot find $BINARY binary"
    echo "  Expected: $SCRIPT_DIR/$BINARY or $SCRIPT_DIR/usr/local/bin/$BINARY"
    exit 1
fi

# Install binary
echo "Installing binary..."
cp "$BINARY_SRC" "$BINDIR/$BINARY"
chmod 755 "$BINDIR/$BINARY"

# Install configs
echo "Installing configs..."
mkdir -p "$ETCDIR/configs"
if [ -d "$SCRIPT_DIR/etc/$BINARY/configs" ]; then
    cp "$SCRIPT_DIR/etc/$BINARY/configs/"*.yaml "$ETCDIR/configs/" 2>/dev/null || true
elif [ -d "$SCRIPT_DIR/configs" ]; then
    cp "$SCRIPT_DIR/configs/"*.yaml "$ETCDIR/configs/" 2>/dev/null || true
fi

# Install data directory
mkdir -p "$VARDIR"
chown root:root "$VARDIR"

# Install systemd unit
echo "Installing systemd service..."
cat > "$UNITFILE" << EOF
[Unit]
Description=Cyber Red Team Simulator
After=network.target

[Service]
Type=simple
ExecStart=$BINDIR/$BINARY --config $ETCDIR/configs/enterprise-ad.yaml
Restart=on-failure
RestartSec=5
WorkingDirectory=$VARDIR

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload

# Install shell completions
echo "Installing shell completions..."
mkdir -p /etc/bash_completion.d
cat > /etc/bash_completion.d/$BINARY << 'BASHCOMP'
_cyber_redteam_sim_completion() {
    local cur opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts="--config --no-server --rest-port --version --help"
    
    if [[ "${cur}" == --* ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    fi
    
    if [[ ${COMP_CWORD} -eq 1 ]] || [[ "${COMP_WORDS[COMP_CWORD-1]}" == "--config" ]]; then
        COMPREPLY=( $(compgen -f -- "${cur}") )
    fi
}
complete -F _cyber_redteam_sim_completion cyber-redteam-sim
BASHCOMP

# Install man page (if we have it)
if [ -f "$SCRIPT_DIR/$BINARY.1.gz" ]; then
    mkdir -p /usr/share/man/man1
    cp "$SCRIPT_DIR/$BINARY.1.gz" /usr/share/man/man1/
fi

echo ""
echo "✅ Installation complete!"
echo ""
echo "Binary:     $BINDIR/$BINARY"
echo "Configs:    $ETCDIR/configs/"
echo "Data:       $VARDIR/"
echo "Service:    $UNITFILE"
echo ""
echo "Quick start:"
echo "  $BINARY --version"
echo "  $BINARY --config $ETCDIR/configs/enterprise-ad.yaml --no-server"
echo "  sudo systemctl enable --now $BINARY    # Run as service"
echo ""
echo "  $BINARY --version"
$BINDIR/$BINARY --version