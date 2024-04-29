FROM rust:slim AS builder

RUN apt-get update && apt-get install -y protobuf-compiler && rm -rf /var/lib/apt/lists/*

COPY ./proto ./proto
COPY ./build.rs ./build.rs
COPY ./Cargo.toml ./Cargo.toml
COPY ./Cargo.lock ./Cargo.lock
COPY ./src ./src


RUN cargo build --release
RUN mv "target/release/notification" /notification

FROM scratch
COPY --from=builder /notification /bin/notification
EXPOSE 3000
COPY .env.sample .env
RUN ["/bin/notification"]