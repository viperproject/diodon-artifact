#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

echo "Verifying I/O independence for the SSM Agent case study."
echo "This might take about 1.5min. Please wait..."

docker run --platform linux/amd64 --rm --entrypoint "/bin/bash" $IMAGE_NAME -c "cp -r /dh-orig/. dh/; cp -r /ssm-agent-orig/. ssm-agent/; /gobra/ssm-agent/verify-io-independence.sh"
