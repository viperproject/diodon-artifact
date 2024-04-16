#!/bin/bash

cwd=$(pwd)
cd /Users/arquintlinard/ETH/PhD/tamigloo-compiler/tamarin-prover || exit
stack exec -- tamarin-prover --tamigloo-compiler "$cwd"/secure-sessions-gobra-config.txt
cd "$cwd" || exit
