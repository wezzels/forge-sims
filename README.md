# FORGE-Sims

Pre-built BMDS simulator binaries for FORGE-C2.

## Quick Start

```bash
# Clone the repo
git clone git@github.com:wezzels/forge-sims.git
cd forge-sims

# Make binaries executable
chmod +x binaries/linux-x86/*

# Run any simulator
./binaries/linux-x86/bmd-sim-sbirs
./binaries/linux-x86/bmd-sim-icbm
./binaries/linux-x86/bmd-sim-c2bmc
```

## Directory Structure

```
forge-sims/
├── binaries/
│   ├── linux-x86/      # Linux x86_64 binaries ✅
│   └── linux-arm/      # Linux ARM64 binaries
├── docker-compose.yml  # Docker Compose for all 25 simulators
├── Dockerfile.template  # Common Dockerfile template
├── deploy-dockerfiles.sh
├── FORGE-SIMS-ROADMAP.md
├── FORGE-SIMS-TODO.md
├── BINARY_STATUS.md
└── README.md
```

## Simulators

### 🔴 Space-Based Sensors

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-sbirs** | SBIRS Satellite Constellation | 4 GEO + 2 HEO, scan/stare modes, boost-phase detection |
| **bmd-sim-stss** | Space Tracking & Surveillance System | LEO dual-band IR/visible, midcourse discrimination |
| **bmd-sim-dsp** | Defense Support Program (Legacy) | 3 GEO spinning scanner, 30-90s latency |

### 🔴 Radar Systems

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-uewr** | Upgraded Early Warning Radar | UHF phased array, 4-site network, 42 dBi gain |
| **bmd-sim-lrdr** | Long Range Discrimination Radar | S-band, 5000+ km, 1000+ tracks, midcourse discrimination |
| **bmd-sim-cobra-judy** | Ship-Based Radar | S-band, ship-mounted, mission scheduling |

### 🔴 Interceptors

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-gmd** | Ground-Based Midcourse Defense | 3-stage EKV, proportional navigation |
| **bmd-sim-sm3** | Standard Missile-3 | IIA/IB/IIA variants, ship launch |
| **bmd-sim-sm6** | Standard Missile-6 | Endo-atmospheric, Mach 3.5, Mk 41 VLS |

### 🔴 Threat Simulators

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-icbm** | ICBM Threat | 3-stage, 10,000 km, MIRV deployment, countermeasures |
| **bmd-sim-irbm** | IRBM Threat | 2-stage, 4,000 km, road-mobile TEL |
| **bmd-sim-hgv** | Hypersonic Glide Vehicle | Mach 5-25, boost-glide, maneuver scheduler |

### 🟡 Countermeasures & EW

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-decoy** | Decoy/Penetration Aids | Balloons, inflatable RVs, RCS enhancers, jammers |
| **bmd-sim-jamming** | Electronic Warfare | Noise/deception jamming, JNR, burnthrough, ECCM |

### 🟡 C2 & Battle Management

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-c2bmc** | C2BMC Battle Manager | Multi-sensor fusion, track correlation, engagement planning |
| **bmd-sim-gfcb** | Ground Fire Control | Request/authorize/fire/assess workflow |
| **bmd-sim-ifxb** | Integrated Fire Control Exchange | Engage-on-remote, data routing |
| **bmd-sim-jrsc** | Joint Regional Security Center | Multi-sensor feeds, alert levels, data fusion |

### 🟡 Communications

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-link16** | Link 16 TDMA | 1535 slots, 238 kbps, J-series messaging |
| **bmd-sim-jreap** | JREAP Protocol | MIL-STD-3011, CRC-16, A/B/C variants |

### 🟢 Extended Range

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-thaad-er** | THAAD Extended Range | 200 km ceiling, 600 km range, Mach 10 |

### 🟢 Environmental

| Binary | Description | Key Features |
|--------|-------------|--------------|
| **bmd-sim-space-weather** | Space Weather Effects | Kp index, scintillation, radar/IR degradation |
| **bmd-sim-atmospheric** | Atmospheric Propagation | Rain attenuation, refraction, turbulence |

