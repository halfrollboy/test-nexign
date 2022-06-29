FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run cmd/test-nexign/main.go
EXPOSE 8080

