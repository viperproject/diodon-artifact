#!/bin/bash

# exit when any command fails
set -e

# Script to run the pass-through analysis bug on the DH implementation

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCH=passthru_bug.patch

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

echo "Applying bug patch: $PATCH"
git apply "$PATCH" || exit
git diff

echo "Running pass-through analysis on DH implementation in directory $(pwd)"
if "$ARGOT_BIN" diodon-passthru -config "$SCRIPT_DIR"/argot-config-dh.yaml; then
    echo "Expected analysis to fail"
    git apply --reverse "$PATCH"
    exit 1
fi

echo "Reverting bug patch: $PATCH"
git apply --reverse "$PATCH" || exit
