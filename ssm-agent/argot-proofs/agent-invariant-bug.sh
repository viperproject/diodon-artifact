#!/bin/bash

# Script to run the immutability analysis bugs on the SSM agent

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCHES="write_core_state_bug.patch read_core_state_bug.patch"

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

for PATCH in $PATCHES; do
    echo "Applying bug patch: $PATCH"
    patch -s -p1 < "$PATCH" || exit

    echo "Running immutability analysis on SSM agent in directory $(pwd)"

    if "$ARGOT_BIN" immutability -config "$SCRIPT_DIR"/argot-config-agent.yaml; then
        echo "Expected analysis to fail"
        patch -sR -p1 < "$PATCH"
        exit 1
    fi

    echo "Reverting bug patch: $PATCH"
    patch -sR -p1 < "$PATCH" || exit
done
