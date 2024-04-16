#!/bin/bash

# exit when any command fails
set -e

SCRIPT_DIR=$(dirname "$0")
AGENT_DIR="$SCRIPT_DIR/implementation"
DATACHANNEL_DIR="$AGENT_DIR/agent/session/datachannel"
GOBRA_JAR="/gobra/gobra.jar"
GOBRA_REPORT_DIR="$AGENT_DIR/.gobra"

mkdir -p "$GOBRA_REPORT_DIR"

INPUT_FILES="\
    $DATACHANNEL_DIR/datachannel.go\
    $DATACHANNEL_DIR/handshake_complete.go\
    $DATACHANNEL_DIR/handshake_request.go\
    $DATACHANNEL_DIR/handshake_response.go\
    $DATACHANNEL_DIR/helper.go\
    $DATACHANNEL_DIR/init.go\
    $DATACHANNEL_DIR/recv.go\
    $DATACHANNEL_DIR/send_recv_channel.go\
    $DATACHANNEL_DIR/send.go\
    $DATACHANNEL_DIR/state.go"

java -Xss128m -jar $GOBRA_JAR \
    --module "github.com/aws/amazon-ssm-agent" \
    --include "$AGENT_DIR/.verification" --include "$AGENT_DIR" \
    --input "$INPUT_FILES" \
    --gobraDirectory "$GOBRA_REPORT_DIR" \
    --parseAndTypeCheckMode PARALLEL \
    --parallelizeBranches
