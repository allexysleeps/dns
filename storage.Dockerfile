# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ENV GO111MODULE=on

WORKDIR .

RUN mkdir app
RUN mkdir app/dns-proto
RUN mkdir app/dns-storage

COPY go.mod ./app
COPY go.sum ./app

ADD /dns-proto/*.go ./app/dns-proto/
ADD /dns-storage/ ./app/dns-storage

WORKDIR app

RUN go mod download

WORKDIR dns-storage

RUN go build -o /dns-storage

CMD [ "/dns-storage"]
EXPOSE 50052

