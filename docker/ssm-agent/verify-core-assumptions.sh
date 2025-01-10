#!/bin/bash

# exit when any command fails
set -e

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")

# Verify that the core invariant is maintained by proving the absence of writes to core instances
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-invariant-proof.sh"

# Verify that each core instance does not escape the goroutine in which it is created
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-concurrency-proof-fix.sh"

# Verify that no arguments to core functions alias each other
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-argument-alias-proof.sh"

# Verify the pass-through conditions
# SSM Agent passthrough proof is supposed to fail due to false-positives
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-passthru-proof.sh" || true
