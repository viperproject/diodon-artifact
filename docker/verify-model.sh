#!/bin/bash

time tamarin-prover --prove --derivcheck-timeout=0 model/protocol-model.spthy 2> /dev/null 
