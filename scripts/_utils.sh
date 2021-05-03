#!/bin/bash
function loadENV {
    export $(grep -v '^#' .env | xargs)
}

function unloadENV {
    unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
}

function header {
    echo ""
    echo -e "\x1b[1m\x1b[94m$1\x1b[0m"
}

function warning {
    echo -e " \x1b[93m▶ $1\x1b[0m"
}

function bullet {
    echo -e " • $1"
}

function underline {
    echo -en "\x1b[4m$1\x1b[0m"
}

function codeBlock {
    echo -en "\x1b[1m\x1b[2m\x1b[3m$1\x1b[0m"
}

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     OS=linux;;
    Darwin*)    OS=mac;;
    CYGWIN*)    OS=windows;;
    MINGW*)     OS=windows;;
    *)          OS="UNKNOWN:${unameOut}"
esac

VERSION=$(jq -r .version meta.json) ;
SUFFIX="-$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')" ;
VERSION_COMPILER_FLAG="-X anime-skip.com/backend/internal/utils/constants.VERSION=$VERSION"
SUFFIX_COMPILER_FLAG="-X anime-skip.com/backend/internal/utils/constants.VERSION_SUFFIX=$SUFFIX"
BUILD_ARGS="-ldflags=\"$VERSION_COMPILER_FLAG $SUFFIX_COMPILER_FLAG\""

loadENV
trap unloadENV EXIT
