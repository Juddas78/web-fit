FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download && go build .

ENTRYPOINT go run .