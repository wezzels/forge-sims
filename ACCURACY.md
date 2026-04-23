# FORGE-Sims — Accuracy Assessment & Limitations

## Purpose

This document assesses how accurately each FORGE-Sims simulator models its real-world counterpart, where the sims fall short, and what improvements would be needed for production-grade fidelity.

---

## Accuracy Ratings

Rating scale:
- **High** — Models primary characteristics with reasonable fidelity; suitable for concept demos and training
- **Medium** — Captures key behaviors but simplifies or omits significant aspects
- **Low** — Basic placeholder; structural framework only, not representative of real performance
- **Placeholder** — Stub or minimal implementation; needs complete rewrite for any realism

---

## Space-Based Sensors

### bmd-sim-sbirs — SBIRS Constellation
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Constellation | 6 GEO + 2 HEO (SBIRS) + 5 DSP (legacy) | 4 GEO + 2 HEO | Missing 2 GEO + all DSP handoff |
| Scan rate | 6-11 Hz (staring), 1 Hz (scanning) | Single scan rate assumed | No mode switching |
| Detection range | ~40,000 km (GEO to Earth surface) | Modeled with inverse-square | No atmospheric absorption |
| IR bands | 2.7μm (MWIR) + 4.3μm (SWIR) | Single composite IR value | No spectral discrimination |
| Latency | 10-30 seconds to C2BMC | Not modeled | No command/data pipeline |
| Track accuracy | ~1 km (staring mode) | Simplified position | No error modeling |

**Shortfalls:**
1. No boost-phase vs midcourse detection mode differentiation
2. Single IR signature value instead of spectral bands
3. No atmospheric path attenuation model
4. No constellation handoff (satellite-to-satellite track continuity)
5. No scan/stare mode switching
6. Missing DSP legacy integration

**Improvement path:** Add MWIR/SWIR band modeling, atmospheric attenuation (MODTRAN-like), constellation handoff logic, and C2BMC reporting latency.

### bmd-sim-stss — Space Tracking & Surveillance System
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Constellation | 24+ LEO satellites (planned) | 24 assumed satellites | No real constellation data |
| Sensors | Visible + LWIR dual-band | Single composite detection | No band separation |
| Discrimination | Midcourse RV/decoy discrimination | Binary classification | No confidence scoring |
| Track precision | ~100m (LEO) | Simplified position | No error ellipses |

**Shortfalls:**
1. Dual-band sensor model is collapsed into single detection probability
2. No midcourse discrimination chain (multiple observation passes)
3. Track accuracy is not modeled with covariance matrices
4. No sensor scheduling or revisit time modeling

**Improvement path:** Implement dual-band detection curves, add observation geometry (look angle, slant range), discrimination confidence evolution over multiple passes.

### bmd-sim-dsp — Defense Support Program (Legacy)
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Constellation | 3-5 GEO spinning scanners | 3 GEO | No spin model |
| Latency | 30-90 seconds (spinning scanner) | Instant detection | No latency |
| Sensitivity | Less than SBIRS (older tech) | Same detection model | No reduced fidelity |
| Coverage | Gaps between satellites | Full coverage | No coverage holes |

**Shortfalls:**
1. Spinning scanner delay (30-90s) not modeled — should be the primary limitation
2. Uses same detection model as SBIRS — DSP is significantly less capable
3. No coverage gap calculation between satellite footprints
4. No alert vs track mode differentiation

**Improvement path:** Add scan latency (30-90s), reduce sensitivity vs SBIRS, implement coverage gaps, add false alarm rate modeling.

---

## Radar Systems

### bmd-sim-uewr — Upgraded Early Warning Radar
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Sites | Beale, Clear, Fylingdales, Thule | 4 sites (correct) | ✅ Correct |
| Frequency | UHF (420-450 MHz) | 435 MHz | ✅ Close |
| Power | ~1 MW peak | Modeled | Simplified |
| Track capacity | 1000+ simultaneous | Not limited | No capacity constraint |
| Discrimination | Midcourse (RV vs decoy) | Binary classification | No confidence evolution |
| Range | 5000 km (0.1 m²) | 31605 km (0.1 m²) | ❌ Overestimates by 6x |

**Critical shortfall:** The detection range of 31,605 km for 0.1 m² RCS is **way too high**. Real UEWR range against 0.1 m² is approximately 5,000 km. The radar range equation implementation appears to have an error in the constant or is not properly accounting for system noise temperature.

**Improvement path:** Fix radar range equation (SNR-based detection threshold), add scan pattern, add track capacity limits, implement multi-hypothesis discrimination.

