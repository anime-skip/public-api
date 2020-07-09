#!/bin/bash
set -e
INPUT="cmd/send-email/main.go"
OUTPUT="bin/send-email"

go build -o $OUTPUT $INPUT
SIZE="$(ls -lah $OUTPUT | awk '{print $5}')"

echo "$OUTPUT"
echo "Size: $SIZE"
