#!/bin/bash
source scripts/_utils.sh

loadENV
go run "cmd/anime-skip-api/main.go"
unloadENV
