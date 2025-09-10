#!/bin/bash
set -e

IMAGE_NAME="ghcr.io/viperproject/diodon-artifact:latest"

docker build --platform linux/amd64 -t "$IMAGE_NAME" -f docker/Dockerfile .
