#!/usr/bin/env bash
set -euo pipefail
docker build --platform linux/amd64 -f goletalab.Dockerfile -t goletalab-reflow .
