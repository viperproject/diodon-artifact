# Tamarin model for Diffie-Hellman

## Files
- The `dh.spthy` file contains the Tamarin model, including the security properties and the auxiliary lemmas.


## Prerequisite
To verify our model of DH, you need **Tamarin**
which can be obtained from [its website](https://tamarin-prover.github.io).
Version 1.10.0 is known to work.


## Instructions
To verify the model with Tamarin, use the following command:

`tamarin-prover --prove dh.spthy --derivcheck-timeout=0`


## Generate I/O specification
Adapt the absolute paths in `generate-spec.sh` and `generate-spec-config.txt` to point to the files in this repository and the specification generator from the [`viperproject/protocol-verification-refinement` repository](https://github.com/viperproject/protocol-verification-refinement) before running `generate-spec.sh`.
The resulting files will be stored in the `generated_iospecs` directory.
