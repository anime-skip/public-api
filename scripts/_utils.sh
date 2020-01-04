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
