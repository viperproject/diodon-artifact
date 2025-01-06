#!/bin/bash

# Script to run the taint analysis bug on the DH implementation

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCHES="leak_private_key_bug.patch leak_nonce_bug.patch"

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

for PATCH in $PATCHES; do
    echo "Applying bug patch: $PATCH"
    git apply "$PATCH" || exit

    echo "Running taint analysis on DH implementation in directory $(pwd)"
    if "$ARGOT_BIN" taint -config "$SCRIPT_DIR"/argot-config-dh.yaml; then
        echo "Expected analysis to fail"
        git apply --reverse "$PATCH"
        exit 1
    fi

    echo "Reverting bug patch: $PATCH"
    git apply --reverse "$PATCH" || exit
done
