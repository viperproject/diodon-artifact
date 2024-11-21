#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")
IMPLEMENTATION_DIR="$SCRIPT_DIR/implementation"
GOBRA_JAR="/gobra/gobra.jar"
GOBRA_REPORT_DIR="$IMPLEMENTATION_DIR/.gobra"

mkdir -p "$GOBRA_REPORT_DIR"

java -Xss128m -jar "$GOBRA_JAR" \
    --module "dh-gobra" \
    --include "$IMPLEMENTATION_DIR" --include "$IMPLEMENTATION_DIR/.verification"  \
    --input "$IMPLEMENTATION_DIR/initiator/initiator.go" \
    --gobraDirectory "$GOBRA_REPORT_DIR" \
    --parseAndTypeCheckMode PARALLEL \
    --parallelizeBranches