## Real-World Simulators

| Binary | Description | Key Features |
|--------|-------------|-------------|
| **satellite-tracker** | Live Satellite Tracking | 14 Celestrak groups, Keplerian propagation, NORAD IDs |
| **space-debris** | Orbital Debris Field | 25K+ objects, collision detection, LEO/MEO/GEO/HIGH |
| **air-traffic** | Air Traffic Simulation | 40 airports, 17 aircraft types, great-circle routing |

---

### From Binary (Linux x86_64)

```bash
# Download
git clone git@github.com:wezzels/forge-sims.git
cd forge-sims

# Run directly
./binaries/linux-x86/bmd-sim-sbirs
```

### From Source

Each simulator has its own Go module. Build any of them:

```bash
# Clone a simulator
git clone git@idm.wezzel.com:crab-meat-repos/bmd-sim-sbirs.git
cd bmd-sim-sbirs

# Build
go build -o bmd-sim-sbirs ./cmd/sbirs/

# Run
./bmd-sim-sbirs
```

### With Docker Compose

```bash
# Start all 25 simulators
docker-compose up -d

# Start specific simulator
docker-compose up -d bmd-sim-sbirs bmd-sim-c2bmc bmd-sim-gmd

# View logs
docker-compose logs -f bmd-sim-sbirs

# Stop
docker-compose down
```

### Cross-Compile

```bash
# macOS
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bmd-sim-sbirs-mac ./cmd/sbirs/

# ARM64 (Raspberry Pi, Apple Silicon)
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bmd-sim-sbirs-arm64 ./cmd/sbirs/

# Windows
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bmd-sim-sbirs.exe ./cmd/sbirs/
```

---

## Usage Examples

### SBIRS — Detect an ICBM Launch

```go
package main

import (
    "fmt"
    sbirs "github.com/forge-c2/bmd-sim-sbirs/pkg/sensor"
    constellation "github.com/forge-c2/bmd-sim-sbirs/pkg/constellation"
)

func main() {
    // Create SBIRS sensor and constellation
    sensor := sbirs.NewSBIRSSensor()
    c := constellation.NewConstellation()
    
    // Detect ICBM boost phase (IR signature = 1e8 W/sr at GEO distance)
    detected, snr := sensor.DetectIR(1e8, 35786000)
    fmt.Printf("SBIRS detection: %v, SNR: %.1f dB\n", detected, snr)
    
    // Check satellite coverage over launch area
    visible := c.VisibleSatellites(35, 130, 200000)
    fmt.Printf("Visible satellites over launch: %v\n", visible)
}
```

### UEWR — Track a Target

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-uewr/pkg/radar"
    "github.com/forge-c2/bmd-sim-uewr/pkg/sites"
)

func main() {
    // Create UEWR at Beale AFB
    r := radar.NewUEWRRadar("BEALE")
    
    // Detect 0.1 m² target at 3000 km
    result := r.Detect(0.1, 3000000)
    fmt.Printf("UEWR detection: %v, SNR: %.1f dB\n", result.Detected, result.SNR)
    
    // Check which sites can see a target
    net := sites.NewSiteNetwork()
    visible := net.CanSee(35, 130, 200000)
    fmt.Printf("Visible sites: %v\n", visible)
}
```

### ICBM — Generate Trajectory

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-icbm/pkg/threat"
    "github.com/forge-c2/bmd-sim-icbm/pkg/mirv"
)

func main() {
    // Create ICBM threat
    icbm := threat.NewICBM()
    icbm.SetLaunchTarget(35, 130, 38.9, -77) // NK to DC
    
    // Generate trajectory
    points := icbm.Trajectory(100)
    fmt.Printf("ICBM: range=%.0f km, apogee=%.0f km, flight=%.0f min\n",
        icbm.Range/1000, icbm.Apogee()/1000, icbm.FlightTime()/60)
    
    // MIRV deployment
    ms := mirv.NewMIRVSystem(3, 55, -30, 800000)
    targets := ms.DispersionPattern(55, -30, 100)
    for i, t := range targets {
        fmt.Printf("  RV %d: (%.1f°, %.1f°)\n", i+1, t[0], t[1])
    }
}
```

