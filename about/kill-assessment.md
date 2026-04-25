# kill-assessment — Post-Intercept Damage Assessment

## Purpose
Models post-intercept kill assessment: warhead damage, structural, fuzing, and guidance system integrity.

## Physics & Fidelity
- 7 warhead types with vulnerability profiles
- 4 kill vehicle types (EKV, SM-3 KW, THAAD, PAC-3)
- Hit-to-kill vs blast-frag assessment
- Assessment sensors: IR, radar, optical
- 5 assessment phases with timing
- False kill rate: 1-5%, false miss rate: 2-10%

### Scenarios
`ekv-vs-icbm`, `sm3-vs-irbm`, `thaad-vs-mrbm`, `pac3-vs-srbm`, `multi-assessment`

## Accuracy: Medium-High
**Reliable**: Warhead vulnerability profiles, kill vehicle comparison, assessment timing
**Needs improvement**: Debris modeling after intercept, detailed fuzing assessment, forensic analysis
