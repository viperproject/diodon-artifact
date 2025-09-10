# Intended Use
This artifact supports the claims of our S&P '26 paper by providing the source code of our static analyses, applying Diodon to two case studies, explaining our claims, providing a Docker image containing all employed tools, and providing scripts to support our claims.
Diodon is in no way limited to just these two case studies and can be applied to other protocols and protocol implementations.
However, doing so requires modeling the protocol in Tamarin, generating the corresponding I/O specification for Gobra, proving the Core using Gobra, adapting the configurations for our static analyses, and running them.
