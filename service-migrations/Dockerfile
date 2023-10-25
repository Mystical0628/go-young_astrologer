FROM ubuntu:22.04

RUN apt-get update \
    && apt-get install -y curl gnupg lsb-release \
    && apt-get clean all

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
    && echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list \
    && apt-get update && apt-get install -y migrate


VOLUME /usr/src/migrations

WORKDIR /usr/src/migrations


ARG DB_DRIVER=postgres
ARG DB_HOST=localhost
ARG DB_PORT=5432
ARG DB_USER=astrologer
ARG DB_PASSWORD=123456768
ARG DB_DATABASE=postgres
ARG DB_DIRECTORY=postgres

ENV DB_URL="$DB_DRIVER://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable"
ENV DB_DIRECTORY=./$DB_DIRECTORY


ARG DB_NASA_DATABASE=nasa
ARG DB_NASA_DIRECTORY=nasa

ENV DB_NASA_URL="$DB_DRIVER://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NASA_DATABASE?sslmode=disable"
ENV DB_NASA_DIRECTORY=./$DB_NASA_DIRECTORY


CMD migrate -path $DB_DIRECTORY -database $DB_URL up \
    && migrate -path $DB_NASA_DIRECTORY -database $DB_NASA_URL up