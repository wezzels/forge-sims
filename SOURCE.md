# SOURCE.md — FORGE-Sims Source Repository Map

All source code lives on **idm.wezzel.com** (GitLab) under the `crab-meat-repos` group.
Build server: **darth** (10.0.0.117), working directory `/home/wez/`.
Production server: **miner** (207.244.226.151).
GitHub mirror: `github.com:wezzels/forge-sims.git` (binaries only).

## SSH URLs

All repos: `git@idm.wezzel.com:crab-meat-repos/<name>.git`

## BMD Sensor Simulators (source on darth)

| Binary | IDM Repo | Branch | Local Path on darth |
|---|---|---|---|
| bmd-sim-aegis | bmd-sim-aegis | master | /home/wez/bmd-sim-aegis |
| bmd-sim-atmospheric | bmd-sim-atmospheric | main | /home/wez/bmd-sim-atmospheric |
| bmd-sim-cobra-judy | bmd-sim-cobra-judy | master | /home/wez/bmd-sim-cobra-judy |
| bmd-sim-decoy | bmd-sim-decoy | master | /home/wez/bmd-sim-decoy |
| bmd-sim-dsp | bmd-sim-dsp | master | /home/wez/bmd-sim-dsp |
| bmd-sim-emp-effects | bmd-sim-emp-effects | master | /home/wez/bmd-sim-emp-effects |
| bmd-sim-gbr | bmd-sim-gbr | master | /home/wez/bmd-sim-gbr |
| bmd-sim-gfcb | bmd-sim-gfcb | main | /home/wez/bmd-sim-gfcb |
| bmd-sim-hub | bmd-sim-hub | master | /home/wez/bmd-sim-hub |
| bmd-sim-ifxb | bmd-sim-ifxb | main | /home/wez/bmd-sim-ifxb |
| bmd-sim-jamming | bmd-sim-jamming | master | /home/wez/bmd-sim-jamming |
| bmd-sim-jreap | bmd-sim-jreap | main | /home/wez/bmd-sim-jreap |
| bmd-sim-jrsc | bmd-sim-jrsc | main | /home/wez/bmd-sim-jrsc |
| bmd-sim-lrdr | bmd-sim-lrdr | master | /home/wez/bmd-sim-lrdr |
| bmd-sim-patriot | bmd-sim-patriot | master | /home/wez/bmd-sim-patriot |
| bmd-sim-rcs | bmd-sim-rcs | master | /home/wez/bmd-sim-rcs |
| bmd-sim-sbirs | bmd-sim-sbirs | main | /home/wez/bmd-sim-sbirs |
| bmd-sim-space-weather | bmd-sim-space-weather | main | /home/wez/bmd-sim-space-weather |
| bmd-sim-stss | bmd-sim-stss | main | /home/wez/bmd-sim-stss |
| bmd-sim-thaad | bmd-sim-thaad | master | /home/wez/bmd-sim-thaad |
| bmd-sim-thaad-er | bmd-sim-thaad-er | master | /home/wez/bmd-sim-thaad-er |
| bmd-sim-tpy2 | bmd-sim-tpy2 | master | /home/wez/bmd-sim-tpy2 |
| bmd-sim-uewr | bmd-sim-uewr | master | /home/wez/bmd-sim-uewr |

## BMD Threat/Weapon Simulators (source on darth)

| Binary | IDM Repo | Branch | Local Path on darth |
|---|---|---|---|
| bmd-sim-hgv | bmd-sim-hgv | main | /home/wez/bmd-sim-hgv |
| bmd-sim-icbm | bmd-sim-icbm | main | /home/wez/bmd-sim-icbm |
| bmd-sim-irbm | bmd-sim-irbm | main | /home/wez/bmd-sim-irbm |
| bmd-sim-mrbm | bmd-sim-mrbm | master | /home/wez/bmd-sim-mrbm |
| bmd-sim-slcm | bmd-sim-slcm | master | /home/wez/bmd-sim-slcm |
| bmd-sim-sm3 | bmd-sim-sm3 | master | /home/wez/bmd-sim-sm3 |
| bmd-sim-sm6 | bmd-sim-sm6 | master | /home/wez/bmd-sim-sm6 |
| bmd-sim-gmd | bmd-sim-gmd | main | /home/wez/bmd-sim-gmd |

## BMD Integration/C2 Simulators (source on darth)

| Binary | IDM Repo | Branch | Local Path on darth |
|---|---|---|---|
| bmd-sim-boost-intercept | bmd-sim-boost-intercept | master | /home/wez/bmd-sim-boost-intercept |
| bmd-sim-c2bmc | bmd-sim-c2bmc | main | /home/wez/bmd-sim-c2bmc |
| bmd-sim-debris-field | bmd-sim-debris-field | master | /home/wez/bmd-sim-debris-field |
| bmd-sim-electronic-attack | bmd-sim-electronic-attack | master | /home/wez/bmd-sim-electronic-attack |
| bmd-sim-engagement-chain | bmd-sim-engagement-chain | master | /home/wez/bmd-sim-engagement-chain |
| bmd-sim-kill-assessment | bmd-sim-kill-assessment | master | /home/wez/bmd-sim-kill-assessment |
| bmd-sim-kill-chain-rt | bmd-sim-kill-chain-rt | master | /home/wez/bmd-sim-kill-chain-rt |
| bmd-sim-link16 | bmd-sim-link16 | main | /home/wez/bmd-sim-link16 |
| bmd-sim-nuclear-efx | bmd-sim-nuclear-efx | master | /home/wez/bmd-sim-nuclear-efx |
| bmd-sim-satellite-weapon | bmd-sim-satellite-weapon | master | /home/wez/bmd-sim-satellite-weapon |
| bmd-sim-tactical-net | bmd-sim-tactical-net | master | /home/wez/bmd-sim-tactical-net |
| bmd-sim-wta | bmd-sim-wta | master | /home/wez/bmd-sim-wta |
| bmd-sim-ufo | bmd-sim-ufo | master | /home/wez/bmd-sim-ufo |
| bmd-sim-air-traffic | bmd-sim-air-traffic | master | /home/wez/bmd-sim-air-traffic |
| bmd-sim-satellite-tracker | bmd-sim-satellite-tracker | master | /home/wez/bmd-sim-satellite-tracker |
| bmd-sim-space-debris | bmd-sim-space-debris | master | /home/wez/bmd-sim-space-debris |
| bmd-sim-ir-signature | bmd-sim-ir-signature | master | /home/wez/bmd-sim-ir-signature |

