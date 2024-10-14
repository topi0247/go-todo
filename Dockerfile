FROM golang:1.22

RUN apt-get update && apt-get install -y libpq-dev git

WORKDIR /app
COPY . .
