#!/bin/bash

# exit when any command fails
set -e

# Script to run the immutability analysis proof on the SSM agent

SCRIPT_DIR=$(dirname "$0")
DIODON_DIR=$(realpath "$SCRIPT_DIR"/..)
IMPL_DIR="$DIODON_DIR"/dh/implementation
ARGOT_DIR=/Users/samarkis/workplace/argot
PROOF_DIR="$DIODON_DIR"/argot-proofs
REPORT_DIR="$PROOF_DIR"/reports
ARGOT_BIN="$ARGOT_DIR"/bin/argot

# Compile Argot tool
cd "$ARGOT_DIR" || exit
echo "Compiling Argot tool in directory $(pwd)"
make argot-build || exit

# Run the immutability analysis
if [ ! -e "$REPORT_DIR" ]; then
    mkdir -p "$REPORT_DIR"
fi

if [ ! -e "$IMPL_DIR" ]; then
    echo Error: "$IMPL_DIR" does not exist
    exit 1
fi
cd "$IMPL_DIR" || exit

IMMUTABILITY_LOG_FILE="$REPORT_DIR"/immutability-log

echo "Running immutability analysis on DH implementation in directory $(pwd)"
echo "Saving log to ${IMMUTABILITY_LOG_FILE}"

"$ARGOT_BIN" immutability -config "$PROOF_DIR"/argot-config-dh.yaml
# >"$IMMUTABILITY_LOG_FILE" 2>&1 # redirect stdout and stderr to log file
