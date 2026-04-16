#!/bin/bash
# Deploy all FORGE-Sims Dockerfiles to each repo
# Usage: bash deploy-dockerfiles.sh

set -e

REPOS=(
  bmd-sim-sbirs bmd-sim-stss bmd-sim-dsp
  bmd-sim-uewr bmd-sim-lrdr bmd-sim-cobra-judy
  bmd-sim-gmd bmd-sim-sm3 bmd-sim-sm6 bmd-sim-thaad-er
  bmd-sim-icbm bmd-sim-irbm bmd-sim-hgv bmd-sim-slcm
  bmd-sim-decoy bmd-sim-jamming
  bmd-sim-c2bmc bmd-sim-gfcb bmd-sim-ifxb bmd-sim-jrsc
  bmd-sim-link16 bmd-sim-jreap
  bmd-sim-space-weather bmd-sim-atmospheric
)

DOCKERFILE='FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 go build -o /sim ./cmd/...

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /sim /sim
EXPOSE 8080
ENTRYPOINT ["/sim"]'

for repo in "${REPOS[@]}"; do
  if [ -d "/home/wez/$repo" ]; then
    echo "$DOCKERFILE" > "/home/wez/$repo/Dockerfile"
    echo "✓ $repo/Dockerfile"
  else
    echo "✗ $repo not found"
  fi
done

echo ""
echo "Done. Run 'docker-compose up -d' from the docker-compose directory."