## Shared Libraries

| Library | IDM Repo | Branch | Local Path on darth |
|---|---|---|---|
| bmd-sim-core | bmd-sim-core | master | /home/wez/bmd-sim-core |
| sim-cli | sim-cli | main | /home/wez/sim-cli |
| bmd-sim-integration | bmd-sim-integration | main | /home/wez/bmd-sim-integration |

## Standalone Simulators (source on darth)

| Binary | IDM Repo | Branch | Local Path on darth |
|---|---|---|---|
| launch-veh-sim | launch-veh-sim | main | /home/wez/launch-veh-sim |
| population-impact-sim | population-impact-sim | main | /home/wez/population-impact-sim |
| electronic-war-sim | electronic-war-sim | main | /home/wez/electronic-war-sim |
| missile-defense-sim | missile-defense-sim | main | /home/wez/missile-defense-sim |
| submarine-war-sim | submarine-war-sim | main | /home/wez/submarine-war-sim |
| air-combat-sim | bmd-sim-air-combat | master | /home/wez/bmd-sim-air-combat |

## IDM Repos WITHOUT Source on darth (not yet cloned/built)

These repos exist on IDM but don't have local clones on darth yet:

| IDM Repo | Branch | Notes |
|---|---|---|
| bmd-sim-dashboard | main | Web dashboard |
| cbrn-sim | master | CBRN warfare |
| cyber-kinetic-sim | main | Cyber-kinetic effects |
| cyber-redteam-sim | main | Red team ops (binary exists, no CLI) |
| information-ops-sim | main | Info operations |
| land-combat-sim | main | Ground warfare |
| logistics-sustainment-sim | main | Logistics chain |
| maritime-sim | main | Maritime warfare (binary exists, needs config) |
| nuclear-effects-sim | main | Nuclear effects |
| space-data-network-sim | main | Space comms |
| space-war-sim | main | Space warfare (binary exists, needs config) |
| c4isr-killchain-sim | main | C4ISR kill chain |
| boost-intercept-sim | main | Boost-phase intercept |
| agentic-ai | main | AI agent framework |
| human-interaction-sim | master | Human factors |
| human-test-sim | main | Human testing |
| forge-c2 | main | C2 framework (original) |
| forge-c2-go | master | C2 framework (Go rewrite) |
| forge-security | main | Security framework |
| forge-sensors | main | Sensor framework |
| forge-test | main | Testing framework |
| forge-track-correlator | main | Track correlation |
| nos3-mdpaf | main | NOS3 MDPAF |
| nos3-mirror | master | NOS3 mirror |
| nos3-sim | main | NOS3 simulation |
| dod-simulator | main | DOD simulator |
| dod-simulator-new | main | DOD simulator v2 |
| argocd-config | main | ArgoCD config |
| assistant-area | main | Assistant area |
| cicerone-goclaw | main | Cicerone GoClaw |
| market-analysis-ml | main | Market ML |
| news-site | main | News site |
| seissentry | main | SeisSentry |
| snl-gms-simulator-gateway | main | SNL GMS gateway |
| stsgym-tools | master | STS Gym tools |
| trade-site | main | Trade site |
| vigil | master | Vigil |
| vimic2 | main | VIMIC2 |
| vimic2-sts | main | VIMIC2 STS |
| weavess | main | Weavess |
| weavess-core | main | Weavess core |

## Build Instructions

```bash
# SSH to darth
ssh -i ~/.ssh/id_rsa wez@10.0.0.117

# Build a single sim
cd /home/wez/<sim-dir>
go build -o /home/wez/forge-sims/binaries/linux-x86/<binary-name> ./cmd/<cmd-dir>/

# Deploy to miner
rsync -avz --chmod=F755 -e "ssh -i ~/.ssh/id_ed25519" \
  /home/wez/forge-sims/binaries/linux-x86/<binary> \
  wez@207.244.226.151:/home/wez/forge-sims/binaries/linux-x86/

# Push to GitHub + IDM
cd /home/wez/forge-sims
git push origin main
git push idm main:master
```

## Go Module Conventions

- Module path: `gitlab.com/crab-meat-repos/<name>` (note: `gitlab.com` not `idm.wezzel.com`)
- go.mod replace directives needed for `bmd-sim-core` and `sim-cli`
- Most sims use `cmd/<name>/main.go` as entry point
- Shared CLI pattern: `-json`, `-scenario <name>`, `-scenario list`