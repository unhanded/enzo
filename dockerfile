FROM golang:1.22-bookworm as BUILDER

WORKDIR /app

COPY . .

RUN go build ./app/enzod

FROM debian:latest as RUNNER

WORKDIR /app

COPY --from=BUILDER /app/enzod ./enzod

ENTRYPOINT [ "/app/enzod" ]
