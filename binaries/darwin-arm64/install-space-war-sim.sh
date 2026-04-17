#!/bin/bash
set -e

PREFIX="${PREFIX:-/usr/local}"
ETC_DIR="/etc/space-war-sim"

echo "╔════════════════════════════════════════════╗"
echo "║   SPACE WAR SIM — macOS Installer           ║"
echo "╚════════════════════════════════════════════╝"
echo ""

# Install binary
mkdir -p "$PREFIX/bin"
cp space-war-sim "$PREFIX/bin/space-war-sim" 2>/dev/null || true
chmod +x "$PREFIX/bin/space-war-sim"

# Remove quarantine
xattr -d com.apple.quarantine "$PREFIX/bin/space-war-sim" 2>/dev/null || true
echo "✓ Binary installed: $PREFIX/bin/space-war-sim"

# Install configs
mkdir -p "$ETC_DIR/configs"
cp -n configs/*.yaml "$ETC_DIR/configs/" 2>/dev/null || true
echo "✓ Configs installed: $ETC_DIR/configs/"

# Zsh completions
mkdir -p /usr/local/share/zsh/site-functions 2>/dev/null
echo "#compdef space-war-sim
_maritime_sim() {
    local -a commands
    commands=(
        '--config:scenario config file'
        '--aar:AAR JSON export path'
        '--web:web UI address'
        '--version:print version'
        '--doctor:health check'
        '--init:initialize configs'
        '--validate:validate config'
        '--list-scenarios:list scenarios'
    )
    _arguments -C \$commands
}
_maritime_sim" > /usr/local/share/zsh/site-functions/_space-war-sim 2>/dev/null || true
echo "✓ Zsh completions installed"

# Bash completions
mkdir -p /usr/local/etc/bash_completion.d 2>/dev/null
echo "_maritime_sim() {
    local cur
    COMPREPLY=()
    cur=\${COMP_WORDS[COMP_CWORD]}
    if [[ \$cur == -* ]]; then
        COMPREPLY=(\$(compgen -W '--config --aar --web --version --doctor --init --validate --list-scenarios' -- \$cur))
    fi
}
complete -F _maritime_sim space-war-sim" > /usr/local/etc/bash_completion.d/space-war-sim 2>/dev/null || true
echo "✓ Bash completions installed"

# launchd plist (if root)
if [ "$(id -u)" -eq 0 ]; then
    cat > /Library/LaunchDaemons/com.maritimesim.plist << 'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.maritimesim</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/space-war-sim</string>
        <string>--config</string>
        <string>/etc/space-war-sim/configs/scenario.yaml</string>
        <string>--web</string>
        <string>:8080</string>
    </array>
    <key>RunAtLoad</key>
    <false/>
    <key>KeepAlive</key>
    <false/>
    <key>StandardOutPath</key>
    <string>/var/log/space-war-sim.log</string>
    <key>StandardErrorPath</key>
    <string>/var/log/space-war-sim.err</string>
</dict>
</plist>
PLIST
    echo "✓ launchd plist installed"
fi

echo ""
echo "✅ Installation complete!"
echo ""
echo "Quick start:"
echo "  space-war-sim --version"
echo "  space-war-sim --doctor"
echo "  space-war-sim --config $ETC_DIR/configs/scenario.yaml"
echo ""
echo "To run as a service:"
echo "  sudo launchctl load -w /Library/LaunchDaemons/com.maritimesim.plist"