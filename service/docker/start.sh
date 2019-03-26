#!/bin/sh

CMD="dockerize -timeout 30s"

echo Dockerize PGSQL: $DOCKERIZE_PGSQL
if [ ! -z "$DOCKERIZE_PGSQL" ]; then
    CMD="$CMD -wait $DOCKERIZE_PGSQL"
fi

echo Starting...
echo cmd: $CMD

$CMD /app/wyrdwest
