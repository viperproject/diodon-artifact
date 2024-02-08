# Gobra commit a2915f9
# FROM ghcr.io/viperproject/gobra@sha256:74c99e35b5e51eefe38931717b373173cddf7931ec13e97cb202b7d739a0664b
FROM ghcr.io/viperproject/gobra:latest

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
    zip

# install go
ENV GO_VERSION 1.22.0
RUN mkdir tmp && \
    cd tmp && \
    wget "https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz" && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
    cd ../ && \
    rm -rf tmp

# add Go to path:
ENV PATH="${PATH}:/usr/local/go/bin"

# `/.local/bin` is the installation location of Tamarin and its dependencies
RUN mkdir -p /.local/bin
ENV PATH=/.local/bin:$PATH

# install Maude
ENV TAMARIN_RELEASE_VERSION_FOR_MAUDE 1.6.0
ENV MAUDE_VERSION 2.7.1
ENV MAUDE_URL https://github.com/tamarin-prover/tamarin-prover/releases/download/$TAMARIN_RELEASE_VERSION_FOR_MAUDE/Maude-$MAUDE_VERSION-linux.zip
RUN curl -q -s -S -L --create-dirs -o maude.zip $MAUDE_URL && \
    unzip maude.zip -d /.local/bin/ && \
    rm maude.zip && \
    mv /.local/bin/maude.linux64 /.local/bin/maude && \
    chmod a+x /.local/bin/maude

# install tamarin-prover
ENV TAMARIN_VERSION 1.8.0
RUN curl -q -s -S -L --create-dirs -o tamarin.tar.gz https://github.com/tamarin-prover/tamarin-prover/releases/download/$TAMARIN_VERSION/tamarin-prover-$TAMARIN_VERSION-linux64-ubuntu.tar.gz && \
    tar -x -C /.local/bin/ -f tamarin.tar.gz && \
    rm tamarin.tar.gz && \
    chmod +x /.local/bin/tamarin-prover

# change locale to `en_US.UTF-8`:
USER root
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && \
    locale-gen
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8

RUN mkdir model

# copy model
copy model ./model

# copy shell scripts
COPY docker/*.sh ./

# remove some unneeded folders that come with the base image:
RUN rm -r evaluation
RUN rm -r tutorial-examples

RUN mv model model-orig

RUN mkdir model

# disable entry point specified by the Gobra base container:
# ENTRYPOINT []
ENTRYPOINT cp -r model-orig/. model/ && \
    /bin/bash
