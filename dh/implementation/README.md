# Go Diffie-Hellman Implementation

## Implementation Overview
`main.go` represents the App that performs all I/O.
The Core is implemented in `initiator/initiator.go`.
More specifically, the `Initiator` struct stores all state related to a protocol session.
In addition, this file provides methods with the `Initiator` struct being the receiver to produce outgoing and consume incoming messages.
We prove (using Gobra) that producing and consuming these messages correspond to steps in the protocol model.
This allows us to treat outgoing messages as being untainted from a taint analysis point of view, such that the corresponding I/O operation in the App satisfies I/O independence.
For this purpose, `library/io.go` provides a function `PerformVirtualOutputOperation` that enforces (via its specification) that a caller gives up an I/O permission for sending a message.
Thus, we configure the taint analysis (in `../argot-proofs/argot-config-dh.yaml`) to treat this function as a sanitizer, i.e., that this returns untainted data.

Protocol steps in the Core are easy to locate as each step requires justification by the I/O specification. I.e., `unfold io.P_Alice(...)` (`unfold iospec.P_Agent(...)` in the SSM Agent case study) applies the I/O specification to obtain the I/O permission for the subsequent operation such as sending or receiving a message or performing an internal operation. These internal operations directly correspond to a transition in the Tamarin model.

`dh/implementation/initiator/initiator.go` provides detailed comments explaining the application of the I/O specification for the `ProduceHsMsg1` method in the DH case study.


## Building & Running the Initiator Role
Build:
```
go build
```

Run:
```
./dh-gobra --isInitiator --endpoint localhost:12345 --privateKey "ACvCw0fb1mqTQikmXOas+YEbJnC9O/N4H12k4w/ADVRVg6YAptHsQO57FNzeeS2BtGwHas51wRruj62+y4WpjQ==" --peerPublicKey "H4omvaajENeqxbRiOVCLZoGUrEWIVrVAtJPk5JgoEV8="
```

The command above configures the Initiator to use Keypair 1 for its own secret key and tries to communicate with the Responder by using Keypair 2's public key.

Note that the Responder has to be started first.


Keypair 1:
    - sk: "ACvCw0fb1mqTQikmXOas+YEbJnC9O/N4H12k4w/ADVRVg6YAptHsQO57FNzeeS2BtGwHas51wRruj62+y4WpjQ=="
    - pk: "VYOmAKbR7EDuexTc3nktgbRsB2rOdcEa7o+tvsuFqY0="

Keypair 2:
    - sk: "k2gUarJExuxjji+KlwD8NfclZ+ZCZ8xZk3NGzN3ypwgfiia9pqMQ16rFtGI5UItmgZSsRYhWtUC0k+TkmCgRXw=="
    - pk: "H4omvaajENeqxbRiOVCLZoGUrEWIVrVAtJPk5JgoEV8="

## Verifying the Initiator Role
Gobra can be used as follows to successfully verify this implementation.
The Docker image provides a Gobra JAR and a shell script to do so.
Alternatively, the following command manually configures Gobra:
```
java -Xss128m -jar <path to gobra.jar> --recursive -I ./ -I .verification --module dh-gobra --includePackages initiator --parallelizeBranches
```
