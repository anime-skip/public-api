#!/bin/bash
source scripts/_utils.sh

loadENV
go build ./...
go test ./...
unloadENV
