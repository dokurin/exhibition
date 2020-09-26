#!/usr/bin/env bash

cd $(dirname $0)

docker image build -t dokurin/exibition-devenv:1.0.0 .