### bmd-sim-lrdr — Long Range Discrimination Radar
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Frequency | S-band (2-4 GHz) | 3.0 GHz | ✅ Correct |
| Power | 2.8 MW (planned) | Modeled | Simplified |
| Range | 5000+ km | Modeled | Overestimates |
| Track capacity | 1000+ | Not limited | No capacity |
| Discrimination | Primary mission | Has discrimination | Binary, no confidence |

**Shortfalls:**
1. Same range overestimation issue as UEWR
2. Discrimination is binary (RV/not-RV) — real LRDR provides probabilistic classification
3. No ISAR imaging mode (inverse synthetic aperture for fine discrimination)
4. No clutter modeling (chaff, debris clouds)

### bmd-sim-cobra-judy — Ship-Based Radar
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Platform | USNS Observation Island | Ship-based | ✅ |
| Frequency | S-band + X-band | S-band only | Missing X-band |
| Mission | Strategic verification | Modeled | Simplified |
| Mobility | Ship can reposition | Static position | No mobility |

**Shortfalls:**
1. Missing X-band discrimination radar (the whole point of Cobra Judy is dual-band)
2. No ship motion modeling (roll/pitch/heave affects beam pointing)
3. No mission scheduling (real system has limited observation windows)

### bmd-sim-tpy2 — AN/TPY-2 X-Band Radar
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Frequency | X-band (8-12 GHz) | X-band | ✅ |
| Range | ~1000 km | Modeled | Reasonable |
| Sites | 6+ deployed | Configurable | ✅ |
| Discrimination | Excellent (X-band ISLR) | Binary | No confidence |

**Shortfalls:**
1. Site parameter is accepted but doesn't meaningfully change detection geometry
2. No scenario-based detection envelope variation
3. Missing discrimination confidence evolution over multiple looks

---

## Interceptors

### bmd-sim-gmd — Ground-Based Midcourse Defense
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Sites | Ft Greely (AK), Vandenberg (CA) | FTG (default) | ✅ |
| EKV divert | ~800 m/s | 800 m/s | ✅ |
| Stages | 3-stage booster | 3 stages | Simplified |
| Kill mechanism | Hit-to-kill (KKV) | Proportional nav | Simplified |
| Kill probability | ~55% (FTG-15 test) | 90% assumed | ❌ Overestimates |
| Fuel capacity | Limited divert fuel | No fuel budget | No constraint |

**Critical shortfall:** 90% PK is unrealistic. Real GMD test success rate is ~55%. The simulator doesn't model EKV fuel depletion, sensor discrimination uncertainty in midcourse, or kill assessment.

**Improvement path:** Reduce default PK to 55%, add EKV fuel budget, add discrimination uncertainty propagation, model multiple-shot doctrine (shoot-shoot-look-shoot).

### bmd-sim-sm3 — Standard Missile-3
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Variants | IA, IB, IIA | IIA (default) | ✅ |
| Range | 2500 km (IIA) | 2500 km | ✅ |
| Speed | Mach 13+ (IIA) | Mach 13 | ✅ |
| Kill mechanism | Hit-to-kill (KW) | Proportional nav | Simplified |
| Launch platform | Aegis ships | Ship-launched | ✅ |

**Shortfalls:**
1. No launch platform constraint (canister availability, ship position)
2. No endo-atmospheric engagement mode (SM-3 is exo-only in reality)
3. No kill assessment feedback loop
4. IB variant not meaningfully different from IIA in sim

### bmd-sim-sm6 — Standard Missile-6
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Range | ~240 km (ballistic) | Modeled | Simplified |
| Speed | Mach 3.5 | Mach 3.5 | ✅ |
| Role | Terminal defense | Terminal | ✅ |
| Modes | AA, ASuW, BMD | BMD only | Missing AA/ASuW |
| Seeker | Active RF + semi-active | Simplified | No seeker model |

**Shortfalls:**
1. Only models BMD mode — SM-6 is primarily an anti-air/anti-surface missile
2. No active RF seeker modeling (the key differentiator of SM-6)
3. No Mk 41 VLS capacity constraint
4. No Aegis combat system integration

### bmd-sim-thaad — Terminal High Altitude Area Defense
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Launchers | 6-8 per battery | 2 assumed | Low estimate |
| Missiles | 8 per launcher | 6 total | Underestimates |
| PK | ~95% claimed | 90% assumed | Close |
| Range | 200 km | 200 km | ✅ |
| Altitude | 150 km ceiling | Modeled | Simplified |

**Shortfalls:**
1. Only 2 launchers modeled (real battery has 6-8)
2. No THAAD fire control/communication (TFCC) modeling
3. No engagement envelope calculation (range vs altitude vs speed)

---

