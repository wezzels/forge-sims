# air-combat-sim — Air Combat Simulation

## Purpose
Models air combat engagement between fighter aircraft.

## Physics & Fidelity
- Aircraft performance data (thrust, wing loading, turn rates)
- BVR (beyond visual range) and WVR (within visual range) engagements
- Missile PK modeling (AMRAAM, Sidewinder, etc.)
- Fuel consumption and endurance

### Scenarios
Run with `-scenario <name> -json -duration 300s`

## Accuracy: Low-Medium
**Reliable**: Basic engagement geometry, missile PK
**Needs improvement**: Simplified flight dynamics, no detailed radar modeling, limited aircraft database
