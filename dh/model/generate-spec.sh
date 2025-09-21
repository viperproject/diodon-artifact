#!/bin/bash

cwd=$(pwd)
cd /Users/arquintlinard/ETH/PhD/protocol-verification-refinement/specification-generator/src || exit
echo $PWD
stack build
stack exec -- tamarin-prover --tamigloo-compiler "$cwd"/generate-spec-config.txt
cd "$cwd" || exit
