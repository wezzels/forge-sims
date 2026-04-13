# FORGE-Sims

Pre-built BMDS simulator binaries for FORGE-C2.

## Directory Structure

```
forge-sims/
├── binaries/
│   ├── mac/           # macOS binaries (cross-compile needed)
│   ├── linux-x86/      # Linux x86_64 binaries ✅
│   ├── linux-arm/      # Linux ARM64 binaries (cross-compile needed)
│   └── windows/        # Windows binaries (cross-compile needed)
└── README.md
```

## Binaries

| Platform | Architecture | Status |
|----------|--------------|--------|
| macOS | amd64/arm64 | Placeholder |
| Linux | x86_64 | ✅ Built |
| Linux | ARM64 | Placeholder |
| Windows | amd64 | Placeholder |

## Simulators

| Binary | Description | Status |
|--------|-------------|--------|
| bmd-sim-core | Core physics (library only) | N/A |
| bmd-sim-tpy2 | AN/TPY-2 radar simulator | ✅ Built |
| bmd-sim-hub | C2BMC hub / engagement coordinator | ✅ Built |
| bmd-sim-aegis | Aegis BMD simulator | ✅ Built |
| bmd-sim-thaad | THAAD battery simulator | ✅ Built |
| bmd-sim-patriot | Patriot battery simulator | ✅ Built |
| bmd-sim-gbr | GBR radar simulator | ✅ Built |

## Build Status

Currently only Linux x86_64 binaries are built natively. Other platforms
require cross-compilation (not available on this machine).

To build for other platforms when you have the proper environment:

```bash
# macOS
GOOS=darwin GOARCH=amd64 go build -o bmd-sim-hub ./cmd/hub/

# Windows
GOOS=windows GOARCH=amd64 go build -o bmd-sim-hub.exe ./cmd/hub/

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o bmd-sim-hub ./cmd/hub/
```