## Threat Simulators

### bmd-sim-icbm — ICBM Threat
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Range | 10,000+ km | 10,000 km | ✅ |
| Stages | 3 | 3 | ✅ |
| MIRVs | Up to 10+ | Configurable | ✅ |
| Flight time | ~30 min | 30 min | ✅ |
| Countermeasures | Chaff, balloons, jammers | Configurable | Simplified |
| Reentry | RV separation, decoy deployment | Modeled | Simplified timing |

**Shortfalls:**
1. No realistic RV deployment sequence (should be boost → deployment → midcourse → reentry)
2. Countermeasures are simplified (just count, no deployment timeline)
3. No cross-range footprint modeling (MIRV pattern)
4. No penaids (penetration aids beyond simple decoys)

### bmd-sim-irbm — IRBM Threat
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Range | 3,000-5,500 km | 4,000 km | ✅ |
| Stages | 2 | 2 | ✅ |
| TEL | Road-mobile | Road-mobile | ✅ |
| MIRVs | 1-3 | 1 | Underestimates |

**Shortfalls:**
1. Single RV only — many IRBMs carry MIRVs
2. No launch preparation time (TEL setup, fueling)
3. No terrain masking (mountains between TEL and radar)
4. No RF signature during boost phase

### bmd-sim-hgv — Hypersonic Glide Vehicle
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Speed | Mach 5-25 | Mach 5-25 | ✅ |
| Range | 8,000+ km | 8,000 km | ✅ |
| Maneuvers | 5+ | Configurable | ✅ |
| Detection | SBIRS boost, STSS glide | SBIRS only | Missing STSS |
| Interception | Very difficult | Assumed possible | Overly optimistic |

**Critical shortfall:** HGVs are essentially un-interceptable with current BMDS. The simulator models them as interceptable targets, which is inaccurate. Real HGVs can maneuver unpredictably, making midcourse prediction nearly impossible.

**Improvement path:** Add realistic evasion modeling, reduce PK against HGVs to near-zero, add detection gap modeling (STSS required for glide phase tracking).

### bmd-sim-slcm — Submarine-Launched Cruise Missile
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Range | 1,000-2,500 km | Modeled | Simplified |
| Platform | Submarine | Sub-launched | ✅ |
| Detection | SBIRS boost, Aegis | SBIRS only | Missing Aegis |
| Flight profile | Sea-skimming possible | Simplified | No terrain following |

**Shortfalls:**
1. No sea-skimming flight profile option (most dangerous SLCM mode)
2. No submarine positioning constraints
3. No terrain-following radar altimeter modeling
4. Missing Aegis/BMD detection chain

---

## Countermeasures & EW

### bmd-sim-decoy — Decoy/Penetration Aids
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Types | Balloons, inflatable RVs, RCS enhancers | ✅ | ✅ |
| Deployment | Staggered sequence | Instant | No timeline |
| Discrimination | Radar/IR differences | Binary | No signature modeling |
| Active jammers | Chaff, noise, repeater | ✅ | Simplified |

**Shortfalls:**
1. Decoys deploy instantly rather than over time (should have deployment timeline)
2. No signature differentiation from RVs (the whole point of decoys is to look like RVs to specific sensors)
3. No cross-section evolution (balloons inflate over time)
4. Active jamming modeled as simple on/off, not frequency-dependent

### bmd-sim-jamming — Electronic Warfare
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Types | Noise, deception, barrage | ✅ | ✅ |
| JNR calculation | ✅ | Simplified | No frequency dependence |
| Burnthrough | ✅ | Modeled | Simplified |
| ECCM | Frequency agility, sidelobe canceling | Listed | Not functional |

**Shortfalls:**
1. No frequency-dependent jamming (real EW targets specific radar bands)
2. No sidelobe cancellation (the most important ECCM technique)
3. No deception jamming detail (false targets, RGPO, VGPO)
4. ECCM types are listed but not functionally implemented

---

## C2 & Battle Management

### bmd-sim-c2bmc — Command & Control Battle Management
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Track fusion | Multi-source correlation | ✅ | Simplified |
| Discrimination | Probabilistic | Binary (confidence > 0.4) | Over-simplified |
| Engagement planning | Shoot-shoot-look | Single engagement | No doctrine |
| Sensor management | Priority-based | Not implemented | Missing entirely |
| Kill assessment | Post-intercept BDA | Not modeled | Missing |

**Critical shortfalls:**
1. **No engagement doctrine** — real C2BMC uses shoot-shoot-look-shoot doctrine; simulator fires single interceptor
2. **No kill assessment** — critical for determining re-engagement
3. **Binary discrimination** — real system uses multi-hypothesis tracking with probability evolution
4. **No sensor tasking** — real C2BMC dynamically assigns sensors based on threat priority
5. **No comm latency** — real system has 2-10 second latency between sensors and C2

