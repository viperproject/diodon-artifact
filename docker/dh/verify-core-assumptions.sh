#!/bin/bash

# exit when any command fails
set -e

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")

# Verify that the core invariant is maintained by proving the absence of writes to core instances
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-invariant-proof.sh"

# Verify that the core invariant is maintained when all pointers written to in
# the app that are allocated in the core pass through one of the core api
# function's return parameters.
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-passthru-proof.sh"

# Verify that each core instance does not escape the goroutine in which it is created
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-concurrency-proof.sh"

# Verify that no arguments to core functions alias each other
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-argument-alias-proof.sh"
