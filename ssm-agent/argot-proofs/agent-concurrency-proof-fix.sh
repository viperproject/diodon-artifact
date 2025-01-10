#!/bin/bash

# exit when any command fails
set -e

# Script to run the goroutine analysis proof on the SSM agent

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

echo "Running goroutine analysis on SSM agent in directory $(pwd)"

"$ARGOT_BIN" goroutine -config "$SCRIPT_DIR"/argot-config-agent.yaml | tee /tmp/argot-output.txt || true

if grep -F '[INFO]  Analysis took ' /tmp/argot-output.txt; then
    echo "Found end of results"
else
    echo "Analysis did not complete"
    exit 1
fi

if grep -F '[ERROR]' /tmp/argot-output.txt \
   | grep -Ev 'Parameter dc of .* has escaped: argument to go at .*(/datastream\.go:(432|360)|/websocketchannel\.go:(171|222))' \
   | grep -v 'Parameter log of .* has escaped: argument to go at' \
   | grep -Fv '[ERROR] Analysis for ssm-session-worker failed: goroutine analysis found problems, inspect logs for more information' \
   | grep -F '[ERROR]'; then
    echo "Other errors found"
    exit 1
else
    echo "Only allowlisted errors found"
fi
