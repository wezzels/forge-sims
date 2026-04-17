#!/bin/bash
set -e

PREFIX="${PREFIX:-/usr/local}"
ETC_DIR="/etc/maritime-sim"
BIN="$PREFIX/bin/maritime-sim"

echo "╔════════════════════════════════════════════╗"
echo "║   MARITIME SIM — Installer                 ║"
echo "╚════════════════════════════════════════════╝"
echo ""

# Check root
if [ "$(id -u)" -ne 0 ]; then
    PREFIX="$HOME/.local"
    BIN="$PREFIX/bin/maritime-sim"
    echo "Not running as root. Installing to $PREFIX"
fi

# Install binary
mkdir -p "$PREFIX/bin"
cp usr/local/bin/maritime-sim "$BIN" 2>/dev/null || cp maritime-sim "$BIN" 2>/dev/null
chmod +x "$BIN"
echo "✓ Binary installed: $BIN"

# Install configs
mkdir -p "$ETC_DIR/configs"
cp -n etc/maritime-sim/configs/*.yaml "$ETC_DIR/configs/" 2>/dev/null || cp -n configs/*.yaml "$ETC_DIR/configs/" 2>/dev/null || true
echo "✓ Configs installed: $ETC_DIR/configs/"

# Data directory
mkdir -p /var/lib/maritime-sim 2>/dev/null || true

# Systemd unit (if root and systemd available)
if [ "$(id -u)" -eq 0 ] && command -v systemctl &>/dev/null; then
    if [ -f lib/systemd/system/maritime-sim.service ]; then
        cp lib/systemd/system/maritime-sim.service /lib/systemd/system/
        systemctl daemon-reload
        echo "✓ Systemd unit installed"
    fi
fi

# Bash completions
if [ "$(id -u)" -eq 0 ]; then
    mkdir -p /etc/bash_completion.d
    echo "# Maritime Sim completions
_maritime_sim() {
    local cur prev
    COMPREPLY=()
    cur=\${COMP_WORDS[COMP_CWORD]}
    prev=\${COMP_WORDS[COMP_CWORD-1]}
    if [[ \$cur == -* ]]; then
        COMPREPLY=(\$(compgen -W '--config --aar --web --version --doctor --init --validate --list-scenarios' -- \$cur))
    fi
    if [[ \$prev == '--config' || \$prev == '--validate' ]]; then
        COMPREPLY=(\$(compgen -f -- \$cur))
        compopt -o filenames
    fi
}
complete -F _maritime_sim maritime-sim" > /etc/bash_completion.d/maritime-sim
    echo "✓ Bash completions installed"
fi

# Man page
if [ "$(id -u)" -eq 0 ] && [ -f usr/share/man/man1/maritime-sim.1.gz ]; then
    cp usr/share/man/man1/maritime-sim.1.gz /usr/share/man/man1/
    mandb -q 2>/dev/null || true
    echo "✓ Man page installed"
fi

echo ""
echo "✅ Installation complete!"
echo ""
echo "Quick start:"
echo "  maritime-sim --version"
echo "  maritime-sim --doctor"
echo "  maritime-sim --config $ETC_DIR/configs/scenario.yaml"