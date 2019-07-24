#!/bin/bash

export GOPATH=$(pwd):$(pwd)/vendor

mkdir bin/
cd src/
go build -o qqrobot main.go
mv qqrobot ../bin/