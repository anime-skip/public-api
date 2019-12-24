#!/bin/bash
function loadENV {
    export $(grep -v '^#' .env | xargs)
}

function unloadENV {
    unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
}
