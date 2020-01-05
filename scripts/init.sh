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
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=<username>
POSTGRES_PASSWORD=<password>
POSTGRES_DBNAME=anime_skip
POSTGRES_DISABLE_SSL=true
POSTGRES_ENABLE_SEEDING=true

# Secrets
JWT_SECRET=<get_from_aaron>
EMAIL_PASSWORD=<get_from_aaron>" > .env
fi

./scripts/help.sh
