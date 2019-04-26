#!/bin/sh
set -x

START_CMD="dockerize -timeout 30s"

if [ ! -z "$DOCKERIZE_WAIT" ]; then
    START_CMD="$START_CMD $DOCKERIZE_WAIT"
fi

$START_CMD yarn start:main
