FROM postgres:14.5 as db
WORKDIR /trackmywealth
COPY ./scripts/dbinit.sh /docker-entrypoint-initdb.d
COPY ./scripts/tables.sql ./scripts/db/dump.sql