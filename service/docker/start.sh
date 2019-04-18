#!/bin/sh

set -x

CMD=
if [ ! -z "$DOCKERIZE_WAIT" ]; then
    CMD="dockerize $DOCKERIZE_WAIT -timeout 30s"
fi

$CMD /wyrdwest/service/wyrdwest_service
