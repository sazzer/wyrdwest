FROM cypress/base:10

ENV CI=1

# Install Dockerize
RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /wyrdwest/e2e

# Download dependencies. This means that layer caching helps us out
COPY package.json yarn.lock /wyrdwest/e2e/
RUN yarn install
RUN npx cypress install

ENV CYPRESS_SERVICE_URI=

# Next do a build and test run. This ensures that the code is ready to execute
COPY . /wyrdwest/e2e/
RUN npx cypress verify
RUN yarn test

CMD dockerize -timeout 30s -wait $CYPRESS_SERVICE_URI/health yarn start:cypress
