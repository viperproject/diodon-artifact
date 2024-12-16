#!/bin/bash

# exit when any command fails
set -e

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")

# Verify that the core invariant is not maintained when a core instance is written to or read from
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-invariant-bug.sh"

# TODO Verify that a core instance escapes the goroutine in which it is created
# /bin/bash "$SCRIPT_DIR/argot-proofs/dh-concurrency-bug.sh"

# Verify that arguments to a core function alias each other
/bin/bash "$SCRIPT_DIR/argot-proofs/dh-argument-alias-bug.sh"
