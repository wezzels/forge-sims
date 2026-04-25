# launch-veh-sim — Launch Vehicle Simulator

## Purpose
Models orbital launch vehicles with realistic ascent physics, atmospheric drag, and orbit insertion.

## Physics & Fidelity

### Spherical Gravity Model
- Polar coordinates (r, θ, Vr, Vt) with centripetal term Vt²/r
- Earth rotation: Ω = 7.292e-5 rad/s
- Multi-stage boost with altitude-dependent thrust (ThrustVac field)
- DragArea per vehicle (F9=3.5, FH=10, Starship=9, etc.)
- Atmosphere-relative velocity for drag and Max Q: `Vrel = sqrt(Vr² + (Vt - ΩR·cos(i))²)`

### Vehicle Fleet (9 vehicles)
| Vehicle | Payload (kg) | Orbit Achieved | Max Q (kPa) | Status |
|---------|-------------|----------------|--------------|--------|
| Falcon 9 | 22,800 | 546×492 km | 30.7 | ✅ |
| Falcon 9-Exp | 18,000 | 580×415 km | 30.4 | ✅ |
| Falcon Heavy | 26,000 | 858×788 km | 41.1 | ✅ |
| Starship | 100,000 | 935×569 km | 29.1 | ✅ (fuel-limited) |
| New Glenn | 13,000 | 380×318 km | 25.0 | ✅ |
| Vulcan | 10,000 | 1,487×1,414 km | 71.1 | ✅ (high orbit) |
| Atlas V | 8,200 | 1,997×1,914 km | 59.3 | ✅ (high orbit) |
| Delta IV | 6,400 | 1,816×150 km | 62.3 | ✅ (eccentric) |
| Ariane 6 | 5,000 | 2,246×2,162 km | 53.7 | ✅ (high orbit) |
| Long March 5 | 14,000 | — | 23.8 | ⚠️ (fuel-limited) |

### Key Physics
- **Max Q**: 30.7 kPa (F9, real ≈ 35 kPa) — atmosphere-relative velocity
- **Energy-on-target**: PeakPower × PulseWidth × NPulses (not AvgPower/Bandwidth)
- **Perigee threshold**: 100 km for orbit detection
- **Multi-burn S2**: Low-TWR stages (Delta IV, LM5) use burn-coast-burn

### Scenarios
Run with `-vehicle <name> -json -duration 9000s`

## Accuracy: High (ascent), Medium (orbit insertion)
**Reliable**: Spherical gravity, atmospheric drag, Max Q, stage separation
**Needs improvement**: Delta IV eccentric orbit (needs better S2 guidance), LM5 doesn't orbit (fuel-limited), Vulcan/Atlas/Ariane overshoot target altitude, no Max Q throttle (rejected — crashes vehicles)
