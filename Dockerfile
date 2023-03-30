FROM golang:1.16.4-buster AS builder

ARG VERSION=dev

FROM golang:alpine AS build
ARG VERSION=dev
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go build -o /tmp/http-server -ldflags=-X=main.version=${VERSION} ./main.go

FROM alpine:edge
COPY --from=build /tmp/http-server /sbin/http-server
CMD /sbin/http-server