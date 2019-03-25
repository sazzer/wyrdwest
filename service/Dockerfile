# build stage
FROM golang:alpine AS build-env

RUN apk add --no-cache git mercurial

ADD . /go/src/github.com/sazzer/service/wyrdwest
WORKDIR /go/src/github.com/sazzer/service/wyrdwest

ENV GO111MODULE=on
ENV GOPATH=/go
ENV GOFLAGS=-mod=vendor

RUN go install github.com/sazzer/wyrdwest/service/cmd/service

# final stage
FROM alpine

ENV PORT=3000

WORKDIR /app
COPY --from=build-env /go/bin/service /app/wyrdwest
ENTRYPOINT ./wyrdwest

