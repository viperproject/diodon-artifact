#!/bin/bash

# exit when any command fails
set -e

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")

# Verify that the core invariant not maintained by writing to and reading from core instances
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-invariant-bug.sh"

# TODO Verify that a core instance escapes the goroutine in which it is created
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-concurrency-bug.sh"

# Verify that arguments to a core function alias each other
/bin/bash "$SCRIPT_DIR/argot-proofs/agent-argument-alias-bug.sh"
