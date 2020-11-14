#!/bin/bash
source scripts/_utils.sh

header "Creating .env"
if [ -f ".env" ]; then
    warning ".env already exists, skipping"
else
  echo "# General
IS_DEV=true
ENABLE_PLAYGROUND=true
ENABLE_INTROSPECTION=true
LOG_LEVEL=0
LOG_SQL=false
ENABLE_COLOR_LOGS=true

# Web Server
HOST=localhost
PORT=8000
GIN_MODE=release

# Postgres
DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<dbname>
DATABASE_DISABLE_SSL=true
DATABASE_ENABLE_SEEDING=true

# Secrets
JWT_SECRET=<get_from_aaron>" > .env
fi

./scripts/help.sh
