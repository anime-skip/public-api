#!/bin/bash
source scripts/_utils.sh

go run "cmd/generate-api-client/main.go" $1
