#!/usr/bin/env bash

cd $(dirname $0)/..

go generate ./...
