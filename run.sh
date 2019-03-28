#!/usr/bin/env bash

#remove exe
rm -rf locrep-go

#build
go build .

#run
PORT=8888 BUILD_MODE=debug  ./locrep-go.exe

