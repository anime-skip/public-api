#!/bin/bash
source scripts/_utils.sh

if [ "$OS" == "linux" ]; then
  docker run --network=host --env-file .env -p 8081:8081 anime-skip/backend/api:dev
elif [ "$OS" == "mac" ]; then
  docker run --env-file .env -p 8081:8081 anime-skip/backend/api:dev
else
  echo "TODO: $OS run script"
fi