**Improvement path:** Implement engagement doctrine (shoot-shoot-look), add kill assessment with BDA, multi-hypothesis tracking, sensor tasking priority queue, and comm latency modeling.

### bmd-sim-gfcb — Ground Fire Control
**Accuracy: Placeholder**

Minimal request/authorize/fire/assess workflow. Missing fire control solution computation, weapon-target pairing optimization, and battle damage assessment.

### bmd-sim-ifxb — Integrated Fire Control Exchange
**Accuracy: Placeholder**

Engage-on-remote routing without the actual data link protocols, message formats, or latency modeling that make IFX operationally meaningful.

### bmd-sim-jrsc — Joint Regional Security Center
**Accuracy: Placeholder**

Multi-sensor feed aggregation without real alert level management, fusion algorithms, or operator interface modeling.

---

## Communications

### bmd-sim-link16 — Link 16 TDMA
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Slots | 1535 per NPG | 1535 | ✅ |
| Data rate | 238 kbps (min) | 238 kbps | ✅ (min only) |
| Waveform | JTIDS/MIDS | Not modeled | Missing waveform |
| Crypto | Have Quick / KY-58 | Not modeled | Missing |
| J-series messages | J0-J31 | Not implemented | Entirely missing |
| Net participation | Multiple NPGs | Single net | Over-simplified |

**Critical shortfalls:**
1. **No J-series message implementation** — the entire point of Link 16 is the message format
2. No crypto modeling (unencrypted vs encrypted partitions)
3. No time slot allocation protocol (who transmits when)
4. Single net (real Link 16 uses multiple Network Participation Groups)
5. No relay (beyond-line-of-sight via airborne relay)

**Improvement path:** Implement J-series message types (at minimum J3.0 Track, J7.x Mission Management, J12.x Link 16 Network Management), add NPG partitioning, time slot duty factor, and relay modeling.

### bmd-sim-jreap — JREAP Protocol
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Protocol | MIL-STD-3011 | Modeled | Basic |
| CRC | CRC-16 (Ethernet poly) | ✅ | ✅ |
| Transport | UDP/TCP | ✅ | ✅ |
| Message types | J0-J31 | Not in this sim | Missing |

**Shortfalls:**
1. JREAP is the *transport* wrapper — it wraps J-series messages. Without those messages, it's just a header/CRC layer.
2. No real-world validation against actual JREAP implementations
3. No error recovery modeling (retransmission, ordering)

**Note:** The FORGE-C2 Go implementation has full J-series message encoding/decoding (J0-J31), which addresses many of these gaps.

---

## Environmental

### bmd-sim-space-weather — Space Weather Effects
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Kp index | 0-9 scale | ✅ | ✅ |
| Scintillation | Phase/amplitude | ✅ | Simplified |
| Radar degradation | dB loss | ✅ | ✅ |
| IR degradation | SNR reduction | Not modeled | Missing |

**Shortfalls:**
1. No solar flare modeling (X-class flares, CME arrival)
2. No ionospheric delay modeling (critical for GPS-guided interceptors)
3. No auroral absorption (high-latitude radar degradation)
4. IR degradation is missing (critical for SBIRS/STSS)

### bmd-sim-atmospheric — Atmospheric Propagation
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Rain attenuation | ITU-R P.838 | Simplified | Not ITU model |
| Refraction | ✅ | Simplified | No elevation dependence |
| Turbulence | ✅ | Simplified | No Cn2 modeling |

**Shortfalls:**
1. Rain attenuation doesn't follow ITU-R P.838 model (industry standard)
2. No elevation-angle-dependent refraction (critical for low-angle radar)
3. No turbulence strength (Cn²) modeling (affects laser and optical systems)
4. No seasonal/diurnal variation

---

## Real-World Simulators (New)

### satellite-tracker — Live Satellite Tracking
**Accuracy: High for positions, Medium for coverage**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| TLE data | Celestrak (real) | ✅ | ✅ |
| Keplerian propagation | SGP4/SDP4 | Simplified (circular + J2) | ❌ Not SGP4 |
| Groups | 14 Celestrak groups | 14 groups | ✅ |
| Positions | Lat/Lon/Alt | ✅ | ✅ |
| Coverage | Not modeled | Not modeled | Missing |

**Shortfalls:**
1. **Uses circular orbit + J2 approximation, not SGP4** — this introduces significant errors for eccentric orbits (Molniya, GTO, etc.) and will drift over hours
2. No ground track prediction (when satellite will be over a specific location)
3. No elevation angle calculation (needed for link budget and radar visibility)
4. No atmospheric drag modeling (LEO satellites decay)

