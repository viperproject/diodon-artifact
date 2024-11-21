#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")

time tamarin-prover --prove --derivcheck-timeout=0 "$SCRIPT_DIR"/model/protocol-model.spthy 2> /dev/null
