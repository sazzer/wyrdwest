FROM node:12.0.0-alpine

# Install Dockerize
RUN apk add --no-cache openssl gettext

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

ENV PORT=3000

WORKDIR /wyrdwest/service

# First download dependencies. This means that layer caching helps us out
COPY package.json yarn.lock /wyrdwest/service/
RUN yarn install

# Next do a build and test run. This ensures that the code is ready to execute
COPY . /wyrdwest/service/
RUN yarn test

COPY docker/start.sh /wyrdwest/service

CMD /wyrdwest/service/start.sh
