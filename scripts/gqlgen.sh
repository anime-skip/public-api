#!/bin/bash
source scripts/_utils.sh

rm -f internal/gql/main.go \
    internal/gql/models/generated.go \
    internal/gql/resolvers/update_then_delete_me.go
go run -v github.com/99designs/gqlgen $1
