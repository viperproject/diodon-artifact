#!/bin/bash

# exit when any command fails
set -e

for i in {1..10}
do
    echo "Run $i"
    time tamarin-prover --prove --derivcheck-timeout=0 protocol-model.spthy &> /dev/null
done
