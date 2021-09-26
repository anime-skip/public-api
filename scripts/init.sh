#!/bin/bash
source scripts/_utils.sh

if [ "$OS" == "linux" ]; then
  HOST="localhost"
elif [ $OS == "mac" ]; then
  HOST="host.docker.internal"
fi

header "Creating .env"
if [ -f ".env" ]; then
    warning ".env already exists, skipping"
else
  echo "# Server
#
PORT=8081
LOG_LEVEL=0
IS_DEV=true
GIN_MODE=release
JWT_SECRET=can-be-anything

# Features
#
LOG_SQL=false
ENABLE_COLOR_LOGS=true
ENABLE_INTROSPECTION=true
ENABLE_PLAYGROUND=false
DISABLE_SHOW_ADMIN_DIRECTIVE=true

# Database
#
AWS_DATABASE_URL=postgres://postgres:password@localhost:9000/anime_skip_local
DATABASE_DISABLE_SSL=true
DATABASE_ENABLE_SEEDING=true
# DATABASE_MIGRATION=

# Emails
#
DISABLE_EMAILS=true
EMAIL_SERVICE_HOST=localhost:8082
EMAIL_SERVICE_SECRET=password

# reCAPTCHA
#
RECAPTCHA_SECRET=can-be-anything
RECAPTCHA_RESPONSE_ALLOWLIST=password1,password2

# Third Party Services
#
BETTER_VRV_APP_ID=<get-from-aaron>
BETTER_VRV_API_KEY=<get-from-aaron>" > .env
fi

./scripts/help.sh
