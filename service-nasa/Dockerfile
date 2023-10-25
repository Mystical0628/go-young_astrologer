FROM ubuntu:22.04

RUN apt-get update \
    && apt-get install -y ca-certificates

VOLUME /usr/local/bin
WORKDIR /usr/local/bin

CMD app