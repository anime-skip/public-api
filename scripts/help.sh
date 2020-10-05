#!/bin/bash
source scripts/_utils.sh

header "Available \x1b[3mmake\x1b[0m\x1b[1m\x1b[94m rules"

function makeRule {
  bullet "$(codeBlock "$1") - $2"
}
makeRule "make build" "Build the source code to $(underline bin/anime-skip-api) (default, can be ran with just $(codeBlock make))"
makeRule "make run  " "Run the server"
makeRule "make watch" "Run the server and restart it when a file changes"
makeRule "make gen  " "Run gqlgen to generate GraphQL models and resolvers"
makeRule "make init " "Initialize a $(underline .env) file and show $(codeBlock make) commands"
makeRule "make help " "show $(codeBlock make) commands"

echo ""
