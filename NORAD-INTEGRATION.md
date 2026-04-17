# FORGE-Sims — NORAD Command Center Integration Guide

## Overview

The NORAD Command Center at **norad.stsgym.com** is a live Flask+gevent+SocketIO web application that integrates FORGE-Sims binaries into a real-time BMDS simulation dashboard. It runs on miner (207.244.226.151) as a Docker container on port 5008, with a read-only legacy mirror on port 5011.

**Architecture:**
```
norad.stsgym.com (nginx → :5008)
  ├── Flask + gevent + SocketIO
  ├── Scenario Engine (7 BMDS scenarios)
  ├── 3D Globe (Three.js r128)
  ├── Leaflet Map (CartoDB Dark Matter)
  ├── Satellite Data → sat-data-service (:5020)
  ├── Space Debris → bmd-sim-space-debris binary
  └── Air Traffic → bmd-sim-air-traffic binary

norad-legacy.stsgym.com (nginx → :5011)
  └── Read-only mirror + Leaflet map
```

---

## Integration Examples

### 1. Starting a BMDS Scenario

The NORAD app uses FORGE-Sims binaries to power its scenarios. Each scenario (ICBM, HGV, Multi-Axis, etc.) invokes multiple simulators through the scenario engine.

```python
# Example: Starting a scenario via the NORAD API
import requests

API_KEY = "norad-forge-api-2026"
BASE = "https://norad.stsgym.com"

# List available scenarios
r = requests.get(f"{BASE}/api/scenarios", headers={"X-API-Key": API_KEY})
scenarios = r.json()
# Returns: {decoy-stress, hypersonic-glide, multi-axis, nuclear-exchange, ...}

# Start the "hypersonic-glide" scenario
r = requests.post(f"{BASE}/api/scenario/start/hypersonic-glide",
                  headers={"X-API-Key": API_KEY})
state = r.json()
# Returns: {scenario: "hypersonic-glide", status: "running", defcon: 2, ...}

# Get current state
r = requests.get(f"{BASE}/api/state", headers={"X-API-Key": API_KEY})
state = r.json()
# Returns: threats, interceptors, sensors, radar_coverage, defcon_level

# Fire interceptor
r = requests.post(f"{BASE}/api/intercept", headers={"X-API-Key": API_KEY},
                  json={"threat_id": "THREAT-001"})

# Adjust DEFCON level
r = requests.post(f"{BASE}/api/defcon/1", headers={"X-API-Key": API_KEY})

# Stop scenario
r = requests.post(f"{BASE}/api/scenario/stop", headers={"X-API-Key": API_KEY})
```

### 2. Real-Time Satellite Tracking

The satellite data comes from **sat-data-service** (port 5020), which fetches TLEs from Celestrak once/day and uses Keplerian propagation to compute positions in real-time.

```python
# Example: Fetching satellite positions from NORAD
import requests

# Public endpoint (no auth required)
r = requests.get("https://norad.stsgym.com/api/public/satellites?count=10")
data = r.json()
# Returns:
# {
#   "group": "gps-ops",
#   "count": 10,
#   "fetched_at": 1713381600,
#   "interpolated": true,
#   "satellites": [
#     {"name": "GPS BIIR-2 (PRN 13)", "norad_id": 24876,
#      "category": "GPS", "alt": 20040.3, "lat": 0.015,
#      "lon": 101.0, "period": 718.0, "inclination": 55.5,
#      "vis": "visible"}
#   ]
# }

# Sat-data-service direct access (internal)
r = requests.get("http://127.0.0.1:5020/api/positions/gps-ops")
# Available groups: gps-ops, stations, weather, geo, military,
#                   intelsat, resource, sarsat, argos, planet,
#                   dmc, orbcomm, x-comm, starlink

# Force refresh (rate limited to 1/hour)
r = requests.post("http://127.0.0.1:5020/api/refresh")
```

