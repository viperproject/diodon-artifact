#!/bin/bash

# exit when any command fails
set -e

# Script to run the immutability analysis bugs on the DH implementation

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCHES="read_core_state_bug.patch write_core_state_bug.patch"

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

for PATCH in $PATCHES; do
    echo "Applying bug patch: $PATCH"
    git apply "$PATCH" || exit
    git diff

    echo "Running immutability analysis on DH implementation in directory $(pwd)"

    if "$ARGOT_BIN" immutability -config "$SCRIPT_DIR"/argot-config-dh.yaml; then
        echo "Expected analysis to fail"
        git apply --reverse "$PATCH"
        exit 1
    fi

    echo "Reverting bug patch: $PATCH"
    git apply --reverse "$PATCH" || exit
done
