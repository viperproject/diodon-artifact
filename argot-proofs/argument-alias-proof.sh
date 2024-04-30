#!/bin/bash

# exit when any command fails
set -e

# Script to run the argalias analysis proof on the SSM agent

SCRIPT_DIR=$(dirname "$0")
DIODON_DIR=$(realpath "$SCRIPT_DIR"/..)
AGENT_DIR="$DIODON_DIR"/implementation
ARGOT_DIR="$DIODON_DIR"/ar-go-tools
PROOF_DIR="$DIODON_DIR"/argot-proofs
REPORT_DIR="$PROOF_DIR"/reports
ARGALIAS_BIN="$ARGOT_DIR"/bin/argalias

# Compile Argot argalias tool
cd "$ARGOT_DIR" || exit
echo "Compiling Argot argalias tool in directory $(pwd)"
make argalias || exit

# Run the argalias analysis
if [ ! -e "$REPORT_DIR" ]; then
    mkdir -p "$REPORT_DIR"
fi

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit

ARGALIAS_LOG_FILE="$REPORT_DIR"/argalias-log

echo "Running argalias (${ARGALIAS_BIN}) analysis on SSM agent in directory $(pwd)"
echo "Saving log to ${ARGALIAS_LOG_FILE}"

"$ARGALIAS_BIN" -config "$PROOF_DIR"/argot-config.yaml \
    "$AGENT_DIR"/agent/framework/processor/executer/outofproc/sessionworker/main.go \
    > "$ARGALIAS_LOG_FILE" 2>&1 # redirect stdout and stderr to log file
