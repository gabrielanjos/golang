FROM postgres:11.5-alpine

COPY ./basedados/bdgo.sql /docker-entrypoint-initdb.d/

ENV PGDATA=/data