### C2BMC — Multi-Sensor Fusion

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-c2bmc/pkg/c2bmc"
)

func main() {
    c := c2bmc.NewC2BMC()
    
    // SBIRS detects boost
    c.ProcessReport(c2bmc.SensorReport{
        SensorType: c2bmc.SensorSBIRS, SensorID: "GEO-1", TrackID: 1,
        Latitude: 35, Longitude: 130, Altitude: 200000, Speed: 5000,
        RCS: 0.1, IRSignature: 1e8, Confidence: 0.9,
    })
    
    // UEWR confirms
    c.ProcessReport(c2bmc.SensorReport{
        SensorType: c2bmc.SensorUEWR, SensorID: "BEALE", TrackID: 1,
        Latitude: 35.01, Longitude: 129.99, Altitude: 199000, Speed: 4950,
        RCS: 0.11, IRSignature: 1e7, Confidence: 0.85,
    })
    
    // Discriminate
    c.Discriminate()
    fmt.Printf("Track classified: IsRV=%v, Confidence=%.2f\n",
        c.Tracks[1].IsRV, c.Tracks[1].Confidence)
    
    // Engagement plan
    plan := c.EngagementPlan()
    fmt.Printf("Engagement plan: %d threats\n", len(plan))
}
```

### HGV — Hypersonic Threat Assessment

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-hgv/pkg/glide"
    "github.com/forge-c2/bmd-sim-hgv/pkg/detection"
)

func main() {
    hgv := glide.NewHGV()
    hgv.SetLaunchTarget(35, 130, 38, -77)
    
    // Generate trajectory
    points := hgv.Trajectory(100)
    fmt.Printf("HGV: range=%.0f km, apogee=%.0f km, flight=%.0f min\n",
        hgv.Range/1000, hgv.Apogee()/1000, hgv.FlightTime()/60)
    
    // Detection assessment
    boost, glide := detection.AssessSBIRS(hgv.IRSignatureBoost, hgv.IRSignatureGlide, 3.0)
    fmt.Printf("SBIRS: boost=%v, glide=%v\n", boost, glide)
}
```

### LRDR — Discrimination

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-lrdr/pkg/radar"
    "github.com/forge-c2/bmd-sim-lrdr/pkg/discrimination"
)

func main() {
    lr := radar.NewLRDR()
    
    // Detect RV at 3000 km
    result := lr.Detect(0.1, 3000000)
    fmt.Printf("LRDR detection: %v, SNR=%.1f dB\n", result.Detected, result.SNR)
    
    // Discriminate RV vs decoy
    d := discrimination.NewDiscriminator()
    rvResult := d.Discriminate(0.1, 5000, 6000, 500000)
    decoyResult := d.Discriminate(0.5, 100, 2000, 100000)
    fmt.Printf("RV: %s (conf=%.2f)\n", rvResult.IsRV, rvResult.Confidence)
    fmt.Printf("Decoy: IsDecoy=%v\n", decoyResult.IsDecoy)
}
```

### EW — Jamming Assessment

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-jamming/pkg/jamming"
)

func main() {
    // Noise jammer: 1 MW, S-band, 500 km away
    j := jamming.NewNoiseJammer(1e6, 3e9, 100e6, 500000)
    
    // Calculate JNR at radar
    jnr := j.JNR(50, 1e6)
    fmt.Printf("JNR: %.1f dB, Effect: %s\n", jnr, jamming.JammingEffect(jnr))
    
    // ECCM effectiveness
    fmt.Printf("Frequency agility: %.0f dB reduction\n",
        jamming.ECCMEffectiveness(jamming.ECCMFreqAgility))
}
```

### Space Weather Impact