```bash
# Using the satellite-tracker binary directly
./binaries/linux-x86/satellite-tracker -group=gps-ops -json
./binaries/linux-x86/satellite-tracker -group=stations -v
./binaries/linux-x86/satellite-tracker -group=geo -count=10 -json
./binaries/linux-x86/satellite-tracker -i -duration=30
```

### 3. Space Debris Monitoring

```python
# Example: Space debris from NORAD
r = requests.get("https://norad.stsgym.com/api/public/debris?count=100")
data = r.json()
# Returns:
# {
#   "total_count": 5000,
#   "by_category": {"LEO": 3030, "GEO": 510, "MEO": 507, "HIGH": 953},
#   "debris": [
#     {"id": 1, "name": "DEBRIS-0001", "object_type": "DEBRIS",
#      "altitude": 850.3, "lat": 12.45, "lon": -67.8,
#      "inclination": 45.0, "category": "LEO"}
#   ]
# }
```

```bash
# Using the space-debris binary directly
./binaries/linux-x86/space-debris -count=5000 -json
./binaries/linux-x86/space-debris -count=10000 -collisions -threshold=500 -json
./binaries/linux-x86/space-debris -i -duration=60
```

### 4. Air Traffic Simulation

```python
# Example: Air traffic from NORAD
r = requests.get("https://norad.stsgym.com/api/public/air-traffic?count=50")
data = r.json()
# Returns:
# {
#   "total_count": 50,
#   "by_type": {"COMMERCIAL": 20, "CARGO": 10, "MILITARY": 15, "PRIVATE": 5},
#   "flights": [
#     {"id": "FL001", "callsign": "AAL123", "type": "COMMERCIAL",
#      "aircraft": "B738", "lat": 34.05, "lon": -118.25,
#      "alt": 35000, "speed": 450, "heading": 90,
#      "origin": "KLAX", "dest": "KJFK", "phase": "CRUISE"}
#   ]
# }
```

```bash
# Using the air-traffic binary directly
./binaries/linux-x86/air-traffic -count=1000 -json
./binaries/linux-x86/air-traffic -count=500 -v
./binaries/linux-x86/air-traffic -i -duration=60 -seed=42
```

### 5. 3D Globe Rendering

The NORAD globe uses Three.js (r128) with these FORGE-Sims layers:

```javascript
// Globe layer toggles (from globe.html)
const LAYERS = {
    SBIRS:     { icon: "🛰️",  label: "SBIRS Constellation",  count: 6  },
    STSS:      { icon: "🔭",  label: "STSS Tracking",         count: 24 },
    RADAR:     { icon: "📡",  label: "Radar Coverage",       count: 10 },
    INTERCEPT: { icon: "🚀",  label: "Interceptors",         count: 0  },
    THREATS:   { icon: "🔴",  label: "Threat Tracks",        count: 0  },
    DECOYS:    { icon: "⚪",  label: "Decoy Objects",        count: 0  },
    C2:        { icon: "🏛️",  label: "C2BMC Nodes",          count: 4  },
    FLEET:     { icon: "🚢",  label: "US Navy Fleet",        count: 24 },
    BASES:     { icon: "📍",  label: "Military Bases",       count: 27 },
    GRID:      { icon: "🌐",  label: "Lat/Lon Grid",         count: 0  },
    TRAILS:    { icon: "➡️",  label: "Threat Trails",        count: 0  },
    COVERAGE:  { icon: "⭕",  label: "Radar Circles",        count: 10 },
};

// Real-time scenario data via Socket.IO
const socket = io();
socket.on('state_update', (state) => {
    // state.threats, state.interceptors, state.sensors, state.defcon_level
    updateGlobeMarkers(state);
});
```

### 6. Scenario Engine Internals

The NORAD scenario engine maps threat types to FORGE-Sims capabilities:

