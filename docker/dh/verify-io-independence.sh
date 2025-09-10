#!/bin/bash

# exit when any command fails
set -e

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")

time /bin/bash "$SCRIPT_DIR/argot-proofs/dh-taint-proof.sh"
