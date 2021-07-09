#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER playground WITH PASSWORD 'password';
    CREATE DATABASE playground;
    GRANT ALL PRIVILEGES ON DATABASE playground TO playground;
EOSQL