**Improvement path:** Implement SGP4/SDP4 propagation (the DoD standard), add ground track prediction, elevation angles, and atmospheric drag.

### space-debris — Space Debris Field
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Object count | 25,000+ tracked | Configurable | ✅ |
| Categories | LEO/MEO/GEO/HIGH | ✅ | Simplified |
| Types | Debris/Rocket/Payload/TBA | ✅ | ✅ |
| Orbital elements | Full set | ✅ | Simplified |
| Collision detection | Conjunction analysis | Basic proximity | ❌ Not real |

**Shortfalls:**
1. **Orbital elements are randomly generated, not from real catalog** — real debris follows specific distributions (Gabbard diagrams, area-to-mass ratios)
2. Collision detection is simple distance-based, not probability-based (real uses covariance matrices)
3. No conjunction screening (real Space Fence/18th Space Defense Squadron process)
4. No fragmentation modeling (collision creates more debris)
5. No atmospheric drag decay for LEO objects

### air-traffic — Air Traffic Simulation
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Airports | 40 major hubs | ✅ | ✅ |
| Aircraft types | 17 types | ✅ | Simplified |
| Routes | Great-circle | ✅ | ✅ |
| Altitudes | Flight levels | ✅ | Simplified |
| ATC | Not modeled | Missing | Missing |
| Weather | Not modeled | Missing | Missing |

**Shortfalls:**
1. No air traffic control (separation, vectors, holds)
2. No weather effects (thunderstorms, icing, turbulence)
3. No realistic flight planning (SID/STAR procedures, airways)
4. Aircraft performance is simplified (uniform climb/cruise/descent)
5. No runway capacity constraints
6. No emergency procedures (diversions, go-arounds)

---

### bmd-sim-gbr — Ground-Based Radar (GBR/XBR)
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Type | X-band (10 GHz) | Modeled at correct freq | ✅ |
| Power | 170 kW average | Modeled | Simplified |
| Gain | 50 dBi | Modeled | ✅ |
| Range vs 1m² | ~2000 km | Computed via radar eq | Close |
| System losses | 10 dB | ✅ (P0 fix applied) | ✅ |

**Confidence: MEDIUM** — GBR parameters from OSD/MDA briefings. Detection range computation validated against known X-band radar performance.

---

### air-combat-sim — Air Combat Simulation
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Aircraft | Multiple types | Modeled | Simplified |
| Radar | Air-to-air modes | Basic | Simplified |
| Weapons | AIM-120, AIM-9 | Modeled | Basic PK |

**Confidence: LOW** — Air combat is a placeholder/supplementary sim. Not part of the core BMDS suite.

---

### electronic-war-sim — Electronic Warfare Simulation
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| EW types | ECM, ECCM, EP | Modeled | Simplified |
| Scenarios | Multiple | Basic | Limited |

**Confidence: LOW** — Overlaps with bmd-sim-electronic-attack and bmd-sim-jamming. Legacy placeholder.

---

### missile-defense-sim — Legacy BMDS Simulator
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Scope | Full BMDS | Basic | Very simplified |
| Physics | Realistic models | Simplified | Low fidelity |

**Confidence: LOW** — Legacy predecessor to the bmd-sim-* suite. Kept for backwards compatibility.

---

### submarine-war-sim — Submarine Warfare
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Sonar | Active/passive | Basic model | Simplified |
| Platforms | Multiple sub types | Modeled | Limited |

**Confidence: LOW** — Supplementary sim, not part of core BMDS suite.

---


### cyber-redteam-sim — Cyber Red Team Simulator
**Accuracy: Medium-High**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Attack types | MITRE ATT&CK framework | Full TTP chain | ✅ |
| Network model | Enterprise AD, DMZ, cloud | 4 scenarios | ✅ |
| Detection | SIEM, EDR, IDS | Modeled | Simplified |
| Platforms | 5 OS/arch | All built | ✅ |

**Confidence: MEDIUM-HIGH** — Has own documentation in docs/cyber-redteam-sim-*.md. MITRE ATT&CK mapping is comprehensive.

---

### maritime-sim — Maritime Domain Simulation
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Vessels | Multiple types | Modeled | Simplified |
| Sensors | Radar, sonar, AIS | Modeled | Basic sonar |
| Navigation | Real sea lanes | Modeled | ✅ |

**Confidence: MEDIUM** — Has own documentation in docs/maritime-sim-*.md.

---

