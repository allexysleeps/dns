# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ENV GO111MODULE=on

WORKDIR .

RUN mkdir app
RUN mkdir app/dns-proto
RUN mkdir app/dns-api

COPY go.mod ./app
COPY go.sum ./app

ADD /dns-proto/*.go ./app/dns-proto/
ADD /dns-api/ ./app/dns-api

WORKDIR app

RUN go mod download

WORKDIR dns-api

RUN go build -o /dns-api

CMD [ "/dns-api"]
EXPOSE 8080
