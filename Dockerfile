FROM --platform=linux/amd64 golang:alpine as build

ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --update curl && rm -rf /var/cache/apk/*
RUN mkdir -p /app

WORKDIR /app
COPY . . 
RUN go build -o xm

FROM --platform=linux/amd64 alpine:latest
WORKDIR /app

RUN apk add --no-cache bash

COPY --from=build /app/xm .

COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8090