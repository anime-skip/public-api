#!/bin/bash
source scripts/_utils.sh

loadENV
go run "cmd/anime-skip-backend/main.go"
unloadENV
