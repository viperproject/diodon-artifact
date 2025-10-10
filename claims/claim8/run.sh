#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

echo "Analyzing the Application for the Diffie-Hellman case study."
echo "This might take about 40s. Please wait..."

docker run --platform linux/amd64 --rm --entrypoint "/bin/bash" $IMAGE_NAME -c "cp -r /dh-orig/. dh/; cp -r /ssm-agent-orig/. ssm-agent/; time ( /gobra/dh/verify-core-assumptions.sh && echo 'Conditions were checked successfully.' )"
