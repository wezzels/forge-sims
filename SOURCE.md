# SOURCE.md — FORGE-Sims Source Repository Map

All source code lives on **idm.wezzel.com** (GitLab) under the `crab-meat-repos` group.
Build server: **darth** (10.0.0.117), working directory `/home/wez/`.
Production server: **miner** (207.244.226.151).
GitHub mirror: `github.com:wezzels/forge-sims.git` (binaries only).

## SSH URL Pattern

`git@idm.wezzel.com:crab-meat-repos/<repo>.git`

## Binary → Source Mapping

Every binary in `binaries/linux-x86/` mapped to its source repo:

| # | Binary | Source Repo (IDM) | Branch | Local Path on darth | CLI Status |
|---|--------|-------------------|--------|---------------------|------------|
| 1 | air-combat-sim | bmd-sim-air-combat | master | /home/wez/bmd-sim-air-combat | ✅ -json -scenario |
| 2 | air-traffic | bmd-sim-air-traffic | master | *(not cloned yet)* | ✅ -json -scenario |
| 3 | bmd-sim-aegis | bmd-sim-aegis | master | /home/wez/bmd-sim-aegis | ✅ -json -scenario |
| 4 | bmd-sim-atmospheric | bmd-sim-atmospheric | main | /home/wez/bmd-sim-atmospheric | ✅ -json -scenario |
| 5 | bmd-sim-c2bmc | bmd-sim-c2bmc | main | /home/wez/bmd-sim-c2bmc | ✅ -json -scenario |
| 6 | bmd-sim-cobra-judy | bmd-sim-cobra-judy | master | /home/wez/bmd-sim-cobra-judy | ✅ -json -scenario |
| 7 | bmd-sim-decoy | bmd-sim-decoy | master | /home/wez/bmd-sim-decoy | ✅ -json -scenario |
| 8 | bmd-sim-dsp | bmd-sim-dsp | master | /home/wez/bmd-sim-dsp | ✅ -json -scenario |
| 9 | bmd-sim-electronic-attack | bmd-sim-electronic-attack | master | /home/wez/bmd-sim-electronic-attack | ✅ -json -scenario |
| 10 | bmd-sim-gbr | bmd-sim-gbr | master | /home/wez/bmd-sim-gbr | ✅ -json -scenario |
| 11 | bmd-sim-gfcb | bmd-sim-gfcb | main | /home/wez/bmd-sim-gfcb | ✅ -json -scenario |
| 12 | bmd-sim-gmd | bmd-sim-gmd | main | /home/wez/bmd-sim-gmd | ✅ -json -scenario |
| 13 | bmd-sim-hub | bmd-sim-hub | master | /home/wez/bmd-sim-hub | ✅ -json -scenario |
| 14 | bmd-sim-icbm | bmd-sim-icbm | main | /home/wez/bmd-sim-icbm | ✅ -json -scenario |
| 15 | bmd-sim-ifxb | bmd-sim-ifxb | main | /home/wez/bmd-sim-ifxb | ✅ -json -scenario |
| 16 | bmd-sim-irbm | bmd-sim-irbm | main | /home/wez/bmd-sim-irbm | ✅ -json -scenario |
| 17 | bmd-sim-jamming | bmd-sim-jamming | master | /home/wez/bmd-sim-jamming | ✅ -json -scenario |
| 18 | bmd-sim-jreap | bmd-sim-jreap | main | /home/wez/bmd-sim-jreap | ✅ -json -scenario |
| 19 | bmd-sim-jrsc | bmd-sim-jrsc | main | /home/wez/bmd-sim-jrsc | ✅ -json -scenario |
| 20 | bmd-sim-lrdr | bmd-sim-lrdr | master | /home/wez/bmd-sim-lrdr | ✅ -json -scenario |
| 21 | bmd-sim-mrbm | bmd-sim-mrbm | master | *(not cloned yet)* | ✅ -json -scenario |
| 22 | bmd-sim-nuclear-efx | bmd-sim-nuclear-efx | master | *(not cloned yet)* | ✅ -json -scenario |
| 23 | bmd-sim-patriot | bmd-sim-patriot | master | /home/wez/bmd-sim-patriot | ✅ -json -scenario |
| 24 | bmd-sim-sbirs | bmd-sim-sbirs | main | /home/wez/bmd-sim-sbirs | ✅ -json -scenario |
| 25 | bmd-sim-slcm | bmd-sim-slcm | master | /home/wez/bmd-sim-slcm | ✅ -json -scenario |
| 26 | bmd-sim-sm3 | bmd-sim-sm3 | master | /home/wez/bmd-sim-sm3 | ✅ -json -scenario |
| 27 | bmd-sim-sm6 | bmd-sim-sm6 | master | /home/wez/bmd-sim-sm6 | ✅ -json -scenario |
| 28 | bmd-sim-space-weather | bmd-sim-space-weather | main | /home/wez/bmd-sim-space-weather | ✅ -json -scenario |
| 29 | bmd-sim-stss | bmd-sim-stss | main | /home/wez/bmd-sim-stss | ✅ -json -scenario |
| 30 | bmd-sim-thaad | bmd-sim-thaad | master | /home/wez/bmd-sim-thaad | ✅ -json -scenario |
| 31 | bmd-sim-thaad-er | bmd-sim-thaad-er | master | /home/wez/bmd-sim-thaad-er | ✅ -json -scenario |
| 32 | bmd-sim-tpy2 | bmd-sim-tpy2 | master | /home/wez/bmd-sim-tpy2 | ✅ -json -scenario |
| 33 | bmd-sim-uewr | bmd-sim-uewr | master | /home/wez/bmd-sim-uewr | ✅ -json -scenario |
| 34 | boost-intercept | bmd-sim-boost-intercept | master | /home/wez/bmd-sim-boost-intercept | ✅ -json -scenario |
| 35 | cyber-redteam-sim | cyber-redteam-sim | main | *(not cloned yet)* | ⚠️ no -json flag |
| 36 | debris-field | bmd-sim-debris-field | master | /home/wez/bmd-sim-debris-field | ✅ -json -scenario |
| 37 | electronic-war-sim | electronic-war-sim | main | /home/wez/electronic-war-sim | ✅ -json -scenario |
| 38 | emp-effects | bmd-sim-emp-effects | master | /home/wez/bmd-sim-emp-effects | ✅ -json -scenario |
| 39 | engagement-chain | bmd-sim-engagement-chain | master | *(not cloned yet)* | ✅ -json -scenario |
| 40 | hgv | bmd-sim-hgv | main | /home/wez/bmd-sim-hgv | ✅ -json -scenario |
| 41 | ir-signature | bmd-sim-ir-signature | master | /home/wez/bmd-sim-ir-signature | ✅ -json -scenario |
| 42 | kill-assessment | bmd-sim-kill-assessment | master | *(not cloned yet)* | ✅ -json -scenario |
| 43 | kill-chain-rt | bmd-sim-kill-chain-rt | master | /home/wez/bmd-sim-kill-chain-rt | ✅ -json -scenario |
| 44 | launch-veh-sim | launch-veh-sim | main | /home/wez/launch-veh-sim | ✅ -json -vehicle |
| 45 | maritime-sim | maritime-sim | main | *(not cloned yet)* | ⚠️ needs config file |
| 46 | missile-defense-sim | missile-defense-sim | main | /home/wez/missile-defense-sim | ✅ -json -scenario |
| 47 | population-impact-sim | population-impact-sim | main | /home/wez/population-impact-sim | ✅ -json -scenario |
| 48 | rcs | bmd-sim-rcs | master | /home/wez/bmd-sim-rcs | ✅ -json -scenario |
| 49 | satellite-tracker | bmd-sim-satellite-tracker | master | *(not cloned yet)* | ✅ -json -scenario |
| 50 | satellite-weapon | bmd-sim-satellite-weapon | master | /home/wez/bmd-sim-satellite-weapon | ✅ -json -scenario |
| 51 | space-debris | bmd-sim-space-debris | master | *(not cloned yet)* | ✅ -json -scenario |
| 52 | space-war-sim | space-war-sim | main | *(not cloned yet)* | ⚠️ needs config file |
| 53 | submarine-war-sim | submarine-war-sim | main | /home/wez/submarine-war-sim | ✅ -json -scenario |
| 54 | tactical-net | bmd-sim-tactical-net | master | /home/wez/bmd-sim-tactical-net | ✅ -json -scenario |
| 55 | ufo | bmd-sim-ufo | master | /home/wez/bmd-sim-ufo | ✅ -json -scenario |
| 56 | wta | bmd-sim-wta | master | *(not cloned yet)* | ✅ -json -scenario |

