#!/bin/bash
source scripts/_utils.sh

go build \
    -ldflags "-X anime-skip.com/backend/internal/utils/constants.VERSION=$VERSION -X anime-skip.com/backend/internal/utils/constants.VERSION_SUFFIX=$SUFFIX" \
    -o bin/server \
    cmd/server/main.go
