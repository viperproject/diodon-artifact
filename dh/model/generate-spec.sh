#!/bin/bash

cwd=$(pwd)
cd /Users/arquintlinard/ETH/PhD/tamigloo-compiler/tamarin-prover || exit
echo $PWD
stack build
stack exec -- tamarin-prover --tamigloo-compiler "$cwd"/generate-spec-config.txt
cd "$cwd" || exit
