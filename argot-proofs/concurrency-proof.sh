#!/bin/bash

# exit when any command fails
set -e

# Script to run the goroutine analysis proof on the SSM agent

SCRIPT_DIR=$(dirname "$0")
DIODON_DIR=$(realpath "$SCRIPT_DIR"/..)
AGENT_DIR="$DIODON_DIR"/implementation
ARGOT_DIR="$DIODON_DIR"/ar-go-tools
PROOF_DIR="$DIODON_DIR"/argot-proofs
REPORT_DIR="$PROOF_DIR"/reports
ARGOT_BIN="$ARGOT_DIR"/bin/argot

# Compile Argot goro-check tool
cd "$ARGOT_DIR" || exit
echo "Compiling Argot tool in directory $(pwd)"
make argot-build || exit

# Run the goroutine analysis
if [ ! -e "$REPORT_DIR" ]; then
    mkdir -p "$REPORT_DIR"
fi

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

GORO_CHECK_LOG_FILE="$REPORT_DIR"/goro-check-log

echo "Running goroutine analysis on SSM agent in directory $(pwd)"
echo "Saving log to ${GORO_CHECK_LOG_FILE}"

"$ARGOT_BIN" goroutine -config "$PROOF_DIR"/argot-config.yaml \
    "$AGENT_DIR"/agent/framework/processor/executer/outofproc/sessionworker/main.go
# > "$goro-check_LOG_FILE" 2>&1 # redirect stdout and stderr to log file
