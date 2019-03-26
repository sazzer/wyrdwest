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

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ENV PORT=3000

WORKDIR /app
COPY --from=build-env /go/bin/service /app/wyrdwest
COPY docker/start.sh /app/start.sh
CMD /app/start.sh

