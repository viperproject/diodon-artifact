#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")
IMPLEMENTATION_DIR="$SCRIPT_DIR/implementation"
GOBRA_JAR="/gobra/gobra.jar"
GOBRA_REPORT_DIR="$IMPLEMENTATION_DIR/.gobra"
PATCH="send_secret_key.patch"

mkdir -p "$GOBRA_REPORT_DIR"

echo "Applying bug patch: $PATCH"
$(cd $IMPLEMENTATION_DIR && patch -s -p1 < "$PATCH") || exit

if java -Xss128m -jar "$GOBRA_JAR" \
    --module "dh-gobra" \
    --include "$IMPLEMENTATION_DIR" --include "$IMPLEMENTATION_DIR/.verification"  \
    --input "$IMPLEMENTATION_DIR/initiator/initiator.go" \
    --gobraDirectory "$GOBRA_REPORT_DIR" \
    --parseAndTypeCheckMode PARALLEL \
    --parallelizeBranches; then
    echo "Expected Gobra to fail"
    $(cd $IMPLEMENTATION_DIR && patch -sR -p1 < "$PATCH")
    exit 1
fi

echo "Reverting bug patch: $PATCH"
$(cd $IMPLEMENTATION_DIR && patch -sR -p1 < "$PATCH") || exit
