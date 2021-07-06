#!/bin/bash
echo "PORT=8081
LOG_LEVEL=0
LOG_SQL=false
ENABLE_INTROSPECTION=true
ENABLE_PLAYGROUND=true
IS_DEV=true
JWT_SECRET=1234

AWS_DATABASE_URL=postgres://username:password@host:port/dbname
DATABASE_DISABLE_SSL=true
DATABASE_ENABLE_SEEDING=true

DISABLE_EMAILS=false
EMAIL_SERVICE_HOST=not.anime-skip.com
EMAIL_SERVICE_SECRET=1234

RECAPTCHA_SECRET=1234
RECAPTCHA_RESPONSE_ALLOWLIST=1234

BETTER_VRV_APP_ID=1234
BETTER_VRV_API_KEY=1234" > .env
