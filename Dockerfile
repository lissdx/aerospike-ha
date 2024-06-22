FROM golang:1.22-alpine AS build

COPY . /go/src/github.com/lissdx/aerospike-ha/

ENV GOOS=linux
ENV GO111MODULE=auto

WORKDIR /go/src/github.com/lissdx/aerospike-ha/

RUN go install github.com/lissdx/aerospike-ha/cmd/asrrdservice

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/bin/* .
COPY ./configs ./configs
COPY ./start.sh .
RUN chmod +x start.sh