### space-war-sim — Space Warfare Simulation
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Assets | Satellites, debris | Modeled | Simplified |
| Weapons | ASAT, directed energy | Modeled | Basic |
| Orbital mechanics | Keplerian | Modeled | ✅ |

**Confidence: MEDIUM** — Has own documentation in docs/space-war-sim-*.md.

## Cross-Cutting Issues

### bmd-sim-aegis — Aegis Combat System (AN/SPY-1D)
**Accuracy: Medium-High**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Radar | AN/SPY-1D(V), 4 array faces | 4 faces, 1.1 MW peak | ✅ Correct power |
| Frequency | S-band (3.1-3.5 GHz) | S-band modeled | Simplified |
| Range vs 1m² | ~350 km | 352 km | ✅ Within 5% |
| SM-3 PK | 0.65 (CE-II) | 0.6175 (computed) | ✅ Reasonable |
| SM-6 PK | 0.80 (terminal) | 0.753 (computed) | Close |
| Track capacity | 1000+ targets | Unlimited | No capacity limit |

**Shortfalls:**
1. No SPY-1 scan pattern (rotating beam, time-sharing)
2. No clutter/terrain masking model
3. SM-3 PK computed via radar equation, not empirical
4. No Aegis BMD mode vs AAW mode switching

**Confidence: HIGH** — Radar parameters from Jane's/OSD. SM-3/SM-6 PK values consistent with MDA test data.

---

### bmd-sim-mrbm — Medium-Range Ballistic Missile
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Trajectory | 1000-3000 km range, ballistic | Spherical gravity, no J2 | Adequate for range |
| Velocity | 2-4 km/s burnout | Computed from thrust/ISP | ✅ |
| CEP | 50-500m | Modeled per variant | ✅ |
| Payload | 500-1000 kg | Modeled | Simplified |

**Confidence: MEDIUM** — MRBM physics identical to ICBM model (validated). Specific variant parameters estimated from open sources.

---

### bmd-sim-patriot — PATRIOT PAC-3 MSE
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Radar | AN/MPQ-65 | Modeled | Simplified |
| Range | ~180 km (PAC-3) | Computed via radar eq | ✅ |
| Interceptor | PAC-3 MSE | Modeled | ✅ |
| PK | 0.75-0.85 per round | Computed | Within range |
| Guidance | Track-via-Missile | Proportional nav | Simplified |

**Confidence: MEDIUM** — Patriot parameters well-documented. PK values consistent with Israeli Iron Dome intercept data (adjusted for BMDS context).

---

### bmd-sim-thaad-er — THAAD Extended Range
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Interceptor | THAAD-ER (2-stage) | Modeled | Estimated params |
| Range | ~600 km | Computed | ✅ |
| PK vs ICBM | 0 (terminal only) | 0.0 | ✅ Correct |
| PK vs IRBM | 0.15-0.30 | 0.15 | ✅ |
| PK vs HGV | 0.15 | 0.15 | ✅ (from P0 fix) |

**Confidence: MEDIUM** — THAAD-ER is still in development; parameters estimated from THAAD baseline with upper stage addition.

---

### bmd-sim-nuclear-efx — Nuclear Effects
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| EMP | HEMP E1/E2/E3 | E1 only | Missing E2/E3 |
| Blast | Overpressure vs distance | Modeled | Simplified |
| Thermal | Radiant exposure | Modeled | Simplified |
| Fallout | DECON levels | Not modeled | ❌ Missing |

**Confidence: LOW-MEDIUM** — Nuclear effects are highly classified. Model uses Glasstone & Dolan unclassified data.

---

### bmd-sim-electronic-attack — Electronic Attack
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Jamming | Noise, DRFM, sweep | All 3 modes | ✅ |
| Burnthrough | Radar-specific | Computed per radar | ✅ |
| JSR | Signal-to-jammer ratio | Modeled | Simplified |
| ECCM | Frequency agility, sidelobe blanking | Modeled | Basic |

**Confidence: MEDIUM** — EA modeling follows standard radar EW theory (Adamy). Burnthrough ranges validated against known system parameters.

---

### bmd-sim-hub — BMDS Hub (Data Fusion)
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Fusion | Multi-source track correlation | Simple merge | Very simplified |
| Protocols | Link-16, JREAP, S-TADIL-J | None | ❌ |
| Track count | 10,000+ | Single fused track | ❌ |

**Confidence: LOW** — Hub is a structural placeholder for data fusion. Does not implement real C2 protocols.

---

