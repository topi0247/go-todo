#!/bin/sh
set -o errexit

echo "DB_HOST: $DB_HOST"
echo "DB_USER: $DB_USER"
echo "DB_NAME: $DB_NAME"

until pg_isready -h "$DB_HOST" -U "$DB_USER"; do
  echo "Waiting for database to be ready..."
  sleep 2
done

sql-migrate up -env="production"
go run main.go
