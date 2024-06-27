FROM golang:1.22-bookworm as BUILDER

WORKDIR /app

COPY . .

RUN go build ./cmd/enzod

FROM debian:latest as RUNNER

WORKDIR /app

COPY --from=BUILDER /app/enzod ./enzod

ENTRYPOINT [ "/app/enzod" ]
