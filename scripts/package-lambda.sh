#!/bin/bash
source scripts/_utils.sh

echo "Compiling..."
GOOS=linux CGO_ENABLED=0 eval "go build $BUILD_ARGS -o bin/lambdas/main cmd/lambda/main.go"

echo "Zipping..."
rm -rf bin/lambdas/*.zip
zip bin/lambdas/main.zip bin/lambdas/main

echo "Done!"
