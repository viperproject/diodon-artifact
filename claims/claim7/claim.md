# Claim 7 -- Core Refinement for Diffie-Hellman
As described in Sec. 5.2, we successfully verify the CORE (package `initiator`) the using the Gobra program verifier.

`run.sh` supports this claim by running Gobra to verify the `initiator` package. This proves that all protocol-relevant I/O operations within the CORE are permitted by the protocol model that we verified in Tamarin and that we have a correct proof in separation logic for the verified methods in this package.
