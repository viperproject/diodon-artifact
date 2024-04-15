# Gobra commit 88e0a9b
FROM ghcr.io/viperproject/gobra@sha256:5f81e650fa8e0039e4ef64d8b211e7d8b1814cdbbebf10d22106300b7c28a418
# FROM ghcr.io/viperproject/gobra:latest

RUN apt-get update && \
    apt-get install -y \
    curl \
    # graphviz \
    # libtinfo5 is needed by maude:
    libtinfo5 \
    locales \
    # we need python for our Tamarin oracle:
    python3 \
    # vim is only provided for convenience:
    vim \
    wget \
    zip \
    make \
    git

# install go
ENV GO_VERSION="1.22.2"
RUN mkdir tmp && \
    cd tmp && \
    wget "https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    cd ../ && \
    rm -rf tmp

# add Go to path:
ENV PATH="${PATH}:/usr/local/go/bin"
RUN go env -w GOBIN="/usr/local/go/bin"

# `/.local/bin` is the installation location of Tamarin and its dependencies
RUN mkdir -p /.local/bin
ENV PATH="${PATH}:/.local/bin"

# install Maude
ENV TAMARIN_RELEASE_VERSION_FOR_MAUDE="1.6.0"
ENV MAUDE_VERSION="2.7.1"
ENV MAUDE_URL="https://github.com/tamarin-prover/tamarin-prover/releases/download/$TAMARIN_RELEASE_VERSION_FOR_MAUDE/Maude-$MAUDE_VERSION-linux.zip"
RUN curl -q -s -S -L --create-dirs -o maude.zip $MAUDE_URL && \
    unzip maude.zip -d /.local/bin/ && \
    rm maude.zip && \
    mv /.local/bin/maude.linux64 /.local/bin/maude && \
    chmod a+x /.local/bin/maude

# install tamarin-prover
ENV TAMARIN_VERSION="1.8.0"
RUN curl -q -s -S -L --create-dirs -o tamarin.tar.gz https://github.com/tamarin-prover/tamarin-prover/releases/download/$TAMARIN_VERSION/tamarin-prover-$TAMARIN_VERSION-linux64-ubuntu.tar.gz && \
    tar -x -C /.local/bin/ -f tamarin.tar.gz && \
    rm tamarin.tar.gz && \
    chmod +x /.local/bin/tamarin-prover

# change locale to `en_US.UTF-8`:
USER root
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && \
    locale-gen
ENV LANG="en_US.UTF-8"
ENV LANGUAGE="en_US:en"
ENV LC_ALL="en_US.UTF-8"

# install argot dependencies
RUN go install golang.org/x/vuln/cmd/govulncheck@latest && \
    go install github.com/fzipp/gocyclo/cmd/gocyclo@latest && \
    go install github.com/gordonklaus/ineffassign@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install golang.org/x/tools/cmd/deadcode@latest

# install argot
RUN mkdir ar-go-tools && mkdir argot-proofs
copy ar-go-tools ./ar-go-tools
RUN cd ./ar-go-tools && make release
copy argot-proofs ./argot-proofs

RUN mkdir model-orig && mkdir implementation-orig

# copy model
copy model ./model-orig

# copy implementation
copy implementation ./implementation-orig

# copy shell scripts
COPY docker/*.sh ./

# remove some unneeded folders that come with the base image:
RUN rm -r evaluation
RUN rm -r tutorial-examples

RUN mkdir model && mkdir implementation

# disable entry point specified by the Gobra base container:
# ENTRYPOINT []
ENTRYPOINT cp -r model-orig/. model/ && \
    cp -r implementation-orig/. implementation/ && \
    /bin/bash
