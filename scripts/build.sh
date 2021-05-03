#!/bin/bash
source scripts/_utils.sh

eval "go build $BUILD_ARGS -o bin/server cmd/server/main.go"
