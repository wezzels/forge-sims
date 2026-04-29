# FORGE-Sims Binary Status

**Last Updated:** 2026-04-27 (Round 20)
**Total Binaries:** 66 (64 sim binaries + 2 install scripts)
**Platforms:** linux-x86, darwin-arm64, windows-amd64

## BMD Simulators (36 binaries)

| Binary | Source Repo | Description | CLI | JSON | Data |
|--------|-------------|-------------|-----|------|------|
| bmd-sim-aegis | bmd-sim-aegis | Aegis Combat System (AN/SPY-1D, SM-3, SM-6) | ✅ | ✅ | ✅ |
| bmd-sim-atmospheric | bmd-sim-atmospheric | Atmospheric propagation model | ✅ | ✅ | — |
| bmd-sim-c2bmc | bmd-sim-c2bmc | Command & Control Battle Management | ✅ | ✅ | ✅ |
| bmd-sim-cobra-judy | bmd-sim-cobra-judy | Ship-based radar (AN/SPQ-78) | ✅ | ✅ | — |
| bmd-sim-decoy | bmd-sim-decoy | Decoy/penetration aids simulation | ✅ | ✅ | ✅ |
| bmd-sim-dsp | bmd-sim-dsp | Defense Support Program (legacy) | ✅ | ✅ | — |
| bmd-sim-electronic-attack | bmd-sim-electronic-attack | Electronic warfare/jamming | ✅ | ✅ | — |
| bmd-sim-gbr | bmd-sim-gbr | Ground-Based Radar (GBR/XBR) | ✅ | ✅ | — |
| bmd-sim-gfcb | bmd-sim-gfcb | Ground Fire Control Node | ✅ | ✅ | ✅ |
| bmd-sim-gmd | bmd-sim-gmd | Ground-Based Midcourse Defense | ✅ | ✅ | — |
| bmd-sim-hgv | bmd-sim-hgv | Hypersonic Glide Vehicle threat | ✅ | ✅ | ✅ |
| bmd-sim-hub | bmd-sim-hub | BMDS data fusion hub | ✅ | ✅ | ✅ |
| bmd-sim-icbm | bmd-sim-icbm | ICBM threat (MM-III, DF-41, Sarmat) | ✅ | ✅ | — |
| bmd-sim-ifxb | bmd-sim-ifxb | Integrated Fire Control Exchange | ✅ | ✅ | ✅ |
| bmd-sim-irbm | bmd-sim-irbm | IRBM threat (DF-26, Pershing-II) | ✅ | ✅ | — |
| bmd-sim-jamming | bmd-sim-jamming | Electronic countermeasures | ✅ | ✅ | ✅ |
| bmd-sim-jreap | bmd-sim-jreap | JREAP protocol simulation | ✅ | ✅ | ✅ |
| bmd-sim-jrsc | bmd-sim-jrsc | Joint Regional Security Center | ✅ | ✅ | ✅ |
| bmd-sim-link16 | bmd-sim-link16 | Link-16 TDMA network | ✅ | ✅ | ✅ |
| bmd-sim-lrdr | bmd-sim-lrdr | Long Range Discrimination Radar | ✅ | ✅ | — |
| bmd-sim-mrbm | bmd-sim-mrbm | Medium-Range Ballistic Missile | ✅ | ✅ | — |
| bmd-sim-nuclear-efx | bmd-sim-nuclear-efx | Nuclear effects (EMP, blast, thermal) | ✅ | ✅ | — |
| bmd-sim-patriot | bmd-sim-patriot | PATRIOT PAC-3 MSE | ✅ | ✅ | ✅ |
| bmd-sim-sbirs | bmd-sim-sbirs | SBIRS constellation (IR warning) | ✅ | ✅ | — |
| bmd-sim-slcm | bmd-sim-slcm | Submarine-Launched Cruise Missile | ✅ | ✅ | — |
| bmd-sim-sm3 | bmd-sim-sm3 | Standard Missile-3 (all variants) | ✅ | ✅ | — |
| bmd-sim-sm6 | bmd-sim-sm6 | Standard Missile-6 | ✅ | ✅ | — |
| bmd-sim-space-weather | bmd-sim-space-weather | Space weather effects on sensors | ✅ | ✅ | — |
| bmd-sim-stss | bmd-sim-stss | Space Tracking & Surveillance System | ✅ | ✅ | — |
| bmd-sim-thaad | bmd-sim-thaad | THAAD (terminal defense) | ✅ | ✅ | ✅ |
| bmd-sim-thaad-er | bmd-sim-thaad-er | THAAD Extended Range | ✅ | ✅ | — |
| bmd-sim-tpy2 | bmd-sim-tpy2 | AN/TPY-2 X-band radar | ✅ | ✅ | — |
| bmd-sim-uewr | bmd-sim-uewr | Upgraded Early Warning Radar | ✅ | ✅ | — |
| bmd-sim-kill-assessment | bmd-sim-kill-assessment | Battle damage assessment | ✅ | ✅ | — |
| bmd-sim-tactical-net | bmd-sim-tactical-net | Tactical network (9 radio types) | ✅ | ✅ | — |
| bmd-sim-wta | bmd-sim-wta | Weapon-target assignment (4 solvers) | ✅ | ✅ | — |
| bmd-sim-engagement-chain | bmd-sim-engagement-chain | Kill chain simulation | ✅ | ✅ | — |

