#!/usr/bin/env bash

if [ "$1" = "image" ]
then
    docker build -t locrep .
else
    #run all tests
    PORT=8888 BUILD_MODE=debug ginkgo -v -r

    go build
fi
