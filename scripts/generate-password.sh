#!/bin/bash
source scripts/_utils.sh

loadENV
go run "cmd/generate-password/main.go" $1
unloadENV
