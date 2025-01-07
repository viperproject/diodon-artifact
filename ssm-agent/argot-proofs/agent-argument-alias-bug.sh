#!/bin/bash

# exit when any command fails
set -e

# Script to run the argument alias analysis bug on the SSM agent

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCH=alias_arguments_bug.patch

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

echo "Applying bug patch: $PATCH"
git apply "$PATCH" || exit
git diff

echo "Running argument alias analysis on SSM agent in directory $(pwd)"
if "$ARGOT_BIN" alias -config "$SCRIPT_DIR"/argot-config-agent.yaml; then
    echo "Expected analysis to fail"
    git apply --reverse "$PATCH"
    exit 1
fi

echo "Reverting bug patch: $PATCH"
git apply --reverse "$PATCH" || exit
