FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download && cd src/ && go build .

ENTRYPOINT go run ./src