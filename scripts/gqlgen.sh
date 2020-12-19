#!/bin/bash
rm -f internal/graphql/main.go \
    internal/graphql/models/generated.go \
    internal/graphql/resolvers/update_then_delete_me.go
go run -v github.com/99designs/gqlgen $1