**Summary:** 56 binaries total. 51 have full CLI support. 3 have issues. 2 need config files.

## Sources NOT Yet Built into Binaries

These repos are cloned on darth but don't produce a standalone binary in forge-sims:

| Source | Purpose |
|--------|---------|
| bmd-sim-core | Shared library (imported by other sims) |
| sim-cli | Shared CLI library (imported by other sims) |
| bmd-sim-integration | Integration tests |

## IDM Repos Not Yet Cloned on darth

| IDM Repo | Branch | Category | Notes |
|----------|--------|----------|-------|
| bmd-sim-air-traffic | master | BMD | Source for `air-traffic` binary |
| bmd-sim-dashboard | main | BMD | Web dashboard |
| bmd-sim-mrbm | master | BMD | Source for `bmd-sim-mrbm` binary |
| bmd-sim-nuclear-efx | master | BMD | Source for `bmd-sim-nuclear-efx` binary |
| bmd-sim-engagement-chain | master | BMD | Source for `engagement-chain` binary |
| bmd-sim-kill-assessment | master | BMD | Source for `kill-assessment` binary |
| bmd-sim-satellite-tracker | master | BMD | Source for `satellite-tracker` binary |
| bmd-sim-space-debris | master | BMD | Source for `space-debris` binary |
| bmd-sim-wta | master | BMD | Source for `wta` binary |
| cbrn-sim | master | Warfare | CBRN simulation |
| c4isr-killchain-sim | main | Warfare | C4ISR kill chain |
| cyber-kinetic-sim | main | Warfare | Cyber-kinetic effects |
| cyber-redteam-sim | main | Warfare | Source for `cyber-redteam-sim` binary |
| information-ops-sim | main | Warfare | Info operations |
| land-combat-sim | main | Warfare | Ground warfare |
| logistics-sustainment-sim | main | Warfare | Logistics chain |
| maritime-sim | main | Warfare | Source for `maritime-sim` binary |
| nuclear-effects-sim | main | Warfare | Nuclear effects |
| space-data-network-sim | main | Warfare | Space comms |
| space-war-sim | main | Warfare | Source for `space-war-sim` binary |
| boost-intercept-sim | main | Warfare | Boost-phase intercept |
| agentic-ai | main | Infra | AI agent framework |
| human-interaction-sim | master | Infra | Human factors |
| human-test-sim | main | Infra | Human testing |
| forge-c2 | main | Infra | C2 framework (original) |
| forge-c2-go | master | Infra | C2 framework (Go) |
| forge-security | main | Infra | Security framework |
| forge-sensors | main | Infra | Sensor framework |
| forge-test | main | Infra | Testing framework |
| forge-track-correlator | main | Infra | Track correlation |
| nos3-mdpaf | main | Infra | NOS3 MDPAF |
| nos3-mirror | master | Infra | NOS3 mirror |
| nos3-sim | main | Infra | NOS3 simulation |
| dod-simulator | main | Infra | DOD simulator |
| dod-simulator-new | main | Infra | DOD simulator v2 |
| argocd-config | main | Infra | ArgoCD config |
| cicerone-goclaw | main | Infra | Cicerone GoClaw |
| seissentry | main | Infra | SeisSentry |
| snl-gms-simulator-gateway | main | Infra | SNL GMS gateway |
| stsgym-tools | master | Infra | STS Gym tools |
| market-analysis-ml | main | Web | Market ML |
| news-site | main | Web | News site |
| trade-site | main | Web | Trade site |
| vigil | master | Web | Vigil |
| vimic2 | main | Web | VIMIC2 |
| vimic2-sts | main | Web | VIMIC2 STS |
| weavess | main | Web | Weavess |
| weavess-core | main | Web | Weavess core |
| assistant-area | main | Web | Assistant area |

## Build Instructions

```bash
# SSH to darth
ssh -i ~/.ssh/id_rsa wez@10.0.0.117

# Clone a repo
cd /home/wez
GIT_SSH_COMMAND="ssh -i ~/.ssh/id_rsa" git clone git@idm.wezzel.com:crab-meat-repos/<name>.git

# Build a sim
cd /home/wez/<sim-dir>
go build -o /home/wez/forge-sims/binaries/linux-x86/<binary> ./cmd/<cmd>/

# Deploy to miner
rsync -avz --chmod=F755 -e "ssh -i ~/.ssh/id_ed25519" \
  /home/wez/forge-sims/binaries/linux-x86/<binary> \
  wez@207.244.226.151:/home/wez/forge-sims/binaries/linux-x86/

# Push forge-sims to remotes
cd /home/wez/forge-sims
git push origin main        # GitHub
git push idm main:master     # IDM
```

## Go Module Conventions

- Module path: `gitlab.com/crab-meat-repos/<name>` (note: `gitlab.com` not `idm.wezzel.com`)
- go.mod replace directives needed for `bmd-sim-core` and `sim-cli`
- Most sims use `cmd/<name>/main.go` as entry point
- Shared CLI pattern: `-json`, `-scenario <name>`, `-scenario list`