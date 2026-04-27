# Intel Collection Sim — INSTALL

## Quick Install

```bash
# Linux x86_64
sudo cp binaries/linux-x86/intel-collection-sim /usr/local/bin/
sudo chmod +x /usr/local/bin/intel-collection-sim

# Verify
intel-collection-sim version
```

## Automated Install

```bash
bash binaries/install-intel-collection-sim.sh
```

## Platforms

| OS | Arch | Binary |
|----|------|--------|
| Linux | x86_64 | `binaries/linux-x86/intel-collection-sim` |
| Linux | ARM64 | `binaries/linux-arm64/intel-collection-sim` |
| macOS | x86_64 | `binaries/darwin-amd64/intel-collection-sim` |
| macOS | ARM64 | `binaries/darwin-arm64/intel-collection-sim` |
| Windows | x86_64 | `binaries/windows-amd64/intel-collection-sim.exe` |

## Requirements

- Go 1.22+ (build from source only)
- No runtime dependencies