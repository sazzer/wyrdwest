FROM rust:1.34 AS build

WORKDIR /wyrdwest/service
COPY Cargo.* /wyrdwest/service/
RUN mkdir src && touch src/lib.rs
RUN cargo fetch

RUN rm -rf /wyrdwest/service/src
COPY src /wyrdwest/service/src
COPY benches /wyrdwest/service/benches
RUN cargo build

FROM rust:1.34

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /wyrdwest/service
COPY --from=build /wyrdwest/service/target/debug/wyrdwest_service /wyrdwest/service/
COPY log4rs.yml /wyrdwest/service/
COPY docker/start.sh /wyrdwest/service
COPY migrations/* /wyrdwest/service/migrations/

CMD /wyrdwest/service/start.sh