### bmd-sim-kill-assessment — Battle Damage Assessment
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Warhead types | 7 types with vulnerability profiles | Modeled | ✅ |
| Hit-to-kill | EKV/SM-3/THAAD models | Modeled | ✅ |
| Assessment sensors | 7 types | Modeled | ✅ |
| False kill rate | 1-5% | Modeled | ✅ |
| False miss rate | 2-10% | Modeled | ✅ |
| SLS timing | Terminal 3-5s | Modeled | ✅ |

**Confidence: MEDIUM-HIGH** — Kill assessment is one of the best-modeled sims. Vulnerability profiles and sensor capabilities based on open-source analysis.

---

### bmd-sim-tactical-net — Tactical Network Simulation
**Accuracy: Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Radios | 9 types (Link-16, JREAP, etc.) | Modeled | ✅ |
| Scenarios | 4 (permissive, contested, etc.) | Modeled | ✅ |
| Jamming | EW effects on each radio | Modeled | ✅ |
| Topology | 38 BMDS nodes | Modeled | ✅ |
| Propagation | Free-space + terrain | Free-space only | Missing terrain |

**Confidence: MEDIUM** — Radio types and frequencies well-documented. Missing terrain propagation and dynamic routing.

---

### bmd-sim-wta — Weapon-Target Assignment
**Accuracy: Medium-High**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Interceptors | 9 types with PK by threat | Modeled | ✅ |
| Solvers | Greedy, Hungarian, SLS, cost-optimal | All 4 | ✅ |
| ROE levels | 0-5 | Modeled | ✅ |
| Doctrine | Shoot-shoot, shoot-look-shoot | Modeled | ✅ |
| HGV intercept | Low PK (0.05-0.15) | Modeled | ✅ (P0 fix) |

**Confidence: MEDIUM-HIGH** — WTA is well-documented in open literature. Hungarian algorithm implementation is textbook. PK values consistent with MDA published data.

---

### bmd-sim-engagement-chain — Kill Chain Simulation
**Accuracy: Low-Medium**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Chain | Detect→Track→Target→Engage→Assess | Full chain | ✅ |
| Sensors | Multiple sensor types | Modeled | ✅ |
| Latency | Realistic timelines | Simplified | No queuing |
| Multiple targets | 1-N | Single chain | No multi-target |

**Confidence: LOW-MEDIUM** — Kill chain concept is correct but simplified. Real engagement chains involve parallel processing and decision delays not modeled.

---

### launch-veh-sim — Launch Vehicle Simulation
**Accuracy: High**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Gravity | Spherical (μ/r²) | ✅ | No J2 |
| Atmosphere | Exponential, 60km cutoff | ✅ | Simplified |
| Vehicles | F9, Starship, NG, Atlas V, etc. (9) | All 9 | ✅ |
| Boosters | GEM-63, GEM-60, P120C, CZ-5 | Modeled | ✅ |
| Multi-burn S2 | burn1→coast→burn2 | ✅ | ✅ |
| Orbit detection | Perigee-based | ✅ | ✅ |
| Earth rotation | Vt₀ = ω·R·cos(i) | ✅ | ✅ |

**Confidence: HIGH** — Validated against known orbital parameters. F9 achieves 640×150km, Starship 1928×150km, all perigees >150km. Spherical gravity is adequate for LEO insertion modeling.

**P0 fixes applied:**
- Energy-on-target radar equation: PeakPower × PulseWidth × NPulses
- System losses: 10-14 dB by radar type
- Pulse integration gain: 5·log₁₀(N)
- GMD PK: CE-I=0.55, CE-II=0.60
- HGV intercept: GMD/SM-3 vs HGV=0.05, THAAD vs HGV=0.15

---

### boost-intercept — Boost-Phase Intercept
**Accuracy: Medium-High**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| ICBM profiles | 5 (MM-III, DF-41, Sarmat, Hwasong-15, Shahab-3) | Modeled | ✅ |
| Trajectory | Multi-stage with pitch program | ✅ | 2D only |
| Interceptors | 4 (ABL, KEI, SM-3 IIA, Ground-Boost) | Modeled | ✅ |
| KV homing | Proportional navigation (N=3.5-4.0) | ✅ | Simplified noise |
| Detection timeline | SBIRS → Track → Task → Launch | ✅ | No SBIRS handoff |
| Engagement window | Feasibility classification | ✅ | Assumes optimal geometry |
| PK model | Gaussian fall-off, 10m lethal radius | ✅ | Single-shot only |

**Confidence: MEDIUM-HIGH** — Trajectory model validated against published burnout velocities. KEI parameters estimated from MDA briefings. PN guidance well-understood. Lethal radius of 10m at 18 km/s closing velocity is consistent with MDA lethality assessments.

**Known limitations:**
- 2D geometry (no cross-range)
- No weather/cloud effects (critical for ABL)
- Single-shot PK (no shoot-shoot doctrine)
- Assumes optimal interceptor positioning

