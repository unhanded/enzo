FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY . .

RUN go build ./app/enzod

FROM debian:latest AS runner

WORKDIR /app

COPY --from=builder /app/enzod ./enzod

ENTRYPOINT [ "/app/enzod" ]
