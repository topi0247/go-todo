FROM golang:1.22

RUN apt-get update && apt-get install -y git postgresql-client

WORKDIR /app

COPY go.mod go.sum ./

RUN set -x \
    && go mod download \
    && go install github.com/cosmtrek/air@v1.48.0 \
    && go install github.com/rubenv/sql-migrate/...@latest \
    && go install github.com/volatiletech/sqlboiler/v4@latest \
    && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

COPY . .

ENTRYPOINT ["./entry-point.sh"]
