#!/usr/bin/env bash

if [[ "$1" -eq "image" ]]
then
    PORT=$port \
    BUILD_MODE=$mode \
    docker run -p $port:$port locrep-go
else
    #run all tests
    PORT=$port BUILD_MODE=$mode  ./locrep-go.exe
fi
#rm -rf ~/.m2/repository/ && mvn clean install -Dmaven.repo.remote=http://localhost:8888
#rm -rf ~/.m2/repository/org/hamcrest/ ~/.m2/repository/junit/junit && mvn clean install -Dmaven.repo.remote=http://localhost:8888
