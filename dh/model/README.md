# Tamarin model for Diffie-Hellman

## Files
- The `dh.spthy` file contains the Tamarin model, including the security properties and the auxiliary lemmas.


## Prerequisite
To verify our model of DH, you need **Tamarin**
which can be obtained from [its website](https://tamarin-prover.github.io).
Version 1.8.0 is known to work.


## Instructions
To verify the model with Tamarin, use the following command:

`tamarin-prover --prove dh.spthy --derivcheck-timeout=0`
