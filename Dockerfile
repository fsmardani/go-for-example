FROM golang:1.22.3

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . /usr/src/app
RUN go mod tidy
