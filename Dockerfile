FROM golang:1.19-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY .env.example .env
COPY . .
RUN go mod tidy