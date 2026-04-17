#!/bin/bash
# Daily satellite TLE refresh: fetch from Celestrak on darth, push to miner
# Runs via cron on darth at 3 AM UTC

set -e

SAT_BIN="/home/wez/forge-sims/binaries/linux-x86/satellite-tracker"
OUT_DIR="/tmp/sat-data-upload"
MINER="wez@207.244.226.151"
MINER_DATA="/home/wez/sat-data-service/data"

echo "[$(date -u +%Y-%m-%dT%H:%M:%SZ)] Starting daily satellite refresh"

rm -rf "$OUT_DIR"
mkdir -p "$OUT_DIR"

TOTAL=0

fetch_group() {
    local group="$1"
    echo "  Fetching $group..."
    timeout 30 "$SAT_BIN" -group="$group" -json 2>/dev/null > "$OUT_DIR/$group.json"
    local count
    count=$(python3 -c "import json; d=json.load(open('$OUT_DIR/$group.json')); print(len(d.get('satellites',[])))" 2>/dev/null || echo "0")
    echo "    $group: $count satellites"
    TOTAL=$((TOTAL + count))
    sleep 5
}

fetch_group gps-ops
fetch_group stations
fetch_group weather
fetch_group geo
fetch_group military
fetch_group intelsat
fetch_group resource
fetch_group sarsat
fetch_group argos
fetch_group planet
fetch_group dmc
fetch_group orbcomm
fetch_group x-comm
fetch_group starlink

echo "  Total: $TOTAL satellites"

if [ "$TOTAL" -lt 10 ]; then
    echo "  ERROR: Too few satellites fetched, aborting upload"
    exit 1
fi

# Wrap into sat-data-service format
python3 - "$OUT_DIR" << 'PYEOF'
import json, time, os, sys

out_dir = sys.argv[1]
manifest = {"last_refresh": time.time(), "last_refresh_time": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime()), "groups": {}, "total_count": 0}

for fname in os.listdir(out_dir):
    if not fname.endswith(".json") or fname == "manifest.json":
        continue
    group = fname.replace(".json", "")
    path = os.path.join(out_dir, fname)
    try:
        raw = json.load(open(path))
        sats = raw.get("satellites", [])
        if not sats:
            continue
        data = {"group": group, "fetched_at": time.time(), "count": len(sats), "satellites": sats}
        with open(path, "w") as f:
            json.dump(data, f)
        manifest["groups"][group] = len(sats)
        manifest["total_count"] += len(sats)
    except:
        pass

with open(os.path.join(out_dir, "manifest.json"), "w") as f:
    json.dump(manifest, f, indent=2)
print(f"Prepared {manifest['total_count']} sats in {len(manifest['groups'])} groups")
PYEOF

# Upload to miner
echo "  Uploading to miner..."
scp "$OUT_DIR"/*.json "$MINER:$MINER_DATA/"

# Restart sat-data-service on miner
echo "  Restarting sat-data-service..."
ssh "$MINER" "sudo systemctl restart sat-data"

# Clean up
rm -rf "$OUT_DIR"

echo "[$(date -u +%Y-%m-%dT%H:%M:%SZ)] Daily satellite refresh complete"