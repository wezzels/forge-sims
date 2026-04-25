# bmd-sim-uewr — Upgraded Early Warning Radar

## Purpose
Models UEWR (Pave Paws / BMEWS) UHF phased-array radars for strategic early warning and midcourse tracking.

## Physics & Fidelity

### Radar Parameters
- **Frequency**: 420-450 MHz (UHF)
- **Peak power**: 582.4 kW
- **Antenna gain**: 38 dBi (phased array face)
- **Effective aperture**: ~6,800 m²
- **System losses**: 12 dB (atmospheric, beamshape, scanning)

### Detection Model
- Full radar range equation: `SNR = (Pt × τ × N × G² × λ² × σ) / ((4π)³ × R⁴ × k × T × Ls)`
- **Energy-on-target**: PeakPower × PulseWidth × NPulses (not AvgPower/Bandwidth)
- **Pulse integration gain**: 5·log₁₀(N) for non-coherent
- **Detection probability**: Min(1, Max(0, pd)) — properly clamped
- **RCS model**: Aspect-dependent with Swerling fluctuation
- Multiple site modeling (Beale AFB, Fylingdales, Thule, Clear)

### Scenarios
`icbm-detection`, `multi-threat-tracking`, `arctic-coverage`, `degraded-mode`

## Accuracy: High
**Reliable**: Correct energy-on-target radar equation, realistic system losses (12 dB), detection probability clamping, multi-site geometry
**Needs improvement**: Atmospheric refraction at UHF, ionospheric scintillation, clutter modeling, track accuracy jitter
