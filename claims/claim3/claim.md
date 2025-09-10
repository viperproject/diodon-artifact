# Claim 3 - Core Refinement for SSM Agent
As described in Sec. 5.1.3, we identified the `datachannel` package as the SSM Agent's CORE and successfully verify this package using the Gobra program verifier.

`run.sh` supports this claim by running Gobra to verify the `datachannel` package. This proves that all protocol-relevant I/O operations within the CORE are permitted by the protocol model that we verified in Tamarin and that we have a correct proof in separation logic for the verified methods in this package.
