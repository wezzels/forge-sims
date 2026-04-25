# satellite-tracker — Orbital Satellite Tracker

## Purpose
Tracks satellites using NORAD/Celestrak TLE data with SGP4/SDP4 propagation.

## Physics & Fidelity
- SGP4/SDP4 orbital propagation models
- Live TLE fetch from Celestrak (daily refresh via sat-data-service)
- 12 satellite groups: active, stations, starlink, gps-ops, weather, geo, military, intelsat, resource, sarsat, argos, planet
- Keplerian element interpolation between TLE epochs
- Ground track and visibility calculation

### Scenarios
`gps-ops`, `starlink`, `active`, `stations`, `weather`, `geo`, `military`

## Accuracy: High
**Reliable**: SGP4/SDP4 propagation (standard orbital mechanics), live TLE data, accurate ephemeris
**Needs improvement**: Real-time conjunction analysis, maneuver detection, collision probability
