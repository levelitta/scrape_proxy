FROM golang:1.22-alpine

RUN apk --no-cache add bash

RUN apk add git

WORKDIR /app

RUN export PATH=$PATH:/app/bin

COPY / ./

RUN go mod download -x