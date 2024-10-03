FROM golang:1.23-alpine
RUN apk add make
RUN mkdir app

ADD . /app/

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
CMD ["air", "-c", "/app/.air.toml"]
