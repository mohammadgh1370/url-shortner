FROM golang:1.19-alpine

WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download