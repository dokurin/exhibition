#!/usr/bin/env bash

cd $(dirname $0)/../../

docker image build -f ./build/develop/Dockerfile -t dokurin/exibition-devenv:1.0.0 .
