#!/bin/sh

set -e

echo "run db migration"

while ! /app/migrate -path /app/migration -database "$DB_SOURCE" up; do
  echo "waiting for database..."
  sleep 1
done

echo "start app"

exec "$@"