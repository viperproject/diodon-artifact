#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")
/bin/bash "$SCRIPT_DIR/argot-proofs/taint-proof.sh"
