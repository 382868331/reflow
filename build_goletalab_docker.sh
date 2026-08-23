#!/usr/bin/env bash
set -euo pipefail
docker build -f goletalab.Dockerfile -t reflow-task .
docker run --rm reflow-task "$@"