## Enhanced Simulators (8 binaries)

| Binary | Description | CLI | JSON | Data |
|--------|-------------|-----|------|------|
| boost-intercept | Boost-phase intercept (5 ICBM, 4 interceptors) | ✅ | ✅ | — |
| satellite-tracker | Live satellite tracking (Celestrak) | ✅ | ✅ | — |
| space-debris | Space debris field simulation | ✅ | ✅ | — |
| air-traffic | Air traffic simulation | ✅ | ✅ | — |
| launch-veh-sim | Launch vehicle simulation (9 vehicles) | ✅ | ✅ | — |
| air-combat-sim | Air combat (BVR kill chain) | ✅ | ✅ | ✅ |
| emp-effects | Nuclear EMP effects on BMDS | ✅ | ✅* | — |
| rcs | Radar cross-section modeling | ✅ | ✅* | — |

*\* Requires `-scenario` or `-threat` flag*

## Threat & Weapon Sims (9 binaries)

| Binary | Description | CLI | JSON | Data |
|--------|-------------|-----|------|------|
| hgv | Hypersonic Glide Vehicle | ✅ | ✅ | ✅ |
| ir-signature | IR signature modeling | ✅ | ✅ | ✅ |
| kill-chain-rt | Real-time kill chain | ✅ | ✅ | ✅ |
| kill-assessment | Battle damage assessment | ✅ | ✅ | — |
| satellite-weapon | ASAT engagement | ✅ | ✅* | — |
| population-impact-sim | Population impact modeling | ✅ | ✅ | — |
| missile-defense-sim | Legacy BMDS simulator | ✅ | ✅ | ✅ |
| submarine-war-sim | Submarine warfare | ✅ | ✅ | ✅ |
| ufo | UFO/UAP analysis (30 cases) | ✅ | ✅ | ✅ |

*\* Requires `-scenario` flag*

## Warfare Sims (7 binaries)

| Binary | Description | CLI | JSON | Data |
|--------|-------------|-----|------|------|
| cbrn-sim | Chemical/Biological/Radiological/Nuclear | ✅ | ✅ | — |
| cyber-kinetic-sim | Cyber attacks on kinetic infrastructure | ✅ | ✅ | — |
| information-ops-sim | Information operations/campaigns | ✅ | ✅ | — |
| land-combat-sim | Ground warfare | ✅ | ✅ | — |
| logistics-sustainment-sim | Supply chain/logistics | ✅ | ✅ | — |
| nuclear-effects-sim | Nuclear detonation modeling | ✅ | ✅ | — |
| space-data-network-sim | Satellite network resilience | ✅ | ✅ | — |
| sensor-jreap | JREAP-C sensor simulator (8 sensors, UDP/TCP) | ✅ | ✅ | FORGE-C2 data path |

## Config-Driven Sims (3 binaries)

| Binary | Description | CLI | JSON | Notes |
|--------|-------------|-----|------|-------|
| cyber-redteam-sim | Cyber red team operations | ✅ | ✅* | Needs `-scenario` |
| maritime-sim | Maritime warfare | ✅ | ✅* | Needs `-scenario` |
| space-war-sim | Space warfare | ✅ | ✅* | Needs `-scenario` |

*\* Requires config/scenario flags*

## Other Binaries (3)

| Binary | Description | CLI | JSON | Notes |
|--------|-------------|-----|------|-------|
| electronic-war-sim | Electronic warfare | ✅ | ✅ | ✅ | ✅ |
| wta | Weapon-target assignment | ✅ | ✅ | ✅ | ✅ |
| engagement-chain | Kill chain simulation | ✅ | ✅ | — |

## Install Scripts (excluded from counts)

| File | Description |
|------|-------------|
| install-intel-collection-sim.sh | Installs intel-collection-sim to /usr/local/bin |
| uninstall-intel-collection-sim.sh | Removes intel-collection-sim |

## Platform Support

| Platform | Count | Notes |
|----------|-------|-------|
| linux-x86 | 65 | Primary deployment target |
| darwin-arm64 | 62 | macOS Apple Silicon |
| windows-amd64 | 68 | Windows x86_64 |

## Cross-Compile Status

- ✅ All simulators cross-compile for darwin-arm64 and windows-amd64
- ⚠️ bmd-sim-gmd and bmd-sim-lrdr had build errors in earlier rounds (now resolved)