```go
package main

import (
    "fmt"
    "github.com/forge-c2/bmd-sim-space-weather/pkg/spaceweather"
)

func main() {
    quiet := spaceweather.NewQuietConditions()
    storm := spaceweather.NewStormConditions()
    
    fmt.Printf("Quiet: Kp=%.1f, radar effect=%s\n", 
        quiet.KpIndex, quiet.EffectOnRadar())
    fmt.Printf("Storm: Kp=%.1f, radar effect=%s, degradation=%.1f dB\n",
        storm.KpIndex, storm.EffectOnRadar(), storm.RadarDegradation())
}
```

---

## Integration Testing

The `bmd-sim-integration` repo provides cross-simulator integration tests:

```bash
git clone git@idm.wezzel.com:crab-meat-repos/bmd-sim-integration.git
cd bmd-sim-integration
go test -v ./...
```

### Test Coverage

| Test | Description |
|------|-------------|
| TestFullEngagementChain | SBIRS → C2BMC → GMD complete chain |
| TestC2BMC_Discrimination | RV vs decoy discrimination |
| TestUEWR_Detection | UEWR radar detection |
| TestLRDR_Discrimination | LRDR midcourse discrimination |
| TestICBM_Trajectory | ICBM trajectory phases |
| TestSBIRS_Constellation | SBIRS satellite coverage |
| TestMultiThreat | 3 RVs + 2 decoys engagement |

---

## Docker Compose Reference

| Port | Simulator |
|------|-----------|
| 5010 | SBIRS |
| 5011 | STSS |
| 5012 | DSP |
| 5020 | UEWR |
| 5021 | LRDR |
| 5022 | Cobra Judy |
| 5030 | GMD |
| 5031 | SM-3 |
| 5032 | SM-6 |
| 5033 | THAAD-ER |
| 5040 | ICBM |
| 5041 | IRBM |
| 5042 | HGV |
| 5043 | SLCM |
| 5050 | Decoy |
| 5051 | Jamming |
| 5060 | C2BMC |
| 5061 | GFCB |
| 5062 | IFXB |
| 5063 | JRSC |
| 5070 | Link 16 |
| 5071 | JREAP |
| 5080 | Space Weather |
| 5081 | Atmospheric |
| 5090 | Satellite Tracker |
| 5091 | Space Debris |
| 5092 | Air Traffic |

---

## Source Repos

All simulators are available on IDM GitLab:

```
git@idm.wezzel.com:crab-meat-repos/bmd-sim-<name>.git
```

Replace `<name>` with: sbirs, stss, dsp, uewr, lrdr, cobra-judy, gmd, sm3, sm6, thaad-er, icbm, irbm, hgv, slcm, decoy, jamming, c2bmc, gfcb, ifxb, jrsc, link16, jreap, space-weather, atmospheric

New real-world simulators: satellite-tracker, space-debris, air-traffic

Core library: `git@idm.wezzel.com:crab-meat-repos/bmd-sim-core.git`

Shared CLI: `git@idm.wezzel.com:crab-meat-repos/sim-cli.git`

Integration tests: `git@idm.wezzel.com:crab-meat-repos/bmd-sim-integration.git`

---

## Statistics

| Metric | Value |
|--------|-------|
| Total Repos | 30 (27 sims + sim-cli + integration + forge-sims + satellite-tracker + space-debris + air-traffic) |
| Total Tests | ~320 |
| Total Binaries | 33 |
| Languages | Go 1.22.2 |
| Platforms | Linux x86_64, ARM64, macOS, Windows (cross-compile) |

---

## Documentation

| File | Description |
|------|-------------|
| [EVALUATION.md](EVALUATION.md) | Full 33-sim evaluation with test results |
| [ACCURACY.md](ACCURACY.md) | Accuracy assessment, limitations, and improvement roadmap |
| [NORAD-INTEGRATION.md](NORAD-INTEGRATION.md) | norad.stsgym.com integration guide with examples |
| [DOCS.md](DOCS.md) | Comprehensive API and usage documentation |
| [FORGE-SIMS-ROADMAP.md](FORGE-SIMS-ROADMAP.md) | Development roadmap |
| [scripts/daily-sat-refresh.sh](scripts/daily-sat-refresh.sh) | Daily Celestrak TLE refresh cron |