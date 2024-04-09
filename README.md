# DIODON: Scaling Verification of Security Protocol Implementations to Large Codebases

[![SSM Agent Verification](https://github.com/ArquintL/diodon-artifact/actions/workflows/artifact.yml/badge.svg?branch=main)](https://github.com/ArquintL/diodon-artifact/actions/workflows/artifact.yml?query=branch%3Amain)

This is the artifact for the paper "DIODON: Scaling Verification of Security Protocol Implementations to Large Codebases" containing the protocol model and codebase of the SSM Agent.


## Structure
- `model` contains the Tamarin model of the protocol used by the SSM Agent to establish interactive shell sessions
- `implementation` contains the entire SSM Agent codebase.
    - `implementation/agent/session/datachannel` contains the Go package representing the CORE.
- `ar-go-tools` contains the static analysis tools (Argot) to analyze the SSM Agent.
- `argot-proofs` contains the scripts used to verify the SSM Agent with Argot.
- `docker` contains the Dockerfile to build a Docker image containing all tools, the model, and the implementation, allowing seamless verification and, thus, reproduction of our results.


## Artifact Docker image
The artifact docker image includes both the protocol model and SSM Agent implementation. Furthermore, it contains all dependencies to verify the model and implementation.

### Set-up
We require an installation of Docker. The following steps have been tested on macOS 14.0 with the latest version of Docker Desktop, which is at time of writing 4.24.2 and comes with version 24.0.6 of the Docker CLI.

#### Installation
- We recommend to adapt the Docker settings to provide sufficient resources to Docker. We have tested our artifact on a 2019 16-inch MacBook Pro with 2.3 GHz 8-Core Intel Core i9 running macOS Sonoma 14.0 and configured Docker to allocate up 16 cores (which includes 8 virtual cores), 6 GB of memory, and 1 GB of swap memory. In case you are using an ARM-based Mac, enable the option "Use Rosetta for x86/amd64 emulation on Apple Silicon" in the Docker Desktop Settings, which is available on macOS 13 or newer. Measurements on an Apple M1 Pro Silicon have shown that performing this additional emulation results in 20-25\% longer verification times compared to those reported in the remainder of this artifact appendix.
- Navigate to a convenient folder, in which directories can be created for the purpose of running this artifact.
- Open a shell at this folder location.
- Create two new folders named `model-sync` and `implementation-sync` by executing:
	```
    mkdir model-sync && mkdir implementation-sync
    ```
- Download and start the Docker image containing our artifact by executing the following command:
    ```
    docker run -it --platform linux/amd64 --volume $PWD/model-sync:/gobra/model --volume $PWD/implementation-sync:/gobra/implementation ghcr.io/arquintl/diodon-artifact:latest
    ```
    > ⚠️
    > Note that this command results in the Docker container writing files to the two folders `model-sync` and `implementation-sync` on your host machine.
    > Thus, make sure that these folders are indeed empty and previous modifications that you have made to files in these folders have been saved elsewhere!
- The Docker command above not only starts a Docker container and provides you with a shell within this container but it also synchronizes all files constituting our artifact with the two folders `model-sync` and `implementation-sync` on your host machine. I.e., the local folders `model-sync` and `implementation-sync` are synchronized with `/gobra/model` and `/gobra/implementation` within the Docker container, respectively.

#### Usage
The Docker image provides several ready-to-use scripts in the `/gobra` directory:
- `verify-model.sh`: Verifies the protocol model using Tamarin
- `verify-core.sh`: Verifies the SSM Agent's CORE using Gobra
