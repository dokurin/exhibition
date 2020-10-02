#!/usr/bin/env bash

cd $(dirname $0)/../..

docker container run -it --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd):/root/app \
  -p 8080:8080 \
  -w /root/app \
  dokurin/exibition-devenv:1.0.0 bash
