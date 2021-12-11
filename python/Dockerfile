 
FROM ubuntu:18.04

ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
ENV PATH /opt/conda/bin:$PATH

ENV NGINX_VERSION 1.17.8

RUN apt-get -y update

RUN apt-get install -y \
    build-essential git wget \
    cmake pkg-config ca-certificates openssl \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

RUN wget -