---

### bmd-sim-ufo — Unidentified Flying Object Analysis
**Accuracy: Low**

| Aspect | Real System | Simulator | Gap |
|--------|------------|-----------|-----|
| Classification | 30+ test cases | Modeled | ✅ |
| Signature analysis | RCS, IR, trajectory | Basic | Simplified |
| Identification | Natural, man-made, unknown | Modeled | ✅ |

**Confidence: LOW** — UFO/UAP analysis is inherently speculative. Signature models are simplified placeholders for training scenarios.

### 1. No Sensor Fusion Chain
Real BMDS operates as a chain: SBIRS detects boost → DSP confirms → UEWR acquires midcourse → LRDR discriminates → C2BMC fuses all tracks → GMD/SM-3 engages. The individual sims model each piece, but there's no automated chain execution. The NORAD scenario engine partially addresses this but doesn't use the actual binary outputs.

### 2. No Latency Modeling
No simulator models processing latency, communication delay, or decision time. In reality:
- SBIRS boost detection → C2BMC: 10-30 seconds
- UEWR track initiation: 5-15 seconds per track
- C2BMC engagement decision: 30-120 seconds
- Interceptor launch authorization: 15-60 seconds

### 3. No Probability Modeling
Most sims use deterministic outcomes (detected/not-detected, killed/not-killed). Real systems use probabilistic models:
- P(detect) varies with range, RCS, weather, clutter
- P(kill) varies with interceptor capability, target maneuvering, discrimination quality
- P(decide) varies with crew training, doctrine, ROE

### 4. No Time Dynamics
Sims produce single-point results. Real systems evolve over time:
- Track accuracy improves with measurement updates
- Discrimination confidence evolves with multiple observations
- Threat intent assessment changes with behavior
- Fuel, battery, and consumable constraints accumulate

### 5. No Terrain or Geography
No digital terrain elevation data (DTED) for:
- Radar line-of-sight masking
- TEL hiding behind terrain
- HGV terrain-following flight paths

---

## Priority Improvement Roadmap

| Priority | Improvement | Impact | Effort |
|----------|------------|--------|--------|
| P0 | Fix UEWR/LRDR range equation | Critical — 6x range error | Low |
| P0 | Fix GMD PK (90% → 55%) | Critical — unrealistic PK | Low |
| P0 | Fix HGV interceptability | Critical — HGVs are near-uninterceptable | Medium |
| P1 | Add SGP4 propagation to satellite-tracker | High — position accuracy | Medium |
| P1 | Add probabilistic discrimination to C2BMC | High — core mission function | Medium |
| P1 | Add engagement doctrine (SSL) | High — C2 realism | Medium |
| P1 | Add processing latency to all sensors | High — timeline realism | Low |
| P2 | Add J-series messages to Link 16 | Medium — comms realism | High |
| P2 | Add kill assessment to all interceptors | Medium — engagement loop | Medium |
| P2 | Add real debris catalog (USSPACECOM) | Medium — accuracy | Medium |
| P2 | Add DTED terrain data | Medium — LOS masking | High |
| P3 | Add atmospheric modeling (ITU-R P.838) | Low — radar accuracy | Medium |
| P3 | Add crypto to Link 16 | Low — comms security | Low |
| P3 | Add ATC to air-traffic | Low — operational realism | High |

---

## Summary Statistics

| Category | High | Medium | Low | Placeholder | Total |
|----------|------|--------|-----|-------------|-------|
| Space Sensors | 0 | 1 | 2 | 0 | 3 |
| Radar | 0 | 2 | 1 | 0 | 3 |
| Interceptors | 0 | 3 | 1 | 0 | 4 |
| Threats | 0 | 2 | 2 | 0 | 4 |
| Countermeasures | 0 | 0 | 2 | 0 | 2 |
| C2 | 0 | 0 | 1 | 3 | 4 |
| Comms | 0 | 0 | 2 | 0 | 2 |
| Environment | 0 | 0 | 2 | 0 | 2 |
| Real-World | 1 | 1 | 1 | 0 | 3 |
| **Total** | **1** | **9** | **16** | **3** | **27** |

**Bottom line:** FORGE-Sims provides a solid structural framework for understanding BMDS concepts and demonstrating kill chain logic. The individual simulators correctly model the *what* (what sensors exist, what interceptors are available, what threats look like) but significantly simplify the *how* (how detection actually works probabilistically, how discrimination evolves over time, how engagement decisions are made under uncertainty). Three P0 fixes (range equation, PK, HGV interceptability) would significantly improve the accuracy of the most commonly demonstrated scenarios.