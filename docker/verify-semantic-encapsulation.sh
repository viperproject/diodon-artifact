#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")

# Verify that the core invariant is maintained by proving the absence of writes to core instances
/bin/bash "$SCRIPT_DIR/argot-proofs/invariant-proof.sh"

# Verify that no arguments to core functions alias each other
/bin/bash "$SCRIPT_DIR/argot-proofs/argument-alias-proof.sh"
