#!/bin/bash
# Cyber Red Team Simulator — macOS installer
set -e

BINARY="cyber-redteam-sim"
BINDIR="/usr/local/bin"
ETCDIR="/etc/$BINARY"

echo "╔════════════════════════════════════════════╗"
echo "║   CYBER RED TEAM SIM — macOS Installer     ║"
echo "╚════════════════════════════════════════════╝"
echo ""

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

# Find binary
if [ -f "$SCRIPT_DIR/usr/local/bin/$BINARY" ]; then
    BINARY_SRC="$SCRIPT_DIR/usr/local/bin/$BINARY"
elif [ -f "$SCRIPT_DIR/$BINARY" ]; then
    BINARY_SRC="$SCRIPT_DIR/$BINARY"
else
    echo "Error: Cannot find $BINARY binary in $SCRIPT_DIR"
    exit 1
fi

# Install binary
echo "Installing binary to $BINDIR..."
sudo cp "$BINARY_SRC" "$BINDIR/$BINARY"
sudo chmod 755 "$BINDIR/$BINARY"

# Install configs
echo "Installing configs to $ETCDIR..."
sudo mkdir -p "$ETCDIR/configs"
if [ -d "$SCRIPT_DIR/etc/$BINARY/configs" ]; then
    sudo cp "$SCRIPT_DIR/etc/$BINARY/configs/"*.yaml "$ETCDIR/configs/" 2>/dev/null || true
elif [ -d "$SCRIPT_DIR/configs" ]; then
    sudo cp "$SCRIPT_DIR/configs/"*.yaml "$ETCDIR/configs/" 2>/dev/null || true
fi

# Install launchd plist for server mode (optional)
echo "Installing launchd plist..."
sudo tee /Library/LaunchDaemons/com.cyberredteamsim.plist > /dev/null << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.cyberredteamsim</string>
    <key>ProgramArguments</key>
    <array>
        <string>$BINDIR/$BINARY</string>
        <string>--config</string>
        <string>$ETCDIR/configs/enterprise-ad.yaml</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/var/log/$BINARY.log</string>
    <key>StandardErrorPath</key>
    <string>/var/log/$BINARY-error.log</string>
</dict>
</plist>
EOF

# Install zsh completions
echo "Installing zsh completions..."
sudo mkdir -p /usr/local/share/zsh/site-functions
sudo tee /usr/local/share/zsh/site-functions/_cyber-redteam-sim > /dev/null << 'ZSHCOMP'
#compdef cyber-redteam-sim
_cyber-redteam-sim() {
    local -a opts
    opts=(--config --no-server --rest-port --version --help)
    _arguments '*:: :->args'
    case $words[1] in
        --config) _files ;;
        *) _describe 'option' opts ;;
    esac
}
_cyber-redteam-sim
ZSHCOMP

# Install bash completions
echo "Installing bash completions..."
sudo mkdir -p /usr/local/etc/bash_completion.d
sudo cp /etc/bash_completion.d/$BINARY /usr/local/etc/bash_completion.d/$BINARY 2>/dev/null || \
sudo tee /usr/local/etc/bash_completion.d/$BINARY > /dev/null << 'BASHCOMP'
_cyber_redteam_sim_completion() {
    local cur opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts="--config --no-server --rest-port --version --help"
    if [[ "${cur}" == --* ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    fi
}
complete -F _cyber_redteam_sim_completion cyber-redteam-sim
BASHCOMP

# Quarantine attribute (macOS)
sudo xattr -d com.apple.quarantine "$BINDIR/$BINARY" 2>/dev/null || true

echo ""
echo "✅ Installation complete!"
echo ""
echo "Binary:     $BINDIR/$BINARY"
echo "Configs:    $ETCDIR/configs/"
echo "LaunchDaemon: /Library/LaunchDaemons/com.cyberredteamsim.plist"
echo ""
echo "Quick start:"
echo "  $BINARY --version"
echo "  $BINARY --config $ETCDIR/configs/enterprise-ad.yaml --no-server"
echo "  sudo launchctl load -w /Library/LaunchDaemons/com.cyberredteamsim.plist  # Run as service"
echo ""
$BINDIR/$BINARY --version