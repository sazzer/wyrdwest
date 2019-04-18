FROM rust:1.34 AS build

WORKDIR /wyrdwest/service
COPY Cargo.* /wyrdwest/service/
RUN mkdir src && touch src/lib.rs
RUN cargo fetch
RUN cargo build --release

RUN rm -rf /wyrdwest/service/src
COPY src /wyrdwest/service/src
RUN cargo build --release

FROM rust:1.34

WORKDIR /wyrdwest/service
COPY --from=build /wyrdwest/service/target/release/wyrdwest_service /wyrdwest/service/
COPY log4rs.yml /wyrdwest/service/

CMD /wyrdwest/service/wyrdwest_service