```python
# Scenario threat types → FORGE-Sims mapping
THREAT_MAP = {
    "ICBM": {
        "simulator": "bmd-sim-icbm",
        "detection_chain": ["SBIRS (boost)", "DSP (boost)", "UEWR (midcourse)"],
        "interceptors": ["GMD (midcourse)", "SM-3 (midcourse)"],
        "countermeasures": "bmd-sim-decoy",
    },
    "IRBM": {
        "simulator": "bmd-sim-irbm",
        "detection_chain": ["SBIRS (boost)", "TPY-2 (terminal)"],
        "interceptors": ["THAAD (terminal)", "Patriot (terminal)"],
    },
    "HGV": {
        "simulator": "bmd-sim-hgv",
        "detection_chain": ["SBIRS (boost)", "STSS (glide)"],
        "interceptors": ["THAAD-ER (terminal)"],
        "countermeasures": "Maneuvering (bmd-sim-hgv -maneuvers)",
    },
    "SLCM": {
        "simulator": "bmd-sim-slcm",
        "detection_chain": ["SBIRS (boost)", "Aegis (terminal)"],
        "interceptors": ["SM-6 (terminal)"],
    },
}

# Available scenarios (7 total)
SCENARIOS = [
    "decoy-stress",       # Single ICBM + 20 decoys
    "hypersonic-glide",   # 2x HGV threats
    "multi-axis",         # 6 threats from DPRK + Iran
    "nuclear-exchange",   # 15 ICBMs from Russia + DPRK
    "single-launch",      # 1 ICBM from DPRK
    "regional-conflict",  # 3x IRBM from Iran
    "satellite-crisis",   # Space event + debris
]
```

### 7. Daily Satellite Data Refresh (Cron)

```bash
# /home/wez/forge-sims/scripts/daily-sat-refresh.sh
# Runs at 3 AM UTC on darth (10.0.0.117), pushes to miner (207.244.226.151)
#
# Crontab entry:
# 0 3 * * * /home/wez/forge-sims/scripts/daily-sat-refresh.sh >> /var/log/sat-refresh.log 2>&1
#
# What it does:
# 1. Fetches TLE data from Celestrak for 14 groups (gps-ops, stations, weather, geo, etc.)
# 2. 5-second delay between groups (polite to Celestrak)
# 3. Converts to sat-data-service format
# 4. SCPs to miner
# 5. Restarts sat-data-service
# 6. Aborts if fewer than 10 sats fetched (Celestrak down)
```

---

## Deployment Reference

### Docker (NORAD)

```yaml
# docker-compose.yml on miner
norad-sim:
  image: norad-forge
  container_name: norad-sim
  network_mode: host
  volumes:
    - /home/wez/norad-sim/host-bin:/host-bin:ro
  environment:
    - FLASK_ENV=production
  restart: unless-stopped
```

### Systemd (sat-data-service)

```ini
# /etc/systemd/system/sat-data.service
[Unit]
Description=Satellite Data Service
After=network.target

[Service]
User=wez
WorkingDirectory=/home/wez/sat-data-service
ExecStart=/home/wez/sat-data-service/venv/bin/gunicorn \
    --bind 127.0.0.1:5020 --workers 2 --timeout 30 \
    app:app
Restart=always

[Install]
WantedBy=multi-user.target
```

### Systemd (norad-legacy)

```ini
# /etc/systemd/system/norad-legacy.service
[Unit]
Description=NORAD Legacy Mirror
After=network.target

[Service]
User=wez
WorkingDirectory=/home/wez/norad-legacy
ExecStart=/home/wez/norad-legacy/venv/bin/gunicorn \
    --bind 127.0.0.1:5011 --workers 2 \
    app:app
Restart=always

[Install]
WantedBy=multi-user.target
```

### Nginx (Sites)

```nginx
# /etc/nginx/sites-available/norad.stsgym.com
server {
    listen 443 ssl;
    server_name norad.stsgym.com;

    # CSP headers for Three.js + Leaflet + Socket.IO
    add_header Content-Security-Policy "
        default-src 'self';
        script-src 'self' 'unsafe-inline' 'unsafe-eval'
            cdnjs.cloudflare.com cdn.jsdelivr.net unpkg.com;
        connect-src 'self' wss:;
        img-src 'self' data: *.cartocdn.com;
        style-src 'self' 'unsafe-inline';
    ";

    location / {
        proxy_pass http://127.0.0.1:5008;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```