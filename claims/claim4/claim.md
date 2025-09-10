# Claim 4 -- Analyzing the Application for SSM Agent
As described in Sec. 5.1.4, we implemented static analyses checking conditions (C1)--(C4) and (C6)--(C8) (as stated in Sec. 4.3).
The conditions (C1)--(C8) ensure that callers of the CORE (which is verified by Gobra, cf. claim 3) satisfy the preconditions of CORE methods.
In particular, we implemented a pass-through analysis to check conditions (C1), (C3), and (C8) and an escape analysis for conditions (C4) and (C6).
Furthermore, we use a pointer analysis to check condition (C7) and an immutability analysis based on the pointer analysis to check condition (C2).
As mentioned in Sec. 5.1.4, our escape analysis fails as it detects that a CORE instance and a thread-safe logger object escapes the thread in which it got created, and our pass-through analysis is a prototype reporting many false positives.

`run.sh` supports this claim by running these analyses.
In line with the claims stated in the paper, we ignore the escaping CORE instance and logger object in our escape analysis, and the false positives of the pass-through analysis.
The other analyses, i.e., pointer and immutability analyses checking (C7) and (C2), succeed.

In more detail, `run.sh` invokes `docker/ssm-agent/verify-core-assumptions.sh`, which internally invokes `ssm-agent/argot-proofs/agent-concurrency-proof-fix.sh` that executes the escape analysis and checks that only objects mentioned above escape.
