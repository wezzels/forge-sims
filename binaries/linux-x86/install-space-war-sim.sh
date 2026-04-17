#!/bin/bash
set -e

PREFIX="${PREFIX:-/usr/local}"
ETC_DIR="/etc/space-war-sim"
BIN="$PREFIX/bin/space-war-sim"

echo "╔════════════════════════════════════════════╗"
echo "║   SPACE WAR SIM — Installer                 ║"
echo "╚════════════════════════════════════════════╝"
echo ""

# Check root
if [ "$(id -u)" -ne 0 ]; then
    PREFIX="$HOME/.local"
    BIN="$PREFIX/bin/space-war-sim"
    echo "Not running as root. Installing to $PREFIX"
fi

# Install binary
mkdir -p "$PREFIX/bin"
cp usr/local/bin/space-war-sim "$BIN" 2>/dev/null || cp space-war-sim "$BIN" 2>/dev/null
chmod +x "$BIN"
echo "✓ Binary installed: $BIN"

# Install configs
mkdir -p "$ETC_DIR/configs"
cp -n etc/space-war-sim/configs/*.yaml "$ETC_DIR/configs/" 2>/dev/null || cp -n configs/*.yaml "$ETC_DIR/configs/" 2>/dev/null || true
echo "✓ Configs installed: $ETC_DIR/configs/"

# Data directory
mkdir -p /var/lib/space-war-sim 2>/dev/null || true

# Systemd unit (if root and systemd available)
if [ "$(id -u)" -eq 0 ] && command -v systemctl &>/dev/null; then
    if [ -f lib/systemd/system/space-war-sim.service ]; then
        cp lib/systemd/system/space-war-sim.service /lib/systemd/system/
        systemctl daemon-reload
        echo "✓ Systemd unit installed"
    fi
fi

# Bash completions
if [ "$(id -u)" -eq 0 ]; then
    mkdir -p /etc/bash_completion.d
    echo "# Space War Sim completions
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
complete -F _maritime_sim space-war-sim" > /etc/bash_completion.d/space-war-sim
    echo "✓ Bash completions installed"
fi

# Man page
if [ "$(id -u)" -eq 0 ] && [ -f usr/share/man/man1/space-war-sim.1.gz ]; then
    cp usr/share/man/man1/space-war-sim.1.gz /usr/share/man/man1/
    mandb -q 2>/dev/null || true
    echo "✓ Man page installed"
fi

echo ""
echo "✅ Installation complete!"
echo ""
echo "Quick start:"
echo "  space-war-sim --version"
echo "  space-war-sim --doctor"
echo "  space-war-sim --config $ETC_DIR/configs/scenario.yaml"