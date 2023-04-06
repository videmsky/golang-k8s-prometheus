FROM golang:alpine AS build
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go build -o /tmp/http-server ./main.go

FROM alpine:edge
COPY --from=build /tmp/http-server /sbin/http-server
ADD src/views ./views
CMD /sbin/http-server