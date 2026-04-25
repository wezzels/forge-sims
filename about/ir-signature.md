# ir-signature — IR Signature Modeling

## Purpose
Models infrared signatures of ballistic threats for BMDS IR sensor detection.

## Physics & Fidelity
- Boost, midcourse, and terminal phase IR signatures
- Background radiance (day/night, Earth limb, cold space)
- Detection SNR: IR SNR model with range, signature, and background
- HGV midcourse: Pd = 0% (SNR -18.6 dB — undetectable)
- Stealth Sarmat boost: still Pd = 100% (bright booster plume)

### Scenarios
`boost-detection`, `midcourse-tracking`, `terminal-detection`, `hgv-stealth`

## Accuracy: Medium-High
**Reliable**: IR detection SNR model, phase-dependent signatures, background radiance
**Needs improvement**: Detailed spectral band modeling, atmospheric path radiance, solar angle effects
