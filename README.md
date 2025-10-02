# The Secrets Must Not Flow: Scaling Security Verification to Large Codebases

[![Artifact Image](https://github.com/viperproject/diodon-artifact/actions/workflows/artifact.yml/badge.svg?branch=main)](https://github.com/viperproject/diodon-artifact/actions/workflows/artifact.yml?query=branch%3Amain)
[![Artifact Claims](https://github.com/viperproject/diodon-artifact/actions/workflows/artifact-evaluation.yml/badge.svg?branch=main)](https://github.com/viperproject/diodon-artifact/actions/workflows/artifact-evaluation.yml?query=branch%3Amain)

This is the artifact for the paper "The Secrets Must Not Flow: Scaling Security Verification to Large Codebases", which will appear at the IEEE Symposium on Security and Privacy (S&P), 2026. This repository contains the protocol model, the forked SSM Agent's codebase, a DH implementation codebase, and the static analysis tools.

> **Note to artifact reviewers:**
> We will publish this repository (including all submodules) after artifact evaluation on [Zenodo](https://zenodo.org).
> Section [Artifact Evaluation](#artifact-evaluation) describes the initial setup and provides pointers to our claims.
> You can also find this artifact on GitHub, for, e.g., rendering of our Markdown files: [`viperproject/diodon-artifact`](https://github.com/viperproject/diodon-artifact).


## Initialization
The `ar-go-tools` and `ssm-agent` directories are git submodules, which must be initialized before proceeding if you cloned this repository using Git.
If you obtained the artifact from a tarball (e.g., via Zenodo), then you can skip this step.

``` shell
git submodule update --init --recursive -- ar-go-tools ssm-agent
```


## Structure
- `ar-go-tools` contains the static analysis tools (Argot) to analyze the implementations.
- `docker` contains the Dockerfile to build a Docker image containing all tools, the model, and the implementations, allowing seamless verification and, thus, reproduction of our results.


### SSM Agent
Whenever we refer to the SSM Agent, we mean the fork of the SSM Agent codebase that implements the protocol for establishing encrypted interactive shell sessions.

- `ssm-agent/model` contains the Tamarin model of the protocol used by the SSM Agent to establish interactive shell sessions
- `ssm-agent/implementation` contains the entire SSM Agent codebase and the corresponding [README](https://github.com/ArquintL/amazon-ssm-agent/tree/update-diodon?tab=readme-ov-file#secure-sessions-overview) provides an overview.
    - `ssm-agent/implementation/agent/session/datachannel` contains the Go package representing the CORE.
- `ssm-agent/argot-proofs` contains the scripts used to verify the SSM Agent with Argot.


### Diffie-Hellman (DH) Implementation
- `dh/model` contains the Tamarin model of the protocol used by the DH implementation to perform a DH key exchange
- `dh/implementation` contains the entire DH codebase and the corresponding [README](dh/implementation/README.md) provides an overview over the codebase.
    - `dh/implementation/library` and `dh/implementation/initiator` contain the Go packages representing the CORE.
- `dh/argot-proofs` contains the scripts used to verify the DH implementation with Argot.


## Artifact Docker Image
This repository builds and provides a docker image that includes the protocol model and implementation for both case studies. Furthermore, it contains all dependencies to verify the model and implementation.


### Set-up
We require an installation of Docker. The following steps have been tested on macOS 15.6.1 with the latest version of Docker Desktop, which is at time of writing 4.45.0 and comes with version 28.3.3 of the Docker CLI.

We recommend adapting the Docker settings to provide sufficient resources to Docker. We have tested our artifact on a 2023 MacBook Pro with a M3 Pro processor running macOS Sequoia 15.6.1 and configured Docker to allocate up 12 cores, 6 GB of memory, and 1 GB of swap memory. In case you are using an ARM-based Mac, enable the option "Use Rosetta for x86/amd64 emulation on Apple Silicon" in the Docker Desktop Settings, which is available on macOS 13 or newer.

Continuous integration of this repository builds a ready-to-use Docker image labeled `ghcr.io/viperproject/diodon-artifact:latest`. Alternatively, the `install.sh` script builds a Docker image with the same label locally.


### Artifact Evaluation
The `claims` folder contains for each claim of our paper a description of the claim, a script for running an experiment supporting such a claim, and the experiment's expected output.
Each script automatically starts and stops a Docker container based on the Docker image mentioned above. If this image is not available locally (by running `install.sh`, the image will automatically be downloaded.)

Note that the execution times in the paper's Fig. 9 and in Sec. 5 were obtained by running the tools natively (i.e., without using a Docker image and virtualization) on an Apple M3 Pro processor. Each claim's `run.sh` script mentions the execution time that we observed on the very same hardware but using our Docker image.

More specifically, our claims are as follows:
- [Claim 1](claims/claim1/claim.md): The Protocol Model for the SSM Agent satisfies secrecy & injective agreement
- [Claim 2](claims/claim2/claim.md): The SSM Agent case study satisfies I/O independence
- [Claim 3](claims/claim3/claim.md): The Gobra program verifier successfully verifies the CORE in the SSM Agent case study
- [Claim 4](claims/claim4/claim.md): We analyse the APPLICATION in the SSM Agent case study as stated in the paper
- [Claim 5](claims/claim5/claim.md): The Protocol Model for the signed Diffie-Hellman key exchange satisfies forward secrecy & injective agreement
- [Claim 6](claims/claim6/claim.md): The Diffie-Hellman case study satisfies I/O independence
- [Claim 7](claims/claim7/claim.md): The Gobra program verifier successfully verifies the CORE in the Diffie-Hellman case study
- [Claim 8](claims/claim8/claim.md): The APPLICATION in the Diffie-Hellman case study satisfies conditions (C1)-(C4) and (C6)-(C8)
- [Claim 9](claims/claim9/claim.md): Our tools detect deliberately seeded bugs in both case studies

Alternatively, we describe next how to manually run the Docker image and interact with the Docker container.


### Directly Using the Docker Image
> Note that the interactions described in the following are subsumed by our claims such that the rest of this README can be ignored for artifact evaluation.
- Navigate to a convenient folder, in which directories can be created for the purpose of running this artifact.
- Open a shell at this folder location.
- Create two new folders named `dh-sync` and `ssm-agent-sync` by executing:
	```
    mkdir dh-sync && mkdir ssm-agent-sync
    ```
- Download and start the Docker image containing our artifact by executing the following command (alternatively, `install.sh` builds a Docker image locally that can be run in the same way):
    ```
    docker run -it --platform linux/amd64 --volume $PWD/dh-sync:/gobra/dh --volume $PWD/ssm-agent-sync:/gobra/ssm-agent ghcr.io/viperproject/diodon-artifact:latest
    ```
    > ⚠️
    > Note that this command results in the Docker container writing files to the two folders `dh-sync` and `ssm-agent-sync` on your host machine.
    > Thus, make sure that these folders are indeed empty and previous modifications that you have made to files in these folders have been saved elsewhere!
- The Docker command above not only starts a Docker container and provides you with a shell within this container, but it also synchronizes all files constituting our artifact with the two folders `dh-sync` and `ssm-agent-sync` on your host machine. I.e., the local folders `dh-sync` and `ssm-agent-sync` are synchronized with `/gobra/dh` and `/gobra/ssm-agent` within the Docker container, respectively.


#### Directly Using the Docker Image: Finch
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
- `alias_arguments_bug.patch`: leaks a pointer to a Core instance field to the App, so it can be passed as an argument to a Core API function, thus resulting in non-disjoint arguments to the Core API function call
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


## Argot extensions for Diodon
We instantiate Diodon for Go by extending the Go static analysis tool [Argot](https://github.com/awslabs/ar-go-tools).
In this subsection, we describe which static analyses we use from or added to Argot, grouped as introduced in the paper, and end with a short overview over Argot's API and configuration.


### I/O independence
As described in the Sec. 4.1 in the paper, we check I/O independence by running a taint analysis that considers implicit taint flows and checks that taint does not escape a thread.
Argot provides such a taint analysis, and we did not make any major changes for Diodon.
See the [taint analysis documentation](https://github.com/awslabs/ar-go-tools/blob/dffa660a6c22232f878d8d9dd5e998f5f52cdbf6/doc/01_taint.md) for more details.


### Analyzing the App
We analyze the App using a combination of static analyses to ensure that the Core's refinement proof is valid (Sec. 4.3 in the paper).
In the following, we describe these analyses and how they relate to the conditions (C1) - (C8) in the paper's Fig. 8.


####  Escape analysis
While Argot's escape analysis was originally designed to be used by the taint analysis, we added a stand-alone `argot escape` command which runs the escape analysis by itself.
We modified the escape analysis to improve context sensitivity, and the escape analysis implementation is located in the `ar-go-tools/analysis/escape` directory.

Diodon uses the escape analysis to check:
- Condition (C4): Core instances are used only in the thread they are created in.
- Condition (C6): Parameters to Core APIs are local.


#### Pointer analysis
The Go standard library provides a [first-party Go pointer analysis](https://pkg.go.dev/golang.org/x/tools/go/pointer) that Argot uses.
This pointer analysis is vendored directly in the codebase in the directory `ar-go-tools/internal/pointer`.
Argot uses the may-alias information computed by this pointer analysis for the taint and escape analyses.
For this purpose and to improve performance, the pointer analysis was modified to expose some pointer analysis internals and to cache may-alias results.
We use this pointer analysis to implement the below custom analyses for Diodon.


##### Pass-through analysis
We implemented the pass-through analysis from scratch by using the pointer analysis' alias information and the SSA program representation.
The pass-through (called `passthru` for short) analysis implementation is located in the `ar-go-tools/analysis/passthru` directory.

Diodon uses the pass-through analysis to check:
- Condition (C1): Core instances are created in a function ensuring the Core invariant in its postcondition.
- Condition (C3): Core instances are passed only to Core functions that preserve the invariant
- Condition (C8): Reads and writes in the Application occur to memory allocated in the Application or transferred from the Core.


#### Immutability analysis
We implemented an immutability analysis from scratch, which records all reads, writes, and allocations of a given SSA value using the results from the pointer analysis.
Note that we do not consider this to be a separate analysis in the paper and call it a "pointer analysis".
The immutability analysis implementation is located in the `ar-go-tools/analysis/immutability` directory.

Diodon uses the immutability analysis to check condition (C2): The Application does not write to Core instances' internal state, even through an alias.


#### Alias analysis
We implemented an alias analysis from scratch, which computes all of the SSA values that may alias a given SSA value using the results from the pointer analysis.
Note that we do not consider this to be a separate analysis in the paper and call it a "pointer analysis".
The alias analysis implementation is located in the `ar-go-tools/analysis/alias` directory.

Diodon uses the alias analysis to check condition (C7): Parameters to the same Core API call do not alias one another.


### Argot API & documentation
To see the documentation and public API for any of our analyses, use the `go doc` command with the analysis' package name. E.g.:
``` shell
cd ar-go-tools
go doc github.com/awslabs/ar-go-tools/analysis/passthru
```

In addition, the `ar-go-tools/doc` directory contains detailed descriptions of the analyses that Argot provides.


### Argot configuration
We configured Argot to run the above analyses on both our case studies, namely the signed DH key exchange and the forked SSM Agent implementations.
The configuration files are located in `dh/argot-proofs/argot-config-dh.yaml` and `ssm-agent/argot-proofs/argot-config-agent.yaml`, respectively.
The escape analysis is configured via `dh/argot-proofs/escape-config.json` and `ssm-agent/argot-proofs/escape-config.json`, respectively.
