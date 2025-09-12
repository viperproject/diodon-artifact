# Claim 2 - I/O Independence for SSM Agent
As described in Sec. 5.1.2, we successfully execute a taint analysis on the entire SSM Agent codebase if we ignore taint escaping the current thread.
We use a taint analysis to establish I/O independence, which states that all I/O operations within the APPLICATION are independent of protocol-relevant secrets.

`run.sh` supports this claim by executing the taint analysis. The taint analysis uses the configuration `ssm-agent/argot-proofs/argot-config-agent.yaml`.
`use-escape-analysis: false` instructs the taint analysis to ignore taint escaping the current thread.
The configuration under key `dataflow-problems` specifies that the analysis considers implicit flows, that the `GenerateKey` function in the `crypto/elliptic` package is a source of taint, and all functions that are treated as sinks.
The taint analysis errors if a flow of taint from the source to any sink is detected.
We use the taint analysis for checking I/O independence by configuring protocol secrets as sources of taint and all I/O operations in the APPLICATION and protocol-irrelevant I/O operations in the CORE as sinks.
