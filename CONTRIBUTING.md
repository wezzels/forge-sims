# Contributing to Forge-Sims

This repository contains **compiled binaries and documentation only**. Source code lives in separate repositories on IDM GitLab.

## Source Repositories

All source code is hosted on IDM GitLab:

```
git@idm.wezzel.com:crab-meat-repos/<sim-name>.git
```

For example:
- `git@idm.wezzel.com:crab-meat-repos/bmd-sim-gmd.git`
- `git@idm.wezzel.com:crab-meat-repos/ufo.git`
- `git@idm.wezzel.com:crab-meat-repos/engagement-chain.git`

## Build Process

1. **Clone** the source repo:
   ```bash
   git clone git@idm.wezzel.com:crab-meat-repos/<sim-name>.git
   cd <sim-name>
   ```

2. **Build** for Linux x86_64:
   ```bash
   GOOS=linux GOARCH=amd64 go build -o <sim-name> ./cmd/<sim-name>/
   ```

3. **Copy** the binary to this repo:
   ```bash
   cp <sim-name> /path/to/forge-sims/binaries/linux-x86/
   ```

## Testing

Every binary must pass these checks before merging:

1. **Clean exit**: Binary exits with code 0
   ```bash
   ./<sim-name> -json -duration 1s; echo $?
   ```

2. **Valid JSON**: Output is parseable JSON (for binaries with `-json` flag)
   ```bash
   ./<sim-name> -json -duration 1s | python3 -m json.tool > /dev/null
   ```

3. **Flag check**: All documented flags work without errors
   ```bash
   ./<sim-name> -help
   ./<sim-name> -v          # if supported
   ./<sim-name> -seed 42     # if supported
   ```

4. **No crashes**: Run with default parameters for at least 30 seconds
   ```bash
   ./<sim-name> -duration 30s
   ```

## Pull Request Process

1. Create a feature branch from `main`
2. Add/update binaries in `binaries/linux-x86/`
3. Update `BINARY_STATUS.md` if adding a new binary or changing flags
4. Update `API_REFERENCE.md` with any JSON schema changes
5. Verify all 47 binaries still exit 0
6. Submit PR with description of changes

## Repository Structure

```
forge-sims/
├── binaries/linux-x86/    # Compiled Linux x86_64 binaries
├── configs/               # Configuration files for config-driven sims
├── docs/                  # Additional documentation
├── scripts/               # Build/deploy scripts
├── API_REFERENCE.md       # JSON output schema for all binaries
├── BINARY_STATUS.md       # Flag and status table for all binaries
├── CONTRIBUTING.md        # This file
├── README.md              # Project overview
└── EVALUATION.md          # Evaluation history
```

## Code of Conduct

Be professional. This is a defense simulation project — accuracy and reliability matter.