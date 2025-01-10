#!/bin/bash

# exit when any command fails
set -e

# Script to run the taint analysis bugs on the SSM agent

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCHES="logging_shared_secret_bug.patch logging_eph_priv_key_bug.patch"

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

for PATCH in $PATCHES; do
    echo "Applying bug patch: $PATCH"
    patch -s -p1 < "$PATCH" || exit

    echo "Running taint analysis on SSM agent in directory $(pwd)"
    if "$ARGOT_BIN" taint -config "$SCRIPT_DIR"/argot-config-agent.yaml; then
        echo "Expected analysis to fail"
        patch -sR -p1 < "$PATCH"
        exit 1
    fi

    echo "Reverting bug patch: $PATCH"
    patch -sR -p1 < "$PATCH" || exit
done
