FROM postgres:10.5

WORKDIR /

ENV VERSION=1017
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres_password
ENV PGDATA=/database/data

COPY create-db-data.sh /docker-entrypoint-initdb.d/init-db-data.sh
COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY postgresql.conf /database/data/postgresql.conf
