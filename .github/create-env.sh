#!/bin/bash
cat > .env << EOL
PORT=8081
LOG_LEVEL=0
DATABASE_URL=postgres://postgres:password@postgres:5432/timestamps
DATABASE_DISABLE_SSL=true
DATABASE_VERSION=20
DATABASE_ENABLE_SEEDING=true
ENABLE_INTROSPECTION=true
ENABLE_PLAYGROUND=true
JWT_SECRET=some-secret
EMAIL_SERVICE_HOST=email-service
EMAIL_SERVICE_SECRET=some-secret
EMAIL_SERVICE_ENABLED=true
RECAPTCHA_SECRET=some-secret
RECAPTCHA_RESPONSE_ALLOWLIST=password1,password2
IS_SHOW_ADMIN_DISABLED=true
EOL
