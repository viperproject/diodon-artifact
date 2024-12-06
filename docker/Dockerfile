# Gobra commit 97f0988
FROM ghcr.io/viperproject/gobra@sha256:8d9f2b83c49a9a19f5a27167f138f23a2c4aa6677e7e9dbf157d83a110d6a088
# FROM ghcr.io/viperproject/gobra:latest

RUN apt-get update && \
    apt-get install -y \
    curl \
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
ENV GO_VERSION="1.23.1"
RUN mkdir tmp && \
    cd tmp && \
    wget "https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    cd ../ && \
    rm -rf tmp

# add Go to path:
ENV PATH="${PATH}:/usr/local/go/bin"
RUN go env -w GOBIN="/usr/local/go/bin"

# bypass Go proxy
RUN go env -w GOPROXY=direct

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

# copy argot source code
RUN mkdir ar-go-tools
COPY ar-go-tools ./ar-go-tools

# install argot binary: this puts the "argot" command on the PATH
RUN cd ./ar-go-tools && make argot-install


RUN mkdir /dh-orig && mkdir /ssm-agent-orig

# copy dh
COPY dh /dh-orig

# copy ssm-agent
COPY ssm-agent /ssm-agent-orig

# copy shell scripts
COPY docker/dh/*.sh /dh-orig
COPY docker/ssm-agent/*.sh /ssm-agent-orig

# remove some unneeded folders that come with the base image:
RUN rm -r evaluation
RUN rm -r tutorial-examples

RUN mkdir dh && mkdir ssm-agent

# disable entry point specified by the Gobra image and make sure to setup the directories
# by creating the following shell script and setting it as the image's entrypoint:
COPY --chmod=755 <<EOT /entrypoint.sh
#!/usr/bin/env bash
set -e
cp -r /dh-orig/. dh/
cp -r /ssm-agent-orig/. ssm-agent/
/bin/bash
EOT

ENTRYPOINT ["/entrypoint.sh"]
