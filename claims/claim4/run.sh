#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

echo "Analyzing the Application for the SSM Agent case study."
echo "This might take about 4min. Please wait..."

docker run --platform linux/amd64 --rm --entrypoint "/bin/bash" $IMAGE_NAME -c "cp -r /dh-orig/. dh/; cp -r /ssm-agent-orig/. ssm-agent/; time /gobra/ssm-agent/verify-core-assumptions.sh && echo 'Conditions were checked successfully modulo the stated exceptions.'"
