FROM node:11.14.0-alpine

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /wyrdwest/integration
COPY package.json yarn.lock /wyrdwest/integration/
RUN yarn install

COPY tests /wyrdwest/integration/tests

ENV SERVICE_URI=

CMD dockerize -wait $SERVICE_URI/health -timeout 30s yarn run jest
