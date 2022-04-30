#!/bin/sh
set -e

shift

until PGPASSWORD=dbPassword psql -h "$dbAddress:$dbPort" -U "$dbUser" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec "$@"
