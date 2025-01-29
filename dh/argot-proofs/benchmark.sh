#!/bin/bash

# exit when any command fails
set -e

for i in {1..10}
do
    echo "Run $i"
    time ./dh-taint-proof.sh &> /dev/null
done
