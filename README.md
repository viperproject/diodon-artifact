# The Secrets Must Not Flow: Scaling Security Verification to Large Codebases

[![SSM Agent Verification](https://github.com/ArquintL/diodon-artifact/actions/workflows/artifact.yml/badge.svg?branch=main)](https://github.com/ArquintL/diodon-artifact/actions/workflows/artifact.yml?query=branch%3Amain)

This is the artifact for the paper "The Secrets Must Not Flow: Scaling Security Verification to Large Codebases" containing the protocol model, the SSM Agent's codebase, a DH implementation codebase, and the static analysis tools.


## Structure
- `ar-go-tools` contains the static analysis tools (Argot) to analyze the implementations.
- `docker` contains the Dockerfile to build a Docker image containing all tools, the model, and the implementations, allowing seamless verification and, thus, reproduction of our results.

### SSM Agent
- `ssm-agent/model` contains the Tamarin model of the protocol used by the SSM Agent to establish interactive shell sessions
- `ssm-agent/implementation` contains the entire SSM Agent codebase.
    - `ssm-agent/implementation/agent/session/datachannel` contains the Go package representing the CORE.
- `ssm-agent/argot-proofs` contains the scripts used to verify the SSM Agent with Argot.

### Diffie-Hellman (DH) Implementation
- `dh/model` contains the Tamarin model of the protocol used by the DH implementation to perform a DH key exchange
- `dh/implementation` contains the entire DH codebase.
    - `dh/implementation/library` and `dh/implementation/initiator` contain the Go packages representing the CORE.
- `dh/argot-proofs` contains the scripts used to verify the DH implementation with Argot.


## Artifact Docker image
The artifact docker image includes the protocol model and implementation for both case studies. Furthermore, it contains all dependencies to verify the model and implementation.

### Set-up
We require an installation of Docker. The following steps have been tested on macOS 15.1.1 with the latest version of Docker Desktop, which is at time of writing 4.35.1 and comes with version 27.3.1 of the Docker CLI.

#### Installation
- We recommend to adapt the Docker settings to provide sufficient resources to Docker. We have tested our artifact on a 2023 MacBook Pro with a M3 Pro processor running macOS Sequoia 15.1.1 and configured Docker to allocate up 12 cores, 6 GB of memory, and 1 GB of swap memory. In case you are using an ARM-based Mac, enable the option "Use Rosetta for x86/amd64 emulation on Apple Silicon" in the Docker Desktop Settings, which is available on macOS 13 or newer. Measurements on an Apple M1 Pro Silicon have shown that performing this additional emulation results in 20-25\% longer verification times compared to those reported in the remainder of this artifact appendix.
- Navigate to a convenient folder, in which directories can be created for the purpose of running this artifact.
- Open a shell at this folder location.
- Create two new folders named `dh-sync` and `ssm-agent-sync` by executing:
	```
    mkdir dh-sync && mkdir ssm-agent-sync
    ```
- Download and start the Docker image containing our artifact by executing the following command:
    ```
    docker run -it --platform linux/amd64 --volume $PWD/dh-sync:/gobra/dh --volume $PWD/ssm-agent-sync:/gobra/ssm-agent ghcr.io/arquintl/diodon-artifact:latest
    ```
    > ⚠️
    > Note that this command results in the Docker container writing files to the two folders `dh-sync` and `ssm-agent-sync` on your host machine.
    > Thus, make sure that these folders are indeed empty and previous modifications that you have made to files in these folders have been saved elsewhere!
- The Docker command above not only starts a Docker container and provides you with a shell within this container but it also synchronizes all files constituting our artifact with the two folders `dh-sync` and `ssm-agent-sync` on your host machine. I.e., the local folders `dh-sync` and `ssm-agent-sync` are synchronized with `/gobra/dh` and `/gobra/ssm-agent` within the Docker container, respectively.

#### Installation: Finch
If you prefer to use Finch instead of Docker, replace the above Docker command with the following:

``` shell
$ cd diodon-artifact
$ finch vm init
$ finch vm start
$ finch build --file docker/Dockerfile --platform linux/amd64 -t diodon-artifact .
$ finch run -it --volume $PWD/dh-sync:/gobra/dh --volume $PWD/ssm-agent-sync:/gobra/ssm-agent diodon-artifact
$ finch vm stop # stop the vm when you're done using the container
```

#### Usage
The Docker image provides several ready-to-use scripts in the `/gobra` directory:
- `ssm-agent/verify-model.sh`: Verifies the protocol model using Tamarin
- `ssm-agent/verify-core.sh`: Verifies the SSM Agent's CORE using Gobra
- `ssm-agent/verify-io-independence.sh`: Runs the taint analysis and verifies that all I/O operations are I/O independent except those for which we prove using Gobra that they refine the protocol model's SSM Agent role.
- `ssm-agent/verify-core-assumptions.sh`: Runs the pointer and escape analyses to verify that the APPLICATION satisfies most assumptions that the CORE refinement proof makes.

Replace `ssm-agent` with `dh` in the above commands to run the same proofs on the Diffie-Hellman implementation.

#### Proof Script Notes
Some proof scripts may have unexpected output.

- `ssm-agent/argot-proofs/agent-concurrency-proof.sh`: The concurrency proof for the SSM Agent has error output because we expect that the `inputData` parameter of `(*dataChannel).SendStreamDataMessage` escapes to a new thread. Applying the patch file `ssm-agent/implementation/datastream-internal-go-fix.patch` removes the goroutines and results in the proof succeeding.
- `ssm-agent/argot-proofs/agent-passthru-proof.sh` The pass-through analysis for the SSM Agent fails due to false-positives. The [pointer analysis](https://pkg.go.dev/golang.org/x/tools/go/pointer) we use is not context-sensitive (most functions are only analyzed once) so it is impossible to precisely distinguish between Core and App calling contexts for large programs such as the SSM Agent.

#### Bug Patches
Each "proof" script in the `argot-proofs` directory for the SSM Agent and DH implementation has a corresponding "bug" script that apply one or more bug patches and expect the corresponding proof to fail.

The git patch files in the `ssm-agent/implementation` directory introduce bugs to the SSM Agent codebase that our verification tools identify.

- `refinement_bug.patch`: adds an extra I/O operation to send the agent secret over the network which violates the refinement of the security protocol
- `logging_eph_priv_key_bug.patch`: logs the Core instance containing the private key field to show a violation of I/O independence
- `logging_shared_secret_bug.patch`: logs the Core instance containing the shared secret field to show a violation of I/O independence
- `alias_arguments_bug.patch`: leaks a pointer to a Core instance field to the App so it can be passed as an argument to a Core API function, thus resulting in non-disjoint arguments to the Core API function call
- `write_core_state_bug.patch`: leaks a pointer to a core instance field to the App which is then modified in the App, violating the Core instance invariant
- `read_core_state_bug.patch`: leaks a pointer to a Core instance field to the App which is then read in the App, violating the Core instance invariant
- `concurrency_leak_core.patch`: leaks a parameter to a Core API function to a new thread spawned in the Core
- `concurrency_leak_app.patch`: leaks a parameter to a Core API function to a new thread spawned in the App

We also have patches in the `dh/implementation` directory which introduce bugs to the DH codebase.

- `leak_nonce_bug.patch`: leaks nonce data to the App via a getter method which is then written to a network connection, violating I/O independence
- `leak_private_key_bug.patch`: leaks private key data to attacker-accessible method `PerformVirtualInputOperation`, violating I/O independence
- `alias_arguments_bug.patch`: aliases a Core API function parameter to a Core instance field, resulting in non-disjoint arguments to the Core API function call
- `write_core_state_bug.patch`: leaks a pointer to a Core instance field which is then written to in the App, violating the Core instance invariant
- `read_core_state_bug.patch`: leaks a pointer to a Core instance field which is then read in the App, violating the Core instance invariant
- `passthru_escape_in_corealloc_func_bug.patch`: leaks memory allocated in the Core allocation function to the App via a callback, violating the pass-through requirement that all memory allocated in the Core allocation function must only be accessible to the App via the function's return value
- `passthru_escape_in_coreapi_func_bug.patch`: leak permissions to access a slice allocated in the Core via a global variable which is then accessed in the App, resulting in accessing memory allocated in the Core which does not pass through the Core API function's return value
