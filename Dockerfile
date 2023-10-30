#!/bin/sh
FROM golang:1.16 as builder
MAINTAINER MKP Mobile Production <mkpproduction@gmail.com>
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN mkdir /app
#ADD . /app
WORKDIR /app
COPY . .

RUN go build -o main .

FROM alpine:3.13.1

WORKDIR /app
RUN touch .env

COPY --from=builder /app/main .
CMD ["/app/main"]
