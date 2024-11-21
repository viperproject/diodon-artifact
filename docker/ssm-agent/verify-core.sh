#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")
IMPLEMENTATION_DIR="$SCRIPT_DIR/implementation"
DATACHANNEL_DIR="$IMPLEMENTATION_DIR/agent/session/datachannel"
GOBRA_JAR="/gobra/gobra.jar"
GOBRA_REPORT_DIR="$IMPLEMENTATION_DIR/.gobra"

mkdir -p "$GOBRA_REPORT_DIR"

java -Xss128m -jar "$GOBRA_JAR" \
    --module "github.com/aws/amazon-ssm-agent" \
    --include "$IMPLEMENTATION_DIR/.verification" --include "$IMPLEMENTATION_DIR" \
    --input "$DATACHANNEL_DIR/datachannel.go" \
    --input "$DATACHANNEL_DIR/handshake_complete.go" \
    --input "$DATACHANNEL_DIR/handshake_request.go" \
    --input "$DATACHANNEL_DIR/handshake_response.go" \
    --input "$DATACHANNEL_DIR/helper.go" \
    --input "$DATACHANNEL_DIR/init.go" \
    --input "$DATACHANNEL_DIR/recv.go" \
    --input "$DATACHANNEL_DIR/send_recv_channel.go" \
    --input "$DATACHANNEL_DIR/send.go" \
    --input "$DATACHANNEL_DIR/state.go" \
    --gobraDirectory "$GOBRA_REPORT_DIR" \
    --parseAndTypeCheckMode PARALLEL \
    --parallelizeBranches
