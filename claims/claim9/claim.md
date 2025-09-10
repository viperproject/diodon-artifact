# Claim 9 -- Soundness
As described in Sec. 5.3, we deliberately introduce bugs in our case studies and check that our tools fail as expected.

`run.sh` supports this claim by seeding bugs in
- the SSM Agent case study:
    - 2 bugs violating I/O independence
        - logging a shared secret
        - logging an ephemeral secret key
    - 2 bugs violating conditions (C2) and (C8)
        - writing to a CORE instance's internal memory in the APPLICATION
        - reading from a CORE instance's internal memory in the APPLICATION
    - 2 bugs violating condition (C6)
        - letting a parameter escape the current thread in the CORE
        - letting a parameter escape the current thread in the APPLICATION
    - 1 bug violating condition (C7)
        - using a parameter that alias a CORE instance's internal memory
- the Diffie-Hellman case study:
    - 1 bug violating CORE refinement
        - including the DH secret key instead of public key in a message in the CORE violates the protocol model
    - 2 bugs violating I/O independence
        - adding a part of the secret key in the APPLICATION to a message violates I/O independence
        - sending the DH secret key in the APPLICATION violates I/O independence
    - 2 bugs violating conditions (C2) and (C8)
        - writing to a CORE instance's internal memory in the APPLICATION
        - reading from a CORE instance's internal memory in the APPLICATION
    - 2 bugs violating conditions (C1) and (C8)
        - passing a CORE instance to the APPLICATION for which the postcondition does not guarantee that the CORE invariant holds
        - passing CORE-controlled memory via a global variable to the APPLICATION for which the postcondition does not guarantee full permissions
    - 1 bug violating condition (C7)
        - using a parameter that alias a CORE instance's internal memory
