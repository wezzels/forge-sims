# FORGE-Sims

Pre-built BMDS simulator binaries for FORGE-C2.

## Binaries

| Platform | Architecture | Path |
|----------|--------------|------|
| macOS | amd64 | `binaries/mac/` |
| Linux | x86_64 | `binaries/linux-x86/` |
| Linux | ARM64 | `binaries/linux-arm/` |
| Windows | amd64 | `binaries/windows/` |

## Simulators

Each directory contains binaries for:

- **bmd-sim-core** - Core physics, trajectory, radar detection
- **bmd-sim-tpy2** - AN/TPY-2 radar simulator
- **bmd-sim-hub** - C2BMC hub / engagement coordinator
- **bmd-sim-aegis** - Aegis BMD simulator
- **bmd-sim-thaad** - THAAD battery simulator
- **bmd-sim-patriot** - Patriot battery simulator
- **bmd-sim-gbr** - GBR radar simulator

## Build Matrix

| Binary | macOS | Linux x86 | Linux ARM | Windows |
|--------|-------|-----------|-----------|---------|
| bmd-sim-core | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-tpy2 | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-hub | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-aegis | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-thaad | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-patriot | ✅ | ✅ | ✅ | ✅ |
| bmd-sim-gbr | ✅ | ✅ | ✅ | ✅ |

## Usage

Download the appropriate binary for your platform and run directly.

```bash
# Example: Run TPY-2 radar simulator
./binaries/linux-x86/bmd-sim-tpy2
```
