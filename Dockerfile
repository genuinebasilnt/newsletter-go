# Builder stage
FROM golang:1.23.1 AS builder

WORKDIR /app

COPY . .

RUN go build -o ./target/newsletter ./cmd/newsletter/main.go

# Runtime stage
FROM debian:bullseye-slim AS runtime

WORKDIR /app

COPY --from=builder /app/target/newsletter newsletter

COPY config config

ENV APP_ENVIRONMENT production

ENTRYPOINT [ "./target/newsletter" ]