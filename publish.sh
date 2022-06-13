#!/bin/bash
function title {
    echo -e "\n\x1b[0m\x1b[1m\x1b[94m$1...\x1b[0m"
}

title "Logging into heroku"
heroku auth:whoami
set -e
if [[ "$?" != "0" ]]; then
    heroku login
fi
heroku container:login
docker login --username=_ --password=$(heroku auth:token) registry.heroku.com

title "Getting version"
VERSION="$(jq -r .version meta.json)-$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')"
echo "$VERSION"

title "Building image"
docker build . \
        --build-arg VERSION=$VERSION \
        --build-arg STAGE=production \
        -t registry.heroku.com/prod-public-api/web \

title "Deploying"
docker push registry.heroku.com/prod-public-api/web
heroku container:release -a prod-public-api web
