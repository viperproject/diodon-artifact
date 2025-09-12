# Claim 6 - I/O Independence for Diffie-Hellman
As described in Sec. 5.2, we successfully execute a taint analysis on the Diffie-Hellman codebase to prove that all I/O operations within the APPLICATION are independent of protocol-relevant secrets.

`run.sh` supports this claim by executing the taint analysis. The taint analysis uses the configuration `dh/argot-proofs/argot-config-dh.yaml`, in particular the taint analysis considers taint escaping the current thread.
The configuration under key `dataflow-problems` specifies that the analysis considers implicit flows, that the `parsePrivateKey` and `createNonce` functions are sources of taint, and all functions that are treated as sinks.
Since the taint analysis errors if taint flows from the source to any sink, successful execution of the analysis proves that there are no taint flows, and, because of our configuration that I/O independence holds.
