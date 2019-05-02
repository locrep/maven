#!/usr/bin/env bash

if [[ $1 -eq "image" ]]; then
    docker build -t locrep .
else
echo "local"
    #run all tests
    PORT=8888 BUILD_MODE=debug ginkgo -v -r

    go build
fi