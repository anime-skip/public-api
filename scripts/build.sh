#!/bin/bash
set -e
INPUT="cmd/anime-skip-backend/main.go"
OUTPUT="bin/anime-skip-backend"

go build -o $OUTPUT $INPUT
SIZE="$(ls -lah $OUTPUT | awk '{print $5}')"

echo "$OUTPUT"
echo "Size: $SIZE"
