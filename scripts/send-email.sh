#!/bin/bash
source scripts/_utils.sh

loadENV
go run "cmd/send-email/main.go"
unloadENV
