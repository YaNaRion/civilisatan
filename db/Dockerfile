FROM postgres:latest

ENV POSTGRES_USER=inf3995
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=server-db

COPY init.sql /docker-entrypoint-initdb.d/

