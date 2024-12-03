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
The artifact docker image includes both the protocol model and SSM Agent implementation. Furthermore, it contains all dependencies to verify the model and implementation.

### Set-up
We require an installation of Docker. The following steps have been tested on macOS 14.0 with the latest version of Docker Desktop, which is at time of writing 4.24.2 and comes with version 24.0.6 of the Docker CLI.

#### Installation
- We recommend to adapt the Docker settings to provide sufficient resources to Docker. We have tested our artifact on a 2019 16-inch MacBook Pro with 2.3 GHz 8-Core Intel Core i9 running macOS Sonoma 14.0 and configured Docker to allocate up 16 cores (which includes 8 virtual cores), 6 GB of memory, and 1 GB of swap memory. In case you are using an ARM-based Mac, enable the option "Use Rosetta for x86/amd64 emulation on Apple Silicon" in the Docker Desktop Settings, which is available on macOS 13 or newer. Measurements on an Apple M1 Pro Silicon have shown that performing this additional emulation results in 20-25\% longer verification times compared to those reported in the remainder of this artifact appendix.
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
    > Note that this command results in the Docker container writing files to the two folders `model-sync` and `implementation-sync` on your host machine.
    > Thus, make sure that these folders are indeed empty and previous modifications that you have made to files in these folders have been saved elsewhere!
- The Docker command above not only starts a Docker container and provides you with a shell within this container but it also synchronizes all files constituting our artifact with the two folders `model-sync` and `implementation-sync` on your host machine. I.e., the local folders `model-sync` and `implementation-sync` are synchronized with `/gobra/model` and `/gobra/implementation` within the Docker container, respectively.

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

#### Bug Patches
The git patch files in the `ssm-agent/implementation` directory introduce bugs to the SSM Agent codebase that our verification tools identify.

- `refinement_bug.patch`: caught by `verify-core.sh`
- `logging_eph_priv_key_bug.patch`: caught by `verify-io-independence.sh`
- `logging_shared_secret_bug.patch`: caught by `verify-io-independence.sh`
- `alias_arguments_bug.patch`: caught by `verify-core-assumptions.sh`
- `modify_core_state_bug.patch`: caught by `verify-core-assumptions.sh`

TODO make bug patches for DH
