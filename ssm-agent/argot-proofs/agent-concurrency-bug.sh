#!/bin/bash

# exit when any command fails
set -e

# Script to run the argument alias analysis bug on the SSM agent

PWD=$(dirname "$0")
SCRIPT_DIR=$(realpath "$PWD")
AGENT_DIR="$SCRIPT_DIR"/../implementation
ARGOT_BIN=argot
PATCH=concurrency_leak_app.patch
PATCH2=concurrency_leak_core.patch

if [ ! -e "$AGENT_DIR" ]; then
    echo Error: "$AGENT_DIR" does not exist
    exit 1
fi
cd "$AGENT_DIR" || exit 1

echo "Applying bug patch: $PATCH"
patch -s -p1 < "$PATCH" || exit 1

echo "Running concurrency analysis on SSM agent in directory $(pwd)"

"$ARGOT_BIN" goroutine -config "$SCRIPT_DIR"/argot-config-agent.yaml > /tmp/diodon-bug-test.out || true

if grep -Fq 'Parameter inputData of (*github.com/aws/amazon-ssm-agent/agent/session/datachannel.dataChannel).SendStreamDataMessage has escaped' /tmp/diodon-bug-test.out; then
    echo "Expected result appears"
else
    echo "Expected analysis to fail with specific error message"
    echo "Reverting bug patch: $PATCH"
    patch -sR -p1 < "$PATCH" || exit 1
    exit 1
fi

echo "Reverting bug patch: $PATCH"
patch -sR -p1 < "$PATCH" || exit 1


### Second patch
PATCH=$PATCH2
echo "Applying bug patch: $PATCH"
patch -s -p1 < "$PATCH" || exit 1

echo "Running concurrency analysis on SSM agent in directory $(pwd)"

"$ARGOT_BIN" goroutine -config "$SCRIPT_DIR"/argot-config-agent.yaml > /tmp/diodon-bug-test.out || true

if grep -Fq 'Parameter inputData of (*github.com/aws/amazon-ssm-agent/agent/session/datachannel.dataChannel).SendStreamDataMessage has escaped' /tmp/diodon-bug-test.out; then
    echo "Expected result appears"
else
    echo "Expected analysis to fail with specific error message"
    echo "Reverting bug patch: $PATCH"
    patch -sR -p1 < "$PATCH" || exit 1
    exit 1
fi

echo "Reverting bug patch: $PATCH"
patch -sR -p1 < "$PATCH" || exit 1

echo "Success"
exit 0