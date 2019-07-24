#!/bin/bash

export GOPATH=${pwd}:${pwd}/vendor

cd src/
go build main.go