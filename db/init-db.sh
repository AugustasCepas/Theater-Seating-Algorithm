#!/bin/bash

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER admin;
    CREATE DATABASE theater_seating_database;
    GRANT ALL PRIVILEGES ON DATABASE newdb TO admin;
EOSQL

psql -f /db-dumps/db.dump.sql theater_seating_database