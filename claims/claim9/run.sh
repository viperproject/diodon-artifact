#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

echo "Seeding bugs and checking that they are detected for the Diffie-Hellman and SSM Agent case studies."
echo "This might take about 10min. Please wait..."

docker run --platform linux/amd64 --rm --entrypoint "/bin/bash" $IMAGE_NAME -c "cp -r /dh-orig/. dh/; cp -r /ssm-agent-orig/. ssm-agent/; /gobra/ssm-agent/verify-io-independence-bug.sh && /gobra/ssm-agent/verify-core-assumptions-bug.sh && /gobra/dh/verify-core-bug.sh && /gobra/dh/verify-io-independence-bug.sh && /gobra/dh/verify-core-assumptions-bug.sh && echo 'All bugs were successfully detected.'"
