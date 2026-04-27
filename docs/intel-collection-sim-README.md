# Intel Collection Sim — Distribution README

## Version: v1.1.0

High-fidelity intelligence collection and warfare simulation engine.

## What It Does

Models the full intelligence cycle across 6 INT disciplines with dynamic adversary counter-collection, Bayesian multi-INT fusion, and game-theoretic asset tasking.

## Disciplines

| INT | Models | Fidelity |
|-----|--------|----------|
| SIGINT | FSPL, radar equation, Swerling targets, COMINT/ELINT/FISINT, freq hop tracking, TDOA/FDOA/AOA geolocation, rain fade, multipath | High |
| HUMINT | MICE recruitment, ABCDEF reliability, source aging, compromise cascades, walk-in assessment, meeting security | High |
| IMINT | GSD/NIIRS, orbital revisit, photogrammetry, change detection, weather degradation | High |
| MASINT | Seismic discrimination (Mb-Ms), CTBT IMS, hydroacoustic, ACINT, RADINT, CBINT, IR | High |
| OSINT | Social media volume, NLP entity extraction, source credibility, news cycle | Medium |
| FININT | Layering detection, crypto tracing, shell companies, hawala, sanctions evasion | Medium |

## Plus

- **Counterintelligence**: Mole detection, double-agent games, damage assessment
- **Covert Action**: Influence operations, propaganda, deception, front companies
- **Adversary D&D**: OPSEC, maskirovka, counter-collection, game-theoretic strategies
- **Dynamic Simulation**: Time-stepped cycle with adversary reaction, source degradation, fusion convergence

## Scenarios

1. Cold War SIGINT
2. Counter-Proliferation
3. Counterterrorism
4. Near-Peer Competition
5. Gray Zone
6. Mole Hunt

## Quick Start

```bash
# Run a scenario
intel-collection-sim run --scenario near-peer

# Validate physics models
intel-collection-sim validate

# List platforms
intel-collection-sim platforms

# Single discipline
intel-collection-sim collect --discipline sigint --target radar-site-1
```

## Stats

- 551 tests, all green
- 90+ physics validations against published references
- 13 packages, 7-phase architecture
- Cross-sim integration with forge-sims constellation

## License

Proprietary — Crab Meat Repos