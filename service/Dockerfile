FROM node:12.0.0-alpine

ENV PORT=3000

WORKDIR /wyrdwest/service

# First download dependencies. This means that layer caching helps us out
COPY package.json yarn.lock /wyrdwest/service/
RUN yarn install

# Next do a build and test run. This ensures that the code is ready to execute
COPY . /wyrdwest/service/
RUN yarn test

CMD yarn start:main
