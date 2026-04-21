# API_REFERENCE.md — Forge-Sims JSON Output Schema

Complete reference for all 47 binaries. Each entry documents invocation flags, JSON output fields, and example output captured from live runs.

## Table of Contents

- [BMDS Sensors](#bmds-sensors)
- [BMDS Threats](#bmds-threats)
- [BMDS Interceptors](#bmds-interceptors)
- [C2/Infrastructure](#c2infrastructure)
- [Support Sims](#support-sims)
- [Config-Driven Sims](#config-driven-sims)
- [Engine Libraries](#engine-libraries)

---

## BMDS Sensors

### bmd-sim-aegis

**Description**: Aegis BMD Combat System Simulator (AN/SPY-1D(V))

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "aegis",
  "status": "complete",
  "parameters": {
    "array_faces": 4,
    "peak_power_kw": 1100,
    "radar": "AN/SPY-1D(V)",
    "range_01m2_km": 351.87,
    "range_1m2_km": 625.72,
    "sm2_pk": 0.74,
    "sm2_range_km": 170,
    "sm3_pk": 0.62,
    "sm3_range_km": 500,
    "sm6_pk": 0.75,
    "sm6_range_km": 370
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `array_faces` | int | Number of SPY-1D array faces |
| `peak_power_kw` | int | Peak power per face (kW) |
| `radar` | string | Radar designation |
| `range_01m2_km` | float | Detection range vs 0.01 m² RCS |
| `range_1m2_km` | float | Detection range vs 1 m² RCS |
| `sm2_pk` | float | SM-2 kill probability |
| `sm2_range_km` | int | SM-2 max range (km) |
| `sm3_pk` | float | SM-3 kill probability |
| `sm3_range_km` | int | SM-3 max range (km) |
| `sm6_pk` | float | SM-6 kill probability |
| `sm6_range_km` | int | SM-6 max range (km) |

---

### bmd-sim-atmospheric

**Description**: Atmospheric Effects Simulator (refraction, attenuation, ducting)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "atmospheric",
  "status": "complete",
  "parameters": {
    "humidity": 50,
    "rain_mmhr": 0,
    "refractivity": 313,
    "wind_ms": 5
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `humidity` | int | Relative humidity (%) |
| `rain_mmhr` | int | Rain rate (mm/hr) |
| `refractivity` | int | Atmospheric refractivity (N-units) |
| `wind_ms` | int | Wind speed (m/s) |

---

### bmd-sim-cobra-judy

**Description**: Cobra Judy (AN/SPQ-11) Radar Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "cobra-judy",
  "status": "complete",
  "parameters": {
    "gain_dbi": 48,
    "max_range_1m2_km": 1860.08
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `gain_dbi` | int | Antenna gain (dBi) |
| `max_range_1m2_km` | float | Max range vs 1 m² RCS (km) |

---

### bmd-sim-dsp

**Description**: Defense Support Program Satellite Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "dsp",
  "status": "complete",
  "parameters": {
    "aperture_m": 0.9,
    "max_range_boost_km": 4180343.14,
    "nefd": 2e-14,
    "satellites": 4,
    "spin_rpm": 6
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `aperture_m` | float | Sensor aperture (m) |
| `max_range_boost_km` | float | Max range vs boost plume (km) |
| `nefd` | float | Noise Equivalent Flux Density (W/m²) |
| `satellites` | int | Number of DSP satellites |
| `spin_rpm` | int | Sensor spin rate (RPM) |

---

### bmd-sim-gbr

**Description**: Ground-Based Radar (X-band) Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "gbr",
  "status": "complete",
  "parameters": {
    "gain_dbi": 50,
    "max_range_0p01_km": 327.87
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `gain_dbi` | int | Antenna gain (dBi) |
| `max_range_0p01_km` | float | Max range vs 0.01 m² RCS (km) |

---

### bmd-sim-lrdr

**Description**: Long Range Discrimination Radar (S-band, 5000+ km)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "lrdr",
  "status": "complete",
  "parameters": {
    "frequency_ghz": 3,
    "power_kw": 280,
    "range_rv_km": 4870.63,
    "track_capacity": 1000
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `frequency_ghz` | int | Operating frequency (GHz) |
| `power_kw` | int | Peak power (kW) |
| `range_rv_km` | float | Discrimination range vs RV (km) |
| `track_capacity` | int | Simultaneous track capacity |

---

### bmd-sim-sbirs

**Description**: Space-Based Infrared System Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "sbirs",
  "status": "complete",
  "parameters": {
    "aperture_m": 1.2,
    "health": 1,
    "max_range_body_km": 42295.93,
    "max_range_boost_km": 45000,
    "nefd": 5e-15,
    "satellites": 6
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `aperture_m` | float | Sensor aperture (m) |
| `health` | float | Constellation health (0-1) |
| `max_range_body_km` | float | Max range vs cold body (km) |
| `max_range_boost_km` | int | Max range vs boost plume (km) |
| `nefd` | float | Noise Equivalent Flux Density (W/m²) |
| `satellites` | int | Number of SBIRS satellites |

---

### bmd-sim-stss

**Description**: Space Tracking & Surveillance System Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "stss",
  "status": "complete",
  "parameters": {
    "aperture_m": 0.6,
    "max_range_boost_km": 45000,
    "nefd": 1e-14,
    "period_s": 6751.83,
    "satellites": 2
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `aperture_m` | float | Sensor aperture (m) |
| `max_range_boost_km` | int | Max range vs boost plume (km) |
| `nefd` | float | Noise Equivalent Flux Density (W/m²) |
| `period_s` | float | Orbital period (s) |
| `satellites` | int | Number of STSS satellites |

---

### bmd-sim-tpy2

**Description**: AN/TPY-2 X-band Transportable Radar Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-site uint`, `-scenario string`, `-help`

**JSON Output**:
```json
{
  "simulator": "tpy2",
  "status": "complete",
  "parameters": {
    "frequency_ghz": 9.5,
    "mode": 0,
    "power_kw": 100,
    "search_range_km": 6303.55,
    "site_id": 1
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `frequency_ghz` | float | Operating frequency (GHz) |
| `mode` | int | Radar mode (0=FBM, 1=TM) |
| `power_kw` | int | Peak power (kW) |
| `search_range_km` | float | Search volume range (km) |
| `site_id` | int | Site identifier |

---

### bmd-sim-uewr

**Description**: Upgraded Early Warning Radar (UHF phased array, 4-site network)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "uewr",
  "status": "complete",
  "parameters": {
    "frequency_mhz": 435,
    "power_kw": 150,
    "sites": 4,
    "track_capacity": 500
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `frequency_mhz` | int | Operating frequency (MHz) |
| `power_kw` | int | Peak power (kW) |
| `sites` | int | Number of UEWR sites |
| `track_capacity` | int | Simultaneous track capacity |

---

## BMDS Threats

### bmd-sim-decoy

**Description**: Decoy/Penetration Aid Simulator (balloons, inflatable RVs, RCS enhancers, jammers)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "decoy",
  "status": "complete",
  "parameters": {
    "active_count": 2,
    "rcs": 0.1,
    "type": "balloon"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `active_count` | int | Number of active decoys |
| `rcs` | float | Decoy RCS (m²) |
| `type` | string | Decoy type (balloon, inflatable_rv, rcs_enhancer, jammer) |

---

### bmd-sim-electronic-attack

**Description**: Electronic Attack (ECM) Simulator — jamming, DRFM, RGPO/VGPO, chaff vs BMDS radars

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-ecm string`, `-eccm string`, `-radar string`, `-power int`, `-range float`, `-help`

**JSON Output**:
```json
{
  "simulator": "electronic-attack",
  "status": "complete",
  "parameters": {
    "burnthrough_km": 1.15,
    "detection_range_km": 1.15,
    "eccm_net_gain_db": 0,
    "eccm_range_recovery_pct": 0,
    "eccm_suppression_db": 0,
    "ecm_power_w": 10000,
    "ecm_type": "NOISE_BARRAGE",
    "jammer_effective": false,
    "jsr_db": -125.44,
    "radar": "AN/TPY-2",
    "radar_only_range_km": 401.15
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `burnthrough_km` | float | Radar burnthrough range (km) |
| `detection_range_km` | float | Detection range under jamming (km) |
| `eccm_net_gain_db` | float | Net ECCM gain (dB) |
| `eccm_range_recovery_pct` | float | Range recovery from ECCM (%) |
| `eccm_suppression_db` | float | ECCM suppression of jamming (dB) |
| `ecm_power_w` | int | Jammer power (W) |
| `ecm_type` | string | ECM type (NOISE_BARRAGE, NOISE_SPOT, DRFM, RGPO, VGPO, CROSSPOL, CHAFF) |
| `jammer_effective` | bool | Whether jamming denies radar |
| `jsr_db` | float | J/S ratio (dB) |
| `radar` | string | Target radar designation |
| `radar_only_range_km` | float | Unjammed radar range (km) |

---

### bmd-sim-hgv

**Description**: Hypersonic Glide Vehicle Threat Simulator (DF-ZF/DF-17)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "hgv",
  "status": "complete",
  "parameters": {
    "altitude_km": 80,
    "category": "HGV",
    "cep_m": 10,
    "flight_time_s": 576,
    "glide_speed_ms": 5000,
    "ir_boost_wsr": 80000000,
    "ir_glide_wsr": 50000,
    "lift_drag_ratio": 3.5,
    "range_km": 1800,
    "rcs": 0.001,
    "terminal_speed_ms": 3500,
    "variant": "DF-ZF (DF-17)"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `altitude_km` | int | Deployment altitude (km) |
| `category` | string | Threat category |
| `cep_m` | int | Circular Error Probable (m) |
| `flight_time_s` | int | Total flight time (s) |
| `glide_speed_ms` | int | Glide speed (m/s) |
| `ir_boost_wsr` | int | IR signature boost phase (W/sr) |
| `ir_glide_wsr` | int | IR signature glide phase (W/sr) |
| `lift_drag_ratio` | float | L/D ratio |
| `range_km` | int | Maximum range (km) |
| `rcs` | float | Radar cross section (m²) |
| `terminal_speed_ms` | int | Terminal speed (m/s) |
| `variant` | string | Designation |

---

### bmd-sim-icbm

**Description**: ICBM Threat Simulator (Minuteman-III class)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "icbm",
  "status": "complete",
  "parameters": {
    "apogee_km": 1370,
    "burnout_ms": 4901.54,
    "cep_m": 120,
    "flight_time_s": 2400,
    "ir_boost_wsr": 120000000,
    "mirv_count": 1,
    "range_km": 13000,
    "reentry_ms": 4509.42,
    "rv_rcs": 0.05,
    "stages": 3,
    "variant": "Minuteman-III"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `apogee_km` | int | Trajectory apogee (km) |
| `burnout_ms` | float | Burnout speed (m/s) |
| `cep_m` | int | Circular Error Probable (m) |
| `flight_time_s` | int | Total flight time (s) |
| `ir_boost_wsr` | int | Boost IR signature (W/sr) |
| `mirv_count` | int | Number of MIRVs |
| `range_km` | int | Maximum range (km) |
| `reentry_ms` | float | Reentry speed (m/s) |
| `rv_rcs` | float | Reentry vehicle RCS (m²) |
| `stages` | int | Number of stages |
| `variant` | string | Designation |

---

### bmd-sim-irbm

**Description**: IRBM Threat Simulator (DF-26 class)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "irbm",
  "status": "complete",
  "parameters": {
    "apogee_km": 580,
    "burnout_ms": 3726.87,
    "category": "IRBM",
    "cep_m": 150,
    "flight_time_s": 1500,
    "ir_boost_wsr": 70000000,
    "mobile": true,
    "range_km": 4000,
    "reentry_ms": 3354.18,
    "rv_rcs": 0.1,
    "stages": 2,
    "variant": "DF-26"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `apogee_km` | int | Trajectory apogee (km) |
| `burnout_ms` | float | Burnout speed (m/s) |
| `category` | string | Threat category |
| `cep_m` | int | CEP (m) |
| `flight_time_s` | int | Flight time (s) |
| `ir_boost_wsr` | int | Boost IR signature (W/sr) |
| `mobile` | bool | TEL-launched |
| `range_km` | int | Range (km) |
| `reentry_ms` | float | Reentry speed (m/s) |
| `rv_rcs` | float | RV RCS (m²) |
| `stages` | int | Number of stages |
| `variant` | string | Designation |

---

### bmd-sim-jamming

**Description**: EW/Jamming Simulator — noise/deception jamming with JNR, burnthrough, ECCM

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "jamming",
  "status": "complete",
  "parameters": {
    "effect": "DENIED",
    "freq_ghz": 3,
    "jnr_db": 109.01,
    "power_w": 1000000
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `effect` | string | Jamming effect (DENIED, DEGRADED, NONE) |
| `freq_ghz` | int | Jamming frequency (GHz) |
| `jnr_db` | float | Jammer-to-Noise ratio (dB) |
| `power_w` | int | Jammer power (W) |

---

### bmd-sim-mrbm

**Description**: MRBM Threat Simulator (DF-21A class, TEL-capable)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-mobile`, `-cms`, `-help`

**JSON Output**:
```json
{
  "simulator": "mrbm",
  "status": "complete",
  "parameters": {
    "apogee_km": 252.0,
    "burnout_ms": 5196.72,
    "category": "MRBM",
    "cep_m": 300,
    "flight_time_s": 735.97,
    "ir_boost_wsr": 8000000,
    "mobile": true,
    "range_km": 1800,
    "rcs": 0.25,
    "reentry_ms": 3741.64,
    "stages": 2,
    "variant": "DF-21A",
    "warhead_kg": 500
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `apogee_km` | float | Trajectory apogee (km) |
| `burnout_ms` | float | Burnout speed (m/s) |
| `category` | string | Threat category |
| `cep_m` | int | CEP (m) |
| `flight_time_s` | float | Flight time (s) |
| `ir_boost_wsr` | int | Boost IR signature (W/sr) |
| `mobile` | bool | TEL-launched |
| `range_km` | int | Range (km) |
| `rcs` | float | RCS (m²) |
| `reentry_ms` | float | Reentry speed (m/s) |
| `stages` | int | Number of stages |
| `variant` | string | Designation |
| `warhead_kg` | int | Warhead mass (kg) |

---

### bmd-sim-nuclear-efx

**Description**: Nuclear Detonation Effects Simulator (EMP, radar degradation, satellite kill, ionospheric blackout)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "nuclear-efx",
  "status": "complete",
  "parameters": {
    "altitude_km": 50,
    "blast_psi": 0.32,
    "burst_type": "HIGH_ALTITUDE",
    "emp_damage": false,
    "emp_downtime_s": 0.1,
    "emp_peak_vm": 20000,
    "fallout_rad_hr": 0.01,
    "ground_range_km": 300,
    "hf_blackout": true,
    "iono_s4": 0.17,
    "iono_tec": 16.67,
    "prompt_dose_rad": 0.01,
    "radar_range_pct": 73.1,
    "radar_recovery_s": 2400,
    "radar_snr_loss_db": 5.44,
    "sat_blackout_s": 600,
    "sat_fluence_jm2": 0.15,
    "sat_kill": false,
    "sat_kill_radius_km": 152.67,
    "sat_seu": true,
    "thermal_cal_cm2": 2.65e-10,
    "vhf_blackout": true,
    "yield_kt": 1000
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `altitude_km` | int | Burst altitude (km) |
| `blast_psi` | float | Overpressure at ground range (psi) |
| `burst_type` | string | HIGH_ALTITUDE, SURFACE, AIR |
| `emp_damage` | bool | Whether EMP causes permanent damage |
| `emp_downtime_s` | float | EMP-induced downtime (s) |
| `emp_peak_vm` | int | Peak EMP field (V/m) |
| `fallout_rad_hr` | float | Fallout dose rate (rad/hr) |
| `ground_range_km` | int | Ground range to assessment point (km) |
| `hf_blackout` | bool | HF communications blackout |
| `iono_s4` | float | Ionospheric scintillation index |
| `iono_tec` | float | Total Electron Content (TECU) |
| `prompt_dose_rad` | float | Prompt radiation dose (rad) |
| `radar_range_pct` | float | Radar range as % of nominal |
| `radar_recovery_s` | int | Radar recovery time (s) |
| `radar_snr_loss_db` | float | SNR loss (dB) |
| `sat_blackout_s` | int | Satellite blackout duration (s) |
| `sat_fluence_jm2` | float | Satellite radiation fluence (J/m²) |
| `sat_kill` | bool | Whether satellite is killed |
| `sat_kill_radius_km` | float | Satellite kill radius (km) |
| `sat_seu` | bool | Single-event upsets expected |
| `thermal_cal_cm2` | float | Thermal flux (cal/cm²) |
| `vhf_blackout` | bool | VHF communications blackout |
| `yield_kt` | int | Weapon yield (kT) |

---

### bmd-sim-slcm

**Description**: Submarine-Launched Cruise Missile Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "slcm",
  "status": "complete",
  "parameters": {
    "altitude_m": 30,
    "cruise_ms": 270,
    "range_km": 1500,
    "rcs_m2": 0.01,
    "variant": 0
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `altitude_m` | int | Cruise altitude (m) |
| `cruise_ms` | int | Cruise speed (m/s) |
| `range_km` | int | Range (km) |
| `rcs_m2` | float | RCS (m²) |
| `variant` | int | Variant identifier |

---

### ufo

**Description**: Unknown Anomalous Phenomena Threat Simulator — physics-violating detection/engagement

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-anomaly-score float`, `-ecm-level int`, `-help`

**JSON Output**:
```json
{
  "simulator": "ufo",
  "status": "complete",
  "parameters": {
    "anomaly_score": 0.5,
    "ecm_level": 3,
    "maneuver_count": 8,
    "resurgence": false
  },
  "data": {
    "anomaly_scores": {
      "craft": 0.5,
      "detection_pd": 0.83,
      "track_conflict": false,
      "GBR-P": { "detected": true, "detection_prob": 0.83, "effective_range": 1616, "failure_reason": "" },
      "LRDR": { "detected": true, "detection_prob": 0.83, "effective_range": 2424, "failure_reason": "" },
      "TPY-2": { "detected": true, "detection_prob": 0.83, "effective_range": 808, "failure_reason": "" },
      "UEWR": { "detected": true, "detection_prob": 0.83, "effective_range": 4040, "failure_reason": "" }
    },
    "detection_failures": [
      { "sensor": "UEWR", "detected": true, "detection_prob": 0.83, "effective_range": 4040, "failure_reason": "" }
    ],
    "engagement_results": [
      { "weapon": "GMD", "base_pk": 0.55, "adjusted_pk": 0.07, "anomaly_score": 0.5, "kill_uncertain": false, "maneuver_count": 8, "failure_reason": "Trajectory discontinuities prevent midcourse intercept prediction" }
    ]
  }
}
```

**Key Fields**:

| Field | Type | Description |
|-------|------|-------------|
| `anomaly_score` | float | Physics anomaly threshold (0-1) |
| `ecm_level` | int | ECM sophistication (1-5) |
| `maneuver_count` | int | Number of physics-violating maneuvers |
| `resurgence` | bool | Whether craft reappears after engagement |
| `anomaly_scores.craft` | float | Overall craft anomaly score |
| `anomaly_scores.detection_pd` | float | Detection probability |
| `anomaly_scores.track_conflict` | bool | Track data conflicts with known physics |
| `anomaly_scores.<sensor>` | object | Per-sensor detection details |
| `engagement_results[].adjusted_pk` | float | PK adjusted for anomaly effects |
| `engagement_results[].failure_reason` | string | Reason for engagement failure |

---

## BMDS Interceptors

### bmd-sim-gmd

**Description**: Ground-Based Midcourse Defense (CE-II GBI) Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "gmd",
  "status": "complete",
  "parameters": {
    "available_silos": 44,
    "burnout_ms": 5810.74,
    "ekv_divert_ms": 800,
    "ekv_lethal_m": 0.5,
    "ekv_seeker_km": 700,
    "max_alt_km": 1700,
    "pk_hgv": 0.029,
    "pk_icbm": 0.535,
    "pk_irbm": 0.543,
    "pk_slcm": 0,
    "site": "FTG",
    "variant": "CE-II GBI @ FTG"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `available_silos` | int | Number of available GBI silos |
| `burnout_ms` | float | GBI burnout speed (m/s) |
| `ekv_divert_ms` | int | EKV divert capability (m/s) |
| `ekv_lethal_m` | float | EKV lethal radius (m) |
| `ekv_seeker_km` | int | EKV seeker range (km) |
| `max_alt_km` | int | Maximum altitude (km) |
| `pk_hgv` | float | PK vs HGV |
| `pk_icbm` | float | PK vs ICBM |
| `pk_irbm` | float | PK vs IRBM |
| `pk_slcm` | float | PK vs SLCM |
| `site` | string | Deployment site |
| `variant` | string | GBI variant designation |

---

### bmd-sim-patriot

**Description**: Patriot Battery Simulator (PAC-2/PAC-3)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "patriot",
  "status": "complete",
  "parameters": {
    "launchers": 1,
    "range_01m2_km": 139.62,
    "range_1m2_km": 248.28,
    "total_missiles": 8
  },
  "data": {
    "pac2_pk": 0.7,
    "pac2_range_km": 100,
    "pac3_pk": 0.75,
    "pac3_range_km": 150
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `launchers` | int | Number of launchers |
| `range_01m2_km` | float | Detection range vs 0.1 m² (km) |
| `range_1m2_km` | float | Detection range vs 1 m² (km) |
| `total_missiles` | int | Total missile inventory |
| `pac2_pk` | float | PAC-2 kill probability |
| `pac2_range_km` | int | PAC-2 range (km) |
| `pac3_pk` | float | PAC-3 kill probability |
| `pac3_range_km` | int | PAC-3 range (km) |

---

### bmd-sim-sm3

**Description**: SM-3 Block IIA Interceptor Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "sm3",
  "status": "complete",
  "parameters": {
    "burnout_ms": 6126.46,
    "leap_divert_ms": 600,
    "leap_mass_kg": 18,
    "leap_seeker_km": 500,
    "max_alt_km": 1500,
    "pk_hgv": 0.044,
    "pk_icbm": 0.314,
    "pk_irbm": 0.610,
    "range_km": 2500,
    "variant": "SM-3 Block IIA"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `burnout_ms` | float | Burnout speed (m/s) |
| `leap_divert_ms` | int | LEAP KV divert (m/s) |
| `leap_mass_kg` | int | LEAP KV mass (kg) |
| `leap_seeker_km` | int | LEAP seeker range (km) |
| `max_alt_km` | int | Maximum altitude (km) |
| `pk_hgv` | float | PK vs HGV |
| `pk_icbm` | float | PK vs ICBM |
| `pk_irbm` | float | PK vs IRBM |
| `range_km` | int | Maximum range (km) |
| `variant` | string | Variant designation |

---

### bmd-sim-sm6

**Description**: SM-6 Dual Missile Simulator (SAM/SSM)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "sm6",
  "status": "complete",
  "parameters": {
    "burnout_ms": 1500,
    "cells": 8,
    "lethal_radius_m": 5,
    "loaded": 8,
    "max_alt_km": 33,
    "pk_aircraft": 0.887,
    "pk_hgv": 0.019,
    "pk_slcm": 0.838,
    "range_km": 370,
    "seeker_km": 50,
    "variant": "SM-6 Baseline"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `burnout_ms` | int | Burnout speed (m/s) |
| `cells` | int | VLS cells available |
| `lethal_radius_m` | int | Blast-frag lethal radius (m) |
| `loaded` | int | VLS cells loaded |
| `max_alt_km` | int | Maximum altitude (km) |
| `pk_aircraft` | float | PK vs aircraft |
| `pk_hgv` | float | PK vs HGV |
| `pk_slcm` | float | PK vs SLCM |
| `range_km` | int | Maximum range (km) |
| `seeker_km` | int | Active seeker range (km) |
| `variant` | string | Variant designation |

---

### bmd-sim-thaad

**Description**: THAAD Battery Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "thaad",
  "status": "complete",
  "parameters": {
    "launchers": 2,
    "range_01m2_km": 335.75,
    "range_1m2_km": 597.05
  },
  "data": {
    "missile_pk": 0.9,
    "missile_range_km": 300,
    "missile_velocity_ms": 6000
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `launchers` | int | Number of launchers |
| `range_01m2_km` | float | Detection range vs 0.1 m² (km) |
| `range_1m2_km` | float | Detection range vs 1 m² (km) |
| `missile_pk` | float | THAAD missile PK |
| `missile_range_km` | int | Missile range (km) |
| `missile_velocity_ms` | int | Missile velocity (m/s) |

---

### bmd-sim-thaad-er

**Description**: THAAD-ER (Extended Range) Interceptor Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "thaad-er",
  "status": "complete",
  "parameters": {
    "burnout_ms": 4555.68,
    "divert_ms": 600,
    "max_alt_km": 200,
    "min_alt_km": 40,
    "pk_hgv": 0.076,
    "pk_mrbm": 0.76,
    "range_km": 600,
    "seeker_km": 300
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `burnout_ms` | float | Burnout speed (m/s) |
| `divert_ms` | int | KKVI divert capability (m/s) |
| `max_alt_km` | int | Maximum altitude (km) |
| `min_alt_km` | int | Minimum altitude (km) |
| `pk_hgv` | float | PK vs HGV |
| `pk_mrbm` | float | PK vs MRBM |
| `range_km` | int | Maximum range (km) |
| `seeker_km` | int | KKVI seeker range (km) |

---

## C2/Infrastructure

### bmd-sim-c2bmc

**Description**: C2BMC (Command & Control Battle Management) Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "c2bmc",
  "status": "complete",
  "parameters": {
    "correlated": 0,
    "threats": 0,
    "tracks": 0
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `correlated` | int | Correlated tracks |
| `threats` | int | Active threats |
| `tracks` | int | Total tracks |

---

### bmd-sim-gfcb

**Description**: GFCB (Global Fire Control & Battle) Simulator

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "gfcb",
  "status": "complete",
  "parameters": {
    "node_id": "GFCB-1",
    "region": "PACIFIC"
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `node_id` | string | GFCB node identifier |
| `region` | string | Operational region |

---

### bmd-sim-hub

**Description**: C2BMC Integration Hub (track fusion and engagement routing)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "hub",
  "status": "complete",
  "parameters": {
    "fused_tracks": 1,
    "routes": 3
  },
  "data": {
    "engagement_id": 1,
    "engagement_phase": "0",
    "pk": 0.85,
    "tti": 0.28
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `fused_tracks` | int | Number of fused tracks |
| `routes` | int | Number of configured routes |
| `engagement_id` | int | Current engagement ID |
| `engagement_phase` | string | Engagement phase |
| `pk` | float | Engagement kill probability |
| `tti` | float | Time to intercept (s) |

---

### bmd-sim-ifxb

**Description**: IFXB (Integrated Fire Exchange) Simulator — data routing between BMDS elements

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "ifxb",
  "status": "complete",
  "parameters": {
    "assets": 0,
    "node_id": "IFXB-1",
    "routes": 0
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `assets` | int | Connected assets |
| `node_id` | string | IFXB node identifier |
| `routes` | int | Configured routes |

---

### bmd-sim-jreap

**Description**: JREAP-C Transport Simulator (UDP/TCP message transport)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "jreap",
  "status": "complete",
  "parameters": {
    "port": 4000,
    "protocol": 0
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `port` | int | JREAP-C port |
| `protocol` | int | Protocol (0=UDP, 1=TCP) |

---

### bmd-sim-jrsc

**Description**: JRSC (Joint Regional Security Center) Simulator — regional sensor fusion

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "jrsc",
  "status": "complete",
  "parameters": {
    "alert": 3,
    "region": "JRSC-PAC",
    "sensors": 0
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `alert` | int | Alert level |
| `region` | string | Regional designation |
| `sensors` | int | Connected sensors |

---

### bmd-sim-link16

**Description**: Link 16 Tactical Data Link Simulator (TDMA, J-series messages)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-v`, `-help`

**JSON Output**:
```json
{
  "simulator": "link16",
  "status": "complete",
  "parameters": {
    "net_id": "N-1",
    "participants": 0,
    "time_slots": 1535
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `net_id` | string | Network identifier |
| `participants` | int | Active participants |
| `time_slots` | int | TDMA time slots per frame |

---

### engagement-chain

**Description**: BMDS Engagement Chain Simulator — full kill chain from boost detection through intercept

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "engagement-chain",
  "status": "complete",
  "parameters": {
    "detections": [
      { "sensor": "SBIRS-GEO-1", "type": "SBIRS", "phase": "BOOST", "range_km": 11940, "alt_km": 140, "snr": 50, "confidence": 0.95, "time_s": 15 },
      { "sensor": "UEWR-Beale", "type": "UEWR", "phase": "POST-BOOST", "range_km": 10615, "alt_km": 1393, "snr": 20, "confidence": 0.85, "time_s": 315 },
      { "sensor": "TPY-2-Guam", "type": "TPY-2", "phase": "POST-BOOST", "range_km": 10429, "alt_km": 1386, "snr": 15, "confidence": 0.8, "time_s": 330 }
    ],
    "doctrine_latency_s": 8,
    "doctrine_max_int": 2,
    "doctrine_sls": true,
    "engagements": [
      { "interceptor": "GMD-GBI", "launch_time_s": 340, "intercept_time_s": 411.19, "intercept_alt_km": 1346.59, "pk": 0.318, "killed": true },
      { "interceptor": "GMD-GBI", "launch_time_s": 419.19, "intercept_time_s": 493.31, "intercept_alt_km": 1307.13, "pk": 0.222, "killed": false }
    ],
    "killed": true,
    "overall_pk": 0.469,
    "threat_apogee_km": 1400,
    "threat_burnout_ms": 7200,
    "threat_class": "ICBM",
    "threat_range_km": 12000,
    "threat_rcs_m2": 0.05,
    "timeline": [
      { "event": "BOOST-DETECT", "phase": "BOOST", "sensor": "SBIRS-GEO-1", "detail": "SBIRS-GEO-1 detected boost plume", "confidence": 0.95, "time_s": 15 },
      { "event": "TRACK-INIT", "phase": "POST-BOOST", "sensor": "UEWR-Beale", "detail": "UEWR established track", "confidence": 0.85, "time_s": 315 }
    ],
    "total_time_s": 1800
  }
}
```

**Key Fields**:

| Field | Type | Description |
|-------|------|-------------|
| `detections[]` | array | Sensor detection events |
| `detections[].sensor` | string | Detecting sensor |
| `detections[].type` | string | Sensor type |
| `detections[].phase` | string | Threat phase at detection |
| `detections[].confidence` | float | Detection confidence (0-1) |
| `doctrine_sls` | bool | Shoot-Look-Shoot doctrine |
| `doctrine_max_int` | int | Max interceptors per threat |
| `doctrine_latency_s` | int | C2 decision latency (s) |
| `engagements[]` | array | Interceptor engagement events |
| `engagements[].interceptor` | string | Interceptor type |
| `engagements[].pk` | float | Single-shot PK |
| `engagements[].killed` | bool | Whether threat was killed |
| `killed` | bool | Overall kill result |
| `overall_pk` | float | Cumulative PK |
| `timeline[]` | array | Chronological event log |

---

### kill-assessment

**Description**: Kill Assessment Simulator — BMDS post-intercept damage assessment

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-interceptor string`, `-warhead string`, `-scenario string`, `-closing-velocity float`, `-help`

**JSON Output**:
```json
{
  "simulator": "kill-assessment",
  "status": "complete",
  "parameters": {
    "closing_velocity": 8000,
    "interceptor": "EKV",
    "miss_distance": 0,
    "scenario": "single-icbm-ekv",
    "seed": 0,
    "sensors": ["X_BAND_RADAR", "IR_SPACE", "UEWR"],
    "warhead_type": "nuclear_fission"
  },
  "intercepts": [
    {
      "threat_id": "T-1",
      "interceptor": "EKV",
      "kv_type": "HIT_TO_KILL",
      "miss_distance_m": 0,
      "hit_zone": "fuze",
      "warhead_damage": 0.9,
      "structural_damage": 1,
      "fuze_damage": 1,
      "guidance_damage": 0.8,
      "is_killed": true,
      "is_disrupted": false,
      "closing_velocity_ms": 8000
    }
  ],
  "assessments": [
    {
      "threat_id": "T-1",
      "decision": "KILL_PROBABLE",
      "kill_confidence": 0.731,
      "miss_confidence": 0.038,
      "disruption_confidence": 0.231,
      "time_to_decision_s": 7.5,
      "debris_fragments": 153,
      "track_status": "terminated",
      "recommended_action": "observe",
      "sensors_used": ["TPY-2/GBR", "SBIRS/STSS", "UEWR"]
    }
  ],
  "statistics": {
    "total_intercepts": 1,
    "kills": 1,
    "misses": 0,
    "unknowns": 0,
    "avg_assessment_time_s": 7.5,
    "avg_confidence": 0.731
  }
}
```

**Key Sections**:

| Section | Description |
|---------|-------------|
| `parameters` | Scenario configuration |
| `intercepts[]` | Per-intercept damage assessment |
| `intercepts[].hit_zone` | Impact zone: fuze, guidance, structural, warhead |
| `intercepts[].warhead_damage` | Warhead damage fraction (0-1) |
| `intercepts[].is_killed` | Kill determination |
| `assessments[]` | Post-intercept assessment |
| `assessments[].decision` | KILL_PROBABLE, KILL_CERTAIN, MISS, UNKNOWN |
| `assessments[].kill_confidence` | Confidence in kill (0-1) |
| `assessments[].recommended_action` | observe, reengage, escalate |
| `statistics` | Aggregate statistics |

---

### tactical-net

**Description**: Tactical Network Simulator — BMDS RF propagation, mesh networking, and EW effects

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-scenario string`, `-freq string`, `-duct string`, `-jammers int`, `-rain float`, `-terrain string`, `-help`

**JSON Output** (non-JSON preamble present):
```json
{
  "simulator": "tactical-net",
  "status": "complete",
  "parameters": {
    "duct": "none",
    "duration_s": 1,
    "freq_band": "mixed",
    "jammers": 0,
    "nodes": 38,
    "rain_mm_hr": 0,
    "scenario": "single-launch",
    "seed": 0,
    "terrain": "flat"
  },
  "links": [
    {
      "From": "GBR",
      "To": "SBIRS-GEO-2",
      "Radio": "LINK16",
      "MarginDB": -92.16,
      "ThroughputKbps": 500,
      "LatencyMs": 38132.66,
      "PacketLoss": 0.5,
      "Status": "down",
      "Quality": 0
    }
  ],
  "statistics": {
    "ActiveLinks": 76,
    "DegradedLinks": 22,
    "DownLinks": 2486,
    "AvgLatencyMs": 18064.56,
    "MaxLatencyMs": 57243.50,
    "TotalThroughputKbps": 2695718
  }
}
```

**Key Fields**:

| Field | Type | Description |
|-------|------|-------------|
| `parameters.nodes` | int | Network nodes |
| `parameters.jammers` | int | Active jammers |
| `parameters.freq_band` | string | uhf, sband, xband, mixed |
| `parameters.duct` | string | none, evaporation, surface, elevated |
| `links[]` | array | Per-link RF analysis |
| `links[].MarginDB` | float | Link margin (dB) |
| `links[].Status` | string | up, degraded, down |
| `links[].Quality` | float | Link quality (0-1) |
| `statistics` | object | Aggregate network stats |

---

### wta

**Description**: Weapon-Target Assignment Simulator — BMDS interceptor allocation optimization

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-doctrine string`, `-roe-level int`, `-cost-limit float`, `-discrimination float`, `-scenario string`, `-help`

**JSON Output**:
```json
{
  "simulator": "wta",
  "status": "complete",
  "parameters": {
    "doctrine": "SHOOT-LOOK-SHOOT",
    "roe_level": 3,
    "scenario": "single-icbm",
    "seed": 0,
    "sites": 1,
    "threats": 1
  },
  "assignments": [
    { "time": 0, "threat_id": "T-1", "threat_type": "", "site_id": "FORT_GREELY", "interceptor": "GMD_GBI", "pk": 0.55, "launch_time": 0, "intercept_time": 0.21 }
  ],
  "kill_results": [
    { "time": 0.21, "threat_id": "T-1", "interceptor": "GMD_GBI", "result": "MISS", "time_from_launch": 0.21 }
  ],
  "leaked_threats": [
    { "threat_id": "T-1", "threat_type": "ICBM", "target_asset": "SEA", "impact_time": 1800 }
  ],
  "statistics": {
    "total_threats": 1,
    "threats_engaged": 1,
    "threats_killed": 0,
    "threats_leaked": 1,
    "interceptors_fired": 1,
    "total_cost_millions": 75,
    "overall_pk": 0
  }
}
```

**Key Fields**:

| Field | Type | Description |
|-------|------|-------------|
| `doctrine` | string | shoot-shoot, shoot-look-shoot, adaptive |
| `roe_level` | int | 0=peacetime, 5=autonomous |
| `assignments[]` | array | Weapon-target pairings |
| `assignments[].pk` | float | Assigned PK |
| `kill_results[]` | array | Intercept outcomes |
| `kill_results[].result` | string | KILL, MISS, PENDING |
| `leaked_threats[]` | array | Threats not intercepted |
| `statistics.overall_pk` | float | Overall engagement PK |
| `statistics.total_cost_millions` | float | Total interceptor cost |

---

## Support Sims

### air-traffic

**Description**: Air Traffic Simulator — commercial/military flights across 40 major hubs

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-count int`, `-help`

**JSON Output** (non-JSON preamble: "Generating flights..."):
```json
{
  "flights": [
    {
      "id": "FL0001",
      "callsign": "EG238",
      "type": "MILITARY",
      "aircraft": "C17",
      "origin": "EGKK",
      "dest": "LEMD",
      "lat": 51.14,
      "lon": -0.19,
      "alt": 416.02,
      "speed": 450,
      "heading": 193.72,
      "phase": "CLIMB",
      "progress": 0.0015,
      "route_index": 0
    }
  ],
  "summary": {
    "total": 5000,
    "CLIMB": 825,
    "CRUISE": 3422,
    "DESCENT": 512
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `flights[].id` | string | Flight identifier |
| `flights[].callsign` | string | ATC callsign |
| `flights[].type` | string | COMMERCIAL, CARGO, MILITARY |
| `flights[].aircraft` | string | Aircraft type code |
| `flights[].origin` | string | ICAO origin |
| `flights[].dest` | string | ICAO destination |
| `flights[].lat/lon` | float | Current position |
| `flights[].alt` | float | Altitude (m) |
| `flights[].speed` | int | Speed (kts) |
| `flights[].heading` | float | Heading (deg) |
| `flights[].phase` | string | CLIMB, CRUISE, DESCENT |
| `flights[].progress` | float | Route progress (0-1) |
| `summary.total` | int | Total flights |
| `summary.CLIMB/CRUISE/DESCENT` | int | Per-phase count |

---

### bmd-sim-space-weather

**Description**: Space Weather Simulator (Kp index, TEC, ionospheric scintillation)

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-help`

**JSON Output**:
```json
{
  "simulator": "space-weather",
  "status": "complete",
  "parameters": {
    "quiet_kp": 1,
    "quiet_tecu": 10,
    "storm_kp": 7,
    "storm_s4": 0.6,
    "storm_tecu": 80
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `quiet_kp` | int | Kp index (quiet conditions) |
| `quiet_tecu` | int | TEC (quiet, TECU) |
| `storm_kp` | int | Kp index (storm conditions) |
| `storm_s4` | float | S4 scintillation index (storm) |
| `storm_tecu` | int | TEC (storm, TECU) |

---

### satellite-tracker

**Description**: Satellite Tracker — live TLE data from Celestrak, SGP4/SDP4 propagation

**Flags**: `-json`, `-i`, `-duration`, `-group string`, `-count int`, `-help`

**JSON Output** (non-JSON preamble: "Fetching TLEs from Celestrak..."):
```json
{
  "parameters": {
    "group": "gps-ops",
    "time": "2026-04-21T14:13:47Z",
    "tle_count": 32
  },
  "satellites": [
    {
      "name": "GPS BIIR-2 (PRN 13)",
      "id": 24876,
      "lat": -0.005,
      "lon": 100.81,
      "alt": 20040.20,
      "vel": 3.90,
      "vis": "visible",
      "category": "GPS"
    }
  ],
  "simulator": "satellite-tracker",
  "status": "complete"
}
```

| Field | Type | Description |
|-------|------|-------------|
| `parameters.group` | string | Celestrak group name |
| `parameters.time` | string | ISO 8601 timestamp |
| `parameters.tle_count` | int | Number of TLEs loaded |
| `satellites[].name` | string | Satellite name |
| `satellites[].id` | int | NORAD catalog number |
| `satellites[].lat/lon` | float | Sub-satellite point |
| `satellites[].alt` | float | Altitude (km) |
| `satellites[].vel` | float | Velocity (km/s) |
| `satellites[].vis` | string | Visibility status |
| `satellites[].category` | string | Satellite category |

---

### space-debris

**Description**: Space Debris Simulator — orbital debris field with collision detection

**Flags**: `-json`, `-i`, `-duration`, `-seed`, `-count int`, `-collisions`, `-help`

**JSON Output** (non-JSON preamble: "Generating N debris objects..."):
```json
{
  "debris": [
    {
      "id": 1,
      "name": "ROCKET_BODY-1",
      "object_type": "ROCKET_BODY",
      "inclination": 38.64,
      "raan": 211.96,
      "eccentricity": 0.361,
      "arg_perigee": 38.93,
      "mean_anomaly": 210.72,
      "mean_motion": 5.884,
      "altitude": 6589.63,
      "perigee": 1908.31,
      "apogee": 11270.94,
      "category": "HIGH",
      "size": "LARGE",
      "rcs": 7.23,
      "lat": -30.60,
      "lon": 79.68
    }
  ],
  "summary": {
    "LEO": 15020,
    "MEO": 2516,
    "HIGH": 5020
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `debris[].object_type` | string | ROCKET_BODY, DEBRIS, PAYLOAD |
| `debris[].inclination` | float | Orbital inclination (deg) |
| `debris[].eccentricity` | float | Orbital eccentricity |
| `debris[].altitude` | float | Current altitude (km) |
| `debris[].category` | string | LEO, MEO, HIGH |
| `debris[].size` | string | SMALL, MEDIUM, LARGE |
| `debris[].rcs` | float | Radar cross section (m²) |
| `summary` | object | Count by orbital category |

---

### launch-veh-sim

**Description**: Launch Vehicle Simulator — orbital insertion trajectory

**Flags**: `-json`, `-duration float`, `-vehicle string`, `-target-alt float`, `-inclination float`, `-help`

**Note**: `-duration` takes float seconds, not Go duration format. No `-seed` or `-i` flag.

**JSON Output**: Outputs JSON when run with `-json` flag (produces empty output with default short duration — use `-duration 1000` for full simulation).

---

## Config-Driven Sims

These sims use YAML configuration files and export AAR (After Action Review) JSON rather than streaming `-json`.

### cyber-redteam-sim

**Description**: Enterprise network cyber attack/defense simulation

**Flags**: `-config string`, `-doctor`, `-init`, `-no-server`, `-rest-port int`, `-version`

**No `-json` flag** — use `-aar <path>` for JSON output (not yet implemented in CLI flags).

**Version**: 0.2.0 (commit 5bef09b)

---

### maritime-sim

**Description**: Naval surface/subsurface warfare scenarios

**Flags**: `-config string`, `-aar string`, `-web string`, `-doctor`, `-init`, `-list-scenarios`, `-validate string`, `-version`

**No `-json` flag** — use `-aar <path>` to write AAR JSON.

**Version**: 0.1.0 (commit babc3e4)

---

### space-war-sim

**Description**: Orbital warfare and ASAT scenarios

**Flags**: `-config string`, `-aar string`, `-doctor`, `-init`, `-list-scenarios`, `-validate string`, `-version`

**No `-json` flag** — use `-aar <path>` to write AAR JSON.

**Version**: 0.1.0 (commit ce2b52c)

---

## Engine Libraries

These are Go packages with no standalone CLI simulation mode. They print version/package info and exit.

### electronic-war-sim

**Description**: Electromagnetic warfare simulation engine

**Packages**: radar, ecm, esm, spectrum, core

**Flags**: None (library only)

**Version**: 0.1.0

---

### missile-defense-sim

**Description**: Ballistic missile defense simulation engine

**Packages**: ballistic, interceptor, seeker, track, countermeasures, layer, core

**Flags**: None (library only)

**Version**: 0.1.0

---

### submarine-war-sim

**Description**: Underwater warfare simulation engine

**Packages**: hydro, sonar, torpedo, asw, mine, submarine, core

**Flags**: None (library only, use `go test ./...` and `go run ./cmd/sim/`)

**Version**: 0.1.0