#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

echo "Verifying Core Refinement for the Diffie-Hellman case study."
echo "This might take about 1min. Please wait..."

docker run --platform linux/amd64 --rm --entrypoint "/bin/bash" $IMAGE_NAME -c "cp -r /dh-orig/. dh/; cp -r /ssm-agent-orig/. ssm-agent/; /gobra/dh/verify-core.sh"
